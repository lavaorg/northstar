/*
Copyright (C) 2017 Verizon. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package portal

import (
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/lavaorg/lrtx/management"
	"github.com/lavaorg/lrtx/middleware"
	"github.com/lavaorg/lrtx/mlog"
	"github.com/lavaorg/northstar/portal/handler"
	"github.com/lavaorg/northstar/portal/model"
	"github.com/lavaorg/northstar/portal/portalglobal"
)

//Service defines the base structure to handle the portal service.
type Service struct {
	controller *handler.Controller
	engine     *gin.Engine
}

//NewService intializes the portal service.
func NewService() (*Service, error) {
	mlog.Debug("NewService")

	err := portalglobal.Load()
	if err != nil {
		return nil, err
	}

	controller, err := handler.NewController()
	if err != nil {
		mlog.Error("Error, failed to create portal service controller with error %s.\n", err.Error())
		return nil, err
	}

	management.SetHealth(controller.GetHealth())
	accessControl := middleware.AccessControl{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Accept", "Content-Type", "Authorization", "Etag", "If-Match", "If-No-Match"},
	}

	engine := management.Engine()

	engine.Use(middleware.Cors(accessControl))
	engine.NoRoute(controller.NotFound)

	// Service Web Application Routes

	// route config for webpack
	engine.Any("/northstar", controller.GetIndexHTML)
	engine.Any("/northstar/", controller.GetIndexHTML)
	engine.Static("/northstar/dist/", "./web/dist/")
	engine.Static("/northstar/assets/", "./web/dist/assets/")

	users := engine.Group("/users")
	{
		users.POST("actions/login", controller.Login)
		users.POST("actions/logout", controller.Logout)
		users.POST("actions/verify", controller.Verify)
	}

	// Service Internal RESTful API
	internal := engine.Group(path.Join(model.InternalContext, model.Version))
	{
		internal.POST("/connections/:connectionId/events/:eventId/callbacks/:type", controller.EventCallback)
	}

	// Service RESTful API
	v1 := engine.Group(path.Join(model.Context, model.Version))
	v1.Use(Authorization)
	{
		// Users API
		user := v1.Group("/user")
		{
			user.POST("/actions/query", controller.QueryUsers)
			user.GET("/", controller.GetUser)
		}

		// Websocket API
		connections := v1.Group("/connections")
		{
			connections.GET("/", controller.GetConnection)
		}

		// Nootebooks API
		notebook := v1.Group("/notebooks")
		{
			notebook.POST("/", controller.CreateNotebook)
			notebook.PUT("/", controller.UpdateNotebook)
			notebook.GET("/", controller.ListNotebooks)
			notebook.GET("/:notebookId", controller.GetNotebook)
			notebook.DELETE("/:notebookId", controller.DeleteNotebook)
			notebook.GET("/:notebookId/users", controller.GetNotebookUsers)
			notebook.PUT("/:notebookId/users", controller.UpdateNotebookUsers)
		}

		// Portfolios API
		portfolio := v1.Group("/portfolios")
		{
			portfolio.GET("/", controller.ListPortfolios)
			portfolio.GET("/:portfolioName", controller.ListFiles)
			portfolio.GET("/:portfolioName/*fileName", controller.GetFile)
		}

		// Transformation API
		transformations := v1.Group("transformations")
		{
			transformations.POST("/", controller.CreateTransformation)
			transformations.GET("/", controller.ListTransformations)
			transformations.GET("/:id", controller.GetTransformation)
			transformations.GET("/:id/results", controller.GetTransformationResults)
			transformations.PUT("/", controller.UpdateTransformation)
			transformations.DELETE("/:id", controller.DeleteTransformation)
		}

		// Template API
		templates := v1.Group("templates")
		{
			templates.GET("/", controller.ListTemplates)
		}

		// Schema endpoint
		v1.GET("/events/schemas", controller.GetScheduleEventSchemas)
	}

	service := &Service{
		controller: controller,
		engine:     engine,
	}

	return service, nil
}

//Start starts initializes the HTTP endpoints and management library
func (service *Service) Start() (err error) {
	mlog.Debug("Start")

	return management.Listen(fmt.Sprintf(":%s", portalglobal.Config.Port))
}

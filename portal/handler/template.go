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

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavaorg/lrtx/management"
	"github.com/lavaorg/lrtx/mlog"
	"github.com/lavaorg/northstar/portal/model"
)

// ListTemplates returns templates associated with authenticated user.
func (controller *Controller) ListTemplates(context *gin.Context) {
	mlog.Debug("ListTemplates")

	token, tokenExists := controller.getAuthToken(context)
	if !tokenExists {
		mlog.Error("Failed to get access token value")
		controller.RenderServiceError(context, model.ErrorInvalidCookie)
		return
	}

	templates, mErr := controller.portalProvider.ListTemplates(token.AccessToken)
	if mErr != nil {
		mlog.Error("Failed to get template table with error: %s", mErr.Description)
		controller.RenderServiceError(context, management.ErrorInternal)
		return
	}

	context.JSON(http.StatusOK, templates)
}

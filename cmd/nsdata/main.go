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

package main

import (
	"os"

	"github.com/lavaorg/lrtx/management"
	"github.com/lavaorg/lrtx/mlog"
	"github.com/lavaorg/northstar/data/cron"
	"github.com/lavaorg/northstar/data/datasets"
	"github.com/lavaorg/northstar/data/datasources"
	"github.com/lavaorg/northstar/data/env"
	"github.com/lavaorg/northstar/data/events"
	"github.com/lavaorg/northstar/data/invocations"
	"github.com/lavaorg/northstar/data/mappings"
	"github.com/lavaorg/northstar/data/notebooks"
	"github.com/lavaorg/northstar/data/snippets"
	"github.com/lavaorg/northstar/data/stream"
	"github.com/lavaorg/northstar/data/templates"
)

type DataService interface {
	AddRoutes()
}

func main() {
	dataPort, err := env.GetDataPort()
	if err != nil {
		mlog.Error("Failed to load data port", err)
		os.Exit(-1)
	}

	var dataService DataService

	dataService = new(snippets.SnippetService)
	dataService.AddRoutes()

	dataService = new(invocations.InvocationService)
	dataService.AddRoutes()

	dataService = new(events.EventsService)
	dataService.AddRoutes()

	dataService = new(mappings.MappingsService)
	dataService.AddRoutes()

	dataService = new(cron.CronService)
	dataService.AddRoutes()

	dataService = new(notebooks.NotebookService)
	dataService.AddRoutes()

	dataService = new(datasources.DatasourcesService)
	dataService.AddRoutes()

	dataService = new(datasets.DatasetsService)
	dataService.AddRoutes()

	dataService = new(templates.TemplateService)
	dataService.AddRoutes()

	dataService = new(stream.StreamService)
	dataService.AddRoutes()

	port := ":" + dataPort
	if err := management.Listen(port); err != nil {
		mlog.Error("Error starting api service", err)
		os.Exit(-1)
	}
}

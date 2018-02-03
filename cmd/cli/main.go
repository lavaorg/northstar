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

	"github.com/lavaorg/lrt/x/mlog"
	"github.com/lavaorg/northstar/cli/parser"
	cron "github.com/lavaorg/northstar/cron/client"
	cronData "github.com/lavaorg/northstar/data/cron/client"
	datasetsData "github.com/lavaorg/northstar/data/datasets/client"
	datasourcesData "github.com/lavaorg/northstar/data/datasources/client"
	eventsData "github.com/lavaorg/northstar/data/events/client"
	invocationData "github.com/lavaorg/northstar/data/invocations/client"
	mappingData "github.com/lavaorg/northstar/data/mappings/client"
	snippetsData "github.com/lavaorg/northstar/data/snippets/client"
	kafkaMngr "github.com/lavaorg/northstar/kafkamngr/client"
	object "github.com/lavaorg/northstar/object/client"
	eventsProcessing "github.com/lavaorg/northstar/processing/events/client"
	snippetsProcessing "github.com/lavaorg/northstar/processing/snippets/client"
)

func main() {
	kafkaMngr, mErr := kafkaMngr.NewKafkaMngrClient()
	if mErr != nil {
		mlog.Error("Failed to create kafkaMngr client: %v", mErr)
		os.Exit(-1)
	}

	snippetsProcessing, mErr := snippetsProcessing.NewSnippetsClient()
	if mErr != nil {
		mlog.Error("Failed to create snippet processing client: %v", mErr)
		os.Exit(-1)
	}

	eventsProcessing, mErr := eventsProcessing.NewEventsClient()
	if mErr != nil {
		mlog.Error("Failed to create events processing client: %v", mErr)
		os.Exit(-1)
	}

	snippetsData, mErr := snippetsData.NewSnippetClient()
	if mErr != nil {
		mlog.Error("Failed to create snippets data client: %v", mErr)
		os.Exit(-1)
	}

	invocationData, mErr := invocationData.NewInvocationClient()
	if mErr != nil {
		mlog.Error("Failed to create invocations data client: %v", mErr)
		os.Exit(-1)
	}

	eventsData, mErr := eventsData.NewEventsClient()
	if mErr != nil {
		mlog.Error("Failed to create events data client: %v", mErr)
		os.Exit(-1)
	}

	mappingData, mErr := mappingData.NewMappingsClient()
	if mErr != nil {
		mlog.Error("Failed to create mappings data client: %v", mErr)
		os.Exit(-1)
	}

	datasetsData, mErr := datasetsData.NewDatasetsClient()
	if mErr != nil {
		mlog.Error("Failed to create datasets data client: %v", mErr)
		os.Exit(-1)
	}

	datasourcesData, mErr := datasourcesData.NewDatasourcesClient()
	if mErr != nil {
		mlog.Error("Failed to create datasources data client: %v", mErr)
		os.Exit(-1)
	}

	cron, mErr := cron.NewCronClient()
	if mErr != nil {
		mlog.Error("Failed to create cron client: %v", mErr)
		os.Exit(-1)
	}

	cronData, mErr := cronData.NewCronClient()
	if mErr != nil {
		mlog.Error("Failed to create cron data client: %v", mErr)
		os.Exit(-1)
	}

	object, mErr := object.NewObjectClient()
	if mErr != nil {
		mlog.Error("Failed to create cron client: %v", mErr)
		os.Exit(-1)
	}

	parser.InitParser(kafkaMngr,
		snippetsProcessing,
		eventsProcessing,
		snippetsData,
		invocationData,
		eventsData,
		mappingData,
		datasetsData,
		datasourcesData,
		cron,
		cronData,
		object)
}

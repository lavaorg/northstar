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

package nsapiglobal

import (
	"encoding/json"
	"fmt"
	"github.com/lavaorg/lrt/env"
	"github.com/lavaorg/lrtx/middleware"
	"github.com/lavaorg/lrtx/mlog"
	"os"
)

// Defines the service name.
const ServiceName = "nsapi"

var Config configuration

type configuration struct {
	ServiceName          string `default:"nsapi"`
	Scopes               middleware.AllowedServices
	ScopesFileName       string   `default:"nsapi_scopes"`
	DataHostPort         string   `default:"8090"`
	ProcessingHostPort   string   `require:"true"`
	CronHostPort         string   `require:"true"`
	KafkaManagerHostPort string   `require:"true"`
	KafkaBrokersHostPort []string `require:"true"`
	ZookeeperHostPort    string   `require:"true"`
	MaxMemory            uint64   `default:"6000"`
	DefaultMemory        uint64   `default:"1000"`
	AcctAuthHostPort     string   `require:"true"`
	AcctClientId         string   `require:"true"`
	AcctClientSecret     string   `require:"true"`
	MaxArgCount          int      `default:"10"`
	MaxTimeout           int      `default:"180"`
	MaxCodeSize          int      `default:"68916"`
	EnforceChecksum      bool     `default:"false"`
}

// Loads service configuration from the environment variables.
func Load() error {

	err := env.Load("nsapi", &Config)
	if err != nil {
		return err
	}

	// Open the access configuration file
	scopesFile, err := os.Open(Config.ScopesFileName)
	if err != nil {
		return fmt.Errorf("Failed to load the scopes file %s with error: %v.", Config.ScopesFileName, err)
	}

	// Decode expected json.
	decoder := json.NewDecoder(scopesFile)
	if err := decoder.Decode(&Config.Scopes); err != nil {
		return fmt.Errorf("Failed to decode scopes file %s with error: %v", Config.ScopesFileName, err)
	}

	mlog.Info("environment settings: %v", Config)
	return nil
}

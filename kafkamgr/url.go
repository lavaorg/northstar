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
package kafkamgr

import (
	"errors"
	"os"
)

const KafkaMngrBasePath = "/kafkamngr/v1"

func GetKafkaMngrHostPort() (string, error) {
	hostPort := os.Getenv("KAFKA_MNGR_HOST_PORT")
	if hostPort == "" {
		return "", errors.New("Please set KAFKA_MNGR_HOST_PORT!")
	}

	return hostPort, nil
}

func GetKafkaMngrBaseUrl() (string, error) {
	hostPort, err := GetKafkaMngrHostPort()
	if err != nil {
		return "", err
	}

	return "http://" + hostPort, nil
}

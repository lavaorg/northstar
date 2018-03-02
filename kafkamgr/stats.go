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

import "github.com/lavaorg/lrtx/stats"

var (
	s             = stats.New("kafkamngr")
	CreateTopic   = s.NewCounter("CreateTopic")
	GetTopicNames = s.NewCounter("GetTopicNames")
	UpdateTopic   = s.NewCounter("UpdateTopic")

	ErrCreateTopic   = s.NewCounter("ErrCreateTopic")
	ErrGetTopicNames = s.NewCounter("ErrGetTopicNames")
	ErrUpdateTopic   = s.NewCounter("ErrUpdateTopic")
)

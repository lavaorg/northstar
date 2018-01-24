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

package model

import (
	"encoding/json"
	"errors"
)

const (
	// Defines the string used to identify a resource as a NorthStar Schedule resource.
	ScheduleKind string = "ts.schedule"

	// Defines the Schedule resource schema version. Note that the format
	// of this value is MAJOR.MINOR where:
	//	MAJOR - Matches the version of the Service RESTful API
	//	MINOR - Represents the resource schema version.
	ScheduleSchemaVersion string = "1.0"
)

// Defines the supported event categories. E.g., Device
// events represents an event type generated by a device.
const (
	DeviceEvent string = "Device"
	TimerEvent  string = "Timer"
)

// Defines the type that describes the events and times
// that trigger a transformation.
type Schedule struct {
	Kind        string `json:"kind,omitempty"`
	Id          string `json:"id,omitempty"`
	Version     string `json:"version,omitempty"`
	CreatedOn   string `json:"createdOn,omitempty"`
	LastUpdated string `json:"lastUpdated,omitempty"`
	Event       Event  `json:"event"`
}

// Defines the type that represents an event.
type Event struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

// Define internal type used for serialization.
type typeSchedule Schedule

// Helper method used to marshal schedule while setting default values.
func (schedule *Schedule) MarshalJSON() ([]byte, error) {
	var value typeSchedule

	value = typeSchedule(*schedule)

	// Make sure resource kind and version are set.
	value.Kind = ScheduleKind
	value.Version = ScheduleSchemaVersion

	return json.Marshal(&value)
}

// Helper method used to unmarshal schedule while validating required fields.
func (schedule *Schedule) UnmarshalJSON(data []byte) error {
	var value typeSchedule

	// Unmarshal to the internal type
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	*schedule = Schedule(value)

	// Validate event name.
	if schedule.Event.Name == "" {
		return errors.New("The event name is missing.")
	}

	// Validate event category and value.
	switch schedule.Event.Category {
	case DeviceEvent:
	case TimerEvent:
		if schedule.Event.Value == "" {
			return errors.New("The timer event value is missing.")
		}
	default:
		return errors.New("The category is invalid.")
	}

	return nil
}

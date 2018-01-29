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

package notebooks

import "github.com/lavaorg/lrt/x/stats"

var (
	s              = stats.New("notebooksdata")
	InsertNotebook = s.NewCounter("InsertNotebook")
	UpdateNotebook = s.NewCounter("UpdateNotebook")
	GetNotebook    = s.NewCounter("GetNotebook")
	DelNotebook    = s.NewCounter("DelNotebook")

	ErrInsertNotebook = s.NewCounter("ErrInsertNotebook")
	ErrUpdateNotebook = s.NewCounter("ErrUpdateNotebook")
	ErrGetNotebook    = s.NewCounter("ErrGetNotebook")
	ErrDelNotebook    = s.NewCounter("ErrDelNotebook")

	InsertAccess = s.NewCounter("InsertAccess")
	UpdateAccess = s.NewCounter("UpdateAccess")
	GetAccess    = s.NewCounter("GetAccess")
	DelAccess    = s.NewCounter("DelAccess")

	ErrInsertAccess = s.NewCounter("ErrInsertAccess")
	ErrUpdateAccess = s.NewCounter("ErrUpdateAccess")
	ErrGetAccess    = s.NewCounter("ErrGetAccess")
	ErrDelAccess    = s.NewCounter("ErrDelAccess")
)

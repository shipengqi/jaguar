// Copyright (c) 2022 PengQi Shi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package skeletons

import "embed"

var (
	//go:embed all:v1/api
	V1API embed.FS

	//go:embed all:v1/cli
	V1CLI embed.FS

	//go:embed all:v1/grpc
	V1GRPC embed.FS

	//go:embed all:v1/projectfiles
	V1ProjectFiles embed.FS
)

type Skeleton struct {
	API, CLI, GRPC, ProjectFiles embed.FS
}

// Skeletons represents struct for embed files.
type Skeletons struct {
	V1 Skeleton
}

// New creates a new collection with embed files by Attachments struct.
func New() *Skeletons {
	return &Skeletons{
		V1: Skeleton{
			API:          V1API,
			CLI:          V1CLI,
			GRPC:         V1GRPC,
			ProjectFiles: V1ProjectFiles,
		},
	}
}

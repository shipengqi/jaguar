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

package v1

import "embed"

//go:embed api
var API embed.FS

//go:embed cli
var CLI embed.FS

//go:embed grpc
var GRPC embed.FS

//go:embed projectfiles/.gitignore
var GitIgnore embed.FS

//go:embed projectfiles/.golangci.yaml
var GoCI embed.FS

//go:embed projectfiles/.goreleaser.yaml
var Releaser embed.FS

//go:embed projectfiles/.gsemver.yaml
var Semver embed.FS

//go:embed projectfiles/.github
var GitHubRepoFiles embed.FS

//go:embed projectfiles/hack
var Hack embed.FS

//go:embed projectfiles/Makefile
var Makefile embed.FS

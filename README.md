# Jaguar

A unified AI-assisted scaffold for generating production-ready project skeletons across multiple languages and frameworks.

[![e2e test](https://github.com/shipengqi/jaguar/actions/workflows/e2e.yaml/badge.svg)](https://github.com/shipengqi/jaguar/actions/workflows/e2e.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/shipengqi/jaguar)](https://goreportcard.com/report/github.com/shipengqi/jaguar)
[![release](https://img.shields.io/github/release/shipengqi/jaguar.svg)](https://github.com/shipengqi/jaguar/releases)
[![license](https://img.shields.io/github/license/shipengqi/jaguar)](https://github.com/shipengqi/jaguar/blob/main/LICENSE)

## Overview

Jaguar has evolved from a Go-only project scaffold into a general-purpose scaffold designed for AI-assisted development workflows. Rather than generating opinionated, version-pinned boilerplate, Jaguar produces lean skeleton projects with minimal dependencies — letting AI tools (like Claude Code, GitHub Copilot, etc.) fill in the implementation details tailored to each project's needs.

**Supported project types:**

| Type | Description |
|---|---|
| `go-api` | Go REST API server (Gin) with SQLite/MySQL store |
| `go-cli` | Go CLI application (Cobra) |
| `go-grpc` | Go gRPC service |
| `go-embed` | Go REST API server with embedded React/Vue/Angular frontend |
| `nodejs` | Node.js application |
| `python` | Python application |

## Installation

### From the Binary Releases

Download the pre-compiled binaries from the [releases page](https://github.com/shipengqi/jaguar/releases) and copy them to the desired location.

```
$ jaguar --version
```

### Go Install

You must have a working Go environment:

```
$ go install github.com/shipengqi/jaguar@latest
```

### From Source

You must have a working Go environment:

```
$ git clone https://github.com/shipengqi/jaguar.git
$ make build
```

## Usage

### Create a new project

Interactive mode (prompts for project type, module name, etc.):

```
$ jaguar new <project-name>
```

Non-interactive mode with flags:

```
$ jaguar new <project-name> -t go-api -m github.com/yourorg/yourproject -o ./output
```

**Flags:**

| Flag | Description | Default |
|---|---|---|
| `-t, --type` | Project type (`go-api`, `go-cli`, `go-grpc`, `go-embed`, `nodejs`, `python`) | interactive |
| `-m, --module` | Go module name | interactive |
| `-o, --output` | Output directory | current directory |
| `--frontend-framework` | Frontend framework for `go-embed` (`react`, `vue`, `angular`) | interactive |

### Add license headers

Add copyright license headers to source code files:

```
$ jaguar tool license add <project>
```

Check which files are missing license headers:

```
$ jaguar tool license check <project>
```

### Generate Error Codes

Automatically generate error code constants from iota-typed enums:

```
$ jaguar tool codegen --types int ./<directory>
```


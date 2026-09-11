# Jaguar

A unified AI-assisted scaffold for generating production-ready project skeletons across multiple languages and frameworks.

[![e2e test](https://github.com/shipengqi/jaguar/actions/workflows/e2e.yaml/badge.svg)](https://github.com/shipengqi/jaguar/actions/workflows/e2e.yaml)
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
| `go-embed` | Go REST API server with embedded React/Vue/Angular frontend (via `go:embed`) |
| `frontend-react` | Standalone React SPA (Vite + React Router + TanStack Query + Zustand) |
| `frontend-vue` | Standalone Vue 3 SPA (Vite + Vue Router + Pinia + TanStack Query) |
| `frontend-angular` | Standalone Angular SPA (Vite + Angular Router + NgRx Signals) |
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
| `-t, --type` | Project type (`go-api`, `go-cli`, `go-grpc`, `go-embed`, `frontend-react`, `frontend-vue`, `frontend-angular`, `nodejs`, `python`) | interactive |
| `-m, --module` | Go module name (Go projects only) | interactive |
| `-o, --output` | Output directory | current directory |
| `--frontend-framework` | Frontend framework for `go-embed` (`react`, `vue`, `angular`) | interactive |
| `--use-github-actions` | Generate GitHub Actions workflows | interactive |

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

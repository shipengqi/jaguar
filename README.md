# Jaguar

A scaffold that makes it easy to create amazing Go applications.

[![e2e test](https://github.com/shipengqi/jaguar/actions/workflows/e2e.yaml/badge.svg)](https://github.com/shipengqi/jaguar/actions/workflows/e2e.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/shipengqi/jaguar)](https://goreportcard.com/report/github.com/shipengqi/jaguar)
[![release](https://img.shields.io/github/release/shipengqi/jaguar.svg)](https://github.com/shipengqi/jaguar/releases)
[![license](https://img.shields.io/github/license/shipengqi/jaguar)](https://github.com/shipengqi/jaguar/blob/main/LICENSE)

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

```
$ jaguar new
```

### Add license

Add the copyright license headers for source code files:
```
$ jaguar tool license add <project>
```

### Generate Error Codes

Automatically generate error codes for API skeleton:
```
$ jaguar tool codegen --types int ./<API code directory>
```
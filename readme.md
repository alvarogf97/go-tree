[![Build Status](https://app.travis-ci.com/alvarogf97/go-tree.svg?branch=master)](https://app.travis-ci.com/alvarogf97/go-tree)
[![Coverage Status](https://coveralls.io/repos/github/alvarogf97/go-tree/badge.svg)](https://coveralls.io/github/alvarogf97/go-tree)

<p align="center">
  <img width="270" height="250" src="https://littleml.files.wordpress.com/2013/04/traditional.png">
</p>

# go-tree in a nutshell

go-tree is an implementation of the unix `tree` comand in go. It also exposes
an interface for those ones who wants to use the folder and files scanner in
their own programs.

# Installation

It's posible to install it by using go:

```go
go install github.com/alvarogf97/go-tree@latest
```
# Usage

please take a look of the --help flag for the command by executing

```cmd
>>> go-tree --help
... Usage:
...   go-tree [flags]
... 
... Flags:
...   -d, --directory string   directory (required) (default "./")
...   -h, --help               help for go-tree
...   -s, --show hidden        shows hidden files & folders
... 
```

example:

```cmd
>>> go-tree
... ├── cmd
... │   ├── root.go
... │   └── tree.go
... ├── pkg
... │   ├── fs
... │   │   ├── file.go
... │   │   ├── file_test.go
... │   │   ├── folder.go
... │   │   └── folder_test.go
... │   ├── scan
... │   │   ├── options.go
... │   │   ├── options_test.go
... │   │   ├── path.go
... │   │   ├── path_test.go
... │   │   ├── scanner.go
... │   │   ├── scanner_test.go
... │   │   ├── walker.go
... │   │   └── walker_test.go
... │   └── shortcuts
... │   │   └── tree.go
... ├── go.mod
... ├── go.sum
... ├── license.md
... ├── main.go
... └── readme.md
```

# Comply

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/leaanthony/comply/blob/master/LICENSE)
[![GitHub version](https://badge.fury.io/gh/leaanthony%2Fcomply.svg)](https://github.com/leaanthony/comply)
[![Go Report Card](https://goreportcard.com/badge/github.com/leaanthony/comply)](https://goreportcard.com/report/github.com/leaanthony/comply)
[![Godoc Reference](https://godoc.org/github.com/leaanthony/comply?status.svg)](http://godoc.org/github.com/leaanthony/comply)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/leaanthony/comply/issues)


Pulls in License files from the vendor directory and places them in ./licenses.
It recreates the package structure to ensure correctness.
It looks for files matching "LICENCE*" (Case Insensitive).
**It will wipe your licenses directory if it exists.** The idea is to keep it in sync.

## Install

`go get github.com/leaanthony/comply`

## Usage

In your project directory, run `comply`. It will scan the vendor directory for license files and copy them to a freshly created `licenses` directory in the current directory. 

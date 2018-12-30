# Comply

Pulls in License files from the vendor directory and places them in ./licenses.
It recreates the package structure to ensure correctness.
It looks for files matching "LICENCE*" (Case Insensitive).
**It will wipe your licenses directory if it exists.** The idea is to keep it in sync.

## Install

`go get github.com/leaanthony/comply`

## Usage

In your project directory, run `comply`. It will scan the vendor directory for license files and copy them to a freshly created `licenses` directory in the current directory. 

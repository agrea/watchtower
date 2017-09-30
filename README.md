# Watchtower

## Development

## Requirements

*Required*
- [Go 1.9](https://golang.org)
- [dep](https://github.com/golang/dep)

*Optional*
- [gin](https://github.com/codegangsta/gin) for auto reloading of the application
- [Bower](https://bower.io) if you need to update JS / CSS libraries
- [EditorConfig](http://editorconfig.org/) to automatically align to the project code standard

### Fetching Go dependencies

This project uses `dep` to manage dependencies. Run `dep ensure` to get all
dependencies for the project.

### Building and running

#### Building

Watchtower can be running `make`. This will generate a `watchtower` binary.

#### Running in watch mode

You can run Watchtower locally in watch mode with:

    make watch

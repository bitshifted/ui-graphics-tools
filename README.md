
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=bitshifted_ui-graphics-tools&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=bitshifted_ui-graphics-tools) 
[![Publish Docker Image](https://github.com/bitshifted/ui-graphics-tools/actions/workflows/publish-docker-image.yml/badge.svg)](https://github.com/bitshifted/ui-graphics-tools/actions/workflows/publish-docker-image.yml)

# UI Graphics Tools

Tools for generating graphics for user interfaces from SVG files. Currently, the following is supported:

* generate icons for Windows, Mac OS and Linux from SVG files
* generate splash screen for application from predefined template

## Usage

Tools are packages as Docker image, so they can be run on any platform that supports docker. Run tools with the following command:

```
docker run  -v ${PWD}:/workspace  ghcr.io/bitshifted/ui-graphics-tools:<version>  <arguments>
```

where `<version>` is the version of image you want to run, and `<arguments>` program arguments.

Mounting local working directory with `-v ${PWD}:/workspace` will enable running container as if it is run in current directory.

For more more detailed usage instructions, check out the following pages:

* [Generate icons from SVG file](./docs/generate-icons.md)
* [Generate splash screen from predefined template](./docs/generate-splash-screen.md)

# License

This project is published under Mozilla Public License 2.0 (MPL-2.0)

# Bug reports and feature requests

Please submit and issue in this repository with detailed description of the bug or feature.

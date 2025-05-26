# Development Container for OpenDocDB CTS

This directory contains configuration for a development container that can be used with Visual Studio Code's [Remote - Containers](https://code.visualstudio.com/docs/remote/containers) extension.

## Features

- Uses the existing Docker Compose setup from the repository
- Includes necessary tools for Go development
- Pre-configured with recommended VS Code extensions for Go development
- Correctly mounts source code and cache directories

## Usage

1. Install the [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extension in VS Code
2. Clone this repository
3. Open the repository in VS Code
4. When prompted, click "Reopen in Container" or use the command palette to run "Remote-Containers: Reopen in Container"
5. VS Code will build the container and open the workspace inside it
6. Development tools and extensions will be pre-configured for Go development

## Container Structure

- The main development container is based on `golang:1.24.3`
- MongoDB and FerretDB services are available via the existing Docker Compose setup
- Source code is mounted at `/src` in the container
- Go cache directories are mounted at `/cache/gopath`, `/cache/gocache`, and `/cache/gomodcache`
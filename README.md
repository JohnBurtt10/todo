# TODO 

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

I created this TODO web app as part of my onboarding to the RCAF Flight Deck, where I worked for my second co-op placement during the Fall 2022 term. The purpose of creating this app was to familarize myself with the technologies, used and the repository of the project I would be eventually working on for the RCAF Flight Deck in a way that cannot be achieved by simply reading through the files in the project. 

### Repository Structure

This repository contains Maintainer and its supporting packages and files.

* Main - Loads Config, Start Service Connections, Setup and starts Server
  * Location: `./main.go`
* Configuration
  * Location: `./config/*`
* Database Connection and Migrations
  * Location: `./database/*`
* View Routes
  * Location: `./routes/*`
* API Routes
  * Location: `./routes/api/*`
* Models - Custom Structs, Response Objects
  * Location: `./app/models/*`
* View Controllers - Server Side Rendering
  * Location: `./app/controllers/*`
* API Controllers - REST API, JSON
  * Location: `./app/controllers/api/*`
* Views - Server Side Pages, GO HTML
  * Location: `./resources/views/*`
* Public Static Files - Transpiled JS, Vendor JS, image files, favicon, site.webmanifest, service worker etc
  * Location: `./static/public/*`
* Services - Auth, Critical Chain, File Upload etc
  * Location: `./app/services/*`
* Repository - Persistent storage CRUD
  * Location: `./app/repos/*`
* Utility Functions - Functions that are re-used throughout the app
  * Location: `./app/utils/*`
* NPM Dev Resources - babel and eslint, package.json, package-lock.json
  * Location: `./`
* GitHub Config - Issues Templates, CI/CD actions, PR Template
  * Location: `./.github/*`
* Air Config - Hot Reload
  * Location: `./.air.toml`

 ## Configuration
 Configurations are in a single file called `.env`. You can copy the `.env.example` and change it to your needs.

 ```bash
 cp .env.example .env
 ```

 This `.env` file represents system environment variables on your machine.

 Keep in mind if configurations are not set, they default to Fiber's default settings which can be found [here](https://docs.gofiber.io/).

### Install golangci-lint:

#### macOS:
```bash
brew install golangci-lint
brew upgrade golangci-lint
```

#### windows/linux
```bash
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.1

golangci-lint --version
```

### Install npm

Follow instructions [here](https://nodejs.org/en/download/package-manager/).

#### install node_modules
```bash
npm install
```

#### dev

Build and run with Air (live reload)

```bash
air
```

In Visual Studio Code use Run and Debug launch command 'Delve into Local' to attach the debugger (you probably have to re-connect after each live reload). This is much slower than 'Launch Package.' The live reload probably isn't worth the effort.

##### test

Go test the repo.

```bash
go test ./..
```

### Recomended Service Containers for Dev ###

It is highly recommended that you develop using Postgre as your target database and redis as your cache and message queue.

```bash
docker run -it --name todo-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=todo -p 5432:5432 -d postgres
```

```bash
docker run -it --name todo-redis -p 6379:6379 -d redis
```

You will also need to make the appropriate changes in your .env file

### Start the application 


```bash
go run main.go
```
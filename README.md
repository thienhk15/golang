# Golang

[Golang](https://go.dev/)

<a href="https://github.com/gin-gonic/gin"><img src="https://img.shields.io/badge/gin-gonic-yellow" height="20"></a>
<a href="https://jmoiron.github.io/sqlx/"><img src="https://img.shields.io/badge/sqlx-v1.35-green" height="20"></a>
<a href="https://github.com/swaggo/swag"><img src="https://img.shields.io/badge/swagger-go-brightgreen" height="20"></a>

## Features

- Organising Golang project structure by kind [reference](https://developer20.com/how-to-structure-go-code/)
- API framework using [Gin Gonic](https://github.com/gin-gonic/gin)
- Connect with Postgres database using [sqlx](https://jmoiron.github.io/sqlx/)
- API document with Swagger
- Sample `Dockerfile` to package application as Docker image.

## Directory Structure

```bash
your-app/                        # Your app root
├─ .gitignore
├─ Dockerfile
├─ go.mod
├─ main.go                       # Main run file
├─ .vscode/                      # Launch / debug project when use VS Code
│  └─ launch.json
├─ component/                    # Main project source code
│  └─ api                        # Common functions for RESTful API
│  │  └─ response.go
│  └─ handlers                   # Contain main handler functions of RESTful API
│  │  └─ <name>_handler.go
│  └─ services                   # Contain business logic functions
│  │  └─ <name>_service.go
│  └─ models                     # Define structs
│  │  └─ <name>_model.go
│  └─ repositories               # Contain functions integrate with database
│  │  └─ <name>_repo.go
├─ config/                       # Config folder
│  └─ application.yml            # Main config file
├─ docs/                         # Swagger config for API document
├─ internal/                     # Contain functions connect with internal system
│  └─ postgres.go
│  └─ kafka.go
│  └─ redis.go
│  └─ ...
├─ external/                     # Contain functions connect with third party
│  └─ ...
├─ middleware/                   # Contain middleware functions
│  └─ base.go
│  └─ <name>_middleware.go
│  └─ ...
├─ server/                       # Init API server
│  └─ base.go
│  └─ command.go
│  └─ route.go
├─ utils/                        # Common files / functions for all project
│  └─ config.go
│  └─ logs.go
│  └─ ...
```

## Requirements

- [Golang](https://go.dev/) version from v1.18
- Swag CLI for generating API document. [Install here](https://github.com/swaggo/swag#getting-started)
- (Optional) Install Docker on local for running app on Docker. [Install here](https://docs.docker.com/engine/install/)

## Reference

- https://github.com/gin-gonic/gin
- https://developer20.com/how-to-structure-go-code/
- https://github.com/jmoiron/sqlx
- https://dskrzypiec.dev/gosqlstdsqlx/
- https://medium.com/@rocketlaunchr.cloud/how-to-benchmark-dbq-vs-sqlx-vs-gorm-e814caacecb5

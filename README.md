# Motivation
Write restful API server for react-admin demo app with the simple sqlite backend

## Go Boilerplate
Based on API boilerplate https://github.com/akmamun/gin-boilerplate-examples, but changed to use sqlite driver to simplify the demo server 

### Configuration Manage
#### ENV Manage

- Default ENV Configuration Manage from `.env`. sample file `.env.example`
```text
# Server Configuration
SECRET=h9wt*pasj6796j##w(w8=xaje8tpi6h*r&hzgrz065u&ed+k2)
DEBUG=True # `False` in Production
ALLOWED_HOSTS=0.0.0.0
SERVER_HOST=0.0.0.0
SERVER_PORT=8000

# Database Configuration
DB_NAME=react-admin.sqlite
```
- Server `DEBUG` set `False` in Production
- Database Logger `MASTER_DB_LOG_MODE` and `REPLICA_DB_LOG_MODE`  set `False` in production
- If ENV Manage from YAML file add a config.yml file and configuration [db.go](config/db.go) and [server.go](config/server.go). See More [ENV YAML Configure](#env-yaml-configure)

#### Server Configuration
- Use [Gin](https://github.com/gin-gonic/gin) Web Framework

#### Database Configuration
- Use [GORM](https://github.com/go-gorm/gorm) as an ORM

### Installation
#### Local Setup Instruction
Follow these steps:
- Copy [.env.example](.env.example) as `.env` and configure necessary values
- To add all dependencies for a package in your module `go get .` in the current directory
- Locally run `go run main.go` or `go build main.go` and run `./main`
- Check Application health available on [0.0.0.0:8000/health](http://0.0.0.0:8000/health)

### Middlewares
- Use Gin CORSMiddleware
```go
router := gin.New()
router.Use(gin.Logger())
router.Use(gin.Recovery())
router.Use(middleware.CORSMiddleware())
```

### Use Packages
- [Viper](https://github.com/spf13/viper) - Go configuration with fangs.
- [Gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
- [Logger](https://github.com/sirupsen/logrus) - Structured, pluggable logging for Go.
- [Air](https://github.com/cosmtrek/air) - Live reload for Go apps (Docker Development)


## Installation
```
$ go mod download
```

## How to run

### Required

- Mysql


### Ready

Create a **starling database** and import sql (path:docs/sql/starling.sql)

### Conf

You should modify `conf/app.ini`

```
[database]
Type = mysql
User = root
Password =
Host = 127.0.0.1:3306
Name = starling
TablePrefix = starling_
...
```

### Run
```
$ cd $GOPATH/src/starling

$ go run main.go 
```

Project information and existing API

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /auth                     --> github.com/EDDYCJY/go-gin-example/routers/api.GetAuth (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
[GIN-debug] GET    /api/starling             --> github.com/EDDYCJY/go-gin-example/routers/api.GetStarlingListByCondition (3 handlers)
[GIN-debug] POST   /api/starling             --> github.com/EDDYCJY/go-gin-example/routers/api.CreateStarling (3 handlers)
[GIN-debug] POST   /api/starling/batch       --> github.com/EDDYCJY/go-gin-example/routers/api.BatchCreateStarling (3 handlers)
[GIN-debug] PUT    /api/starling             --> github.com/EDDYCJY/go-gin-example/routers/api.UpdateStarling (3 handlers)
[GIN-debug] DELETE /api/starling/:id         --> github.com/EDDYCJY/go-gin-example/routers/api.DeleteStarling (3 handlers)
[GIN-debug] DELETE /api/starling             --> github.com/EDDYCJY/go-gin-example/routers/api.DeleteBatchStarling (3 handlers)
[GIN-debug] POST   /api/starling/key_lang    --> github.com/EDDYCJY/go-gin-example/routers/api.GetTextByKeyAndLang (3 handlers)
           --> github.com/EDDYCJY/go-gin-example/view.DeleteBatchStarling (3 handlers)

Listening port is 8000
Actual pid is 4393
```
Swagger doc


## Features

- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin

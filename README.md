# About This Template

- [x] vue3+ts, [https://v3.vuejs.org/](https://v3.vuejs.org/)
- [x] gin, [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- [x] systray, [https://github.com/getlantern/systray](https://github.com/getlantern/systray)
- [x] air, [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air)
- [x] gorm, [https://gorm.io/](https://gorm.io/)
- [x] swagger, [https://github.com/swaggo/gin-swagger](https://github.com/swaggo/gin-swagger)
- [x] jwt, [https://github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt)
- [x] ini, [https://github.com/go-ini/ini](https://github.com/go-ini/ini)

## Environment

- Windows 10 x64
- Go 1.17+
- Node.js v14
- NPM v8
- Git Bash

## Setup

``` bash
# install air cli
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
# install swag cli
go install github.com/swaggo/swag/cmd/swag@latest
# install go dependencies
go mod tidy
```

## Scripts

- `front:serve`: run server for frontend development
- `front:build`: build frontend into **folder `/build/www`** for production
- `front:test:unit`: run frontend unit tests
- `front:lint`: run frontend eslint
- `ready:certs`: generate certificate for `localhost` to **folder `/build`**
- `ready:swag`: generate swagger files to **folder `/swagger`**
- `back:air`: run air for backend development
- `back:run`: build backend to **folder `/build`** with **filename `app.exe`**, then run it
- `back:build`: build backend to **folder `/build`** with **filename `app.exe`** and **ldflags `-H=windowsgui`**

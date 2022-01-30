# About This Template

- [x] vue3+ts, [https://v3.vuejs.org/](https://v3.vuejs.org/)
- [ ] gin, [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- [ ] systray, [https://github.com/getlantern/systray](https://github.com/getlantern/systray)
- [ ] air, [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air)
- [ ] gorm, [https://gorm.io/](https://gorm.io/)
- [ ] swagger, [https://github.com/swaggo/gin-swagger](https://github.com/swaggo/gin-swagger)
- [ ] jwt, [https://github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt)
- [ ] ini, [https://github.com/go-ini/ini](https://github.com/go-ini/ini)

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

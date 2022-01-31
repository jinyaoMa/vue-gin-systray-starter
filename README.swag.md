# Swagger Template

## Mime Types for `@Accept` and `@Produce`

- json
- xml
- plain
- html
- mpfd (multipart/form-data)
- x-www-form-urlencoded
- json-api
- json-stream
- octet-stream
- png
- jpeg
- gif

## API Operation

``` go
// @Summary      Summary
// @Description  Description
// @Tags         tag1,tag2
// @Accept       json
// @Produce      json
// @Security     BearerToken
// @param        Authorization header string false "Authorization"
// @Param        qParam query int false "Query Param"
// @Param        pParam path uint true "Path Param"
// @Param        bParam body float32 false "Body Param (allow json)"
// @Param        fParam formData bool true "Form Data Param"
// @Success      200 {object} formats.JSONResult{data=string} "OK"
// @Failure      400 {object} formats.JSONError "Bad Request"
// @Failure      401 {object} formats.JSONError "Unauthorized"
// @Failure      403 {object} formats.JSONError "Forbidden"
// @Failure      404 {object} formats.JSONError "Not Found"
// @Failure      500 {object} formats.JSONError "Internal Server Error"
// @Router       /path/get/{pParam} [get]
// @Router       /path/post [post]
// @Router       /path/put [put]
// @Router       /path/delete [delete]
```

## General API Info

``` go
// @title app
// @version 0.0.0
// @description "app"

// @contact.name Github Issues
// @contact.url https://github.com/jinyaoMa/vue-gin-systray-starter/issues

// @license.name MIT
// @license.url https://github.com/jinyaoMa/vue-gin-systray-starter/blob/main/LICENSE

// @BasePath /api

// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
```

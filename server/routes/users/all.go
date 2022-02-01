package users

import (
	"app/database"
	"app/server/formats"
	"app/server/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      All
// @Description  Get all users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerToken
// @param        Authorization header string false "Authorization"
// @Success      200 {object} formats.JSONResult{data=[]formats.UserDataResult} "OK"
// @Failure      400 {object} formats.JSONError "Bad Request"
// @Failure      401 {object} formats.JSONError "Unauthorized"
// @Failure      500 {object} formats.JSONError "Internal Server Error"
// @Router       /users/all [get]
func All() gin.HandlerFunc {
	const (
		ErrCode int = -10000 - iota
		ErrCodeFindUsers
		// append here, do not insert
	)
	var (
		ErrFindUsers = errors.New("find users error")
	)

	DB := database.DB

	return func(c *gin.Context) {
		var user *models.User
		users, err := user.Find(DB)
		if err != nil {
			if gin.Mode() == gin.DebugMode {
				c.AbortWithStatusJSON(http.StatusInternalServerError, formats.JSONError{
					Code:  ErrCodeFindUsers,
					Error: err.Error(),
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, formats.JSONError{
				Code:  ErrCodeFindUsers,
				Error: ErrFindUsers.Error(),
			})
			return
		}

		usersData := []formats.UserDataResult{}
		for _, u := range users {
			usersData = append(usersData, formats.UserDataResult{
				ID:         u.ID,
				Account:    u.Account,
				Password:   u.Password,
				Name:       u.Name,
				Permission: u.Permission,
				Active:     u.Active,
			})
		}

		c.JSON(http.StatusOK, formats.JSONResult{
			Success: true,
			Data:    usersData,
		})
	}
}

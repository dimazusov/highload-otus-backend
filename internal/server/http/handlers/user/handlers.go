package user

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/minipkg/selection_condition"

	"social/internal/app"
	"social/internal/domain/user"
	"social/internal/pkg/apperror"
	"social/internal/server/http/api_error"
)

// @Summary auth user
// @Description auth user
// @ID auth-user
// @Accept json
// @Produce json
// @Success 200 {object} map[string]
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /auth [get]
func AuthHandler(c *gin.Context, app *app.App) {
	credentials := user.Credentials{}
	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := app.Domain.User.Service.Auth(context.Background(), credentials)
	if err != nil {
		if err == apperror.ErrUserNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// @Summary get users
// @Description get usersby params
// @ID get-users
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /registration [get]
func RegisterHandler(c *gin.Context, app *app.App) {
	regUser := user.User{}
	err := c.ShouldBindJSON(&regUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := app.Domain.User.Service.Create(context.Background(), &regUser)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	credentials := user.Credentials{
		Email: regUser.Email,
		Password: regUser.Password,
	}
	authToken, err := app.Domain.User.Service.Auth(context.Background(), credentials)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": authToken,
		"userId": userID,
	})
}

// @Summary get users
// @Description get usersby params
// @ID get-users
// @Accept json
// @Produce json
// @Param with_organization query boolean false "with organization"
// @Param per_page query int false "per page"
// @Param page query int false "page"
// @Success 200 {object} UsersList
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /users [get]
func GetUsersHandler(c *gin.Context, app *app.App) {
	cond := selection_condition.SelectionCondition{}
	err := c.ShouldBindQuery(&cond)
	if err := c.ShouldBindQuery(&cond); err != nil {
		c.JSON(http.StatusBadRequest, api_error.New(err))
		return
	}

	users, err := app.Domain.User.Service.Query(context.Background(), &cond)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, api_error.New(apperror.ErrInternal))
		return
	}

	count, err := app.Domain.User.Service.Count(context.Background(), &cond)
	if err != nil {
		log.Println(err)

		c.JSON(http.StatusInternalServerError, api_error.New(apperror.ErrInternal))
		return
	}

	c.JSON(http.StatusOK, UsersList{
		Items: users,
		Count: count,
	})
}

// @Summary get user by id
// @Description get user by id
// @ID get-user-by-id
// @Accept json
// @Produce json
// @Param id path boolean true "id"
// @Success 200 {object} User
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /user/{id} [get]
func GetUserHandler(c *gin.Context, app *app.App) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, api_error.New(err))
		return
	}

	bld, err := app.Domain.User.Service.Get(context.Background(), uint(userID))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, api_error.New(apperror.ErrInternal))
		return
	}

	c.JSON(http.StatusOK, bld)
}

// @Summary update user
// @Description update user
// @ID update-user
// @Accept json
// @Produce json
// @Param user body user.User true "updatable user"
// @Success 200 {string} success
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /user [put]
func UpdateUserHandler(c *gin.Context, app *app.App) {
	bdg := user.User{}
	if err := c.ShouldBindJSON(&bdg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.Domain.User.Service.Update(context.Background(), &bdg)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// @Summary create user
// @Description create user
// @ID create-user
// @Accept json
// @Produce json
// @Param user body user.User true "creatable user"
// @Success 200 {object} User
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /user [post]
func CreateUserHandler(c *gin.Context, app *app.App) {
	bdg := user.User{}
	if err := c.ShouldBindJSON(&bdg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := app.Domain.User.Service.Create(context.Background(), &bdg)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}
	bdg.ID = userID

	c.JSON(http.StatusOK, bdg)
}

// @Summary delete user by id
// @Description delete user by id
// @ID delete-user-by-id
// @Accept json
// @Produce json
// @Param id path boolean true "id"
// @Success 200 {string} success
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /user/{id} [delete]
func DeleteUserHandler(c *gin.Context, app *app.App) {
	userId, err := selection_condition.ParseUintParam(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = app.Domain.User.Service.Delete(context.Background(), userId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

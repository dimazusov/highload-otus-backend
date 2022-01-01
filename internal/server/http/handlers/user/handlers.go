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

// @Summary get buildings
// @Description get buildings by params
// @ID get-buildings
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

// @Summary get building by id
// @Description get building by id
// @ID get-building-by-id
// @Accept json
// @Produce json
// @Param id path boolean true "id"
// @Success 200 {object} User
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /user/{id} [get]
func GetUserHandler(c *gin.Context, app *app.App) {
	buildingID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, api_error.New(err))
		return
	}

	bld, err := app.Domain.User.Service.Get(context.Background(), uint(buildingID))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, api_error.New(apperror.ErrInternal))
		return
	}

	c.JSON(http.StatusOK, bld)
}

// @Summary update building
// @Description update building
// @ID update-building
// @Accept json
// @Produce json
// @Param building body user.User true "updatable building"
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

// @Summary create building
// @Description create building
// @ID create-building
// @Accept json
// @Produce json
// @Param building body user.User true "creatable building"
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

	buildingID, err := app.Domain.User.Service.Create(context.Background(), &bdg)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}
	bdg.ID = buildingID

	c.JSON(http.StatusOK, bdg)
}

// @Summary delete building by id
// @Description delete building by id
// @ID delete-building-by-id
// @Accept json
// @Produce json
// @Param id path boolean true "id"
// @Success 200 {string} success
// @Failure 400 {object} api_error.Error
// @Failure 500 {object} api_error.Error
// @Router /user/{id} [delete]
func DeleteUserHandler(c *gin.Context, app *app.App) {
	buildingId, err := selection_condition.ParseUintParam(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = app.Domain.User.Service.Delete(context.Background(), buildingId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

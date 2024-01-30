package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo IUserRepository
}

func NewUserHandler(store IUserRepository) *UserHandler {
	return &UserHandler{repo: store}
}

func (uh *UserHandler) NewUser(c *gin.Context) {
	var u User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = uh.repo.NewUser(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + u.Name,
		"id":      u.Model.ID,
	})
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	ctxName := c.GetString("Name")
	fmt.Println(ctxName)

	users, err := uh.repo.GetUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "can not parse id",
		})
		return
	}

	err = uh.repo.DeleteUser(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted!",
	})
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "can not parse id",
		})
		return
	}

	var u User
	err = c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = uh.repo.UpdateUser(idInt, &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated!",
	})
}

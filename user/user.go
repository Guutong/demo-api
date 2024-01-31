package user

import (
	"fmt"
	"net/http"
)

type UserHandler struct {
	repo IUserRepository
}

func NewUserHandler(store IUserRepository) *UserHandler {
	return &UserHandler{repo: store}
}

func (uh *UserHandler) NewUser(c IContext) {
	var u User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	err = uh.repo.NewUser(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hello " + u.Name,
		"id":      u.ID,
	})
}

func (uh *UserHandler) GetUser(c IContext) {
	ctxName := c.Name()
	fmt.Println(ctxName)

	users, err := uh.repo.GetUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})
}

func (uh *UserHandler) DeleteUser(c IContext) {
	id, err := c.ParamInt("id")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "can not parse id",
		})
		return
	}

	err = uh.repo.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Deleted!",
	})
}

func (uh *UserHandler) UpdateUser(c IContext) {
	id, err := c.ParamInt("id")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "can not parse id",
		})
		return
	}

	var u User
	err = c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	err = uh.repo.UpdateUser(id, &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Updated!",
	})
}

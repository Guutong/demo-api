package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type IContext interface {
	BindJSON(interface{}) error
	JSON(int, interface{})
	Param(string) string
	ParamInt(string) (int, error)

	// from middleware
	Name() string
}

type AppContext struct {
	*gin.Context
}

func NewAppContext(c *gin.Context) *AppContext {
	return &AppContext{Context: c}
}

func (c *AppContext) BindJSON(u interface{}) error {
	return c.Context.ShouldBindJSON(u)
}

func (c *AppContext) JSON(status int, v interface{}) {
	c.Context.JSON(status, v)
}

func (c *AppContext) Param(k string) string {
	return c.Context.Param(k)
}

func (c *AppContext) ParamInt(string) (int, error) {
	id := c.Context.Param("id")
	return strconv.Atoi(id)
}

func (c *AppContext) Name() string {
	return c.Context.GetString("name")
}

func NewGinHandler(handler func(IContext)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewAppContext(c))
	}
}

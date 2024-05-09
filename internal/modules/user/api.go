package user

import (
	"be-assignment/domain"
	"be-assignment/dto"

	"github.com/gin-gonic/gin"
)

type route struct {
	userService domain.UserService
}

func NewRoute(app *gin.Engine, userService domain.UserService) {
	route := route{
		userService,
	}

	v1Public := app.Group("/api/v1")
	{
		v1Public.GET("/users", route.getAllUsers)
		v1Public.POST("/register", route.Register)
	}
}

func (r *route) Register(c *gin.Context) {
	var userReq dto.RegisterRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(400, &dto.Response{
			Message: err.Error(),
			Data:    []string{},
			Status:  400,
		})
		return
	}

	if err := r.userService.Register(userReq); err != nil {
		if err == domain.ErrUserAlreadyExists {
			c.JSON(400, &dto.Response{
				Message: err.Error(),
				Data:    []string{},
				Status:  400,
			})
			return
		}

		c.JSON(500, &dto.Response{
			Message: err.Error(),
			Data:    []string{},
			Status:  500,
		})
		return
	}

	c.JSON(200, &dto.Response{
		Message: "Success",
		Data:    []string{},
		Status:  200,
	})
}

func (r *route) getAllUsers(c *gin.Context) {
	users, err := r.userService.GetAllUsers()
	if err != nil {
		c.JSON(500, &dto.Response{
			Message: err.Error(),
			Data:    nil,
			Status:  500,
		})
		return
	}

	c.JSON(200, &dto.Response{
		Message: "Success",
		Data:    users,
		Status:  200,
	})
}

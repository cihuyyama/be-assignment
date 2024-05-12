package accountmanager

import (
	"be-assignment/domain"
	"be-assignment/dto"
	"be-assignment/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type route struct {
	accountManagerService domain.AccountManagerService
}

func NewRoute(app *gin.Engine, accountManagerService domain.AccountManagerService) {
	route := route{
		accountManagerService,
	}

	v1Public := app.Group("/api/v1")
	{
		v1Public.POST("/register", route.Register)
		v1Public.POST("/login", route.Login)
	}

	v1Private := app.Group("/api/v1")
	v1Private.Use(middleware.Authenticate())
	{
		v1Private.GET("/users", route.GetUser)
		v1Private.GET("/accounts", route.GetAllAccount)
		v1Private.POST("/accounts", route.CreateAccount)
	}
}

// @Summary Register a new user
// @Description Register a new user
// @Tags Account Manager
// @Accept json
// @Produce json
// @Param body body dto.RegisterRequest true "Register Request"
// @Success 200 {object} dto.Response
// @Router /register [post]
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

	if err := validator.New().Struct(userReq); err != nil {
		c.JSON(400, &dto.Response{
			Message: err.Error(),
			Data:    []string{},
			Status:  400,
		})
		return
	}

	if err := r.accountManagerService.Register(userReq); err != nil {
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

// @Summary Login a user
// @Description Login a user
// @Tags Account Manager
// @Accept json
// @Produce json
// @Param body body dto.LoginRequest true "Login Request"
// @Success 200 {object} dto.Response
// @Router /login [post]
func (r *route) Login(c *gin.Context) {
	var userReq dto.LoginRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(400, &dto.Response{
			Message: err.Error(),
			Data:    []string{},
			Status:  400,
		})
		return
	}

	if err := validator.New().Struct(userReq); err != nil {
		c.JSON(400, &dto.Response{
			Message: err.Error(),
			Data:    []string{},
			Status:  400,
		})
		return
	}

	loginResponse, err := r.accountManagerService.Login(userReq)
	if err != nil {
		if err == domain.ErrUserNotFound || err == domain.ErrInvalidPassword {
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
		Data:    loginResponse,
		Status:  200,
	})
}

// @Summary Get user data
// @Description Get user data using token from the authorization header
// @Tags Account Manager
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.Response
// @Router /users [get]
func (r *route) GetUser(c *gin.Context) {
	user, err := r.accountManagerService.GetUser(c)
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
		Data: &dto.UserDataResponse{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Status: 200,
	})
}

// @Summary Get all user's account and user's transactions per account
// @Description Get all user's account using token from the authorization header and transactions per account
// @Tags Account Manager
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.Response
// @Router /accounts [get]
func (r *route) GetAllAccount(c *gin.Context) {
	accounts, err := r.accountManagerService.GetAllAccount(c)
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
		Data:    accounts,
		Status:  200,
	})
}

// @Summary Create a new account
// @Description Create a new account
// @Tags Account Manager
// @Accept json
// @Produce json
// @Param body body dto.CreateAccountRequest true "Create Account Request"
// @Security BearerAuth
// @Success 200 {object} dto.Response
// @Router /accounts [post]
func (r *route) CreateAccount(c *gin.Context) {
	var accountReq dto.CreateAccountRequest
	if err := c.ShouldBindJSON(&accountReq); err != nil {
		c.JSON(400, &dto.Response{
			Message: err.Error(),
			Data:    []string{},
			Status:  400,
		})
		return
	}

	if err := validator.New().Struct(accountReq); err != nil {
		c.JSON(400, &dto.Response{
			Message: err.Error(),
			Data:    []string{},
			Status:  400,
		})
		return
	}

	if err := r.accountManagerService.CreateAccount(c, accountReq); err != nil {
		if err == domain.ErrAccountAlreadyExists {
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

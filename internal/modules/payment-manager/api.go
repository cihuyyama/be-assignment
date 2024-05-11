package paymentmanager

import (
	"be-assignment/domain"
	"be-assignment/dto"
	"be-assignment/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type route struct {
	service domain.PaymentManagerService
}

func NewPaymentHandler(app *gin.Engine, service domain.PaymentManagerService) {
	route := route{
		service,
	}

	v1Public := app.Group("/api/v1")
	{
		v1Public.GET("/transactions", route.GetAllTransaction)
	}

	v1Private := app.Group("/api/v1")
	v1Private.Use(middleware.Authenticate())
	{
		v1Private.POST("/send", route.Transfer)
	}
}

// @Summary Get all transactions
// @Description Get all transactions
// @Tags Payment Manager
// @Accept json
// @Produce json
// @Success 200 {object} dto.Response
// @Router /transactions [get]
func (r *route) GetAllTransaction(c *gin.Context) {
	transactions, err := r.service.GetAllTransaction()
	if err != nil {
		c.JSON(500, &dto.Response{
			Message: err.Error(),
			Data:    []string{},
			Status:  500,
		})
		return
	}

	c.JSON(200, &dto.Response{
		Message: "Success",
		Data:    transactions,
		Status:  200,
	})
}

// @Summary Transfer money
// @Description Transfer money
// @Tags Payment Manager
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.TransferRequest true "Transfer Request"
// @Success 200 {object} dto.Response
// @Router /send [post]
func (r *route) Transfer(c *gin.Context) {
	var tfRequest dto.TransferRequest
	if err := c.ShouldBindJSON(&tfRequest); err != nil {
		c.JSON(400, &dto.Response{
			Message: err.Error(),
			Data:    []string{},
			Status:  400,
		})
		return
	}

	if err := validator.New().Struct(tfRequest); err != nil {
		c.JSON(400, &dto.Response{
			Message: err.Error(),
			Data:    []string{},
			Status:  400,
		})
		return

	}

	if err := r.service.Transfer(tfRequest); err != nil {
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

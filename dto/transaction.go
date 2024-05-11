package dto

type TransferRequest struct {
	SofNumber string `json:"sof_number" validate:"required"` // Source of Fund Number
	DofNumber string `json:"dof_number" validate:"required"` // Destination of Fund Number
	Amount    int    `json:"amount" validate:"required"`
}

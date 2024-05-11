package dto

type TransferRequest struct {
	SofNumber string `json:"sof_number" validate:"required"` // Source of Fund Number
	DofNumber string `json:"dof_number" validate:"required"` // Destination of Fund Number
	Amount    uint   `json:"amount" validate:"required"`
}

type WithdrawRequest struct {
	SofNumber string `json:"sof_number" validate:"required"`
	Amount    uint   `json:"amount" validate:"required"`
}

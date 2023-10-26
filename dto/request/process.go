package request

type ProcessReq struct {
	PID int32 `form:"pid" json:"pid" validate:"required"`
}

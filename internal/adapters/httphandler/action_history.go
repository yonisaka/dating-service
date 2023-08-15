package httphandler

import (
	"net/http"

	"github.com/yonisaka/dating-service/internal/dto"
	"github.com/yonisaka/dating-service/internal/usecases"
	"github.com/yonisaka/dating-service/pkg/logger"
)

// actionHistoryHandler is a struct for login handler
type actionHistoryHandler struct {
	actionHistoryUsecase usecases.ActionHistoryUsecase
}

// NewActionHistoryHandler is a constructor function for login handler
func NewActionHistoryHandler(
	actionHistoryUsecase usecases.ActionHistoryUsecase,
) Handler {
	return &actionHistoryHandler{
		actionHistoryUsecase: actionHistoryUsecase,
	}
}

// Handle is a function to handle login check
func (h *actionHistoryHandler) Handle(req *http.Request) dto.HTTPResponse {
	var (
		ctx = req.Context()
		lf  = logger.NewFields(
			logger.EventName("handler.action.history"),
		)
	)

	result, err := h.actionHistoryUsecase.GetActionHistories(ctx)
	if err != nil {
		logger.ErrorWithContext(req.Context(), err.Error(), lf...)

		return *dto.NewResponse().
			WithCode(http.StatusInternalServerError).
			WithMessage(err.Error())
	}

	return *dto.NewResponse().WithCode(http.StatusOK).WithData(result).WithMessage("success get action history")
}

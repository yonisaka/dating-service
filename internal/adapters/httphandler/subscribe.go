package httphandler

import (
	"net/http"

	"github.com/yonisaka/dating-service/internal/dto"
	"github.com/yonisaka/dating-service/internal/usecases"
	"github.com/yonisaka/dating-service/pkg/logger"
)

// subscribeHandler is a struct for login handler
type subscribeHandler struct {
	subscribeUsecase usecases.SubscribeUsecase
}

// NewSubscribeHandler is a constructor function for login handler
func NewSubscribeHandler(
	subscribeUsecase usecases.SubscribeUsecase,
) Handler {
	return &subscribeHandler{
		subscribeUsecase: subscribeUsecase,
	}
}

// Handle is a function to handle login check
func (h *subscribeHandler) Handle(req *http.Request) dto.HTTPResponse {
	var (
		ctx = req.Context()
		lf  = logger.NewFields(
			logger.EventName("handler.profile"),
		)
	)

	result, err := h.subscribeUsecase.Subscribe(ctx)
	if err != nil {
		logger.ErrorWithContext(req.Context(), err.Error(), lf...)

		return *dto.NewResponse().
			WithCode(http.StatusInternalServerError).
			WithMessage(err.Error())
	}

	return *dto.NewResponse().WithCode(http.StatusOK).WithData(result).WithMessage("success subscribe")
}

package httphandler

import (
	"net/http"

	"github.com/yonisaka/dating-service/internal/dto"
	"github.com/yonisaka/dating-service/internal/usecases"
	"github.com/yonisaka/dating-service/pkg/logger"
)

// queryProfileHandler is a struct for login handler
type queryProfileHandler struct {
	queryProfileUsecase usecases.QueryProfileUsecase
}

// NewQueryProfileHandler is a constructor function for login handler
func NewQueryProfileHandler(
	queryProfileUsecase usecases.QueryProfileUsecase,
) Handler {
	return &queryProfileHandler{
		queryProfileUsecase: queryProfileUsecase,
	}
}

// Handle is a function to handle login check
func (h *queryProfileHandler) Handle(req *http.Request) dto.HTTPResponse {
	var (
		ctx = req.Context()
		lf  = logger.NewFields(
			logger.EventName("handler.query.profile"),
		)
	)

	result, err := h.queryProfileUsecase.GetQueryProfile(ctx)
	if err != nil {
		logger.ErrorWithContext(req.Context(), err.Error(), lf...)

		return *dto.NewResponse().
			WithCode(http.StatusInternalServerError).
			WithMessage(err.Error())
	}

	return *dto.NewResponse().WithCode(http.StatusOK).WithData(result).WithMessage("success get query profile")
}

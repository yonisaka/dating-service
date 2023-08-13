package httphandler

import (
	"net/http"

	"github.com/yonisaka/dating-service/internal/dto"
	"github.com/yonisaka/dating-service/internal/usecases"
	"github.com/yonisaka/dating-service/pkg/logger"
)

// profileHandler is a struct for login handler
type profileHandler struct {
	profileUsecase usecases.ProfileUsecase
}

// NewProfileHandler is a constructor function for login handler
func NewProfileHandler(
	profileUsecase usecases.ProfileUsecase,
) Handler {
	return &profileHandler{
		profileUsecase: profileUsecase,
	}
}

// Handle is a function to handle login check
func (h *profileHandler) Handle(req *http.Request) dto.HTTPResponse {
	var (
		ctx = req.Context()
		lf  = logger.NewFields(
			logger.EventName("handler.profile"),
		)
	)

	result, err := h.profileUsecase.GetProfile(ctx)
	if err != nil {
		logger.ErrorWithContext(req.Context(), err.Error(), lf...)

		return *dto.NewResponse().
			WithCode(http.StatusInternalServerError).
			WithMessage(err.Error())
	}

	return *dto.NewResponse().WithCode(http.StatusOK).WithData(result).WithMessage("success get profile")
}

package httphandler

import (
	"net/http"

	"github.com/yonisaka/dating-service/internal/consts"
	"github.com/yonisaka/dating-service/internal/dto"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/internal/usecases"
	"github.com/yonisaka/dating-service/pkg/logger"
)

// preferenceHandler is a struct for login handler
type preferenceHandler struct {
	profileUsecase usecases.ProfileUsecase
}

// NewPreferenceHandler is a constructor function for login handler
func NewPreferenceHandler(
	profileUsecase usecases.ProfileUsecase,
) Handler {
	return &preferenceHandler{
		profileUsecase: profileUsecase,
	}
}

// Handle is a function to handle login check
func (h *preferenceHandler) Handle(req *http.Request) dto.HTTPResponse {
	var (
		ctx = req.Context()
		lf  = logger.NewFields(
			logger.EventName("handler.preference"),
		)
		param presentations.UserPreferenceRequest
	)

	data := Data{req}
	err := data.Cast(&param)

	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)

		return *dto.NewResponse().
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrorInvalidRequest)
	}

	err = h.profileUsecase.SetPreference(ctx, param)
	if err != nil {
		logger.ErrorWithContext(req.Context(), err.Error(), lf...)

		return *dto.NewResponse().
			WithCode(http.StatusInternalServerError).
			WithMessage(err.Error())
	}

	return *dto.NewResponse().WithCode(http.StatusOK).WithMessage("success set preference")
}

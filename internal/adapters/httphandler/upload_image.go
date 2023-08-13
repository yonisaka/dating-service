package httphandler

import (
	"net/http"

	"github.com/yonisaka/dating-service/internal/consts"
	"github.com/yonisaka/dating-service/internal/dto"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/internal/usecases"
	"github.com/yonisaka/dating-service/pkg/logger"
)

// uploadImageHandler is a struct for login handler
type uploadImageHandler struct {
	profileUsecase usecases.ProfileUsecase
}

// NewUploadImageHandler is a constructor function for login handler
func NewUploadImageHandler(
	profileUsecase usecases.ProfileUsecase,
) Handler {
	return &uploadImageHandler{
		profileUsecase: profileUsecase,
	}
}

// Handle is a function to handle login check
func (h *uploadImageHandler) Handle(req *http.Request) dto.HTTPResponse {
	var (
		ctx = req.Context()
		lf  = logger.NewFields(
			logger.EventName("handler.profile"),
		)
		param presentations.UploadImageRequest
	)

	data := Data{req}
	err := data.Cast(&param)

	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)

		return *dto.NewResponse().
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrorInvalidRequest)
	}

	err = h.profileUsecase.UploadImages(ctx, param)
	if err != nil {
		logger.ErrorWithContext(req.Context(), err.Error(), lf...)

		return *dto.NewResponse().
			WithCode(http.StatusInternalServerError).
			WithMessage(err.Error())
	}

	return *dto.NewResponse().WithCode(http.StatusOK).WithMessage("success upload images")
}

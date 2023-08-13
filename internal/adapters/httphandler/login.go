package httphandler

import (
	"net/http"

	"github.com/thedevsaddam/govalidator"
	"github.com/yonisaka/dating-service/internal/consts"
	"github.com/yonisaka/dating-service/internal/dto"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/internal/usecases"
	"github.com/yonisaka/dating-service/pkg/logger"
)

// loginHandler is a struct for login handler
type loginHandler struct {
	authUsecase usecases.AuthUsecase
}

// NewLoginHandler is a constructor function for login handler
func NewLoginHandler(
	authUsecase usecases.AuthUsecase,
) Handler {
	return &loginHandler{
		authUsecase: authUsecase,
	}
}

// Handle is a function to handle login check
func (h *loginHandler) Handle(req *http.Request) dto.HTTPResponse {
	var (
		ctx = req.Context()
		lf  = logger.NewFields(
			logger.EventName("handler.login"),
		)
		param  presentations.LoginRequest
		errors []presentations.ErrorValidation
	)

	data := Data{req}
	err := data.Cast(&param)

	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)

		return *dto.NewResponse().
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrorInvalidRequest)
	}

	rules := govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"required", "min:8"},
	}

	opt := govalidator.New(
		govalidator.Options{
			Data:  &param,
			Rules: rules,
		},
	)

	errValidate := opt.ValidateStruct()
	if len(errValidate) != 0 {
		for i, e := range errValidate {
			logger.WarnWithContext(ctx, e, lf...)
			errors = append(errors, presentations.ErrorValidation{
				Field:    i,
				Messages: e,
			})
		}
	}

	if len(errors) > 0 {
		return *dto.NewResponse().
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrorValidation).
			WithError(errors)
	}

	result, err := h.authUsecase.Login(ctx, param)
	if err != nil {
		logger.ErrorWithContext(req.Context(), err.Error(), lf...)

		return *dto.NewResponse().
			WithCode(http.StatusInternalServerError).
			WithMessage(err.Error())
	}

	return *dto.NewResponse().WithCode(http.StatusOK).WithData(result).WithMessage("success login")
}

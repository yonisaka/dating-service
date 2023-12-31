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

// registerHandler is a struct for health handler
type registerHandler struct {
	authUsecase usecases.AuthUsecase
}

// NewRegisterHandler is a constructor function for health handler
func NewRegisterHandler(
	authUsecase usecases.AuthUsecase,
) Handler {
	return &registerHandler{
		authUsecase: authUsecase,
	}
}

// Handle is a function to handle health check
//
//nolint:funlen
func (h *registerHandler) Handle(req *http.Request) dto.HTTPResponse {
	var (
		ctx = req.Context()
		lf  = logger.NewFields(
			logger.EventName("handler.register"),
		)
		param  presentations.RegisterRequest
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
		"first_name": []string{"required", "max:50"},
		"last_name":  []string{"required", "max:50"},
		"email":      []string{"required", "email"},
		"phone":      []string{"required", "min:10", "max:15"},
		"password":   []string{"required", "min:8"},
		"dob":        []string{"required", "date"},
		"gender":     []string{"required", "in:man,woman"},
		"intend":     []string{"required", "in:long-term,long-term partner,short-term fun,short-term but long-term ok,new friends,still figuring it out"}, //nolint:lll
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

	result, err := h.authUsecase.Register(req.Context(), param)
	if err != nil {
		logger.ErrorWithContext(req.Context(), err.Error(), lf...)

		return *dto.NewResponse().
			WithCode(http.StatusInternalServerError).
			WithMessage(err.Error())
	}

	return *dto.NewResponse().WithCode(http.StatusOK).WithData(result).WithMessage("Success Register")
}

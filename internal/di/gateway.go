package di

import (
	"context"
	"net/http"

	"github.com/yonisaka/dating-service/internal/adapters/httphandler"
	"github.com/yonisaka/dating-service/internal/consts"
	"github.com/yonisaka/dating-service/internal/dto"
	"github.com/yonisaka/dating-service/pkg/msg"
)

func httpGateway(request *http.Request, handler httphandler.Handler) dto.HTTPResponse {
	cfg := GetConfig()
	if !msg.GetAvailableLang(200, request.Header.Get(consts.HeaderLanguageKey)) {
		request.Header.Set(consts.HeaderLanguageKey, cfg.App.DefaultLang)
	}

	ctx := context.WithValue(request.Context(), consts.CtxLang, request.Header.Get(consts.HeaderLanguageKey))

	req := request.WithContext(ctx)

	return handler.Handle(req)
}

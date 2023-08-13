package helper

import (
	"context"
	"encoding/json"

	"github.com/yonisaka/dating-service/internal/consts"
	"github.com/yonisaka/dating-service/internal/presentations"
)

// AuthInfoFromContext extract auth information from context
func AuthInfoFromContext(ctx context.Context) presentations.AuthInfo {
	var authInfo presentations.AuthInfo

	jsonBody, err := json.Marshal(ctx.Value(consts.CtxAuthInfo))
	if err != nil {
		return authInfo
	}

	err = json.Unmarshal(jsonBody, &authInfo)
	if err != nil {
		return authInfo
	}

	return authInfo
}

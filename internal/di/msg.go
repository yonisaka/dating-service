package di

import (
	"github.com/yonisaka/dating-service/pkg/logger"
	"github.com/yonisaka/dating-service/pkg/msg"
)

func RegistryMessage() {
	err := msg.Setup("msg.yaml", "config/")
	if err != nil {
		logger.Fatal(logger.MessageFormat("file message multi language load error %s", err.Error()))
	}

}

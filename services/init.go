package services

import (
	"github.com/gocastsian/adamak/protocol"
	"github.com/gocastsian/adamak/services/user"
)

func Init(config *protocol.ServiceInitializeConfig) {
	user.Init(config)
}

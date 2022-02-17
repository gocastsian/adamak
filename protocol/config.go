package protocol

import (
	"github.com/gocastsian/adamak/external"

	"github.com/labstack/echo/v4"
)

type ServiceInitializeConfig struct {
	StorageEngine string
	MySQL         external.MySQL
	Echo          *echo.Echo
}

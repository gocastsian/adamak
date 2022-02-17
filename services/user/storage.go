package user

import (
	"github.com/gocastsian/adamak/protocol"
)

// If you want to add more storage engine other than MySQL, comment two first lines and uncomment third one.

var storage *storage_MySQL
var _ protocol.UserServices = &storage_MySQL{}

// var storage protocol.UserServices

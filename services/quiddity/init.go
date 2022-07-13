/* For license and copyright information please see LEGAL file in ChaparKhane repository */

package quiddity

import (
	"../../libgo/protocol"
)

// Init method use to register services and storage structure in the application.
func Init() {
	protocol.App.RegisterService(&RegisterService)
	protocol.App.RegisterService(&GetService)

	protocol.OS.RegisterMediaType(&MediaType)
}

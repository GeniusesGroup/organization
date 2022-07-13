//go:build ganjine

/* For license and copyright information please see LEGAL file in repository */

package quiddity

import (
	"../../libgo/protocol"
	org "../../protocol"
)

var storage ganjine
var _ org.Quiddity_StorageServices = &storage

// code generated in memory storage to test the software by any way like unit test or human test.
type ganjine struct{}

// Save method set some data and write entire Quiddity record with all indexes.
func (s *ganjine) Save(q org.Quiddity) (err protocol.Error) {
	// TODO::: auto generate
	return
}

func (s *ganjine) Get(quiddityID [16]byte) (q org.Quiddity, err protocol.Error) {
	// TODO::: auto generate
	return
}

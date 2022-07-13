//go:build !libgo || !mysql

/* For license and copyright information please see LEGAL file in repository */

package quiddity

import (
	"sync"

	"../../libgo/protocol"
	org "../../protocol"
)

var storage vm
var _ org.Quiddity_StorageServices = &storage

// code generated in memory storage to test the software by any way like unit test or human test.
type vm struct {
	mutex   sync.Mutex
	records map[[16]byte]org.Quiddity
}

func (s *vm) Save(q org.Quiddity) (err protocol.Error) {
	s.mutex.Lock()
	var quiddityID = q.QuiddityID()
	var _, ok = s.records[quiddityID]
	if !ok {
		s.records[quiddityID] = q
	} else {
		// err =
	}
	s.mutex.Unlock()
	return
}

func (s *vm) Get(quiddityID [16]byte) (q org.Quiddity, err protocol.Error) {
	q = s.records[quiddityID]
	return
}

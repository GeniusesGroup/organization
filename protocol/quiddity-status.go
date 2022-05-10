/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type QuiddityStatus interface {
	QuiddityID() [16]byte    // Unique content ID
	Status() Quiddity_Status //
	Time() protocol.Time     // Save time
	RequestID() [16]byte     // user-request domain
}

type Quiddity_StorageServices interface {
	Save(qs QuiddityStatus) (err protocol.Error)

	Count(quiddityID [16]byte) (numbers uint64, err protocol.Error)
	Get(quiddityID [16]byte, versionOffset uint64) (qs QuiddityStatus, err protocol.Error)
	Last(quiddityID [16]byte) (qs QuiddityStatus, err protocol.Error)
}

// Quiddity_Status indicate Quiddity record status
type Quiddity_Status uint8

// Quiddity status
const (
	Quiddity_Status_Unset = iota
	Quiddity_Status_Registered
	Quiddity_Status_UnRegistered
	Quiddity_Status_Blocked
)

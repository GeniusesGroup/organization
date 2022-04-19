/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

/*
	Storage Data Structure
*/

type Quiddity interface {
	ID() [16]byte            // Unique content ID
	OrgID() [16]byte         // Owner of the quiddity
	Status() Quiddity_Status //
	Time() protocol.Time     // Set or update time
	RequestID() [16]byte     // user-request domain
}

type Quiddity_StorageServices interface {
	Save(q Quiddity) (err protocol.Error)

	Count(id [16]byte) (numbers uint64, err protocol.Error)
	Get(id [16]byte, versionOffset uint64) (q Quiddity, err protocol.Error)
	Last(id [16]byte) (q Quiddity, err protocol.Error)

	ListUserQuiddities(orgID [16]byte, offset, limit uint64) (ids [][16]byte, err protocol.Error)
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

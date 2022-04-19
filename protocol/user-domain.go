/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type UserDomain interface {
	Domain() string            // unique name in the computer world
	UserID() [16]byte          // user-status domain
	Status() UserDomain_Status //
	Time() protocol.Time       // Save time
	RequestID() [16]byte       // user-request domain
}

type UserDomain_StorageServices interface {
	Save(ud UserDomain) (err protocol.Error)

	Count(domain string) (numbers uint64, err protocol.Error)
	Get(domain string, versionOffset uint64) (ud UserDomain, err protocol.Error)
	Last(domain string) (ud UserDomain, err protocol.Error)

	FindByUserID(userID [16]byte offset, limit uint64) (domains []string, numbers uint64, err protocol.Error)
}

// UserDomain_Status indicate UserDomain record status
type UserDomain_Status uint8

// UserDoamin status
const (
	UserDomain_Status_Unset UserDomain_Status = iota
	UserDomain_Status_Active
	UserDomain_Status_Inactive
	UserDomain_Status_Blocked
)

/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type UserStatus interface {
	UserID() [16]byte    // Generated unique ID for the user
	Status() User_Status //
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

type UserStatus_StorageServices interface {
	Save(us UserStatus) protocol.Error

	Count(userID [16]byte) (numbers uint64, err protocol.Error)
	Get(userID [16]byte, versionOffset uint64) (us UserStatus, err protocol.Error)
	Last(userID [16]byte) (us UserStatus, numbers uint64, err protocol.Error)

	GetIDs(offset, limit uint64) (userIDs [][16]byte, numbers uint64, err protocol.Error)
	// GetIDsByDateTime(time protocol.Time, offset, limit uint64) (userIDs [][16]byte, numbers uint64, err protocol.Error)
}

// User_Status indicate UserStatus record status
type User_Status uint8

// User status
const (
	User_Status_Unset User_Status = iota
	User_Status_Active
	User_Status_Inactive
	User_Status_Blocked
)

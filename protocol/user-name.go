/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

/*
	Storage Data Structure
*/

type UserName interface {
	UserID() [16]byte       // user-status domain
	Username() string       // It is not replace of user ID. It usually use to find user by their friends in more human friendly manner.
	Status() UserNameStatus //
	RequestID() [16]byte    // user-request domain
}

// UserNameStatus indicate UserName record status
type UserNameStatus uint8

// UserName status
const (
	UserNameStatus_Unset    UserNameStatus = 0
	UserNameStatus_Register UserNameStatus = (1 << iota)
	UserNameStatus_Remove   UserNameStatus //= (1 << 1)
	UserNameStatus_Blocked  UserNameStatus //= (1 << 2)
	UserNameStatus_Reserved UserNameStatus //= (1 << 3)
)

type UserNameService interface {
	Register(username Username) protocol.Error

	Count(userID [16]byte) (numbers uint64, err protocol.Error)
	Get(userID [16]byte, versionOffset uint64) (u Username, err protocol.Error)
	Last(userID [16]byte) (u Username, err protocol.Error)

	Find(username string) (userID [16]byte, err protocol.Error)
}

/*
	Business Services
*/

type RegisterUsernameRequest interface {
	UserID() [16]byte // TODO::: get user id from connection??
	Username() string
}

type GetUsernameRequest interface {
	UserID() [16]byte
	VersionOffset() uint64
}
type GetUsernameResponse interface {
	Username
}

type GetLastUsernameRequest interface {
	UserID() [16]byte
}
type GetLastUsernameResponse interface {
	Username
}

type GetUsernameStatusRequest interface {
	Username() string
}
type GetUsernameStatusResponse interface {
	Status() UserNameStatus
}

// Admins services

type ChangeUsernameStatusRequest interface {
	Username() string
	Status() UserNameStatus
}

/*
	Errors
*/

/* For license and copyright information please see LEGAL file in repository */

package org

/*
	Storage Data Structure
*/

type UserStatus interface {
	UserID() [16]byte    // Generated unique ID for the user
	Status() UserStatus_ //
	RequestID() [16]byte // user-request domain
}

// UserStatus_ indicate UserStatus record status
type UserStatus_ uint8

// User status
const (
	UserStatus_Unset UserStatus_ = iota
	UserStatus_Active
	UserStatus_Inactive
	UserStatus_Blocked
)

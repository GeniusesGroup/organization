/* For license and copyright information please see LEGAL file in repository */

package org

/*
	Storage Data Structure
*/

// UserBlock indicate the user-block domain record data fields.
type UserBlock interface {
	UserID() [16]byte        // user-status domain
	BlockUserID() [16]byte   // user-status domain
	Status() UserBlockStatus //
	RequestID() [16]byte     // user-request domain
}

// UserBlockStatus indicate UserBlock record status
type UserBlockStatus uint8

// User status
const (
	UserBlockStatus_Unset UserBlockStatus = iota
	UserBlockStatus_Blocked
	UserBlockStatus_Unblocked
)

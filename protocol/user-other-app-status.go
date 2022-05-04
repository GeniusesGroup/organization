/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type UserOtherAppStatus interface {
	UserOtherAppID() [16]byte          // user-other-app domain
	Status() UserOtherAppStatus_Status //
	Time() protocol.Time               // Save time
	RequestID() [16]byte               // user-request domain
}

type UserOtherAppStatus_Status uint8

const (
	UserOtherAppStatus_Status_Unset UserOtherAppStatus_Status = iota
	UserOtherAppStatus_Status_Registered
	UserOtherAppStatus_Status_Inactivated
	UserOtherAppStatus_Status_Blocked
)

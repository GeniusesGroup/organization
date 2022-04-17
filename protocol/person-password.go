/* For license and copyright information please see LEGAL file in repository */

package org

// PersonPassword indicate the person-password domain record data fields.
type PersonPassword interface {
	PersonID() [16]byte // user-status domain
	PasswordHash() []byte
	Status() PersonPasswordStatus
	RequestID() [16]byte // user-request domain
}

// PersonPasswordStatus indicate PersonPassword record status
type PersonPasswordStatus uint8

// PersonPasswordStatus status
const (
	PersonPasswordStatus_Unset              PersonPasswordStatus = 0
	PersonPasswordStatus_ForceUse2Factor    PersonPasswordStatus = (1 << iota) // authenticate person with Password + OTP
	PersonPasswordStatus_MustChangePassword                                    // user must change password and otp key(if enabled) for some reason
)

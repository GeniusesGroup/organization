/* For license and copyright information please see LEGAL file in repository */

package org

/*
	Storage Data Structure
*/

// UserPublicKey indicate the user-public-key domain record data fields.
type UserPublicKey interface {
	UserID() [16]byte            // user-status domain
	PublicKey() []byte           // suggest use DER format
	Status() UserPublicKeyStatus //
	RequestID() [16]byte         // user-request domain
}

// UserPublicKeyStatus use to indicate UserPublicKey record status
type UserPublicKeyStatus uint8

// UserPublicKey status
const (
	UserPublicKeyStatus_Unset UserPublicKeyStatus = iota
	UserPublicKeyStatus_Active
	UserPublicKeyStatus_Inactive
	UserPublicKeyStatus_Blocked
	UserPublicKeyStatus_Revoked
)

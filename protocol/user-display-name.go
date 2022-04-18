/* For license and copyright information please see LEGAL file in repository */

package org

// UserDisplayName indicate the user-display-name domain record data fields.
// It is not unique and not replace of user ID. Just use to find user by desire name.
type UserDisplayName interface {
	UserID() [16]byte    // user-status domain
	DisplayName() string //
	RequestID() [16]byte // user-request domain
}

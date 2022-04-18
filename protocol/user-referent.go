/* For license and copyright information please see LEGAL file in repository */

package org

// UserReferent indicate the user-referent domain record data fields.
type UserReferent interface {
	UserID() [16]byte         // user-status domain
	ReferentUserID() [16]byte // user-status domain
	RequestID() [16]byte      // user-request domain
}

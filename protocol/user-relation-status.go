/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// UserRelationStatus indicate the domain record data fields.
type UserRelationStatus interface {
	UserID() [16]byte            // user-status domain
	SideID() [16]byte            // user-status domain
	Status() UserRelation_Status //
	Time() protocol.Time         // Save time
	RequestID() [16]byte         // user-request domain
}

type UserRelation_Status uint8

const (
	UserRelation_Status_Unset   UserRelation_Status = iota
	UserRelation_Status_Blocked UserRelation_Status = (1 << iota)
	UserRelation_Status_Muted
)

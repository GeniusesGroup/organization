/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ContentParticipant indicate the content-participant domain record data fields.
type ContentParticipant interface {
	ContentID() [16]byte           // content domain
	Type() ContentParticipant_Type //
	UserID() [16]byte              // user-status domain
	Time() protocol.Time           // Save time
	RequestID() [16]byte           // user-request domain
}

// TODO::: immutable?? if wrong data save for a content??
type ContentParticipant_StorageServices interface {
	Save(cp ContentParticipant) protocol.Error

	Count(contentID [16]byte) (numbers uint64, err protocol.Error)
	Get(contentID [16]byte, versionOffset uint64) (cp ContentParticipant, err protocol.Error)
	Last(contentID [16]byte) (cp ContentParticipant, numbers uint64, err protocol.Error)

	FindByUserID(userID [16]byte, offset, limit uint64) (contentIDs [][16]byte, numbers uint64, err protocol.Error)
}

// ContentParticipant_Type indicate participant type or role
type ContentParticipant_Type uint16

// ContentParticipant types
const (
	ContentParticipant_Type_Unset ContentParticipant_Type = iota
	ContentParticipant_Type_SongWriter
	ContentParticipant_Type_Composer
	ContentParticipant_Type_Regulator
	ContentParticipant_Type_Director
	ContentParticipant_Type_MusicSupplier
)

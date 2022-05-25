/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// CommentObject indicate the domain record data fields.
// Object can be any such as image, video, voice, ...
type CommentObject interface {
	CommentID() [16]byte // comment domain
	ObjectID() [16]byte  // object domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

type CommentForward_StorageServices interface {
	Save(co CommentObject) protocol.Error

	Count(commentID [16]byte) (numbers uint64, err protocol.Error)
	Get(commentID [16]byte, versionOffset uint64) (co CommentObject, err protocol.Error)
	Last(commentID [16]byte) (co CommentObject, numbers uint64, err protocol.Error)

	FindByObjectID(objectID [16]byte, offset, limit uint64) (commentIDs [][16]byte, numbers uint64, err protocol.Error)
}

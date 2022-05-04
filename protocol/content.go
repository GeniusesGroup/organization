/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type Content interface {
	QuiddityID() [16]byte       // quiddity domain. As content UUID
	UserID() [16]byte           // user-status domain
	Content_Type() Content_Type //
	Time() protocol.Time        // Save time
	RequestID() [16]byte        // user-request domain
}

type Content_StorageServices interface {
	Save(c Content) protocol.Error

	Count(quiddityID [16]byte) (numbers uint64, err protocol.Error)
	Get(quiddityID [16]byte, versionOffset uint64) (c Content, err protocol.Error)
	Last(quiddityID [16]byte) (c Content, numbers uint64, err protocol.Error)

	GetIDs(offset, limit uint64) (quiddityIDs [][16]byte, numbers uint64, err protocol.Error)
	// GetIDsByDateTime(time protocol.Time, offset, limit uint64) (quiddityIDs [][16]byte, numbers uint64, err protocol.Error)

	FindByUserID(userID [16]byte, offset, limit uint64) (quiddityIDs [][16]byte, numbers uint64, err protocol.Error)
}

type Content_Type uint16

const (
	Content_Type_Unset Content_Type = iota

	Content_Type_Article
	Content_Type_BlogPost

	Content_Type_Book
	Content_Type_Course
	Content_Type_Tutorial

	Content_Type_Album
	Content_Type_Music
	Content_Type_MusicVideo

	Content_Type_Movie
	Content_Type_Serial
	Content_Type_SerialEpisode
)

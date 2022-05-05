/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type Picture interface {
	RelatedID() [16]byte     // any domain e.g. user, quiddity, ...
	RelatedDomainID() uint64 //
	ObjectID() [16]byte      //
	Time() protocol.Time     // Save time
	RequestID() [16]byte     // user-request domain
}

type Picture_StorageServices interface {
	Save(p Picture) (err protocol.Error)

	// TODO::: how to change the order of pictures for a RelatedID

	Count(relatedID [16]byte) (numbers uint64, err protocol.Error)
	Get(relatedID [16]byte, versionOffset uint64) (p Picture, err protocol.Error)
	Last(relatedID [16]byte) (p Picture, numbers uint64, err protocol.Error)
}

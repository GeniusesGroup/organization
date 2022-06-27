/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// DiscountOwner indicate the domain record data fields.
type DiscountOwner interface {
	DiscountID() [16]byte // discount domain
	OrgID() [16]byte      // user domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type DiscountOwner_StorageServices interface {
	Save(dp DiscountOwner) (numbers uint64, err protocol.Error)

	Count(discountID [16]byte) (numbers uint64, err protocol.Error)
	Get(discountID [16]byte, versionOffset uint64) (dp DiscountOwner, numbers uint64, err protocol.Error)

	FindByOrg(orgID [16]byte, offset, limit uint64) (discountIDs [][16]byte, numbers uint64, err protocol.Error)
}

/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ProductProvider indicate the domain record data fields.
type ProductProviderStaff interface {
	ProductID() [16]byte // product domain
	StaffID() [16]byte   // staff domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type ProductProviderStaff_StorageServices interface {
	Save(pp ProductProviderStaff) protocol.Error

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pp ProductProviderStaff, err protocol.Error)
	Last(productID [16]byte) (pp ProductProviderStaff, numbers uint64, err protocol.Error)

	FindByStaff(staffID [16]byte, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)
}

// related UserIDs by StaffID use to get Coordinate(domain)

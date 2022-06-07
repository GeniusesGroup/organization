/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ProductProviderRole indicate the domain record data fields.
// ProductProviderRole or product operator is to supply a particular service.
type ProductProviderRole interface {
	ProductID() [16]byte // product domain
	RoleID() [16]byte    // role domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type ProductProviderRole_StorageServices interface {
	Save(pp ProductProviderRole) protocol.Error

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pp ProductProviderRole, err protocol.Error)
	Last(productID [16]byte) (pp ProductProviderRole, numbers uint64, err protocol.Error)

	FindByRoleID(roleID [16]byte, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)
}

// related UserIDs by RoleID use to get Coordinate(domain)

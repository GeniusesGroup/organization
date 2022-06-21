/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// VoucherProduct restrict use voucher on specific product
type VoucherProduct interface {
	VoucherID() [16]byte // voucher domain
	Each() uint8         // Each time use
	ProductID() [16]byte // product domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type VoucherProduct_StorageServices interface {
	Save(vp VoucherProduct) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vp VoucherProduct, numbers uint64, err protocol.Error)

	FindByProduct(productID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, numbers uint64, err protocol.Error)
}

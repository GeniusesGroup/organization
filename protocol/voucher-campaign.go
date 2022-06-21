/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// VoucherCampaign indicate the domain record data fields.
type VoucherCampaign interface {
	VoucherID() [16]byte  // voucher domain
	CampaignID() [16]byte // marketing-campaign domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type VoucherCampaign_StorageServices interface {
	Save(vc VoucherCampaign) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vc VoucherCampaign, numbers uint64, err protocol.Error)

	FindByCampaign(campaignID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, numbers uint64, err protocol.Error)
}

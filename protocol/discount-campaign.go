/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// DiscountCampaign indicate the domain record data fields.
type DiscountCampaign interface {
	DiscountID() [16]byte // discount domain
	CampaignID() [16]byte // marketing-campaign domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type DiscountCampaign_StorageServices interface {
	Save(dc DiscountCampaign) (numbers uint64, err protocol.Error)

	Count(discountID [16]byte) (numbers uint64, err protocol.Error)
	Get(discountID [16]byte, versionOffset uint64) (dc DiscountCampaign, numbers uint64, err protocol.Error)
}

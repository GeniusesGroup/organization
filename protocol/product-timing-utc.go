/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
	"../libgo/time/earth"
	"../libgo/time/utc"
)

type ProductTimingUTC interface {
	ProductID() [16]byte      // product domain
	Day() utc.DayElapsed      //
	DayHours() earth.DayHours // what hours in a day this service active to provide services
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}

type ProductTimingUTC_StorageServices interface {
	Save(pt ProductTimingUTC) protocol.Error

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pt ProductTimingUTC, err protocol.Error)
	Last(productID [16]byte) (pt ProductTimingUTC, numbers uint64, err protocol.Error)
}

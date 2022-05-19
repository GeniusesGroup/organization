/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"../libgo/protocol"
	"../libgo/time/earth"
	"../libgo/time/utc"
)

type ProductTimingUTC interface {
	ProductID() [16]byte         // product-status domain
	Weekdays() utc.Weekdays      // what days in weekday in utc time week allow to use this product
	DayHours() earth.DayHours    // what hours in day in utc time week allow to use this product
	Interval() protocol.Duration //
	Time() protocol.Time         // Save time
	RequestID() [16]byte         // user-request domain
}

/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"../libgo/protocol"
	"../libgo/time/utc"
)

type DiscountTimingUTC interface {
	DiscountID() [16]byte   // discount domain
	Weekdays() utc.Weekdays // what days in weekday in utc time week allow to use this discount
	DayHours() utc.DayHours // what hours in day in utc time week allow to use this discount
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}

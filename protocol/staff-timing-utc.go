/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"../libgo/protocol"
	"../libgo/time/earth"
	"../libgo/time/utc"
)

type StaffTimingUTC interface {
	StaffID() [16]byte        // staff-status domain
	Week() utc.WeekElapsed    //
	Weekdays() utc.Weekdays   // what days in weekday in utc time week allow to use this product
	DayHours() earth.DayHours // what hours in a day this role active to provide services
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}

/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"../libgo/protocol"
	"../libgo/time/earth"
	"../libgo/time/utc"
)

type OrgShiftUTC interface {
	ShiftID() [16]byte        // quiddity domain
	Weekdays() utc.Weekdays   // what days in weekday in utc time week this role active to provide services
	DayHours() earth.DayHours // what hours in day in utc time week this role active to provide services
	Type() OrgShift_Type      //
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}

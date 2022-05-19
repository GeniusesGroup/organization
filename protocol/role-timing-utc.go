/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"../libgo/protocol"
	"../libgo/time/earth"
	"../libgo/time/utc"
)

type RoleTimingUTC interface {
	RoleID() [16]byte         // role domain
	Shift() uint8             //
	Weekdays() utc.Weekdays   // what days in weekday in utc time week this role active to provide services
	DayHours() earth.DayHours // what hours in day in utc time week this role active to provide services
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}

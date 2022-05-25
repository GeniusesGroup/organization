/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"../libgo/protocol"
	"../libgo/time/earth"
	"../libgo/time/utc"
)

type DepartmentTimingUTC interface {
	DepartmentID() [16]byte   // department domain
	Weekdays() utc.Weekdays   // what days in weekday in utc time week this department active to provide services
	DayHours() earth.DayHours // what hours in day in utc time week this department active to provide services
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}

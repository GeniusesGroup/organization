/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"../libgo/protocol"
	"../libgo/time/earth"
	"../libgo/time/utc"
)

type RoleTimingUTC interface {
	RoleID() [16]byte         // role domain
	Weekdays() utc.Weekdays   // what days in weekday in utc time week this role active to provide services
	DayHours() earth.DayHours // what hours in day in utc time week this role active to provide services
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}

type RoleTimingUTC_StorageServices interface {
	Save(rt RoleTimingUTC) protocol.Error

	Count(roleID [16]byte) (numbers uint64, err protocol.Error)
	Get(roleID [16]byte, versionOffset uint64) (rt RoleTimingUTC, err protocol.Error)
	Last(roleID [16]byte) (rt RoleTimingUTC, numbers uint64, err protocol.Error)
}

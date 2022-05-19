/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"../libgo/protocol"
	"../libgo/time/utc"
)

type StaffOvertimeStatus interface {
	StaffID() [16]byte             // staff-status domain
	Day() utc.DayElapsed           //
	Status() StaffOvertime_Status //
	Time() protocol.Time           // Save time
	RequestID() [16]byte           // user-request domain
}

type StaffOvertime_Status uint8

const (
	StaffOvertime_Status_Unset StaffOvertime_Status = iota
	StaffOvertime_Status_Changed
	StaffOvertime_Status_NeedManagerAprove
	StaffOvertime_Status_ManagerApprove
)

/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"../libgo/protocol"
	"../libgo/time/earth"
	"../libgo/time/utc"
)

type StaffTakeLeaveStatus interface {
	StaffID() [16]byte             // staff-status domain
	Day() utc.DayElapsed           //
	Status() StaffTakeLeave_Status //
	Time() protocol.Time           // Save time
	RequestID() [16]byte           // user-request domain
}

type StaffTakeLeave_Status uint8

const (
	StaffTakeLeave_Status_Unset StaffTakeLeave_Status = iota
	StaffTakeLeave_Status_Changed
	StaffTakeLeave_Status_NeedManagerAprove
	StaffTakeLeave_Status_ManagerApprove
)

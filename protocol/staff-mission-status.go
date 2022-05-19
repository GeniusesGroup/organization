/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"../libgo/protocol"
	"../libgo/time/utc"
)

type StaffMissionStatus interface {
	StaffID() [16]byte           // staff-status domain
	Day() utc.DayElapsed         //
	Status() StaffMission_Status //
	Time() protocol.Time         // Save time
	RequestID() [16]byte         // user-request domain
}

type StaffMission_Status uint8

const (
	StaffMission_Status_Unset StaffMission_Status = iota
	StaffMission_Status_Changed
	StaffMission_Status_NeedManagerAprove
	StaffMission_Status_ManagerApproved
)

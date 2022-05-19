/* For license and copyright information please see LEGAL file in repository */

package org

import "../libgo/protocol"

type StaffStatus interface {
	ID() [16]byte         // user-status domain
	Status() Staff_Status //
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}

type Staff_Status uint8

const (
	Staff_Status_Unset Staff_Status = iota

	Staff_Status_Talent
	Staff_Status_NeedHRInterview
	Staff_Status_ApproveForHRInterview
	Staff_Status_NeedTechnicalInterview
	Staff_Status_ApproveForTechnicalInterview
	Staff_Status_NeedForManagerInterview
	Staff_Status_ApproveForManagerInterview
	Staff_Status_NeedForChiefInterview
	Staff_Status_ApproveForChiefInterview
	Staff_Status_Designation

	Staff_Status_Hire
	Staff_Status_Fire
	Staff_Status_End // death, end of contract, ...

	Staff_Status_Active   // Shift
	Staff_Status_Inactive // Shift
)

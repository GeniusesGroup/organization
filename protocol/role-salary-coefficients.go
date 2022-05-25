/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/math"
	"../libgo/protocol"
)

type RoleSalaryCoefficients interface {
	RoleID() [16]byte // role domain

	Importance() math.PerMyriad         // importance coefficient of the role
	Accompanying() math.PerMyriad       //
	ExpertiseAllowance() math.PerMyriad //
	HardshipAllowance() math.PerMyriad  //

	OvertimeWorking() math.PerMyriad //
	WorkingHolidays() math.PerMyriad //
	RotationShift() math.PerMyriad   //
	NightShift() math.PerMyriad      //

	Income() math.PerMyriad          // whole organization income
	DepartmenIncome() math.PerMyriad //

	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

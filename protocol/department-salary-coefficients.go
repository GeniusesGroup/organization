/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/math"
	"../libgo/protocol"
)

// due to organization payment capacity set values to show how many percent add or remove due to organization profit status
type DepartmentSalaryCoefficients interface {
	DepartmentID() [16]byte // department domain

	Importance() math.PerMyriad         // importance coefficient of the department
	Accompanying() math.PerMyriad       // Percent for each week
	ExpertiseAllowance() math.PerMyriad //
	HardshipAllowance() math.PerMyriad  //

	OvertimeWorking() math.PerMyriad //
	WorkingHolidays() math.PerMyriad //
	RotationShift() math.PerMyriad   //
	NightShift() math.PerMyriad      //

	Income() math.PerMyriad           // whole organization income in each week
	DepartmentIncome() math.PerMyriad //

	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

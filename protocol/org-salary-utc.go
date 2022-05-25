/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// OrgSalaryUTC is fixed salary that use in roles salary rules to calculate each staff salary.
type OrgSalaryUTC interface {
	WeeklySalary() protocol.AmountOfMoney // base salary for all staff in organization
	Time() protocol.Time                  // save time
	RequestID() [16]byte                  // user-request domain
}

/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// Department indicate the domain record data fields.
// It can use to show organizational chart. Also staffs can join to the related group.
type Department interface {
	DepartmentID() [16]byte // quiddity domain. use to get and show locale title.
	ParentID() [16]byte     // DepartmentID, Exist if it isn't top of the organization.
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type Department_StorageServices interface {
	Save(d Department) (err protocol.Error)

	Count(departmentID [16]byte) (numbers uint64, err protocol.Error)
	Get(departmentID [16]byte, versionOffset uint64) (d Department, err protocol.Error)
	Last(departmentID [16]byte) (d Department, numbers uint64, err protocol.Error)
}

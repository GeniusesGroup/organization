/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// RoleDepartment indicate the domain record data fields.
type RoleDepartment interface {
	RoleID() [16]byte       // role domain
	DepartmentID() [16]byte // department domain
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}

type RoleDepartment_StorageServices interface {
	Save(rd RoleDepartment) protocol.Error

	Count(roleID [16]byte) (numbers uint64, err protocol.Error)
	Get(roleID [16]byte, versionOffset uint64) (rd RoleDepartment, err protocol.Error)
	Last(roleID [16]byte) (rd RoleDepartment, err protocol.Error)

	FindByDepartmentID(departmentID [16]byte, offset, limit uint64) (roleIDs [][16]byte, numbers uint64, err protocol.Error)
}

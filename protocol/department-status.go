/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type DepartmentStatus interface {
	DepartmentID() [16]byte    // department domain
	Status() Department_Status //
	Time() protocol.Time       // Save time
	RequestID() [16]byte       // user-request domain
}

type DepartmentStatus_StorageServices interface {
	Save(ds DepartmentStatus) (err protocol.Error)

	Count(departmentID [16]byte) (numbers uint64, err protocol.Error)
	Get(departmentID [16]byte, versionOffset uint64) (ds DepartmentStatus, err protocol.Error)
	Last(departmentID [16]byte) (ds DepartmentStatus, numbers uint64, err protocol.Error)
}

type Department_Status uint8

const (
	Department_Status_Unset           Department_Status = 0
	Department_Status_PermanentClosed Department_Status = (1 << iota)
	Department_Status_TemporaryClosed
	Department_Status_Blocked
)

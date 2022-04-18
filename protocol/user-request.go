/* For license and copyright information please see LEGAL file in repository */

package org

/*
	Storage Data Structure
*/

// UserRequest use to prove how other records created as a chain to other records or standalone data.
type UserRequest interface {
	ID() [16]byte               // RequestID
	UserID() [16]byte           // user-status domain
	UserConnectionID() [16]byte // Store to remember request belong to which Connection/AppInstance.
	ServiceID() uint64          //
	CompressID() uint64         // gzip, ...
	Request() []byte            // Decode by service request structure and compress type.
	Signature() []byte          // Store to prove user requested to set||update this record.
}

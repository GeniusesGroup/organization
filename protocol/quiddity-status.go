/* For license and copyright information please see LEGAL file in repository */

package org

// Quiddity_Status indicate any quiddity record status
type Quiddity_Status uint64

// base quiddity status
const (
	Quiddity_Status_Unset                Quiddity_Status = 0
	Quiddity_Status_Edited               Quiddity_Status = (1 << iota)
	Quiddity_Status_Locked                               // by owner
	Quiddity_Status_Blocked                              // by justice(admins)
	Quiddity_Status_PermanentInactivated                 // Due to some reason quiddity was Locked, Closed, ... and must not use or edit
	Quiddity_Status_TemporaryInactivated
	Quiddity_Status_Hidden
	Quiddity_Status_Deleted
	_
	_
	_
	_
	_
	_
	_
	_
	_
	Quiddity_Status_FreeFlag
)

//go:build mysql

/* For license and copyright information please see LEGAL file in repository */

package quiddity

import (
	"database/sql"

	"../../libgo/protocol"
	org "../../protocol"
)

// TODO::: code generated in memory storage to test the software by any way like unit test or human test.
var storage mysql
var _ org.Quiddity_StorageServices = &storage

type mysql struct {
	db *sql.DB
}

func (s *mysql) Save(r org.Quiddity) (err protocol.Error) {
	const query = "INSERT INTO Quiddity VALUES (?, ?, ?, ?)"
	_, err = s.db.Exec(query, r.QuiddityID(), r.DomainID(), r.RequestID(), r.Time())
	return
}

func (s *mysql) Get(quiddityID [16]byte) (q org.Quiddity, err protocol.Error) {
	const query = "SELECT * FROM Quiddity WHERE QuiddityID = ?"
	var rows *sql.Rows
	rows, err = s.db.Query(query, quiddityID)
	// TODO:::
	return
}

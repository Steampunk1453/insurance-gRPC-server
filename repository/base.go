package repository

import (
	"github.com/hashicorp/go-memdb"
)

var db *memdb.MemDB

func init() {
	// Create the DB schema
	schema := &memdb.DBSchema {
		Tables: map[string]*memdb.TableSchema{
			"user": {
				Name: "user",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
					},
					"mobileNumber": {
						Name:    "mobileNumber",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "MobileNumber"},
					},
				},
			},
			"policy": {
				Name: "policy",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
					},
					"mobileNumber": {
						Name:    "mobileNumber",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "MobileNumber"},
					},
				},
			},
		},
	}
	conn, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
	db = conn
}

func GetDB() *memdb.MemDB {
	return db
}


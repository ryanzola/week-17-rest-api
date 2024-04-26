package db

import "os"

var DBNAME = os.Getenv("MONGO_DB_NAME")
var TestDBNAME = os.Getenv("MONGO_DB_TEST_NAME")
var DBURI = os.Getenv("MONGO_URI")

type Store struct {
	Record RecordStore
}

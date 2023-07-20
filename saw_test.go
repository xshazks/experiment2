package gosaw

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "dbmonitor",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertMonitor(t *testing.T) {
	proker := "Sosialisasi"
	status := "Progress"
	about := "Edukasi jasmani rohani" 
	karyawan := "Wawan Hariawan"
	hasil := InsertMonitor(MongoConn, proker, status, about, karyawan)
	fmt.Println(hasil)
}

func TestGetDataMonitor(t *testing.T) {
	status := "Progress"
	dt := GetDataMonitor(status)
	fmt.Println(dt)
}


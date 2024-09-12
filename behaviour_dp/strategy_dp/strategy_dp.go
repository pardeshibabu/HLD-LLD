package strategydp

import "fmt"

type IDBConnection interface {
	Connect()
}

type DBConnection struct {
	DB IDBConnection
}

func (dbc DBConnection) DBConnect() {
	dbc.DB.Connect()
}

type MysqlConnection struct {
	URL string
}

func (mc *MysqlConnection) Connect() {
	fmt.Println("MySql:: ", mc.URL)
}

type MongoDBlConnection struct {
	URL string
}

func (mc *MongoDBlConnection) Connect() {
	fmt.Println("MongoDB:: ", mc.URL)
}

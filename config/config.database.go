package config

const (
	DriverSqlite   Driver = "sqlite"
	DriverMysql    Driver = "mysql"
	DriverPostgres Driver = "postgres"
)

type Driver string

type Database struct {
	Driver   Driver `comment:"Driver options: sqlite, mysql, postgres"`
	Host     string
	Port     uint16
	User     string
	Password string
	Database string
	Tail     string
}

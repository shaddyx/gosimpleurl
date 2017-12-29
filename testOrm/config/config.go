package config

type config struct {
	DriverName     string
	DataSourceName string
}

// Greeet: this is config
var Config = config{
	DriverName:     "sqlite3",
	DataSourceName: "test.db",
}

package config

// DatabaseConfig holds MySQL connection information
var DatabaseConfig = map[string]string{
	"user":     "root",
	"password": "root",
	"host":     "127.0.0.1",
	"port":     "8889",
	"dbname":   "go_todo",
}

// GetDSN formats the MySQL connection string
func GetDSN() string {
	return DatabaseConfig["user"] + ":" +
		DatabaseConfig["password"] + "@tcp(" +
		DatabaseConfig["host"] + ":" +
		DatabaseConfig["port"] + ")/" +
		DatabaseConfig["dbname"]
}

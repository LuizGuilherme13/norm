package models

type DBConn struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (db *DBConn) String() string {
	return ("host=" + db.Host +
		" port=" + db.Port +
		" user=" + db.User +
		" password=" + db.Password +
		" dbname=" + db.DBName +
		" sslmode=" + db.SSLMode)
}

package unidb

type connection interface {
	Connect() error
	Disconnect()
}

type Reader interface {
	Find()
	FindAll(query string, args ...any) ([]map[string]interface{}, error)
}

type Writer interface {
	New()
	Edit()
	Delete()
}

type Repository interface {
	connection
	Reader
	Writer
}

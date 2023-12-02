package unodatabase

func (db *DB) Create() error {
	return db.Repository.Create()
}

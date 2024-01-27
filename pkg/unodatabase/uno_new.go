package unodatabase

func (db *DB) New() error {
	return db.Repository.New()
}

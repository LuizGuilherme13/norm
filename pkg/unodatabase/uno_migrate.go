package unodatabase

func (db *DB) Migrate(model any) {
	db.Repository.Migrate(model)
}

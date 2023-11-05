package postgres

func (r *Repository) Disconnect() {
	r.DB.Close()
}

package unodatabase

import "github.com/LuizGuilherme13/unodatabase/pkg/internal/structutils"

// Migrate por enquanto apenas pega o nome da struct e passa para UniDB.Table.
// Sendo necessário utilizar este método sempre que for realizar alguma solicitação ao banco.
func (udb *UniDB) Migrate(model interface{}) error {
	tableName := structutils.GetName(model)

	udb.Query.T = tableName
	return nil
}

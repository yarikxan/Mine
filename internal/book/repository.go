// repository.go
package book

import "database/sql"

type Repository struct {
	db sql.DB
}

package helper

import "database/sql"

func CategoryTxCommit(tx *sql.Tx) {
	defer func() {
		if r := recover(); r != nil {
			err := tx.Rollback()
			CategoryPanicIfError(err)
		}
	}()
	err := tx.Commit()
	CategoryPanicIfError(err)
}

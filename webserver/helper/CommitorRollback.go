package helper

import "database/sql"

func CommitorRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}

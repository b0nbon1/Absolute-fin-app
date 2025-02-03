package db

type Store struct {
	*Queries
	db DBTX
}




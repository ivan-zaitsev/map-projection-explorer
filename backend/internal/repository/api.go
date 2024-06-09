package repository

import (
	"database/sql"
	"map-projection-explorer-backend/internal/domain/model"
)

func NewDatabase(uri string) (*sql.DB, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

type EpsgExtentRepository interface {
	FindAllAfterCode(search string, afterCode *int, size int) ([]*model.EpsgExtentRecord, error)
	FindByCode(code int) (*model.EpsgExtentRecord, error)
}

type SrsRepository interface {
	FindByCode(code int) (*model.SrsRecord, error)
}

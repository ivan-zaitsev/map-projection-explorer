package repository

import (
	"database/sql"
	"map-projection-explorer-backend/internal/domain/model"
)

type srsRepository struct {
	db *sql.DB
}

func NewSrsRepository(db *sql.DB) SrsRepository {
	return &srsRepository{db: db}
}

func (s *srsRepository) FindByCode(code int) (*model.SrsRecord, error) {
	query := `
          SELECT 
              proj4text
          FROM spatial_ref_sys
          WHERE srid = $1`

	row := s.db.QueryRow(query, code)

	record := model.SrsRecord{}

	err := row.Scan(&record.Proj4text)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

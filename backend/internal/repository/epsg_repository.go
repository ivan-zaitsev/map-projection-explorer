package repository

import (
	"database/sql"
	"map-projection-explorer-backend/internal/domain/model"
)

type epsgExtentRepository struct {
	db *sql.DB
}

func NewEpsgExtentRepository(db *sql.DB) EpsgExtentRepository {
	return &epsgExtentRepository{db: db}
}

func (e *epsgExtentRepository) FindAllAfterCode(
	search string, afterCode *int, size int) ([]*model.EpsgExtentRecord, error) {

	query := `
          SELECT
              ec.coord_ref_sys_name,
              ec.coord_ref_sys_code
          FROM epsg_coordinatereferencesystem ec
          WHERE 
              (LOWER(ec.coord_ref_sys_name) LIKE LOWER(CONCAT('%%',$1::varchar,'%%')) OR 
               LOWER(ec.coord_ref_sys_code::varchar) LIKE LOWER(CONCAT('%%',$1::varchar,'%%'))) AND 
              ($2::integer IS NULL OR ec.coord_ref_sys_code > $2)
          ORDER BY ec.coord_ref_sys_code LIMIT $3`

	rows, err := e.db.Query(query, search, afterCode, size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []*model.EpsgExtentRecord

	for rows.Next() {
		record := &model.EpsgExtentRecord{}

		err := rows.Scan(&record.Name, &record.Code)
		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}
	return records, nil
}

func (e *epsgExtentRepository) FindByCode(code int) (*model.EpsgExtentRecord, error) {
	query := `
          SELECT
              ec.coord_ref_sys_name,
              ec.coord_ref_sys_code,
              ee.bbox_south_bound_lat, 
              ee.bbox_west_bound_lon, 
              ee.bbox_north_bound_lat, 
              ee.bbox_east_bound_lon 
          FROM epsg_coordinatereferencesystem ec
          LEFT JOIN epsg_usage eu ON ec.coord_ref_sys_code = eu.object_code
          LEFT JOIN epsg_extent ee ON eu.extent_code = ee.extent_code
          WHERE ec.coord_ref_sys_code = $1`

	row := e.db.QueryRow(query, code)

	record := model.EpsgExtentRecord{}

	err := row.Scan(&record.Name, &record.Code,
		&record.BboxSouthBoundLat, &record.BboxWestBoundLon, &record.BboxNorthBoundLat, &record.BboxEastBoundLon)

	if err != nil {
		return nil, err
	}
	return &record, nil
}

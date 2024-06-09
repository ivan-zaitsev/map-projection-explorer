package service

import (
	"database/sql"
	"errors"
	"log"
	"map-projection-explorer-backend/internal/domain/dto"
	"map-projection-explorer-backend/internal/domain/model"
	"map-projection-explorer-backend/internal/repository"
	"math"
)

type crsService struct {
	epsgRepository repository.EpsgExtentRepository
	srsRepository  repository.SrsRepository
}

func NewCrsService(epsgRepository repository.EpsgExtentRepository, srsRepository repository.SrsRepository) CrsService {
	return &crsService{epsgRepository: epsgRepository, srsRepository: srsRepository}
}

func (c *crsService) FindCoordinateReferenceSystem(code int) (*dto.CrsRecord, *dto.ServiceError) {
	epsgRecord, err := c.epsgRepository.FindByCode(code)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, &dto.ServiceError{Code: dto.ErrorCodeNotFound, Message: "Extent with such code not found"}
	}
	if err != nil {
		log.Printf("Failed to find epsg record by code: %s", err)
		return nil, &dto.ServiceError{Code: dto.ErrorCodeInternalServer, Message: "Unknown error"}
	}

	srsRecord, err := c.srsRepository.FindByCode(code)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, &dto.ServiceError{Code: dto.ErrorCodeNotFound, Message: "SRS with such code not found"}
	}
	if err != nil {
		log.Printf("Failed to find srs record by code: %s", err)
		return nil, &dto.ServiceError{Code: dto.ErrorCodeInternalServer, Message: "Unknown error"}
	}

	crsRecord := convertToCrsRecordFull(epsgRecord, srsRecord)
	return crsRecord, nil
}

func (c *crsService) FindAllCoordinateReferenceSystems(
	search string, pageCursor *int, pageSize int) (*dto.Page[dto.CrsRecord], *dto.ServiceError) {

	if pageSize <= 0 || pageSize == math.MaxInt {
		return nil, &dto.ServiceError{Code: dto.ErrorInvalidRequest, Message: "Invalid page size"}
	}

	epsgRecords, err := c.epsgRepository.FindAllAfterCode(search, pageCursor, pageSize+1)
	if err != nil {
		log.Printf("Failed to find epsg records by code: %s", err)
		return nil, &dto.ServiceError{Code: dto.ErrorCodeInternalServer, Message: "Unknown error"}
	}

	var crsRecords []*dto.CrsRecord
	for _, epsgRecord := range epsgRecords {
		crsRecords = append(crsRecords, convertToCrsRecord(epsgRecord))
	}

	if len(epsgRecords) > pageSize {
		crsRecords = crsRecords[:pageSize]
	}
	page := dto.Page[dto.CrsRecord]{
		Content:        crsRecords,
		NextPageCursor: findNextPageCursor(crsRecords, len(epsgRecords), pageSize),
	}
	return &page, nil
}

func findNextPageCursor(records []*dto.CrsRecord, recordsLen, pageSize int) *int {
	if len(records) == 0 || recordsLen <= pageSize {
		return nil
	}
	return &records[len(records)-1].Code
}

func convertToCrsRecord(epsgRecord *model.EpsgExtentRecord) *dto.CrsRecord {
	crsRecord := dto.CrsRecord{
		Name: epsgRecord.Name,
		Code: epsgRecord.Code,
	}
	return &crsRecord
}

func convertToCrsRecordFull(epsgRecord *model.EpsgExtentRecord, srsRecord *model.SrsRecord) *dto.CrsRecord {
	crsRecord := dto.CrsRecord{
		Name: epsgRecord.Name,
		Code: epsgRecord.Code,
		Definition: &dto.CrsRecordDefinition{
			Type:  dto.CrsRecordDefinitionTypeProj4,
			Value: srsRecord.Proj4text,
		},
		Extent: &dto.CrsRecordExtent{
			WestLongitude: &epsgRecord.BboxWestBoundLon,
			SouthLatitude: &epsgRecord.BboxSouthBoundLat,
			EastLongitude: &epsgRecord.BboxEastBoundLon,
			NorthLatitude: &epsgRecord.BboxNorthBoundLat,
		},
		ExtentCenter: calculateExtentCenter(epsgRecord),
	}
	return &crsRecord
}

func calculateExtentCenter(epsgRecord *model.EpsgExtentRecord) *dto.CrsRecordExtentCenter {
	n := epsgRecord.BboxNorthBoundLat
	s := epsgRecord.BboxSouthBoundLat
	w := epsgRecord.BboxWestBoundLon
	e := epsgRecord.BboxEastBoundLon

	lon := w + 180 + (360-(w+180)+e+180)/2.0
	lat := (n-s)/2.0 + s

	return &dto.CrsRecordExtentCenter{
		Longitude: &lon,
		Latitude:  &lat,
	}
}

package service

import (
	"database/sql"
	"log"
	"map-projection-explorer-backend/internal/domain/dto"
	"map-projection-explorer-backend/internal/domain/model"
	"map-projection-explorer-backend/internal/repository"
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
	if err == sql.ErrNoRows {
		return nil, &dto.ServiceError{Code: dto.ErrorCodeNotFound, Message: "Extent with such code not found"}
	}
	if err != nil {
		log.Printf("Failed to find epsg record by code: %s", err)
		return nil, &dto.ServiceError{Code: dto.ErrorCodeInternalServer, Message: "Unknown error"}
	}

	srsRecord, err := c.srsRepository.FindByCode(code)
	if err == sql.ErrNoRows {
		return nil, &dto.ServiceError{Code: dto.ErrorCodeNotFound, Message: "SRS with such code not found"}
	}
	if err != nil {
		log.Printf("Failed to find srs record by code: %s", err)
		return nil, &dto.ServiceError{Code: dto.ErrorCodeInternalServer, Message: "Unknown error"}
	}

	crsRecord := convertToCrsRecordFull(epsgRecord, srsRecord)
	return crsRecord, nil
}

func (c *crsService) FindAllCoordinateReferenceSystems(cursorCode *int, pageSize int) ([]*dto.CrsRecord, *dto.ServiceError) {
	epsgRecords, err := c.epsgRepository.FindAllAfterCode(cursorCode, pageSize)
	if err != nil {
		log.Printf("Failed to find epsg records by code: %s", err)
		return nil, &dto.ServiceError{Code: dto.ErrorCodeInternalServer, Message: "Unknown error"}
	}

	var crsRecords []*dto.CrsRecord
	for _, epsgRecord := range epsgRecords {
		crsRecords = append(crsRecords, convertToCrsRecord(epsgRecord))
	}
	return crsRecords, nil
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
			WestLongitude: epsgRecord.BboxWestBoundLon,
			SouthLatitude: epsgRecord.BboxSouthBoundLat,
			EastLongitude: epsgRecord.BboxEastBoundLon,
			NorthLatitude: epsgRecord.BboxNorthBoundLat,
		},
	}
	return &crsRecord
}

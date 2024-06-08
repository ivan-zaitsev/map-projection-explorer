package dto

type CrsRecordDefinitionType string

var (
	CrsRecordDefinitionTypeProj4 CrsRecordDefinitionType = "proj4"
)

type Page[T any] struct {
	Content []*T `json:"content"`
}

type CrsRecord struct {
	Name       string               `json:"name,omitempty"`
	Code       int                  `json:"code,omitempty"`
	Definition *CrsRecordDefinition `json:"definition,omitempty"`
	Extent     *CrsRecordExtent     `json:"extent,omitempty"`
}

type CrsRecordDefinition struct {
	Type  CrsRecordDefinitionType `json:"type,omitempty"`
	Value string                  `json:"value,omitempty"`
}

type CrsRecordExtent struct {
	WestLongitude float64 `json:"west_longitude,omitempty"`
	SouthLatitude float64 `json:"south_latitude,omitempty"`
	EastLongitude float64 `json:"east_longitude,omitempty"`
	NorthLatitude float64 `json:"north_latitude,omitempty"`
}

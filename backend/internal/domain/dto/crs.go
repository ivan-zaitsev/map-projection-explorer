package dto

type CrsRecordDefinitionType string

var (
	CrsRecordDefinitionTypeProj4 CrsRecordDefinitionType = "proj4"
)

type Page[T any] struct {
	Content        []*T `json:"content,omitempty"`
	NextPageCursor *int `json:"nextPageCursor,omitempty"`
}

type CrsRecord struct {
	Name         string                 `json:"name,omitempty"`
	Code         int                    `json:"code,omitempty"`
	Definition   *CrsRecordDefinition   `json:"definition,omitempty"`
	Extent       *CrsRecordExtent       `json:"extent,omitempty"`
	ExtentCenter *CrsRecordExtentCenter `json:"extentCenter,omitempty"`
}

type CrsRecordDefinition struct {
	Type  CrsRecordDefinitionType `json:"type,omitempty"`
	Value string                  `json:"value,omitempty"`
}

type CrsRecordExtent struct {
	WestLongitude *float64 `json:"westLongitude,omitempty"`
	SouthLatitude *float64 `json:"southLatitude,omitempty"`
	EastLongitude *float64 `json:"eastLongitude,omitempty"`
	NorthLatitude *float64 `json:"northLatitude,omitempty"`
}

type CrsRecordExtentCenter struct {
	Longitude *float64 `json:"longitude,omitempty"`
	Latitude  *float64 `json:"latitude,omitempty"`
}

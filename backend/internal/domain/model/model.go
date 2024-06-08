package model

type EpsgExtentRecord struct {
	Name              string
	Code              int
	BboxSouthBoundLat float64
	BboxWestBoundLon  float64
	BboxNorthBoundLat float64
	BboxEastBoundLon  float64
}

type SrsRecord struct {
	Proj4text string
}

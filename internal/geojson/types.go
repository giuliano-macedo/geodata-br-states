package geogjson

type Coordinate [2]float64

type GeometryType string

const (
	MultiPolygon GeometryType = "MultiPolygon"
)

type Geometry struct {
	Type        GeometryType     `json:"type"`
	Coordinates [][][]Coordinate `json:"coordinates"`
}

type FeatureType string

const (
	FeatureTypeFeature        FeatureType = "Feature"
	FeatureTypeCollectionType FeatureType = "FeatureCollection"
)

type Feature[T any] struct {
	Id           string      `json:"id"`
	Geometry     Geometry    `json:"geometry"`
	GeometryName string      `json:"geometry_name"`
	Properties   T           `json:"properties"`
	Type         FeatureType `json:"type"`
}

type FeatureCollection[T any] struct {
	Features      []Feature[T] `json:"features"`
	Type          FeatureType  `json:"type"`
	Crs           interface{}  `json:"crs"`
	TimeStamp     string       `json:"timeStamp"`
	TotalFeatures int          `json:"totalFeatures"`
	Name          string       `json:"name"`
}

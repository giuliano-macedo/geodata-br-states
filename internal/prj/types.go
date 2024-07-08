package prj

type CoordinateSystem struct {
	Name                string
	GeoCoordinateSystem GeoCoordinateSystem
	Projection          string
	FalseEasting        float64
	FalseNorthing       float64
	CentralMeridian     float64
	ScaleFactor         float64
	LatitudeOfOrigin    float64
	Unit                Unit
}

type GeoCoordinateSystem struct {
	Name          string
	Datum         Datum
	PrimeMeridian Primem
	Unit          Unit
}

type Datum struct {
	Name     string
	Spheroid Spheroid
}

type Spheroid struct {
	Name              string
	SemiMajorAxis     float64
	InverseFlattening float64
}

type Primem struct {
	Name  string
	Value float64
}

type Unit struct {
	Name  string
	Value float64
}

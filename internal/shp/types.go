package shp

type (
	Integer int32
	Double  float64
)

type ShapeFile struct {
	Header  Header
	Records []Record
}

type Header struct {
	FirstPart struct {
		Magic      uint32
		_          [5]uint32
		FileLength Integer
	}
	SecondPart struct {
		Version   Integer
		ShapeType ShapeType
		Mbr       [4]Double
		ZRange    [2]Double
		MRange    [2]Double
	}
}

type Record struct {
	Header struct {
		Number        Integer
		ContentLength Integer // This is 16-bit length of the Content
	}
	Polygon PolygonM
}

type Point struct {
	X, Y float64
}

type PolygonM struct {
	Header struct {
		Box       [4]Double // Bounding Box
		NumParts  Integer   // Number of Parts
		NumPoints Integer   // Total Number of Points
	}
	Parts  []Integer // Index to First Point in Part (length=NumParts)
	Points []Point   // Points for All Parts (length=NumPoints)
	MRange [2]Double // Bounding Measure Range
	MArray []Double  // Measures for All Points (length=NumPoints)
}

type ShapeType Integer

var (
	ShapePolygonM ShapeType = 25
)

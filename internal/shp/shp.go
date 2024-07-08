// Reference: https://www.esri.com/content/dam/esrisites/sitecore-archive/Files/Pdfs/library/whitepapers/pdfs/shapefile.pdf
package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

func (p PolygonM) PointsSegments() [][]Point {
	n := int(p.Header.NumParts)
	segments := make([][]Point, n)

	for i, start := range p.Parts {
		var end Integer
		if i == n-1 {
			end = p.Header.NumPoints
		} else {
			end = p.Parts[i+1]
		}

		segments[i] = p.Points[start:end]
	}

	return segments
}

func ReadRecord(reader io.Reader) (record Record, err error) {
	if err = binary.Read(reader, binary.BigEndian, &record.Header); err != nil {
		return
	}

	var recordType ShapeType
	if err = binary.Read(reader, binary.LittleEndian, &recordType); err != nil {
		return
	}

	if recordType != ShapePolygonM {
		return record, fmt.Errorf("shapeType is not PolygonM of record %v (%v)", record.Header.Number, recordType)
	}

	polygon := &record.Polygon

	if err = binary.Read(reader, binary.LittleEndian, &polygon.Header); err != nil {
		return
	}

	polygon.Parts = make([]Integer, polygon.Header.NumParts)
	polygon.Points = make([]Point, polygon.Header.NumPoints)
	polygon.MArray = make([]Double, polygon.Header.NumPoints)

	// NOTE: I don't want to write 4 if err == nil (:
	errs := [4]error{
		binary.Read(reader, binary.LittleEndian, &polygon.Parts),
		binary.Read(reader, binary.LittleEndian, &polygon.Points),
		binary.Read(reader, binary.LittleEndian, &polygon.MRange),
		binary.Read(reader, binary.LittleEndian, &polygon.MArray),
	}
	for _, err = range errs {
		if err != nil {
			return
		}
	}
	return
}

func ReadShp(reader io.Reader) (shape ShapeFile, err error) {
	if err = binary.Read(reader, binary.BigEndian, &shape.Header.FirstPart); err != nil {
		return
	}

	if shape.Header.FirstPart.Magic != 0x0000270a {
		return shape, fmt.Errorf("invalid format (%X)", shape.Header.FirstPart.Magic)
	}

	if err = binary.Read(reader, binary.LittleEndian, &shape.Header.SecondPart); err != nil {
		return
	}

	if shape.Header.SecondPart.Version != 1000 {
		return shape, fmt.Errorf("invalid version (%v)", shape.Header.SecondPart.Version)
	}

	if shape.Header.SecondPart.ShapeType != ShapePolygonM {
		return shape, fmt.Errorf("shapeType is not PolygonM (%v)", shape.Header.SecondPart.ShapeType)
	}

recordParsing:
	for i := 1; true; i++ {
		var record Record
		record, err = ReadRecord(reader)
		switch {
		case err == io.EOF:
			err = nil
			break recordParsing
		case err != nil:
			return
		}
		shape.Records = append(shape.Records, record)
	}

	return
}

package core

import (
	"fmt"

	"github.com/giuliano-oliveira/geodata-br-states/internal/dbf"
	geojson "github.com/giuliano-oliveira/geodata-br-states/internal/geojson"
	"github.com/giuliano-oliveira/geodata-br-states/internal/prj"
	"github.com/giuliano-oliveira/geodata-br-states/internal/shp"
)

type FeatureCollection geojson.FeatureCollection[dbf.Record]

func convertPolygonToCoordinates(p shp.PolygonM) [][][]geojson.Coordinate {
	segments := p.PointsSegments()
	coordinates := make([][][]geojson.Coordinate, len(segments))

	for i, segment := range segments {
		segmentCoordiantes := make([]geojson.Coordinate, len(segment))
		for j, point := range segment {
			segmentCoordiantes[j] = geojson.Coordinate{point.Y, point.X}
		}

		coordinates[i] = [][]geojson.Coordinate{segmentCoordiantes}
	}

	return coordinates
}

func convertToGeoJson(cs prj.CoordinateSystem, shape shp.ShapeFile, db dbf.Dbf) (all FeatureCollection, states []FeatureCollection) {
	numStates := len(shape.Records)

	all.Name = fmt.Sprintf("States of brazil feature collection converted from projection %v to %v datum", cs.Name, prj.WgsSphere.Name)
	all.Type = geojson.FeatureTypeCollectionType
	all.TotalFeatures = numStates
	all.TimeStamp = db.Header.LastUpdated.Date().String()

	allFeatures := make([]geojson.Feature[dbf.Record], numStates)
	for i, record := range shape.Records {
		feature := &(allFeatures[i])
		dbRecord := db.Records[i]

		feature.Id = dbRecord.Sigla
		feature.Type = geojson.FeatureTypeFeature
		feature.GeometryName = dbRecord.Estado
		feature.Properties = dbRecord

		feature.Geometry.Type = geojson.MultiPolygon
		feature.Geometry.Coordinates = convertPolygonToCoordinates(record.Polygon)
	}

	states = make([]FeatureCollection, numStates)
	for i := range states {
		states[i] = all
		states[i].TotalFeatures = 1
		states[i].Name = fmt.Sprintf("state of %v feature collection converted from projection %v to %v datum", db.Records[i].Estado, cs.Name, prj.WgsSphere.Name)
		states[i].Features = []geojson.Feature[dbf.Record]{allFeatures[i]}
	}

	all.Features = allFeatures
	return
}

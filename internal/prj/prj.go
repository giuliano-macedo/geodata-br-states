package prj

import (
	"fmt"
	"io"
)

func ReadCoordinateSystem(reader io.Reader) (cs CoordinateSystem, err error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return
	}

	node, err := parseTree(string(data))
	if err != nil {
		return
	}

	return buildCoordinateSystemFromNode(node)
}

var paramNames = []string{
	"False_Easting",
	"False_Northing",
	"Central_Meridian",
	"Scale_Factor",
	"Latitude_Of_Origin",
}

func buildCoordinateSystemFromNode(node *Node) (cs CoordinateSystem, err error) {
	if err = assertNode("coordinate system", node, NodeProjcs, 8); err != nil {
		return
	}

	cs.Name = node.Name
	var (
		geogcsNode     = node.Children[0]
		projectionNode = node.Children[1]
		parameterNodes = node.Children[2:7]
		unitNode       = node.Children[7]
	)

	cs.GeoCoordinateSystem, err = buildGeoCoordinateSystemFromNode(geogcsNode)
	if err != nil {
		return
	}

	if err = assertNode("projection", projectionNode, NodeProjection, 0); err != nil {
		return
	}

	cs.Projection = projectionNode.Name

	for i, paramNode := range parameterNodes {
		if err = assertNode("coordinate system", paramNode, NodeParameter, 1); err != nil {
			return
		}

		paramName := paramNames[i]
		if paramNode.Name != paramName {
			return cs, invalidNodeNameErr("param", paramNode, paramName)
		}

		value := paramNode.Children[0].LiteralValue
		switch i {
		case 0:
			cs.FalseEasting = value
		case 1:
			cs.FalseNorthing = value
		case 2:
			cs.CentralMeridian = value
		case 3:
			cs.ScaleFactor = value
		case 4:
			cs.LatitudeOfOrigin = value
		}
	}

	cs.Unit, err = buildUnitFromNode(unitNode)

	return
}

func buildGeoCoordinateSystemFromNode(node *Node) (gcs GeoCoordinateSystem, err error) {
	if err = assertNode("geogcs", node, NodeGeogcs, 3); err != nil {
		return
	}

	gcs.Name = node.Name
	var (
		datumNode  = node.Children[0]
		primemNode = node.Children[1]
		unitNode   = node.Children[2]
	)

	gcs.Datum, err = buildDatumFromNode(datumNode)
	if err != nil {
		return
	}

	gcs.PrimeMeridian, err = buildPrimemFromNode(primemNode)
	if err != nil {
		return
	}

	gcs.Unit, err = buildUnitFromNode(unitNode)
	if err != nil {
		return
	}

	return
}

func buildDatumFromNode(node *Node) (datum Datum, err error) {
	if err = assertNode("datum", node, NodeDatum, 1); err != nil {
		return
	}

	spheroidNode := node.Children[0]
	if err = assertNode("spheroid", spheroidNode, NodeSpheroid, 2); err != nil {
		return
	}

	datum.Name = node.Name
	datum.Spheroid.Name = spheroidNode.Name
	datum.Spheroid.SemiMajorAxis = spheroidNode.Children[0].LiteralValue
	datum.Spheroid.InverseFlattening = spheroidNode.Children[1].LiteralValue

	return
}

func buildPrimemFromNode(node *Node) (primem Primem, err error) {
	if err = assertNode("datum", node, NodePrimem, 1); err != nil {
		return
	}

	primem.Name = node.Name
	primem.Value = node.Children[0].LiteralValue
	return
}

func buildUnitFromNode(node *Node) (unit Unit, err error) {
	if err = assertNode("unit", node, NodeUnit, 1); err != nil {
		return
	}

	unit.Name = node.Name
	unit.Value = node.Children[0].LiteralValue
	return
}

func assertNode(entityName string, node *Node, expectedNodeType NodeType, expectedNumberOfChildren int) error {
	if node.Type != expectedNodeType {
		return fmt.Errorf("invalid node type for %v (%v)", entityName, node.Type)
	}

	if len(node.Children) != expectedNumberOfChildren {
		return fmt.Errorf("invalid number of children nodes for %v (%v)", entityName, len(node.Children))
	}

	return nil
}

func invalidNodeNameErr(entityName string, node *Node, expectedName string) error {
	return fmt.Errorf("invalid node name for %v, got %v expected %v", entityName, node.Name, expectedName)
}

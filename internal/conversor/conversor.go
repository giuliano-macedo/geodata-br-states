package conversor

import (
	"fmt"

	"github.com/giuliano-oliveira/geodata-br-states/internal/prj"
)

// Converts Projection coordinates from a projection to geopgrahic coordinates
// given a different spheroid
type Conversor struct {
	Name   string
	srcPc  *ProjectionConstants
	destSc *SphereConstants
}

func NewConversor(srcCoordinateSystem *prj.CoordinateSystem, destSpheroid prj.Spheroid) (c Conversor) {
	return Conversor{
		Name:   fmt.Sprintf("%v -> %v", srcCoordinateSystem.Name, destSpheroid.Name),
		srcPc:  ComputeProjectionConstants(srcCoordinateSystem),
		destSc: ComputeSphereConstants(destSpheroid),
	}
}

func (c Conversor) ProjectedToGeographic(E, N float64) (lat, lng float64) {
	p, l := c.srcPc.ProjectedToGeographic(E, N)

	x, y, z := c.srcPc.sc.GeographicToCartesian(p, l, 0)

	// TODO: Implement EPSG1031 Geocentric translations (3-parameter) by x+=dx; y+=dy; z+=dz, each projection conversion pair has its own dx,dy,dz

	p, l = c.destSc.CartesianToGeographic2d(x, y, z)

	return RadToDeg(p), RadToDeg(l)
}

package conversor

import (
	"math"

	"github.com/giuliano-oliveira/geodata-br-states/internal/prj"
)

type SphereConstants struct {
	Name string

	a  float64 // semi-major axis
	b  float64 // semi-minor axis
	e  float64 // eccentricity
	e2 float64 // e²
	f  float64 // flattening
	// nu  float64
	eta float64
}

func ComputeSphereConstants(sphere prj.Spheroid) *SphereConstants {
	var (
		a = sphere.SemiMajorAxis
		f = 1 / sphere.InverseFlattening
		b = a * (1 - f)
		// e2 = ((a * a) - (b * b)) / (a * a)
		e2 = ((2 * f) - (f * f))
	)
	return &SphereConstants{
		Name: sphere.Name,

		a:   a,
		f:   f,
		b:   b,
		e:   math.Sqrt(e2),
		e2:  e2,
		eta: e2 / (1 - e2),
	}
}

// p,l,h must be in rad
func (sc *SphereConstants) GeographicToCartesian(p, l, h float64) (x, y, z float64) {
	// reference: Geomatics Guidance Note number 7, part 2 - 4.1.1 (ESPG9602)
	var (
		sinP, cosP = math.Sincos(p)
		sinl, cosl = math.Sincos(l)
		sinP2      = sinP * sinP
		nu         = sc.a / math.Sqrt(1-(sc.e2*sinP2))
	)

	x = (nu + h) * cosP * cosl
	y = (nu + h) * cosP * sinl
	z = ((1 - sc.e2) * (nu + h)) * sinP

	return
}

// resulting p,l will be in rad
func (sc *SphereConstants) CartesianToGeographic2d(x, y, z float64) (p, l float64) {
	// reference: Geomatics Guidance Note number 7, part 2 - 4.1.1 (ESPG9602)
	var (
		pp = math.Sqrt((x * x) + (y * y))
		q  = math.Atan((z * sc.a) / (pp * sc.b))

		sinq, cosq = math.Sincos(q)

		sinq3 = sinq * sinq * sinq
		cosq3 = cosq * cosq * cosq
	)

	p = math.Atan2((z + (sc.eta * sc.b * sinq3)), (pp - (sc.e2 * sc.a * cosq3)))
	l = math.Atan2(y, x)
	return
}

// resulting p,l,h will be in rad
func (sc *SphereConstants) CartesianToGeographic(x, y, z float64) (p, l, h float64) {
	// reference: Geomatics Guidance Note number 7, part 2 - 4.1.1 (ESPG9602)
	var (
		pp         = math.Sqrt((x * x) + (y * y))
		q          = math.Atan((z * sc.a) / (pp * sc.b))
		sinq, cosq = math.Sincos(q)

		sinq2 = sinq * sinq
		sinq3 = sinq2 * sinq
		cosq3 = cosq * cosq * cosq
	)

	p = math.Atan2((z + (sc.eta * sc.b * sinq3)), (pp - (sc.e2 * sc.a * cosq3)))
	l = math.Atan2(y, x)
	nu := sc.a / math.Sqrt(1-(sc.e2*sinq2))
	h = (pp / math.Cos(p)) - nu
	return
}

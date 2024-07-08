package conversor

import (
	"math"

	"github.com/giuliano-oliveira/geodata-br-states/internal/prj"
)

// JHS reverse formula constants
type ProjectionConstants struct {
	Name string

	sc SphereConstants

	lambda_0 float64 //this is calculated from (Z−1)*W−180+3, where Z is the zone and W=6 is the zone width
	fe       float64
	bigB     float64
	fn       float64
	k_0      float64
	m_0      float64
	h_1      float64
	h_2      float64
	h_3      float64
	h_4      float64
}

func ComputeProjectionConstants(cs *prj.CoordinateSystem) *ProjectionConstants {
	sc := ComputeSphereConstants(cs.GeoCoordinateSystem.Datum.Spheroid)
	var (
		phi_0 = cs.LatitudeOfOrigin
		n     = sc.f / (2 - sc.f)
		n2    = n * n
		n3    = n2 * n
		n4    = n2 * n2
		bigB  = (sc.a / (1 + n)) * (1 + (n2 / 4) + (n4 / 64.0))

		h_1 = n/2.0 - ((2.0 / 3.0) * n2) + ((37.0 / 96.0) * n3) - ((1.0 / 360.0) * n4)
		h_2 = ((1.0 / 48.0) * n2) + ((1.0 / 15.0) * n3) - ((437.0 / 1440.0) * n4)
		h_3 = ((17.0 / 480.0) * n3) - ((37.0 / 840.0) * n4)
		h_4 = (4397.0 / 161280.0) * n4
	)

	return &ProjectionConstants{
		Name:     cs.Name,
		sc:       *sc,
		lambda_0: DegToRad(cs.CentralMeridian),
		fe:       cs.FalseEasting,
		fn:       cs.FalseNorthing,
		k_0:      cs.ScaleFactor,
		bigB:     bigB,
		m_0:      computeM0(phi_0, bigB, sc.e, n),
		h_1:      h_1,
		h_2:      h_2,
		h_3:      h_3,
		h_4:      h_4,
	}
}

// resulting p,l will be in rad
func (pc *ProjectionConstants) ProjectedToGeographic(E, N float64) (p, l float64) {
	// reference: Geomatics Guidance Note number 7, part 2 - 3.5.3.1 (EPSG9808)
	// Using JHS reverse formulas

	var (
		nu    = (E - pc.fe) / (pc.bigB * pc.k_0)
		gamma = ((N - pc.fn) + (pc.k_0 * pc.m_0)) / (pc.bigB * pc.k_0)

		gamma_1 = pc.h_1 * math.Sin(2*gamma) * math.Cosh(2*nu)
		gamma_2 = pc.h_2 * math.Sin(4*gamma) * math.Cosh(4*nu)
		gamma_3 = pc.h_3 * math.Sin(6*gamma) * math.Cosh(6*nu)
		gamma_4 = pc.h_4 * math.Sin(8*gamma) * math.Cosh(8*nu)
		gamma_0 = (gamma) - (gamma_1 + gamma_2 + gamma_3 + gamma_4)

		nu_1 = pc.h_1 * math.Cos(2*gamma) * math.Sinh(2*nu)
		nu_2 = pc.h_2 * math.Cos(4*gamma) * math.Sinh(4*nu)
		nu_3 = pc.h_3 * math.Cos(6*gamma) * math.Sinh(6*nu)
		nu_4 = pc.h_4 * math.Cos(8*gamma) * math.Sinh(8*nu)
		nu_0 = (nu) - (nu_1 + nu_2 + nu_3 + nu_4)

		beta          = math.Asin(math.Sin(gamma_0) / math.Cosh(nu_0))
		q_prime_prime = computeQ(beta, pc.sc.e)
	)

	p = math.Atan(math.Sinh(q_prime_prime))
	l = pc.lambda_0 + (math.Asin(math.Tanh(nu_0) / math.Cos(beta)))

	return
}

func computeQ(beta, e float64) float64 {
	Q := math.Asinh(math.Tan(beta))
	Qnext := Q
	er := 1e-13
	for i := 0; i < 1000; i++ {
		Qprev := Qnext
		Qnext = Q + (e * math.Atanh(e*math.Tanh(Qnext)))
		currErr := math.Abs(Qprev - Qnext)
		if currErr < er || math.IsNaN(Qnext) {
			break
		}
	}
	return Qnext
}

// phi_0 is in deg
func computeM0(phi_0, B, e, n float64) (m0 float64) {
	switch phi_0 {
	case 0:
		return 0
	case 90:
		return B * (math.Pi / 2)
	case -90:
		return B * (-math.Pi / 2)
	}
	phi_0 = DegToRad(phi_0)

	var (
		n2 = n * n
		n3 = n2 * n
		n4 = n2 * n2

		h1 = n/2.0 - (2.0/3.0)*n2 + (5.0/16.0)*n3 + (41.0/180.0)*n4
		h2 = (13.0/48.0)*n2 - (3.0/5.0)*n3 + (557.0/1440.0)*n4
		h3 = (61.0/240.0)*n3 - (103.0/140.0)*n4
		h4 = (49561.0 / 161280.0) * n4

		Q_0     = math.Asinh(math.Tan(phi_0)) - (e * math.Atanh(e*math.Sin(phi_0)))
		beta_0  = math.Atan(math.Sinh(Q_0))
		gamma_0 = math.Asin(math.Sin(beta_0))

		gamma_1 = h1 * math.Sin(2*gamma_0)
		gamma_2 = h2 * math.Sin(4*gamma_0)
		gamma_3 = h3 * math.Sin(6*gamma_0)
		gamma_4 = h4 * math.Sin(8*gamma_0)
		gamma   = (gamma_0 + gamma_1 + gamma_2 + gamma_3 + gamma_4)
	)

	m0 = B * gamma
	return
}

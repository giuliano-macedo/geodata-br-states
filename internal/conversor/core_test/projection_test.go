package conversor_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/giuliano-oliveira/geodata-br-states/internal/conversor"
	"github.com/giuliano-oliveira/geodata-br-states/internal/prj"
)

func TestProjection(t *testing.T) {
	const tolerance = 6e-6

	britain := conversor.ComputeProjectionConstants(&prj.BritainCs)
	australia := conversor.ComputeProjectionConstants(&prj.AustraliaCs)
	testCases := []struct {
		pc         *conversor.ProjectionConstants
		l, p, e, n float64
	}{
		{britain, 50.5, 0.5, 577274.99, 69740.50},
		{australia, 80, 146, 596813.055, 18885748.708},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%v-%v", tc.pc.Name, i), func(t *testing.T) {
			l, p := tc.pc.ProjectedToGeographic(tc.e, tc.n)

			l, p = conversor.RadToDeg(l), conversor.RadToDeg(p)
			var (
				errX = math.Abs(l - tc.l)
				errY = math.Abs(p - tc.p)
			)
			if (errX > tolerance || errY > tolerance) || (math.IsNaN(l) || math.IsNaN(p)) {
				t.Fatalf("expected (%v, %v) got (%v, %v). errs: (%v, %v)", tc.l, tc.p, l, p, errX, errY)
			}
		})
	}
}

package conversor_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/giuliano-oliveira/geodata-br-states/internal/conversor"
	"github.com/giuliano-oliveira/geodata-br-states/internal/prj"
)

func TestConversor(t *testing.T) {
	const tolerance = 6e-6

	brConversor := conversor.NewConversor(&prj.BrazilCs, prj.WgsSphere)
	testCases := []struct {
		c              conversor.Conversor
		e, n, lat, lng float64
	}{
		{brConversor, -1254473.4447543642, 8873047.5070001, -9.814462179359298, -66.8061931554248},
		{brConversor, 1300294.629115, 7440500.082427, -22.9519075, -43.2104462},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%v-%v", tc.c.Name, i), func(t *testing.T) {
			lat, lng := tc.c.ProjectedToGeographic(tc.e, tc.n)

			var (
				errX = math.Abs(lat - tc.lat)
				errY = math.Abs(lng - tc.lng)
			)
			if (errX > tolerance || errY > tolerance) || (math.IsNaN(lat) || math.IsNaN(lng)) {
				t.Fatalf("expected (%v, %v) got (%v, %v). errs: (%v, %v)", tc.lat, tc.lng, lat, lng, errX, errY)
			}
		})
	}
}

func BenchmarkConversor(b *testing.B) {
	brConversor := conversor.NewConversor(&prj.BrazilCs, prj.WgsSphere)

	rng := rand.New(rand.NewSource(42))

	baseE, baseN := -1254473.4447543642, 8873047.5070001
	variation := 0.5 / 100

	for i := 0; i < b.N; i++ {
		var (
			e = rng.NormFloat64()*variation*baseE + baseE
			n = rng.NormFloat64()*variation*baseN + baseN
		)

		brConversor.ProjectedToGeographic(e, n)
	}
}

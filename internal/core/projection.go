package core

import (
	"runtime"
	"sync"

	"github.com/giuliano-oliveira/geodata-br-states/internal/conversor"
	"github.com/giuliano-oliveira/geodata-br-states/internal/prj"
	"github.com/giuliano-oliveira/geodata-br-states/internal/shp"
)

func convertToWgs(cs prj.CoordinateSystem, shape *shp.ShapeFile) {
	var wg sync.WaitGroup
	numberOfWorkers := runtime.NumCPU()
	conversor := conversor.NewConversor(&cs, prj.WgsSphere)

	runWorker := func(id int) {
		for i := range shape.Records {
			var (
				feature = &(shape.Records[i])

				n       = len(feature.Polygon.Points)
				workerN = n / numberOfWorkers
				start   = id * workerN
				end     = start + workerN
			)
			if id == numberOfWorkers-1 {
				end = n
			}

			for j := start; j < end; j++ {
				point := &(feature.Polygon.Points[j])

				point.X, point.Y = conversor.ProjectedToGeographic(point.X, point.Y)
			}
		}
		wg.Done()
	}

	wg.Add(numberOfWorkers)

	for id := 0; id < numberOfWorkers; id++ {
		go runWorker(id)
	}

	wg.Wait()

}

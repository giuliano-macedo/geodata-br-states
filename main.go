package main

import (
	"flag"

	"github.com/giuliano-oliveira/geodata-br-states/internal/core"
)

var (
	outputDir  = flag.String("output-dir", "geojson", "directory to save the GeoJson files")
	uri        = flag.String("uri", "https://geonode.paranagua.pr.gov.br/download/167", "url to download the shp zip file")
	prettyJson = flag.Bool("pretty-json", false, "indent output json files")
)

func main() {
	flag.Parse()

	if err := core.Main(*outputDir, *uri, *prettyJson); err != nil {
		panic(err)
	}
}

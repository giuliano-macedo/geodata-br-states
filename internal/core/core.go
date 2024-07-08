package core

import (
	"fmt"
	"time"
)

func Main(outputDir, uri string, prettyJson bool) error {
	fmt.Println("downloading zip...")
	start := time.Now()
	zipFile, zipFileSize, err := getZip(uri)
	if err != nil {
		return err
	}

	fmt.Println("downloaded zip in", time.Since(start))

	cs, db, shape, err := getShpFromZip(zipFile, zipFileSize)
	if err != nil {
		return err
	}

	start = time.Now()
	convertToWgs(cs, &shape)
	fmt.Println("converted to Wgs in", time.Since(start))

	all, states := convertToGeoJson(cs, shape, db)

	return saveGeogjsons(outputDir, prettyJson, all, states)
}

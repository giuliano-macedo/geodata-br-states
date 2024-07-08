package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

func writeJsonFile(fname string, value any, pretty bool) (err error) {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	if pretty {
		encoder.SetIndent("", strings.Repeat(" ", 4))
	}

	return encoder.Encode(value)
}

func saveGeogjsons(outputDir string, pretty bool, all FeatureCollection, states []FeatureCollection) (err error) {
	if err = os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return
	}

	allStatesFileOut := path.Join(outputDir, "br_states.json")
	err = writeJsonFile(allStatesFileOut, all, pretty)
	if err != nil {
		return
	}
	fmt.Printf("All states: [%v](%v)\n\n", allStatesFileOut, allStatesFileOut)

	individualBaseDir := path.Join(outputDir, "br_states")

	if err = os.MkdirAll(individualBaseDir, os.ModePerm); err != nil {
		return
	}

	fmt.Println("Individual states:")
	for _, feature := range states {
		props := feature.Features[0].Properties

		fname := fmt.Sprintf("br_%v.json", strings.ToLower(props.Sigla))
		fileOut := path.Join(individualBaseDir, fname)

		err = writeJsonFile(fileOut, feature, pretty)
		fmt.Printf("* %v: [%v](%v)\n", props.Estado, fileOut, fileOut)
		if err != nil {
			return err
		}
	}

	return
}

package core

import (
	"archive/zip"
	"fmt"
	"io"
	"path"
	"strings"

	"github.com/giuliano-oliveira/geodata-br-states/internal/dbf"
	"github.com/giuliano-oliveira/geodata-br-states/internal/prj"
	"github.com/giuliano-oliveira/geodata-br-states/internal/shp"
)

func readAndParseFromZip[T any](z *zip.Reader, basename, ext string, parser func(reader io.Reader) (T, error)) (ans T, err error) {
	f, err := z.Open(basename + ext)
	if err != nil {
		return
	}

	defer f.Close()
	ans, err = parser(f)

	return
}

func getBaseName(file []*zip.File) (baseName string, err error) {
	for _, f := range file {

		switch ext := strings.ToLower(path.Ext(f.Name)); ext {
		case ".shp", ".prj", ".dbf":
			baseName = strings.TrimSuffix(f.Name, ext)
			return
		}
	}
	err = fmt.Errorf("didn't find any useful file in zip")
	return
}

func getShpFromZip(zipFile io.ReaderAt, sipFileSize int64) (cs prj.CoordinateSystem, db dbf.Dbf, shape shp.ShapeFile, err error) {
	var (
		z        *zip.Reader
		baseName string
	)

	if z, err = zip.NewReader(zipFile, sipFileSize); err != nil {
		return
	}

	if baseName, err = getBaseName(z.File); err != nil {
		return
	}

	if cs, err = readAndParseFromZip(z, baseName, ".prj", prj.ReadCoordinateSystem); err != nil {
		return
	}

	if db, err = readAndParseFromZip(z, baseName, ".dbf", dbf.ReadDbf); err != nil {
		return
	}

	if shape, err = readAndParseFromZip(z, baseName, ".shp", shp.ReadShp); err != nil {
		return
	}

	return
}

// reference: http://www.independent-software.com/dbase-dbf-dbt-file-format.html
package dbf

import (
	"encoding/binary"
	"fmt"
	"io"
)

func ReadHeader(reader io.Reader) (header Header, descriptors []FieldDescriptor, err error) {
	if err = binary.Read(reader, binary.LittleEndian, &header); err != nil {
		return
	}

	if header.Version != 0x3 {
		err = fmt.Errorf("invalid version (%X)", header.Version)
		return
	}

	descriptors = make([]FieldDescriptor, header.NumberOfFields())
	if err = binary.Read(reader, binary.LittleEndian, &descriptors); err != nil {
		return
	}

	headerTerminator := []byte{0}
	if _, err = reader.Read(headerTerminator); err != nil {
		return
	}

	if headerTerminator[0] != 0x0d {
		err = fmt.Errorf("invalid header terminator (%X)", headerTerminator[0])
		return
	}
	return

}

func ReadDbf(reader io.Reader) (dbf Dbf, err error) {
	header, descriptors, err := ReadHeader(reader)
	if err != nil {
		return
	}

	decoder, err := NewDecoder(header, descriptors)
	if err != nil {
		return
	}

	dbf.Header = header
	dbf.Records, err = decoder.Decode(reader)

	return
}

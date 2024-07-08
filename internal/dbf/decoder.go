package dbf

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	IdFieldFidExport int = iota
	IdFieldSigla
	IdFieldTotal
	IdFieldHomens
	IdFieldMulheres
	IdFieldUrbana
	IdFieldRural
	IdFieldTxAlfab
	IdFieldFidEstado
	IdFieldEstado
	IdFieldFkMacro
	IdFieldPkSigla
)

var decoderConfig = []struct {
	fieldName string
	fieldType FieldType
}{
	IdFieldFidExport: {"FID_Export", Numeric},
	IdFieldSigla:     {"SIGLA", Character},
	IdFieldTotal:     {"Total", Float},
	IdFieldHomens:    {"Homens", Float},
	IdFieldMulheres:  {"Mulheres", Float},
	IdFieldUrbana:    {"Urbana", Float},
	IdFieldRural:     {"Rural", Float},
	IdFieldTxAlfab:   {"TX_Alfab", Float},
	IdFieldFidEstado: {"FID_estado", Numeric},
	IdFieldEstado:    {"Estado", Character},
	IdFieldFkMacro:   {"FK_macro", Character},
	IdFieldPkSigla:   {"PK_sigla", Character},
}

type Decoder struct {
	header      Header
	descriptors []FieldDescriptor
}

func NewDecoder(header Header, descriptors []FieldDescriptor) (*Decoder, error) {
	for i, descriptor := range descriptors {
		config := decoderConfig[i]

		if config.fieldName != descriptor.Name.String() || config.fieldType != descriptor.Type {
			return nil, fmt.Errorf("invalid schema for field %v %v", descriptor.Name.String(), descriptor.Type)
		}
	}

	return &Decoder{
		header:      header,
		descriptors: descriptors,
	}, nil
}

func (decoder *Decoder) Decode(reader io.Reader) (rows []Record, err error) {
	recordRaw := make([]byte, decoder.header.RecordLength)

	for {
		var n int
		n, err = reader.Read(recordRaw)

		if n == 1 && recordRaw[0] == 0x1a {
			err = io.EOF
		}

		if err == io.EOF {
			err = nil
			break
		}

		if n != len(recordRaw) {
			err = fmt.Errorf("couldn't read record %v", n)
			break
		}

		if err != nil {
			return
		}

		if recordRaw[0] == '*' {
			continue
		}

		var record Record
		record, err = decoder.DecodeRecord(recordRaw[1:])
		if err != nil {
			return
		}

		rows = append(rows, record)
	}
	return
}

func (decoder *Decoder) DecodeRecord(recordRaw []byte) (record Record, err error) {
	offset := 0
	for descriptorId, descriptor := range decoder.descriptors {
		fieldData := recordRaw[offset : offset+int(descriptor.Length)]

		if err = decodeField(descriptorId, fieldData, &record); err != nil {
			return
		}

		offset += int(descriptor.Length)
	}
	return
}

func decodeField(descriptorId int, fieldData []byte, record *Record) (err error) {
	num := func() float64 { return decodeNumericField(fieldData) }
	chr := func() string { return decodeCharacterField(fieldData) }
	switch descriptorId {
	case IdFieldFidExport:
		record.Fid = int(num())
	case IdFieldSigla:
		record.Sigla = chr()
	case IdFieldTotal:
		record.Total = int(num())
	case IdFieldHomens:
		record.Homens = int(num())
	case IdFieldMulheres:
		record.Mulheres = int(num())
	case IdFieldUrbana:
		record.Urbana = int(num())
	case IdFieldRural:
		record.Rural = int(num())
	case IdFieldTxAlfab:
		record.TxAlfab = num()
	case IdFieldFidEstado:
		record.FidEstado = int(num())
	case IdFieldEstado:
		record.Estado = chr()
	case IdFieldFkMacro:
		record.FkMacro = chr()
	case IdFieldPkSigla:
		record.PkSigla = chr()
	default:
		err = fmt.Errorf("field yet not supported %v", descriptorId)
	}
	return

}

func decodeCharacterField(fieldData []byte) string {
	return strings.TrimSpace(string(fieldData))
}

func decodeNumericField(fieldData []byte) float64 {
	data := strings.TrimLeft(string(fieldData), " ")
	ans, _ := strconv.ParseFloat(string(data), 64)
	return ans
}

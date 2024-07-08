package dbf

import (
	"strings"
	"time"
)

type DbfDate [3]uint8

func (d DbfDate) Date() time.Time {
	return time.Date(1900+int(d[0]), time.Month(d[1]), int(d[2]), 0, 0, 0, 0, time.UTC)
}

const (
	HeaderSize          = 32
	FieldDescriptorSize = 32
)

type Header struct {
	Version      byte    // Version byte
	LastUpdated  DbfDate // Date of last update in YYMMDD format (where YY is equal to year minus 1900)
	NumRecords   uint32  // Number of records in table
	HeaderLength uint16  // Number of bytes in the header
	RecordLength uint16  // Number of bytes in a record
	_            [20]byte
}

func (header *Header) NumberOfFields() int {
	return int(header.HeaderLength-HeaderSize) / FieldDescriptorSize
}

type FieldType byte

const (
	Character FieldType = 'C' // A string of characters, padded with spaces if shorter than the field length
	Date      FieldType = 'D' // Date stored as string in YYYYMMDD format
	Float     FieldType = 'F' // Floating point number, stored as string, padded with spaces if shorter than the field length
	Numeric   FieldType = 'N' // Floating point number, stored as string, padded with spaces if shorter than the field length
	Logical   FieldType = 'L' // A boolean value, stored as one of YyNnTtFf. May be set to ? if not initialized
)

func (ft FieldType) String() string {
	return string(ft)
}

const MaxFieldNameLength = 11

type FieldNameType [MaxFieldNameLength]byte

type FieldDescriptor struct {
	Name         FieldNameType // Field name (padded with NULL-bytes)
	Type         FieldType     // Field type
	Address      uint32        // Field data address in memory
	Length       byte          // Field length
	DecimalCount byte          //  Field decimal count
	_            [2]byte
	WorkAreaId   byte // Work area ID
	_            [2]byte
	SetFieldFlag byte //SET FIELDS flag
	_            [8]byte
}

func (fn FieldNameType) String() string {
	var sb strings.Builder
	for _, c := range fn {
		if c == 0 {
			break
		}
		sb.WriteByte(c)
	}
	return sb.String()
}

type Record struct {
	Fid       int     `json:"FID_Export"`
	Sigla     string  `json:"SIGLA"`
	Total     int     `json:"Total"`
	Homens    int     `json:"Homens"`
	Mulheres  int     `json:"Mulheres"`
	Urbana    int     `json:"Urbana"`
	Rural     int     `json:"Rural"`
	TxAlfab   float64 `json:"TX_Alfab"`
	FidEstado int     `json:"FID_estado"`
	Estado    string  `json:"Estado"`
	FkMacro   string  `json:"FK_macro"`
	PkSigla   string  `json:"PK_sigla"`
}

type Dbf struct {
	Header  Header
	Records []Record
}

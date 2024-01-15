package vlc

import (
	"archiver/lib/compression/types/table"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"strings"
)

type EncoderDecoder struct {
	tableGenerator table.Generator
}

func New(tblGenerator table.Generator) EncoderDecoder {
	return EncoderDecoder{
		tableGenerator: tblGenerator,
	}
}

func (end EncoderDecoder) Encode(str string) []byte {

	if len(str) == 0 {
		log.Panic("Empty file to encode\t")
	}

	tbl := end.tableGenerator.NewTable(str)

	encodeData := encodeString(str, tbl)

	return buildEncodedFile(tbl, encodeData)

}

func (end EncoderDecoder) Decode(enData []byte) string {
	tbl, data := parseFile(enData)

	return tbl.DecodeForTbl(data)
}

func parseFile(data []byte) (table.EncodingTable, string) {
	const (
		tableSizeBytesCount = 4
		textSizeBytesCount  = 4
	)
	tableSizeBinary, data := data[:tableSizeBytesCount], data[tableSizeBytesCount:]
	textSizeBinary, data := data[:textSizeBytesCount], data[textSizeBytesCount:]

	tableSize := binary.BigEndian.Uint32(tableSizeBinary)
	dataSize := binary.BigEndian.Uint32(textSizeBinary)

	tblBinary, data := data[:tableSize], data[tableSize:]

	table := fromByteTable(tblBinary)

	body := NewBinChunks(data).Join()

	return table, body[:dataSize]
}

func buildEncodedFile(table table.EncodingTable, data string) []byte {
	encodedTable := toByteTable(table)

	var buf bytes.Buffer

	buf.Write(toByteInt(len(encodedTable)))
	buf.Write(toByteInt(len(data)))
	buf.Write(encodedTable)
	buf.Write(splitByChunk(data, chunkSize).ToBytes())

	return buf.Bytes()
}

func toByteInt(num int) []byte {
	res := make([]byte, 4)

	binary.BigEndian.PutUint32(res, uint32(num))

	return res
}

func toByteTable(table table.EncodingTable) []byte {
	var tableBuf bytes.Buffer

	if err := gob.NewEncoder(&tableBuf).Encode(table); err != nil {
		log.Fatal("cant serialize  table", err)
	}

	return tableBuf.Bytes()
}

func fromByteTable(tblBinary []byte) table.EncodingTable {
	var tbl table.EncodingTable

	r := bytes.NewReader(tblBinary)

	if err := gob.NewDecoder(r).Decode(&tbl); err != nil {
		log.Fatal("cant serialize table", err)
	}

	return tbl

}

func encodeString(str string, enTbl table.EncodingTable) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch, enTbl))
	}

	return buf.String()
}

func bin(ch rune, table table.EncodingTable) string {
	res, ok := table[ch]
	if !ok {
		panic("unknown symbol " + string(ch))
	}

	return res
}

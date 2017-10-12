package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Zip struct {
	Code  string
	City  string
	State string
}

type ZipSlice []*Zip

type ZipIndex map[string]ZipSlice

func LoadZips(filePath string) (ZipSlice, error) {
	fileInputStream, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file. %v", err)
	}

	reader := csv.NewReader(fileInputStream)
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading header row. %v", err)
	}
	// make lets us pre-allocate the memory we may need. Reallocations take time.
	// ZipSlice is a list of pointers t zips.
	iZipSlice := make(ZipSlice, 0, 43000)
	lineNum := 0
	for {
		lineNum++
		fields, err := reader.Read() // One slice per line. Splits it on ,
		if err == io.EOF {
			return iZipSlice, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error reading record on line: %d", lineNum)
		}
		iZip := &Zip{
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}
		iZipSlice = append(iZipSlice, iZip)
	}
}

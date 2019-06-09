package cleaner

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestCleanFile(t *testing.T) {
	var fc FileCleaner
	var c CsvFileCleaner
	fc = &c
	f := fc.CleanFile("./test1/test11.csv")
	fmt.Println("File in test: ", string(f))
	r := csv.NewReader(strings.NewReader(string(f)))
	records, err := r.ReadAll()
	if err != nil {
		log.Println("csv error: ", err)
	}
	fmt.Println("records", records)
	if err != nil {
		t.Fail()
	}
}

func TestCleanFile2(t *testing.T) {
	var fc FileCleaner
	var c CsvFileCleaner
	fc = &c
	f := fc.CleanFile("./test1/test11222.csv")
	fmt.Println("File in bad test: ", string(f))
	if len(f) > 0 {
		t.Fail()
	}
}

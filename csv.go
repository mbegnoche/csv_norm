package main

import (
	"fmt"
	"time"
	"encoding/csv"
	"unicode/utf8"
	"strings"
	"bufio"
	"io"
	"os"
	"log"
	"flag"
)

type CSVRecords [][]string
type CSVRecord []string

func main() {
	var inputFile string
	var outputFile string
	flag.StringVar(&inputFile, "input", "", "Input Filename")
	flag.StringVar(&outputFile, "output", "", "Output Filename")
	flag.Parse()

	var reader io.Reader = bufio.NewReader(os.Stdin)
	var writer io.Writer = bufio.NewWriter(os.Stdout)

	if inputFile != "" {
		f, err := os.Open(inputFile)
		if err == nil {
			reader = f
			defer f.Close()
		}
	}

	if outputFile != "" {
		f, err := os.Create(outputFile)
		if err == nil {
			writer = f
			defer f.Close()
		}
	}


	records := get_records(reader)
	norm_records(&records)
	write_records(writer, records)
}

func write_records(w io.Writer, records CSVRecords) {
	csvWriter := csv.NewWriter(w)
	csvWriter.WriteAll(records)
}


func validate_record(record CSVRecord) bool {
	valid := true
	for _, f := range record {
		if !utf8.ValidString(f) {
			valid = false
			break
		}
	}
	return valid
}

func get_records(r io.Reader) CSVRecords {
	ioReader := bufio.NewReader(r)
    csvReader := csv.NewReader(ioReader)
    records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Validate the records
	validRecs := records[:0]
	for _, rec := range records {
	    if validate_record(rec) {
	         validRecs = append(validRecs, rec)  
	    }
	}

	return validRecs
}

func norm_records(records *CSVRecords) {
	headers := (*records)[0]
	
	for i, rec := range *records {
	    if i == 0 { continue } // skip header
	    total := 0.0
	    for j, f := range rec {
	        switch strings.ToLower(headers[j]) {
	            case "timestamp":
	                t, _ := time.Parse("1/2/06 3:4:05 PM", f)
	                (*records)[i][j] = t.Add(time.Hour*3).Format(time.RFC3339)
	                
	            case "zip":
	                (*records)[i][j] = fmt.Sprintf("%05s", f)
	                
	            case "fullname":
	                (*records)[i][j] = strings.ToUpper(f)
	                
	            case "fooduration", "barduration":
					var h, m int
					s := 0.0
                    _, err := fmt.Sscanf(f, "%d:%d:%f", &h, &m, &s)
                    if err == nil {
						s += float64(h*3600 + m*60)
					}
					total += s
					(*records)[i][j] = fmt.Sprintf("%f", s)

				case "totalduration":
					(*records)[i][j] = fmt.Sprintf("%f", total)
	        }
	    }
	}
}
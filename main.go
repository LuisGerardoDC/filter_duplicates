package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	records, err := loadRecords("Code Assessment - Find Duplicates Input - SampleCodecsv (10).csv")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	duplicates, err := FindDuplicates(records)

	if err != nil {
		fmt.Println("Error finding duplicates:", err)
		return
	}

	fmt.Println("Duplicates found:", duplicates)

}

func FindDuplicates(records []Record) ([]Response, error) {
	response := []Response{}

	for i := 0; i < len(records); i++ {
		record := records[i]

		for j := i + 1; j < len(records); j++ {
			compareRecord := records[j]
			matchPercentage := HasMatch(record, compareRecord)
			if matchPercentage > 20 {
				records = append(records[:j], records[j+1:]...)
				response = append(response, Response{
					ContactIDSource: record.ContactID,
					ContactIDMatch:  compareRecord.ContactID,
					Acurracy:        matchPercentage,
				})
			}
		}
	}

	return response, nil
}

func HasMatch(record1, record2 Record) int {
	result := 0

	result += PercentageMatch(record1.Name, record2.Name)
	result += PercentageMatch(record1.Name1, record2.Name1)
	result += PercentageMatch(record1.Email, record2.Email)
	result += PercentageMatch(record1.PostalZip, record2.PostalZip)
	result += PercentageMatch(record1.Address, record2.Address)

	return result / 5
}

func PercentageMatch(str1, str2 string) int {
	if strings.EqualFold(str1, str2) {
		return 100
	}

	percentageForChar := 100 / max(len(str1), len(str2))
	matchCount := 0

	for i := 0; i < min(len(str1), len(str2)); i++ {
		if strings.EqualFold(string(str1[i]), string(str2[i])) {
			matchCount++
		} else {
			break
		}
	}

	return percentageForChar * matchCount
}

func loadRecords(filename string) ([]Record, error) {
	file, err := os.Open(filename)
	Records := make([]Record, 0)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		Records = append(Records, lineToRecord(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return Records, nil
}

func lineToRecord(line string) Record {
	values := strings.Split(line, ",")
	if len(values) < 6 {
		return Record{}
	}
	return Record{
		ContactID: values[0],
		Name:      values[1],
		Name1:     values[2],
		Email:     values[3],
		PostalZip: values[4],
		Address:   values[5],
	}
}

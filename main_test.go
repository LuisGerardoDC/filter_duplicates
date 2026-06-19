package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPercentageMatch(t *testing.T) {

	tests := []struct {
		name    string
		string1 string
		string2 string
		result  int
	}{
		{name: "Exact match", string1: "JohnDoe", string2: "JohnDoe", result: 100},
		{name: "Partial match", string1: "JohnDoe", string2: "JohnD", result: 70},
		{name: "No match", string1: "JohnDoe", string2: "Fane", result: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(T *testing.T) {
			result := PercentageMatch(test.string1, test.string2)
			assert.Equal(T, result, test.result)
		})
	}
}

func TestHasMatch(t *testing.T) {

	tests := []struct {
		name    string
		record1 Record
		record2 Record
		result  int
	}{
		{
			name: "Exact match",
			record1: Record{
				ContactID: "1",
				Name:      "John",
				Name1:     "Doe",
				PostalZip: "12345",
				Email:     "john.doe@example.com",
				Address:   "123 Main St",
			},
			record2: Record{
				ContactID: "2",
				Name:      "John",
				Name1:     "Doe",
				PostalZip: "12345",
				Email:     "john.doe@example.com",
				Address:   "123 Main St",
			},
			result: 100,
		},
		{
			name: "Partial match",
			record1: Record{
				ContactID: "1",
				Name:      "Joh",
				Name1:     "D",
				PostalZip: "12345",
				Email:     "john.doe@example.com",
				Address:   "123 Main St",
			},
			record2: Record{
				ContactID: "2",
				Name:      "John",
				Name1:     "Doe",
				PostalZip: "1235",
				Email:     "john.doe@example.com",
				Address:   "123 Main St",
			},
			result: 73,
		},
		{
			name: "No match",
			record1: Record{
				ContactID: "1",
				Name:      "John",
				Name1:     "Doe",
				PostalZip: "12345",
				Email:     "john.doe@example.com",
				Address:   "123 Main St",
			},
			record2: Record{
				ContactID: "2",
				Name:      "Fane",
				Name1:     "Cane",
				PostalZip: "54321",
				Email:     "email@example.com",
				Address:   "456 Oak Ave",
			},
			result: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HasMatch(test.record1, test.record2)
			assert.Equal(t, result, test.result)
		})
	}
}

func TestFindDuplicates(t *testing.T) {

	recordsNoDuplicates := []Record{
		{ContactID: "1", Name: "Zara", Name1: "Quinn", PostalZip: "11111", Email: "z.quinn@alpha.com", Address: "1 Alpha St"},
		{ContactID: "2", Name: "Bruce", Name1: "Kimura", PostalZip: "22222", Email: "b.kimura@beta.com", Address: "2 Beta Ave"},
		{ContactID: "3", Name: "Donna", Name1: "Petrov", PostalZip: "33333", Email: "d.petrov@gamma.com", Address: "3 Gamma Rd"},
		{ContactID: "4", Name: "Felix", Name1: "Obasi", PostalZip: "44444", Email: "f.obasi@delta.com", Address: "4 Delta Blvd"},
		{ContactID: "5", Name: "Helen", Name1: "Russo", PostalZip: "55555", Email: "h.russo@epsilon.com", Address: "5 Epsilon Way"},
	}

	// R1(original), R2, R3(variacion de R1: "Carin" faltan 'a'), R4, R5, R6(variacion de R1: "Cari" faltan 'na'), R7
	recordsWithOneDuplicate := []Record{
		{ContactID: "1", Name: "Carina", Name1: "Torres", PostalZip: "39746", Email: "carina.torres@example.com", Address: "449 Tellus Rd"},
		{ContactID: "2", Name: "Bruce", Name1: "Kimura", PostalZip: "22222", Email: "b.kimura@beta.com", Address: "2 Beta Ave"},
		{ContactID: "3", Name: "Carin", Name1: "Torres", PostalZip: "39746", Email: "carina.torres@example.com", Address: "449 Tellus Rd"},
		{ContactID: "4", Name: "Donna", Name1: "Petrov", PostalZip: "33333", Email: "d.petrov@gamma.com", Address: "3 Gamma Rd"},
		{ContactID: "5", Name: "Felix", Name1: "Obasi", PostalZip: "44444", Email: "f.obasi@delta.com", Address: "4 Delta Blvd"},
		{ContactID: "6", Name: "Cari", Name1: "Torres", PostalZip: "39746", Email: "carina.torres@example.com", Address: "449 Tellus Rd"},
		{ContactID: "7", Name: "Helen", Name1: "Russo", PostalZip: "55555", Email: "h.russo@epsilon.com", Address: "5 Epsilon Way"},
	}

	// R1(original 1), R2, R3(variacion de R1), R4(original 2), R5, R6(variacion de R1), R7(variacion de R4), R8, R9(variacion de R4)
	recordsWithMultipleDuplicates := []Record{
		{ContactID: "1", Name: "Carina", Name1: "Torres", PostalZip: "39746", Email: "carina.torres@example.com", Address: "449 Tellus Rd"},
		{ContactID: "2", Name: "Bruce", Name1: "Kimura", PostalZip: "22222", Email: "b.kimura@beta.com", Address: "2 Beta Ave"},
		{ContactID: "3", Name: "Carin", Name1: "Torres", PostalZip: "39746", Email: "carina.torres@example.com", Address: "449 Tellus Rd"},
		{ContactID: "4", Name: "Marco", Name1: "Reyes", PostalZip: "67890", Email: "marco.reyes@example.com", Address: "99 Palm Dr"},
		{ContactID: "5", Name: "Felix", Name1: "Obasi", PostalZip: "44444", Email: "f.obasi@delta.com", Address: "4 Delta Blvd"},
		{ContactID: "6", Name: "Cari", Name1: "Torres", PostalZip: "39746", Email: "carina.torres@example.com", Address: "449 Tellus Rd"},
		{ContactID: "7", Name: "Marc", Name1: "Reyes", PostalZip: "67890", Email: "marco.reyes@example.com", Address: "99 Palm Dr"},
		{ContactID: "8", Name: "Helen", Name1: "Russo", PostalZip: "55555", Email: "h.russo@epsilon.com", Address: "5 Epsilon Way"},
		{ContactID: "9", Name: "Ma", Name1: "Reyes", PostalZip: "67890", Email: "marco.reyes@example.com", Address: "99 Palm Dr"},
		{ContactID: "10", Name: "Carina", Name1: "Torr", PostalZip: "39777", Email: "carina.torres@example.com", Address: "449 Tellus Rd"},
		{ContactID: "11", Name: "Bruce", Name1: "Kimura", PostalZip: "22222", Email: "b.kimura@beta.com", Address: "2 Beta Ave"},
		{ContactID: "12", Name: "Carina", Name1: "Torr", PostalZip: "39777", Email: "carina.torres@example.com", Address: "448 Tellus Rd"},
	}

	Tests := []struct {
		name    string
		records []Record
		results []Response
	}{
		{
			name:    "No duplicates",
			records: recordsNoDuplicates,
			results: []Response{},
		},
		{
			name:    "One duplicate",
			records: recordsWithOneDuplicate,
			results: []Response{
				{ContactIDSource: "1", ContactIDMatch: "3", Acurracy: 96},
				{ContactIDSource: "1", ContactIDMatch: "6", Acurracy: 92},
			},
		},
		{
			name:    "Multiple duplicates",
			records: recordsWithMultipleDuplicates,

			results: []Response{
				{ContactIDSource: "1", ContactIDMatch: "3", Acurracy: 96},
				{ContactIDSource: "1", ContactIDMatch: "6", Acurracy: 92},
				{ContactIDSource: "1", ContactIDMatch: "10", Acurracy: 84},
				{ContactIDSource: "1", ContactIDMatch: "12", Acurracy: 67},
				{ContactIDSource: "2", ContactIDMatch: "11", Acurracy: 100},
				{ContactIDSource: "4", ContactIDMatch: "7", Acurracy: 96},
				{ContactIDSource: "4", ContactIDMatch: "9", Acurracy: 88},
			},
		},
	}

	for _, test := range Tests {
		t.Run(test.name, func(t *testing.T) {
			results, err := FindDuplicates(test.records)
			assert.NoError(t, err)
			assert.Equal(t, len(test.results), len(results))
			assert.Equal(t, test.results, results)
		})
	}

}

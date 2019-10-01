package main

import (
	"bytes"
	"flag"
	"testing"
)

var sorted_test = "Fcab " +
	"Gacd " +
	"aAcd " +
	"aacd " +
	"bFad " +
	"bdGc " +
	"cdcH " +
	"caVc " +
	"cacA " +
	"cdcH " +
	"dcDa " +
	"dSdc"

var test = "aAcd " +
	"dcDa " +
	"dSdc " +
	"Fcab " +
	"Gacd " +
	"bFad " +
	"bdGc " +
	"cdcH " +
	"caVc " +
	"cacA " +
	"aacd " +
	"cdcH"

var fru_sort = "Gacd " +
	"Fcab " +
	"dcDa " +
	"dSdc " +
	"cdcH " +
	"caVc " +
	"cacA " +
	"bdGc " +
	"bFad " +
	"aAcd"

func TestSorting(t *testing.T) {
	p := new(bool)
	*p = false
	o := new(string)
	*o = "stdout"
	k := new(int)
	*k = 0
	flags = params{
		Reverse:  p,
		Numerals: p,
		Unique:   p,
		Column:   k,
		Output:   o,
		Register: p,
	}
	strs := bytes.Split([]byte(test), []byte(" "))
	result := sorting(strs, flags)
	strs = bytes.Split([]byte(sorted_test), []byte(" "))
	for i, v := range strs {
		if !bytes.Equal(result[i], v) {

			t.Error("Test failed results don't match!")
		}
	}

}

func TestSorting_fru(t *testing.T) {
	p := new(bool)
	*p = true
	o := new(string)
	*o = "stdout"
	k := new(int)
	*k = 0
	flags = params{
		Reverse:  p,
		Numerals: p,
		Unique:   p,
		Column:   k,
		Output:   o,
		Register: p,
	}
	flag.Parse()

	strs := bytes.Split([]byte(test), []byte(" "))
	sorted := sorting(strs, flags)
	strs = bytes.Split([]byte(fru_sort), []byte(" "))

	for i, _ := range strs {
		if !bytes.Equal(strs[i], sorted[i]) {
			t.Error("Test failed results don't match!")
		}
	}
}

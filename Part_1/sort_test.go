package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSorting(t *testing.T) {
	var flags = params{
		{Name: "f", Value: flag.Bool("f", false, "register")},
		{Name: "u", Value: flag.Bool("u", false, "unique")},
		{Name: "r", Value: flag.Bool("r", false, "reverse")},
		{Name: "n", Value: flag.Bool("n", false, "numerals")},
	}
	col := 0
	file, err := ioutil.ReadFile("Text/test.txt")
	if err != nil {
		fmt.Println("Couldn't open file! : ", err)
		os.Exit(1)
	}
	strs := bytes.Split(file, []byte("\n"))
	sorted := sorting(strs, flags, col)

	file, err = ioutil.ReadFile("Text/sorted_test.txt")
	if err != nil {
		fmt.Println("Couldn't open file! : ", err)
		os.Exit(1)
	}
	strs = bytes.Split(file, []byte("\n"))

	for i, _ := range strs {
		if !bytes.Equal(strs[i], sorted[i]) {
			t.Error("Test failed results don't match!")
		}
	}

}

func TestSorting_fru(t *testing.T) {
	p := new(bool)
	*p = true
	var flags = params{
		{Name: "f", Value: p},
		{Name: "u", Value: p},
		{Name: "r", Value: p},
		{Name: "n", Value: p},
	}
	col := 0

	flag.Parse()

	file, err := ioutil.ReadFile("Text/test.txt")
	if err != nil {
		fmt.Println("Couldn't open file! : ", err)
		os.Exit(1)
	}

	strs := bytes.Split(file, []byte("\n"))
	sorted := sorting(strs, flags, col)

	file, err = ioutil.ReadFile("Text/sorted_test_f,r,u.txt")
	if err != nil {
		fmt.Println("Couldn't open file! : ", err)
		os.Exit(1)
	}
	strs = bytes.Split(file, []byte("\n"))

	for i, _ := range strs {
		if !bytes.Equal(strs[i], sorted[i]) {
			t.Error("Test failed results don't match!")
		}
	}
}

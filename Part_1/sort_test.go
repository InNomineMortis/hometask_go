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
	p := new(bool)
	*p = false
	o := new(string)
	*o = "stdout"
	k:= new(int)
	*k = 0
	flags = params{
		Reverse: p,
		Numerals: p,
		Unique: p,
		Column: k,
		Output: o,
		Register: p,
	}
	file, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Couldn't open file! : ", err)
		os.Exit(1)
	}
	strs := bytes.Split(file, []byte("\n"))
	sorted := sorting(strs, flags)

	file, err = ioutil.ReadFile("sorted_test.txt")
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
	o := new(string)
	*o = "stdout"
	k:= new(int)
	*k = 0
	flags = params{
		Reverse: p,
		Numerals: p,
		Unique: p,
		Column: k,
		Output: o,
		Register: p,
	}
	flag.Parse()

	file, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Couldn't open file! : ", err)
		os.Exit(1)
	}

	strs := bytes.Split(file, []byte("\n"))
	sorted := sorting(strs, flags)

	file, err = ioutil.ReadFile("sorted_test_f,r,u.txt")
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

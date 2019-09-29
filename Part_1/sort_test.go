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
	params := []struct {
		Name  string
		Value bool
		Usage string
	}{
		{Name: "f", Value: false, Usage: "ignore letters size"},
		{Name: "u", Value: false, Usage: "show only first"},
		{Name: "r", Value: false, Usage: "from biggest to lowest"},
		{Name: "n", Value: false, Usage: "numerals sort"},
	}

	flags := make(map[string]*bool)

	for _, v := range params {
		flags[v.Name] = flag.Bool(v.Name, v.Value, v.Usage)
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

	flags := make(map[string]*bool)
	flags["u"] = flag.Bool("U", true, "unique")
	flags["f"] = flag.Bool("F", true, "register")
	flags["r"] = flag.Bool("R", true, "reverse")
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

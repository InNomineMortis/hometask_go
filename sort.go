package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func sorting(fileName string, flags map[string]*bool, col int) [][]byte {
	file, err := ioutil.ReadFile("Text/" + fileName)
	if err != nil {
		fmt.Println("Couldn't open file! : ", err)
		os.Exit(1)
	}
	order := -1
	if *flags["r"] {
		order = 1
	}

	strs := bytes.Split(file, []byte("\n"))
	if *flags["u"] {
		if *flags["f"] {
			sort.Slice(strs, func(i, j int) bool {
				return strings.Compare(strings.ToLower(string(string(strs[i])[col])), strings.ToLower(string(string(strs[j])[col]))) == order
			})
			strs1 := make([][]byte, len(strs))
			copy(strs1, strs)
			k := 0
			sort.Slice(strs1, func(i, j int) bool {
				if bytes.Compare(bytes.ToLower(strs1[i]), bytes.ToLower(strs1[j])) == 0 {
					strs = append(strs[0:i-k], strs[j+2-k:]...)
					k++
				}
				return false
			})
		} else {
			sort.Slice(strs, func(i, j int) bool {
				return strings.Compare(string(string(strs[i])[col]), string(string(strs[j])[col])) == order
			})
			strs1 := make([][]byte, len(strs))
			copy(strs1, strs)
			k := 0
			sort.Slice(strs1, func(i, j int) bool {
				if bytes.Compare(strs1[i], strs1[j]) == 0 {
					strs = append(strs[0:i-k], strs[j+2-k:]...)
					k++
				}
				return false
			})
		}
	} else {
		if *flags["f"] {
			sort.Slice(strs, func(i, j int) bool {
				return strings.Compare(strings.ToLower(string(string(strs[i])[col])), strings.ToLower(string(string(strs[j])[col]))) == order
			})
		} else {
			sort.Slice(strs, func(i, j int) bool {
				return strings.Compare(string(string(strs[i])[col]), string(string(strs[j])[col])) == order
			})
		}
	}
	return strs
}

func main() {
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

	intPtr := flag.Int("k", 0, "column")
	filePtr := flag.String("o", "stdout", "in file, otherwise in stdout")
	flag.Parse()

	outStrs := sorting(os.Args[len(os.Args)-1], flags, *intPtr)

	if *filePtr != "stdout" {
		outFile, err := os.Create("Text/" + *filePtr)
		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
		defer outFile.Close()
		for _, v := range outStrs {
			outFile.Write(v)
			outFile.WriteString("\n")
		}
	} else {
		for _, v := range outStrs {
			fmt.Println(string(v))
		}
	}

}

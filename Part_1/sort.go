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

type params = []struct {
	Name  string
	Value *bool
}

func sorting(strs [][]byte, flags params, col int) [][]byte {

	order := -1
	if *flags[2].Value {
		order = 1
	}

	if *flags[1].Value {
		if *flags[0].Value {
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
		if *flags[0].Value {
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
	var flags = params{
		{Name: "f", Value: flag.Bool("f", false, "register")},
		{Name: "u", Value: flag.Bool("u", false, "unique")},
		{Name: "r", Value: flag.Bool("r", false, "reverse")},
		{Name: "n", Value: flag.Bool("n", false, "numerals")},
	}
	intPtr := flag.Int("k", 0, "column")
	filePtr := flag.String("o", "stdout", "in file, otherwise in stdout")
	flag.Parse()
	fmt.Println(flags[2].Value, flags[2].Name)
	file, err := ioutil.ReadFile(os.Args[len(os.Args)-1])

	if err != nil {
		fmt.Println("Couldn't open file! : ", err)
		os.Exit(1)
	}

	strs := bytes.Split(file, []byte("\n"))

	outStrs := sorting(strs,flags, *intPtr)

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

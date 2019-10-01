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

type params = struct {
	Reverse  *bool
	Unique *bool
	Register *bool
	Numerals *bool
	Output *string
	Column *int
}

func sorting(strs [][]byte, flags params) [][]byte {
	order := -1
	if *flags.Reverse {
		order = 1
	}

	if *flags.Unique {
		if *flags.Register {
			sort.Slice(strs, func(i, j int) bool {
				return strings.Compare(strings.ToLower(string(string(strs[i])[*flags.Column])), strings.ToLower(string(string(strs[j])[*flags.Column]))) == order
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
				return strings.Compare(string(string(strs[i])[*flags.Column]), string(string(strs[j])[*flags.Column])) == order
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
		if *flags.Register {
			sort.Slice(strs, func(i, j int) bool {
				return strings.Compare(strings.ToLower(string(string(strs[i])[*flags.Column])), strings.ToLower(string(string(strs[j])[*flags.Column]))) == order
			})
		} else {
			sort.Slice(strs, func(i, j int) bool {
				return strings.Compare(string(string(strs[i])[*flags.Column]), string(string(strs[j])[*flags.Column])) == order
			})
		}
	}
	return strs
}

func main() {
	flags := params{
		Reverse: flag.Bool("r", false, "reverse"),
		Numerals: flag.Bool("n", false, "numerals"),
		Unique: flag.Bool("u", false, "unique"),
		Column: flag.Int("k", 0, "column"),
		Output: flag.String("o", "stdout", "outfile"),
		Register: flag.Bool("f", false, "numerals"),
	}
	flag.Parse()
	file, err := ioutil.ReadFile(os.Args[len(os.Args)-1])

	if err != nil {
		fmt.Println("Couldn't open file! : ", err)
		os.Exit(1)
	}

	strs := bytes.Split(file, []byte("\n"))

	outStrs := sorting(strs,flags)

	if *flags.Output != "stdout" {
		outFile, err := os.Create("Text/" + *flags.Output)
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

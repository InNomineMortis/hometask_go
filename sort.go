package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	//"strings"
)

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

	//exist := make(map[string]bool) //Словарь имеющихся слов
	intPtr := flag.Int("k", 0, "column")
	filePtr := flag.String("o", "output.txt", "in file, otherwise in stdout")

	flag.Parse()
	fmt.Println(*flags["f"], *flags["u"], *flags["r"], *flags["n"], *filePtr, *intPtr)

	file, err := ioutil.ReadFile("Text/data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	order := -1
	if *flags["r"] {
		order = 1
	}
	strs := bytes.Split(file, []byte("\n"))
	//if *flags["u"] {
	//	for v,_ := range strs {
	//		lowerStrs = strings.ToLower(string(strs[v]))
	//	}
	//}
	if *flags["f"] {
		sort.Slice(strs, func(i, j int) bool {
			return bytes.Compare(bytes.ToLower(strs[i]), bytes.ToLower(strs[j])) == order
		})
	} else {
		sort.Slice(strs, func(i, j int) bool {
			return bytes.Compare(strs[i], strs[j]) == order
		})
	}

	for _, v := range strs {
		fmt.Println(string(v))
	}

}

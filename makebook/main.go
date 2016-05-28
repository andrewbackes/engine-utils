package main

import (
	"flag"
	"fmt"
	"os"
)

type filters []string

func (f *filters) Set(s string) error {
	(*f) = append((*f), s)
	return nil
}

func (f *filters) String() string {
	s := ""
	for _, filter := range *f {
		s += "\t" + filter + "\n"
	}
	return s
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] pgnfiles ... \n", os.Args[0])
	flag.PrintDefaults()
}

func params() ([]string, *filters) {
	var f filters
	flag.Var(&f, "f", "filters to apply to a pgn")
	flag.Parse()
	flag.Usage = usage
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	return flag.Args(), &f
}

func main() {
	p, f := params()
	fmt.Println("Using:")
	fmt.Println("  PGN File(s):")
	for _, pgn := range p {
		fmt.Print("\t", pgn, "\n")
	}
	fmt.Println("  Filter(s):  \n", f)
}

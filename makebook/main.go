package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func params() ([]string, *filters, int) {
	var f filters
	flag.Var(&f, "f", "filters to apply to a pgn")
	var moves int
	flag.IntVar(&moves, "m", 14, "number of moves to include in the book")

	flag.Parse()
	flag.Usage = usage
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	return flag.Args(), &f, moves
}

func main() {
	p, f, m := params()
	fmt.Println("PGN File(s):")
	for _, pgn := range p {
		fmt.Print("\t", pgn, "\n")
	}
	outfile := strings.TrimSuffix(filepath.Base(p[0]), ".pgn") + ".bin"
	err := makeBook(p, f, m, outfile)
	handle(err)
	fmt.Println("Made", outfile)
}

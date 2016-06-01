package main

import (
	"errors"
	"fmt"
	"github.com/andrewbackes/chess/book"
	"github.com/andrewbackes/chess/pgn"
	"os"
	"time"
)

func handle(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func makeBook(pgnFiles []string, f *filters, moves int, outfile string) {
	start := time.Now()
	filts := parseFilters(f)
	if len(filts) != len(*f) {
		handle(errors.New("could not parse filters"))
	}
	fmt.Println("with", len(filts), "filters:")
	for _, f := range filts {
		fmt.Println("\t", f)
	}

	// open
	fmt.Print("Opening PGN(s)... ")
	pgns, err := openPGNs(pgnFiles)
	handle(err)
	fmt.Println("found", len(pgns), "games.")
	// filter
	fmt.Print("Filtering... ")
	filtered := pgn.Filter(pgns, filts...)
	fmt.Println(len(filtered), "games remain.")
	// convert
	fmt.Print("Converting to polyglot... ")
	b, err := book.FromPGN(filtered, moves*2)
	handle(err)
	fmt.Println(len(b.Positions), "unique positions.")
	// save
	handle(b.Save(outfile))
	lapsed := time.Now().Sub(start)
	fmt.Println("\nCreated", outfile, "in", lapsed)
}

func parseFilters(f *filters) []pgn.Filterer {
	var filts []pgn.Filterer
	for _, filt := range *f {
		tf := pgn.NewTagFilter(filt)
		if validTag(tf) {
			filts = append(filts, tf)
		}
	}
	return filts
}

func validTag(t pgn.TagFilter) bool {
	return t.Tag != "" && t.Operator != "" && t.Operand != ""
}

func openPGNs(pgns []string) ([]*pgn.PGN, error) {
	var r []*pgn.PGN
	for _, filename := range pgns {
		f, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		pgn, err := pgn.Open(f)
		if err != nil {
			return nil, err
		}
		r = append(r, pgn...)
	}
	if len(r) == 0 {
		return r, errors.New("no games found in pgn")
	}
	return r, nil
}

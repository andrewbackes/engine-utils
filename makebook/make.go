package main

import (
	"errors"
	"fmt"
	"github.com/andrewbackes/chess"
	"github.com/andrewbackes/chess/book"
	"os"
)

func handle(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func makeBook(pgn []string, f *filters, moves int, outfile string) error {
	filts := parseFilters(f)
	if len(filts) != len(*f) {
		handle(errors.New("could not parse filters"))
	}
	fmt.Println(len(filts), "filters.")

	pgns, err := openPGNs(pgn)
	handle(err)
	fmt.Println("Found", len(pgns), "games in PGN(s).")

	filtered := chess.FilterPGNs(pgns, filts...)
	b, err := book.FromPGN(filtered, moves*2)
	handle(err)
	handle(b.Save(outfile))
	return nil
}

func parseFilters(f *filters) []chess.TagFilter {
	var filts []chess.TagFilter
	for _, filt := range *f {
		tf := chess.NewTagFilter(filt)
		if validTag(tf) {
			filts = append(filts, tf)
		}
	}
	return filts
}
func validTag(t chess.TagFilter) bool {
	return t.Tag != "" && t.Operator != "" && t.Operand != ""
}

func openPGNs(pgns []string) ([]*chess.PGN, error) {
	var r []*chess.PGN
	for _, filename := range pgns {
		f, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		pgn, err := chess.ReadPGN(f)
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

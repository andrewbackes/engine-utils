package main

import (
	"fmt"
	"github.com/andrewbackes/chess/engines"
	"github.com/andrewbackes/chess/epd"
	"os"
)

func usage() {
	fmt.Println("Runepd expects an epd and engine file to be passed to it.\nFor example: ./runepd Bratko-Kopec.epd stockfish5")
}

func main() {
	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}
	epds := openEpd(os.Args[1])
	engine := openEngine(os.Args[2])
	run(epds, engine)
}

func openEngine(engineFilename string) engines.Engine {
	if _, err := os.Stat(engineFilename); err != nil {
		fmt.Println("Could not open", engineFilename)
		os.Exit(1)
	}
	engine, err := engines.NewUCIEngine(engineFilename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return engine
}

func openEpd(epdFilename string) []*epd.EPD {
	if _, err := os.Stat(epdFilename); err != nil {
		fmt.Println("Could not open", epdFilename)
		os.Exit(1)
	}
	epdFile, err := os.Open(epdFilename)
	if err != nil {
		fmt.Println("Could not open", epdFilename)
	}
	epds, err := epd.Open(epdFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return epds
}

package main

import (
	"flag"
	"fmt"
	"github.com/andrewbackes/chess/engines"
	"github.com/andrewbackes/chess/epd"
	"os"
	"time"
)

func usage() {
	fmt.Println("Runepd expects an epd and engine file to be passed to it.\nFor example: ./runepd Bratko-Kopec.epd stockfish5")
}

func main() {
	var timeout, outfile string
	flag.StringVar(&timeout, "t", "24h", "timeout for each test position")
	flag.StringVar(&outfile, "o", "results.txt", "output file for results")
	flag.Parse()
	fmt.Println(flag.Args())
	if len(flag.Args()) < 0 {
		usage()
		os.Exit(1)
	}
	timeoutDur, err := time.ParseDuration(timeout)
	if err != nil {
		fmt.Println("could not parse timeout")
		os.Exit(1)
	}
	epds := openEpd(flag.Args()[0])
	engine := openEngine(flag.Args()[1])
	run(epds, engine, timeoutDur)
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
	defer epdFile.Close()
	epds, err := epd.Read(epdFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return epds
}

package main

import (
	"fmt"
	"github.com/andrewbackes/chess/engines"
	"github.com/andrewbackes/chess/epd"
	"github.com/andrewbackes/chess/position"
	"strings"
	"time"
)

func run(epds []*epd.EPD, eng engines.Engine) {
	for _, e := range epds {
		omap := make(map[string]string)
		for _, o := range e.Operations {
			omap[o.Code] = o.Operand
		}
		for _, o := range e.Operations {
			switch o.Code {
			case "bm":
				bm(eng, e.Position, o.Operand, omap["id"])
			case "am":
				am(eng, e.Position, o.Operand)
			}
		}
	}
}

func bm(e engines.Engine, p *position.Position, m string, id string) {
	e.NewGame()
	fmt.Print(id, "\t")
	start := time.Now()
	c, err := e.Think(p)
	if err != nil {
		fmt.Println("ERROR", err, "with", p)
	} else {
		failed := waitForBm(c, p, m)
		lapsed := time.Now().Sub(start)
		fmt.Println()
		e.Stop()
		drain(c)
		if failed {
			fmt.Println("FAILED\ttimedout")
		} else {
			fmt.Println("\t", lapsed)
		}
	}
}

func drain(c chan string) {
	for {
		select {
		case <-c:
		// take it
		default:
			return
		}
	}
}

func waitForBm(c chan string, p *position.Position, m string) bool {
	for {
		fmt.Print(".")
		select {
		case info := <-c:
			//fmt.Println(info)
			words := strings.Split(info, " ")
			for i, word := range words {
				if word == "pv" {
					moveOne := words[i+1]
					pcnCm, _ := p.ParseMove(moveOne)
					pcnBm, _ := p.ParseMove(m)
					fmt.Print(pcnCm, "/", pcnBm, " ")
					if pcnCm == pcnBm {
						return false
					}
				}
			}
		case <-time.Tick(1 * time.Minute):
			return true
		}
	}
}

func am(e engines.Engine, p *position.Position, m string) {

}

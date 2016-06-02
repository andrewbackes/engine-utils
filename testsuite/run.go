package main

import (
	"fmt"
	"github.com/andrewbackes/chess/engines"
	"github.com/andrewbackes/chess/epd"
	"github.com/andrewbackes/chess/position"
	"strings"
	"time"
)

func run(epds []*epd.EPD, eng engines.Engine, timeout time.Duration) {
	for _, e := range epds {
		omap := make(map[string]string)
		for _, o := range e.Operations {
			omap[o.Code] = o.Operand
		}
		fmt.Print(omap["id"], "\t")
		for _, o := range e.Operations {
			switch o.Code {
			case "bm":
				bm(eng, e.Position, splitMoves(e.Position, o.Operand), timeout)
			case "am":
				am(eng, e.Position, o.Operand)
			}
		}
	}
}

func splitMoves(p *position.Position, s string) map[position.Move]struct{} {
	ms := strings.Split(s, " ")
	ret := make(map[position.Move]struct{})
	for _, m := range ms {
		pcn, err := p.ParseMove(m)
		if err == nil {
			ret[pcn] = struct{}{}
		}
	}
	return ret
}

func bm(e engines.Engine, p *position.Position, m map[position.Move]struct{}, timeout time.Duration) {
	e.NewGame()
	start := time.Now()

	timedOut := make(chan struct{})
	go func() {
		time.Sleep(timeout)
		close(timedOut)
	}()

	if think(e, p, m, timedOut) == true {
		fmt.Print("PASSED\t")
	} else {
		fmt.Print("FAILED\t")
	}
	lapsed := time.Now().Sub(start)
	fmt.Print(lapsed)
	e.Stop()
	time.Sleep(1 * time.Second)
	fmt.Println()
}

func think(e engines.Engine, p *position.Position, solutions map[position.Move]struct{}, timedOut chan struct{}) bool {
	c, err := e.Think(p)
	if err != nil {
		return false
	}
	for {
		select {
		case info := <-c:
			if foundBm(info, p, solutions) {
				return true
			}
		case <-timedOut:
			return false
		}
	}
}

func foundBm(info string, p *position.Position, solutions map[position.Move]struct{}) bool {
	words := strings.Split(info, " ")
	for i, word := range words {
		if (word == "bestmove" || word == "pv") && len(words) > i {
			moveOne := words[i+1]
			pcnCm, _ := p.ParseMove(moveOne)
			if _, exists := solutions[pcnCm]; exists {
				return true
			}
		}
	}
	return false
}

func am(e engines.Engine, p *position.Position, m string) {

}

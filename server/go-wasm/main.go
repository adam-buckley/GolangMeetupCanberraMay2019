package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
	"syscall/js"
	// "strconv"
	// "strings"
	"runtime"
	"runtime/pprof"
	"flag"
	"os"
	"log"
)
import _ "net/http/pprof"

type Dot struct {
	X int `json:x`
	Y int	`json:y`
	DX int `json:d_x`
	DY int `json:d_y`
}

var DotStruct = []Dot{}

// Upper bounds, lower bounds of 0
var BoundsX = 1920
var BoundsY = 902


func (d *Dot) UpdatePos() {
	// if d.DX >= 0 && d.X + d.DX > BoundsX {
	// 	d.DX *= -1
	// }

	// if d.DX < 0 && d.X + d.DX <= 0 {
	// 	d.DX *= -1
	// }

	// if d.DY >= 0 && d.Y + d.DY > BoundsY {
	// 	d.DY *= -1
	// }

	// if d.DY < 0 && d.Y + d.DY <= 0 {
	// 	d.DY *= -1
	// }
	
	if (d.X <= 0 || d.X >= BoundsX) {
		d.DX *= -1
	}

	if (d.Y <= 0 || d.Y >= BoundsY) {
		d.DY *= -1
	}

	d.X += d.DX
	d.Y += d.DY
}

func AddDot(this js.Value, args []js.Value) interface{} {

	var amount = 1
	if len(args) > 0 {
		amount = args[0].Int()

		if amount <= 0 {
			amount = 1
		}
	}

	for i := 0; i < amount; i++ {
		dot := Dot{}

		// Start at random point on screen
		dot.X = rand.Intn(BoundsX)
		dot.Y = rand.Intn(BoundsY)

		// Random X delta
		rand_delta := rand.Intn(5)
		flip := rand.Intn(2)
		if flip == 1 {
			rand_delta *= -1
		}

		if rand_delta == 0 {
			rand_delta = 1
		}

		dot.DX = rand_delta
		
		// Random Y delta
		rand_delta = rand.Intn(5)
		flip = rand.Intn(2)
		if flip == 1 {
			rand_delta *= -1
		}

		if rand_delta == 0 {
			rand_delta = 1
		}
		
		dot.DY = rand_delta
		
		DotStruct = append(DotStruct, dot)
	}

	fmt.Print("Struct size: ")
	fmt.Println(len(DotStruct))

	js.Global().Get("dots_counter").Set("innerHTML", len(DotStruct))
	
	return nil
}

func UpdateDots(this js.Value, args []js.Value) interface{} {
	return UpdateDotsLocal()
}

func UpdateDotsLocal() interface{} {
	start := time.Now()

	// Get canvas and update image
	canvas := js.Global().Get("document").Call("getElementById", "canvas")
	// canvas = document.getElementById("canvas")
	ctx := canvas.Call("getContext", "2d")
	
	ctx.Call("clearRect", 0, 0, BoundsX, BoundsY)

	// Black background
	ctx.Set("fillStyle", "black")
	ctx.Call("fillRect", 0, 0, BoundsX, BoundsY)
	
	// var str strings.Builder
	// str.WriteString("[")

	if len(DotStruct) > 0 {

		ctx.Set("fillStyle", "white")

		var wg sync.WaitGroup
		dotStructLength := len(DotStruct)
		wg.Add(dotStructLength)

		for index := 0; index < dotStructLength; index++ {
			go func(i int) {
				defer wg.Done()
				DotStruct[i].UpdatePos()
				ctx.Call("fillRect", DotStruct[i].X, DotStruct[i].Y, 1, 1)
			}(index)

			// str.WriteString(strconv.Itoa(DotStruct[index].X))
			// str.WriteString(",")
			// str.WriteString(strconv.Itoa(DotStruct[index].Y))
			
			// if index < len(DotStruct) - 1 {
			// 	str.WriteString(",")
			// }

			// fmt.Println(dot)
		}
		wg.Wait()
	}


	// bufferParam := buffer[:]
	// str.WriteString("]")

	// js.Global().Call("drawCanvas", str.String())
	
	// var t1 = performance.now()
	finish := time.Since(start)
	js.Global().Get("fps").Set("innerHTML", fmt.Sprintf("Frame process time: %s", finish))

	return nil
}

func registerCallbacks() {
	js.Global().Set("addDot", js.FuncOf(AddDot))
// 	js.Global().Set("updateDots", js.FuncOf(UpdateDots))
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	runtime.SetBlockProfileRate(1)
	
	flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close()
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
	}
	
	rand.Seed(time.Now().UTC().UnixNano())
	c := make(chan struct{}, 0)
	registerCallbacks()
	
	go heartBeat()
    time.Sleep(time.Second * 20)
	<-c

	if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        defer f.Close()
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
    }
}

func heartBeat() {
    for range time.Tick(time.Millisecond * (1000/60)) {
        go UpdateDotsLocal()
    }
}
	// timer1 := time.NewTimer(1000/60 * time.Millisecond)
	// ticker := time.NewTicker(1000/60 * time.Millisecond)
    // go func() {
	// 	for t := range ticker.C {
	// 		fmt.Println(t)
	// 		UpdateDotsLocal()
	// 	}
	// }()
	
	// <-c
// }

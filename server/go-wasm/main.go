package main

import (
	"fmt"
	"time"
	"math/rand"
	"syscall/js"
)

type Dot struct {
	X int
	Y int
	DX int
	DY int
}

var DotStruct = []Dot{}

// Upper bounds, lower bounds of 0
var BoundsX = 1920
var BoundsY = 902


func (d *Dot) UpdatePos() {
	if d.DX >= 0 && d.X + d.DX > BoundsX {
		d.DX *= -1
	}

	if d.DX < 0 && d.X + d.DX <= 0 {
		d.DX *= -1
	}

	if d.DY >= 0 && d.Y + d.DY > BoundsY {
		d.DY *= -1
	}

	if d.DY < 0 && d.Y + d.DY <= 0 {
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

	return nil
}

func UpdateDots(this js.Value, args []js.Value) interface{} {
	
	// Get canvas and update image
	canvas := js.Global().Get("document").Call("getElementById", "canvas")
	ctx := canvas.Call("getContext", "2d")
	
	// ctx.Call("clearRect", 0, 0, BoundsX, BoundsY)

	// Black background
	ctx.Set("fillStyle", "black")
	ctx.Call("fillRect", 0, 0, BoundsX, BoundsY)
	
	if len(DotStruct) > 0 {
		ctx.Set("fillStyle", "white")
		for index, dot := range DotStruct {
			// fmt.Println(dot)
			dot.UpdatePos()
			DotStruct[index] = dot
			// fmt.Println(dot)
			ctx.Call("fillRect", dot.X, dot.Y, 1, 1)
		}
	}

	return nil
}

func registerCallbacks() {
	js.Global().Set("addDot", js.FuncOf(AddDot))
	js.Global().Set("updateDots", js.FuncOf(UpdateDots))
}

func main() {
	
	rand.Seed(time.Now().UTC().UnixNano())

	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}

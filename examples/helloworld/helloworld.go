package main

import (
	"fmt"
	"os"

	g "github.com/AllenDang/giu"
)

func onClickMe() {
	fmt.Println("Hello world!")
}

func onImSoCute() {
	fmt.Println("Im sooooooo cute!!")
}

func onQuit() {
	os.Exit(0)
}

func loop() {
	g.SingleWindow("hello world", g.Layout{
		g.Label("Hello world from giu"),
		g.Line(
			g.Button("Click Me", onClickMe),
			g.Button("I'm so cute", onImSoCute),
			g.Button("Quit", onQuit),
		),
	})
}

func main() {
	wnd := g.NewMasterWindow("Hello world", 400, 200, g.MasterWindowFlagsNotResizable, nil)
	wnd.Main(loop)
}

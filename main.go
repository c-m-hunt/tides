package main

import (
	"github.com/c-m-hunt/tides/pkg/tides"
)

func main() {
	dt := tides.GetTides()
	dt.Display()
}

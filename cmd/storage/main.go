package main

import (
	"fmt"

	"github.com/PesotskyVladislav/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	fmt.Println("it works", st)
}

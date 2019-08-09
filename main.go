package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"math/rand"
	"os"
	"hash/fnv"
)

func rainbow(numOfSteps, step float64) (int, int, int) {

	var r, g, b float64

	h := step / numOfSteps
	i := math.Floor(h * 6)
	f := h*6 - i
	q := 1 - f

	os := math.Remainder(i, 6)

	switch os {
	case 0:
		r = 1
		g = f
		b = 0
	case 1:
		r = q
		g = 1
		b = 0
	case 2:
		r = 0
		g = 1
		b = f
	case 3:
		r = 0
		g = q
		b = 1
	case 4:
		r = f
		g = 0
		b = 1
	case 5:
		r = 1
		g = 0
		b = q
	}
	r = r * 255
	g = g * 255
	b = b * 255

	return int(r), int(g), int(b)
}

func main() {
	seedStr := os.Getenv("SEED")
	h := fnv.New64a()
	h.Write([]byte(seedStr))
	rand.Seed(int64(h.Sum64()))
	r, g, b := rainbow(100, rand.Float64() * 100)

	http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, `
<html>
<head><title>Demo</title>
<body style="background: rgb(%v, %v, %v)"></body>
</html>
		`, r, g, b)
    })

	port := os.Getenv("PORT")
	log.Printf("listening on 0.0.0.0:%v", port)
    http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

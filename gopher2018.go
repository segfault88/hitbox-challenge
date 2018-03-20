package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	counters    = map[string]int{}
	lock        = sync.Mutex{}
	numberSheet image.Image
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		get(w, r)
	} else if r.Method == "DELETE" {
		http.NotFound(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	key := parts[len(parts)-1]

	count := getCount(key)

	w.Header().Add("Content-Type", "image/png")

	numberString := fmt.Sprintf("%d", count)

	dest := image.NewRGBA(image.Rectangle{Max: image.Point{X: len(numberString) * 100, Y: 100}})

	for ii, c := range numberString {
		startPoint := numberPoint(c)
		log.Printf("start: %#v", startPoint)

		destRect := image.Rectangle{
			Min: image.Point{X: ii * 100, Y: 0},
			Max: image.Point{X: (ii + 1) * 100, Y: 100},
		}

		draw.Draw(dest, destRect, numberSheet, startPoint, draw.Over)
	}

	err := png.Encode(w, dest)
	if err != nil {
		panic(err)
	}
}

func numberPoint(n rune) image.Point {
	m := map[rune]image.Point{
		'0': image.Point{X: 100, Y: 300},
		'1': image.Point{X: 0, Y: 0},
		'2': image.Point{X: 100, Y: 0},
		'3': image.Point{X: 200, Y: 0},
		'4': image.Point{X: 0, Y: 100},
		'5': image.Point{X: 100, Y: 100},
		'6': image.Point{X: 200, Y: 100},
		'7': image.Point{X: 0, Y: 200},
		'8': image.Point{X: 100, Y: 200},
		'9': image.Point{X: 200, Y: 200},
		',': image.Point{X: 0, Y: 300},
		'.': image.Point{X: 200, Y: 300},
	}

	return m[n]
}

func getCount(key string) int {
	lock.Lock()
	defer lock.Unlock()

	count := counters[key] + 1
	counters[key] = count

	return count
}

func main() {
	data, err := os.Open("images/numbers.png")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	numberSheet, _, err = image.Decode(data)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/counter/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

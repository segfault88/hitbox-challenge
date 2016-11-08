package main

import (
	"os"
	"fmt"
	"image"
	_ "image/png"
	"image/draw"
	"image/png"
	"net/http"
	"strconv"
	"html"
	"strings"
)

var count map[string]int

func main()  {
	count=make(map[string]int)
	http.HandleFunc("/counter/", handler)

	http.ListenAndServe(":8080", nil)
}

func getNumber(num int) image.Image {
	f,e:=os.Open("images/numbers.png")
	if e != nil {
		fmt.Println("Can't open"+e.Error())
	}
	im,st,er:=image.Decode(f)
	if er != nil {
		fmt.Println(er.Error())
		return nil
	}
	defer f.Close()
	switch num {
	case 1:
		fmt.Println(st)
		src := image.NewRGBA(image.Rect(0, 0, 100, 100))
		draw.Draw(src, src.Bounds(), im, image.Point{0,0}, draw.Over)
		return src

	case 2:
		fmt.Println(st)
		src := image.NewRGBA(image.Rect(0, 0, 100, 100))
		draw.Draw(src, src.Bounds(), im, image.Point{100,0}, draw.Over)
		return src

	case 3:

		fmt.Println(3)
		src := image.NewRGBA(image.Rect(0, 0, 100, 100))
		draw.Draw(src, src.Bounds(), im, image.Point{200,0}, draw.Over)
		return src

	case 4:

		fmt.Println(st)
		src := image.NewRGBA(image.Rect(0, 0, 100, 100))
		draw.Draw(src, src.Bounds(), im, image.Point{0,100}, draw.Over)
		return src

	case 5:
		fmt.Println(st)
		src := image.NewRGBA(image.Rect(0, 0, 100, 100))
		draw.Draw(src, src.Bounds(), im, image.Point{100,100}, draw.Over)
		return src

	case 6:
		fmt.Println(st)
		src := image.NewRGBA(image.Rect(0, 0, 100, 100))
		draw.Draw(src, src.Bounds(), im, image.Point{200,100}, draw.Over)
		return src

	case 7:
		fmt.Println(st)
		src := image.NewRGBA(image.Rect(0, 0, 100, 100))
		draw.Draw(src, src.Bounds(), im, image.Point{0,200}, draw.Over)
		return src

	case 8:
		fmt.Println(st)
		src := image.NewRGBA(image.Rect(0, 0, 100, 100))
		draw.Draw(src, src.Bounds(), im, image.Point{100,200}, draw.Over)
		return src

	case 9:
		fmt.Println(st)
		src := image.NewRGBA(image.Rect(0, 0, 100, 100))
		draw.Draw(src, src.Bounds(),im, image.Point{200,200}, draw.Over)
		return src

	case 0:
		fmt.Println(st)
		src := image.NewRGBA(image.Rect(0, 0, 100, 100))
		draw.Draw(src, src.Bounds(), im, image.Point{100,300}, draw.Over)
		return src


	}
	return nil
}

func getDigits(num int) []int {
	var res []int
	for num >= 10 {
		rem:=num%10
		num=num/10
		res=append(res,rem)
	}
	res=append(res,num)
	return res
}

func composeNumbers(nums []int) image.Image{
	src := image.NewRGBA(image.Rect(0, 0, 100 * (len(nums)), 100))
	for i:=len(nums)-1;i>-1;i--{
		j:=len(nums)-1-i
		fmt.Println("Drawing "+strconv.Itoa(nums[i]))
		im:=getNumber(nums[i])
		draw.Draw(src, image.Rect(j*100,0,(j*100)+100,100), im, image.ZP, draw.Over)

	}
	return src

}


func handler(w http.ResponseWriter, r *http.Request) {
	st:=html.EscapeString(r.URL.Path)
	n:=strings.TrimPrefix(st,"/counter/")
	if count[n] != 0 {
		t:=count[n]
		t++
		count[n] = t
	}else{
		count[n] = 1
	}

	nums:=getDigits(count[n])
	png.Encode(w,composeNumbers(nums))
}
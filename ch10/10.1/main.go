package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png" // register PNG decoder
	"io"
	"os"
)

func main() {
	var t = flag.String("t", "", "image output format")
	flag.Parse()
	if err := transform(os.Stdin, os.Stdout, *t); err != nil {
		fmt.Fprintf(os.Stderr, "transform: %v\n", err)
		os.Exit(1)
	}
}

func transform(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	fmt.Fprintln(os.Stderr, "Output format =", format)
	if format == "jpeg" {
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	} else if format == "gif" {
		return gif.Encode(out, img, &gif.Options{})
	} else if format == "png" {
		return png.Encode(out, img)
	} else {
		return errors.New("unsupported output format")
	}

}

/*
$ go run gopl.io/ch3/mandelbrot | go run main.go -t jpeg > x.jpeg
Input format = png
Output format = jpeg
*/

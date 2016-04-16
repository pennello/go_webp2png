// chris 091315

// webp2png converts WebP images to PNG.
//
//	Usage of ./webp2png:
//	  -input string
//	        can be - for standard in (default "-")
//	  -output string
//	        can be - for standard out (default "-")
//
// If the output is not standard out and already exists, it will be
// truncated when opened for writing.
//
// Both the input and output files will be opened before doing any WebP
// decoding or PNG encoding.
package main

import (
	"flag"
	"io"
	"log"
	"os"

	"image/png"

	"golang.org/x/image/webp"
)

func main() {
	inname  := flag.String("input",  "-", "can be - for standard in")
	outname := flag.String("output", "-", "can be - for standard out")
	flag.Parse()

	var (
		infile  io.Reader
		outfile io.Writer
	)

	if *inname == "-" {
		infile = os.Stdin
	} else {
		file, err := os.Open(*inname)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		infile = file
	}

	if *outname == "-" {
		outfile = os.Stdout
	} else {
		file, err := os.Create(*outname)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		outfile = file
	}

	if m, err := webp.Decode(infile); err != nil {
		log.Fatal(err)
	} else {
		if err := png.Encode(outfile, m); err != nil {
			log.Fatal(err)
		}
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lorenzoliver/edi-tools/edifact/directories/d01b"
	"github.com/lorenzoliver/edi-tools/edifact/parser"
)

func main() {
	f, err := os.Open("./edifact/examples/input.edi")
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}
	p := parser.NewParser(f)
	interchange, err := p.Parse()
	if err != nil {
		log.Fatal(err)
	}
	if len(interchange.Messages) == 0 {
		log.Fatal("expected at least 1 message in interchange")
	}
	ap, err := parser.MapEdifact(interchange.Messages[0])
	if err != nil {
		log.Fatal(err)
	}
	aperak := ap.(*d01b.APERAK)
	for _, dtm := range aperak.DTM {
		fmt.Printf("datetime = %+v\n", dtm.C507.E2380)
	}
	for _, sg3 := range aperak.SG3 {
		switch sg3.NAD.E3035 {
		case "VP":
			fmt.Printf("Receiver manager = %s\n", sg3.NAD.C082.E3039)
		}
	}
}

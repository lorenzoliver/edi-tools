package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/lorenzoliver/edi-tools/edifact/directories/d01b"
	"github.com/lorenzoliver/edi-tools/edifact/parser"
)

const input = `UNB+UNOA:2+VALENCIAPORT+MVAL+60913:1143+2006091311458'
UNH+VPRT6000005192+APERAK:D:01B:UN'
BGM+963+VPRT6000005192+27'
DTM+137:200609131145:203'
RFF+ACW:20060622415205'
RFF+AGO:20060622415205'
NAD+VP+VALENCIAPORT'
ERC+1'
FTX+AAO+++El documento identificado no existe'
UNT+9+VPRT6000005192'
UNZ+1+2006091311458'
`

func init() {
	parser.Register("APERAK:D:01B:UN", reflect.TypeFor[d01b.APERAK]())
}

func main() {
	f := strings.NewReader(input)

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

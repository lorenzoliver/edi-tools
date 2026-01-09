package parser_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/lorenzoliver/edi-tools/edifact/directories/d01b"
	"github.com/lorenzoliver/edi-tools/edifact/parser"
)

func init() {
	parser.Register("APERAK:D:01B:UN", reflect.TypeFor[d01b.APERAK]())
}

func TestMappingEdifact(t *testing.T) {
	const sampleAPERAK = `UNB+UNOA:2+VALENCIAPORT+MVAL+60913:1143+2006091311458'
UNH+VPRT6000005192+APERAK:D:01B:UN'
BGM+963+VPRT6000005192+27'
DTM+137:200609131145:203'
RFF+ACW:20060622415205'
RFF+AGO:20060622415205'
NAD+VP+VALENCIAPORT'
ERC+1'
FTX+AAO+++El documento identificado no existe'
UNT+9+VPRT6000005192'
UNZ+1+2006091311458'`
	p := parser.NewParser(strings.NewReader(sampleAPERAK))
	interchange, err := p.Parse()
	if err != nil {
		t.Fatalf("error parsing message: %s", err)
	}
	aperak, err := parser.MapEdifact(interchange.Messages[0])
	if err != nil {
		t.Fatalf("%s", err)
	}
	castAperak := aperak.(*d01b.APERAK)
	if castAperak == nil {
		t.Fatalf("expected d01B APERAK, got=%T", aperak)
	}
	if len(castAperak.DTM) == 0 {
		t.Errorf("expected DTMs to be mapped, got 0")
	}
	if castAperak.DTM[0].C507.E2380 != "200609131145" {
		t.Errorf("expected DTM val to be '200609131145', got=%s", castAperak.DTM[0].C507.E2380)
	}
}

package parser_test

import (
	"strings"
	"testing"

	"github.com/lorenzoliver/edi-tools/edifact/parser"
)

func TestParser(t *testing.T) {
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
	parser := parser.NewParser(strings.NewReader(sampleAPERAK))
	interchange, err := parser.Parse()
	if err != nil {
		t.Fatalf("error parsing message: %s", err)
	}
	if interchange.UNB.Components[0].Elements[0] != "UNOA" {
		t.Errorf("Expected UNB to start with `UNOA`, got=%s", interchange.UNB.Components[0].Elements[0])
	}
	if len(interchange.Messages) != 1 {
		t.Errorf("expected 1 message, got=%d", len(interchange.Messages))
	}
}

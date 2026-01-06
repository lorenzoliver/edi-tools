package d01b

type APERAK struct {
	BGM BGM          `edi:"min=1,max=1,type=s"`
	DTM []DTM        `edi:"min=0,max=9,type=s"`
	FTX []FTX        `edi:"min=0,max=9,type=s"`
	CNT []CNT        `edi:"min=0,max=9,type=s"`
	SG1 []APERAK_SG1 `edi:"min=0,max=99,type=g"`
	SG2 []APERAK_SG2 `edi:"min=0,max=9,type=g"`
	SG3 []APERAK_SG3 `edi:"min=0,max=9,type=g"`
	SG4 []APERAK_SG4 `edi:"min=0,max=99999,type=g"`
}

type APERAK_SG1 struct {
	DOC DOC   `edi:"min=1,max=1,type=s"`
	DTM []DTM `edi:"min=0,max=99,type=s"`
}

type APERAK_SG2 struct {
	RFF RFF   `edi:"min=1,max=1,type=s"`
	DTM []DTM `edi:"min=0,max=9,type=s"`
}

type APERAK_SG3 struct {
	NAD NAD   `edi:"min=1,max=1,type=s"`
	CTA []CTA `edi:"min=0,max=9,type=s"`
	COM []COM `edi:"min=0,max=9,type=s"`
}

type APERAK_SG4 struct {
	ERC ERC          `edi:"min=1,max=1,type=s"`
	FTX []FTX        `edi:"min=0,max=1,type=s"`
	SG5 []APERAK_SG5 `edi:"min=0,max=9,type=g"`
}

type APERAK_SG5 struct {
	RFF RFF   `edi:"min=1,max=1,type=s"`
	FTX []FTX `edi:"min=0,max=9,type=s"`
}

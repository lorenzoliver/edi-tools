package d01b

type BGM struct {
	C002  C002   `edi:"min=0,max=1,type=c"`
	C106  C106   `edi:"min=0,max=1,type=c"`
	E1225 string `edi:"min=0,max=1,len=3,type=e"`
	E4343 string `edi:"min=0,max=1,len=3,type=e"`
}

type CTA struct {
	E3139 string `edi:"min=0,max=1,len=3,type=e"`
	C056  C056   `edi:"min=0,max=1,type=c"`
}

type CNT struct {
	C270 C270 `edi:"min=1,max=1,type=c"`
}

type COM struct {
	C076 C076 `edi:"min=1,max=1,type=c"`
}

type DOC struct {
	C002  C002   `edi:"min=1,max=1,type=c"`
	C503  C503   `edi:"min=0,max=1,type=c"`
	E3153 string `edi:"min=0,max=1,len=3,type=e"`
	E1220 string `edi:"min=0,max=1,len=2,type=e"`
	E1218 string `edi:"min=0,max=1,len=2,type=e"`
}

type DTM struct {
	C507 C507 `edi:"min=1,max=1,type=c"`
}

type ERC struct {
	C901 C901 `edi:"min=1,max=1,type=c"`
}

type FTX struct {
	E4451 string `edi:"min=1,max=1,len=3,type=e"`
	E4453 string `edi:"min=0,max=1,len=3,type=e"`
	C107  C107   `edi:"min=0,max=1,type=c"`
	C108  C108   `edi:"min=0,max=1,type=c"`
	E3453 string `edi:"min=0,max=1,len=3,type=e"`
	E4447 string `edi:"min=0,max=1,len=3,type=e"`
}

type NAD struct {
	E3035 string `edi:"min=1,max=1,len=3,type=e"`
	C082  C082   `edi:"min=0,max=1,type=c"`
	C058  C058   `edi:"min=0,max=1,type=c"`
	C080  C080   `edi:"min=0,max=1,type=c"`
	C059  C059   `edi:"min=0,max=1,type=c"`
	E3164 string `edi:"min=0,max=1,len=35,type=e"`
	C819  C819   `edi:"min=0,max=1,type=c"`
	E3251 string `edi:"min=0,max=1,len=17,type=e"`
	E3207 string `edi:"min=0,max=1,len=3,type=e"`
}

type RFF struct {
	C506 C506 `edi:"min=1,max=1,type=c"`
}

type UNH struct {
	E0062 string `edi:"min=1,max=1,len=14,type=e"`
	S009  S009   `edi:"min=1,max=1,type=c"`
	E0068 string `edi:"min=1,max=1,len=10,type=e"`
	S010  S010   `edi:"min=0,max=1,type=c"`
	S016  S016   `edi:"min=0,max=1,type=c"`
	S017  S017   `edi:"min=0,max=1,type=c"`
	S018  S018   `edi:"min=0,max=1,type=c"`
}

type UNT struct {
	E0074 string `edi:"min=1,max=1,len=10,type=e"`
	E0062 string `edi:"min=1,max=1,len=14,type=e"`
}

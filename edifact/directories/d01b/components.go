package d01b

type C002 struct {
	E1001 string `edi:"min=0,max=1,len=3,type=e"`
	E1131 string `edi:"min=0,max=1,len=17,type=e"`
	E3055 string `edi:"min=0,max=1,len=3,type=e"`
	E1000 string `edi:"min=0,max=1,len=35,type=e"`
}

type C056 struct {
	E3413 string `edi:"min=0,max=1,len=17,type=e"`
	E3412 string `edi:"min=0,max=1,len=35,type=e"`
}

type C058 struct {
	E3124 []string `edi:"min=1,max=5,len=35,type=e"`
}

type C059 struct {
	E3042 []string `edi:"min=1,max=4,len=35,type=e"`
}

type C076 struct {
	E3148 string `edi:"min=1,max=1,len=512,type=e"`
	E3155 string `edi:"min=1,max=1,len=3,type=e"`
}

type C080 struct {
	E3036 []string `edi:"min=1,max=5,len=35,type=e"`
	E3045 string   `edi:"min=0,max=1,len=3,type=e"`
}

type C082 struct {
	E3039 string `edi:"min=1,max=1,len=35,type=e"`
	E1131 string `edi:"min=0,max=1,len=17,type=e"`
	E3055 string `edi:"min=0,max=1,len=3,type=e"`
}

type C106 struct {
	E1004 string `edi:"min=0,max=1,len=35,type=e"`
	E1056 string `edi:"min=0,max=1,len=9,type=e"`
	E1060 string `edi:"min=0,max=1,len=6,type=e"`
}

type C107 struct {
	E4441 string `edi:"min=1,max=1,len=17,type=e"`
	E1131 string `edi:"min=0,max=1,len=17,type=e"`
	E3055 string `edi:"min=0,max=1,len=3,type=e"`
}

type C108 struct {
	E4440 []string `edi:"min=1,max=5,len=512,type=e"`
}

type C270 struct {
	E6069 string `edi:"min=1,max=1,len=3,type=e"`
	E6066 string `edi:"min=1,max=1,len=18,type=e"`
	E6411 string `edi:"min=0,max=1,len=3,type=e"`
}

type C503 struct {
	E1004 string `edi:"min=0,max=1,len=35,type=e"`
	E1373 string `edi:"min=0,max=1,len=3,type=e"`
	E1366 string `edi:"min=0,max=1,len=70,type=e"`
	E3453 string `edi:"min=0,max=1,len=3,type=e"`
	E1056 string `edi:"min=0,max=1,len=9,type=e"`
	E1060 string `edi:"min=0,max=1,len=6,type=e"`
}

type C506 struct {
	E1153 string `edi:"min=1,max=1,len=3,type=e"`
	E1154 string `edi:"min=0,max=1,len=70,type=e"`
	E1156 string `edi:"min=0,max=1,len=6,type=e"`
	E4000 string `edi:"min=0,max=1,len=35,type=e"`
	E1060 string `edi:"min=0,max=1,len=6,type=e"`
}

type C507 struct {
	E2005 string `edi:"min=1,max=1,len=3,type=e"`
	E2380 string `edi:"min=0,max=1,len=35,type=e"`
	E2379 string `edi:"min=0,max=1,len=3,type=e"`
}

type C819 struct {
	E3229 string `edi:"min=0,max=1,len=9,type=e"`
	E1131 string `edi:"min=0,max=1,len=17,type=e"`
	E3055 string `edi:"min=0,max=1,len=3,type=e"`
	E3228 string `edi:"min=0,max=1,len=70,type=e"`
}

type C901 struct {
	E9321 string `edi:"min=1,max=1,len=8,type=e"`
	E1131 string `edi:"min=0,max=1,len=17,type=e"`
	E3055 string `edi:"min=0,max=1,len=3,type=e"`
}

type S009 struct {
	E0065 string `edi:"min=1,max=1,len=6,type=e"`
	E0052 string `edi:"min=1,max=1,len=3,type=e"`
	E0054 string `edi:"min=1,max=1,len=3,type=e"`
	E0051 string `edi:"min=1,max=1,len=3,type=e"`
	E0057 string `edi:"min=0,max=1,len=6,type=e"`
	E0110 string `edi:"min=0,max=1,len=6,type=e"`
	E0113 string `edi:"min=0,max=1,len=6,type=e"`
}

type S010 struct {
	E0070 string `edi:"min=1,max=1,len=2,type=e"`
	E0073 string `edi:"min=0,max=1,len=1,type=e"`
}

type S016 struct {
	E0115 string `edi:"min=1,max=1,len=14,type=e"`
	E0116 string `edi:"min=0,max=1,len=3,type=e"`
	E0118 string `edi:"min=0,max=1,len=3,type=e"`
	E0051 string `edi:"min=0,max=1,len=3,type=e"`
}

type S017 struct {
	E0121 string `edi:"min=1,max=1,len=14,type=e"`
	E0122 string `edi:"min=0,max=1,len=3,type=e"`
	E0124 string `edi:"min=0,max=1,len=3,type=e"`
	E0051 string `edi:"min=0,max=1,len=3,type=e"`
}

type S018 struct {
	E0127 string `edi:"min=1,max=1,len=14,type=e"`
	E0128 string `edi:"min=0,max=1,len=3,type=e"`
	E0130 string `edi:"min=0,max=1,len=3,type=e"`
	E0051 string `edi:"min=0,max=1,len=3,type=e"`
}

package sample

//go:generate easyhide

//easyhide:json
type T1 struct {
	A1 string `json:"a1" easyhide:"show"`
	A2 string `json:"a2" easyhide:"hide"`
	A3 string `easyhide:"hide:HL"`
	A4 string `easyhide:"hide:HR"`
	A5 string `easyhide:"hide:NE"`
	A6 string `easyhide:"hide:HL,NE"`
	A7 string `easyhide:"hide:HR,NE"`
}

type T2 struct {
	B int
}

//easyhide:json
type T3 struct {
	C1 T4 `json:"c1"`
}

type T4 struct {
	D1 string `json:"d1" easyhide:"hide"`
}

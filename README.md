# go-easyhide (WORK IN PROGRESS)

## Example
```
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

func ExamplePrint() {
    v := T{
        A1: "value1", //will be shown
        A2: "value2", //will be hidden
        A3: "value3", //the left half of chars will be hidden
        A4: "value4", //the rigth half of chars will be hidden
        A5: "value5", //will be hidden
        A6: "", //will be shown
        A7: "", //will be shown
    }
    
    b, _ := v.EasyHide()
    fmt.Println(string(b))
}

//EXPECTED OUTPUT:
//{"a1:"value1","a2":"****","a3":"val****","a4":"****val","a5":"****","a5":"","a6":""}
```

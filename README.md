# go-easyhide (WORK IN PROGRESS)

## Example
```
//go:generate easyhide

//easyhide:json
type T struct {
    A1 string `easyhide:"a1,show"`
    A2 string `easyhide:"a2,hide"`
    A3 string `easyhide:"a3,hide:H"`
    A4 string `easyhide:"a4,hide:k"`
    A5 string `easyhide:"-"`
}

func ExamplePrint() {
    v := T{
        A1: "value1", //will be shown
        A2: "value2", //will be hidden
        A3: "value3", //half chars will be hidden with ****
        A4: "value4", //will be hidden, the number of * will be the same as number of hidden chars
        A5: "value5", //will not be included
    }
    
    b, _ := v.EasyHide()
    fmt.Println(string(b))
}

//EXPECTED OUTPUT:
//{"a1:"value1","a2":"****","a3":"val****","a4":"******"}
```

package main

import (
	"bufio"
	"fmt"
	"os"
	
	"strings"
)

type Printable interface {
	ToString() string
}

type ObjectValue struct {
	Data map[string]interface{}
}

type TypeBox struct {
	store map[string]interface{}
}

type ListValue struct {
	Data []interface{}
}

func NewTypeBox(store map[string]interface{}) *TypeBox {
	return &TypeBox{
		store: make(map[string]interface{}),
	}
}

func NewObjectValue(objectStore map[string]interface{}) *ObjectValue{
	return &ObjectValue{
		Data: make(map[string]interface{}),
	}
}


func (tb *TypeBox) SetScalar(key, typ, raw string) {
	tb.store[key] = raw
}

func (ov *ObjectValue) SetData(key, typ, raw string){
	ov.Data[key] = raw
}

func (tb *TypeBox) PrintKey(key string) string {
	if value, ok := tb.store[key].(string); ok {
		fmt.Println(value)
	} else {
		fmt.Println("Неизвестный тип")
	}
	return "че то не то"
}

func main() {
	var k int
	fmt.Println("Введите количество команд: ")
	fmt.Scan(&k)

	store := make(map[string]interface{})
	// objectStore := make(map[string]interface{})

	tb := NewTypeBox(store)
	// ov := NewObjectValue(objectStore)

	

	in := bufio.NewScanner(os.Stdin)

	for i := 0; i <= k && in.Scan(); i++ {
		data := in.Text()
		arrData := strings.Split(data, " ")
		switch arrData[0] {
		case "SET":
			tb.SetScalar(arrData[1], arrData[2], arrData[3])
		case "PRINT":
			tb.PrintKey(arrData[1])
		// case "OBJECT": 
		// 	num, err := strconv.Atoi(arrData[2])
		// 	if err != nil{
		// 		fmt.Println("Ошибка:", err)
		// 	} else{
		// 		for j := 0; j <= num; j++{
		// 			ov.SetData(arrData[0], arrData[1], arrData[2])
		// 		}
		// 	}
			

		}
		

	}
	

}

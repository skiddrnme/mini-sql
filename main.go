package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Printable interface {
	ToString() string
}

type ObjectValue struct {
	Data map[string]interface{}
}

// func (ov ObjectValue) ToString() string{

// }

type TypeBox struct {
	store map[string]interface{}
}

type ListValue struct {
	Data []interface{}
}

func NewTypeBox(store map[string]interface{}) *TypeBox {
	return &TypeBox{
		store: store,
	}
}

func NewObjectValue(objectStore map[string]interface{}) *ObjectValue {
	return &ObjectValue{
		Data: objectStore,
	}
}

func (tb *TypeBox) SetScalar(key, typ, raw string) {
	tb.store[key] = raw
}

func (tb *TypeBox) SetObject(key string, fields ObjectValue) {
	tb.store[key] = fields
}

func (tb *TypeBox) PrintKey(key string) string {
	if value, ok := tb.store[key].(string); ok {
		fmt.Println(value)
	} else {
		fmt.Println("null")
	}
	return "Error from Print key"
}

func main() {
	var k int
	fmt.Println("Введите кол-во команд: ")
	fmt.Scan(&k)

	store := make(map[string]interface{})
	objectStore := make(map[string]interface{})

	tb := NewTypeBox(store)
	ov := NewObjectValue(objectStore)

	var resultPrint []interface{}

	var printFlag bool
	in := bufio.NewScanner(os.Stdin)

	for i := 0; i <= k && in.Scan(); i++ {
		data := in.Text()
		arrData := strings.Split(data, " ")
		switch arrData[0] {
		case "SET":
			tb.SetScalar(arrData[1], arrData[2], arrData[3])
		case "PRINT":
			resultPrint = append(resultPrint, arrData[1])
			printFlag = true
		case "OBJECT":
			num, err := strconv.Atoi(arrData[2])
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				k += num
				ov.Data[arrData[0]] = arrData[2]
				
			}

		}

	}
// Доработать логику с OBJECT, добавить PUSH (list), подкрутить interface{}
	if printFlag {
		for _, v := range resultPrint {
			switch n := v.(type) {
			case string:
				tb.PrintKey(n)
			default:
				fmt.Println("Unknown type")
			}
		}
	}
	fmt.Println(ov.Data)
}

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

func NewObjectValue() *ObjectValue {
	return &ObjectValue{
		Data: make(map[string]interface{}),
	}
}

func (tb *TypeBox) SetScalar(key, typ, raw string) {
	switch typ {
	case "INT":
		v, err := strconv.Atoi(raw)
		if err == nil {
			tb.store[key] = v
		}
	case "STRING":
		tb.store[key] = raw
	case "FLOAT":
		v, err := strconv.ParseFloat(raw, 64)
		if err == nil {
			tb.store[key] = v
		}
	default:
		fmt.Println("Unknown type:", typ)
	}
}

func (ov *ObjectValue) SetField(key, typ, raw string) {
	ov.Data[key] = parseValue(typ, raw)
}

func (tb *TypeBox) SetObject(key string, fields [][3]string) {
	obj := NewObjectValue()

	for _, field := range fields {
		obj.SetField(field[0], field[1], field[2])
	}

	tb.store[key] = obj
}

func (tb *TypeBox) PrintKey(key string) string {
	if value, ok := tb.store[key].(string); ok {
		fmt.Println(value)
	} else {
		fmt.Println("null")
	}
	return "Error from Print key"
}

func (tb *TypeBox) MergeObjects(target, source string) {

}

func (lv ListValue) ToString() string {
	return "aba"
}

func (ov ObjectValue) ToString() string {
	
}

func parseValue(typ, val string) interface{} {
	switch typ {
	case "INT":
		v, _ := strconv.Atoi(val)
		return v
	case "STRING":
		return val
	case "FLOAT":
		v, _ := strconv.ParseFloat(val, 64)
		return v
	default:
		return nil
	}
}

func main() {

	store := make(map[string]interface{})

	tb := NewTypeBox(store)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	var resultPrint []interface{}
	var printFlag bool

	
	for range k {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, " ")

		switch parts[0] {

		case "SET":
			tb.SetScalar(parts[1], parts[2], parts[3])

		case "OBJECT":
			objName := parts[1]
			count, _ := strconv.Atoi(parts[2])
			if count > 20 {
				fmt.Println("Команд очень много!")
				os.Exit(1)
			}

			var fields [][3]string

			for range count {
				scanner.Scan()
				objLine := scanner.Text()
				objParts := strings.Split(objLine, " ")

				fields = append(fields, [3]string{
					objParts[0],
					objParts[1],
					objParts[2],
				})
			}

			tb.SetObject(objName, fields)

		case "PRINT":
			resultPrint = append(resultPrint, parts[1])
			printFlag = true
		}
	}

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
}

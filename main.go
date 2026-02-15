package main

import (
	"bufio"
	"fmt"
	"os"

	// "sort"
	"strconv"
	"strings"
)
// Интерфейс для вывода данных определенных типов
type Printable interface {
	ToString() string
}

// Значения (модель данных)
type ObjectValue struct {
	Data map[string]interface{}
}
func (ov ObjectValue) ToString() string {
	// сортировка ключей
	// формат {key:value,key2:value2}
	// использовать formatValue
	return "aba"
}
func NewObjectValue() *ObjectValue {
	return &ObjectValue{
		Data: make(map[string]interface{}),
	}
}
func (ov *ObjectValue) SetField(key, typ, raw string) {
	ov.Data[key] = parseValue(typ, raw)
}

type ListValue struct {
	Data []interface{}
}
func (lv ListValue) ToString() string {
	return "aba"
}


// Общий форматировщик
func formatValue(v interface{}) string {
	if val, ok := v.(Printable); ok {
		return val.ToString()
	}
	switch value := v.(type){
		case int:
			return strconv.Itoa(value)
		case float64:
			return strconv.FormatFloat(value, 'f', 2, 64)
		case string:
			return value
		default: 
			return "null"
	}
}

// TypeBox (ядро системы)
type TypeBox struct {
	store map[string]interface{}
}
func NewTypeBox() *TypeBox {
	return &TypeBox{
		store: make(map[string]interface{}),
	}
}

// Команды
func (tb *TypeBox) SetScalar(key, typ, raw string) string{
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
		return "Unknown type"
	}
	return ""
}

func (tb *TypeBox) SetObject(key string, fields [][3]string) {
	obj := NewObjectValue()

	for _, field := range fields {
		obj.SetField(field[0], field[1], field[2])
	}

	tb.store[key] = obj
}

func (tb *TypeBox) PushValue(key, typ, raw string) {
	// работать с ListValue
}

// Вывод 
func (tb *TypeBox) PrintKey(key string) string {
	if value, ok := tb.store[key]; ok {
		return formatValue(value)
	}
	return "null"
}

// Слияние двух объектов
func (tb *TypeBox) MergeObjects(target, source string) {

}

// Вспомогательная функция (Преобразование типов из stdin)
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
	tb := NewTypeBox()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	var resultPrint []interface{}
	var printFlag bool

	for i := 0; i < k; i++ {
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

			for j := 0; j < count; j++ {
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

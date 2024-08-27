package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

type Numeric interface {
	int | int8 | int16 | int32 | int64 | float64
}

func main() {
	A := Record{"String value", -12.123, Secret{"Dan", "Password"}}
	r := reflect.ValueOf(A)
	fmt.Println("String value: ", r.String())
	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType)
	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("\t%s ", iType.Field(i).Name)
		fmt.Printf("\twith type: %s ", r.Field(i).Type())
		fmt.Printf("\tand value _%v_\n", r.Field(i).Interface())
		// Check whether there are other structures embedded in Record
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		// Need to convert it to string in order to compare it
		if k.String() == "struct" {
			fmt.Println(r.Field(i).Type())
		}
		// Same as before but using the internal value
		if k == reflect.Struct {
			fmt.Println(r.Field(i).Type())
		}

		// 
		Print(12)
	}

}

func Print(s interface{}) {
	// type switch
	switch s.(type) {
	case int:
		fmt.Println("Type: int")
	case float64:
		fmt.Println("Type: float64")
	default:
		fmt.Println("Unknown data type!")
	}
}

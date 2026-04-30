package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Server struct {
	Host       string   `json:"host"`
	Port       int      `json:"port"`
	Debug      bool     `json:"debug"`
	AllowedIPs []string `json:"allowed_ips"`
}

func ToJSON(v any) (string, error) {
	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Struct {
		return "", errors.New("очікується структура")
	}

	var fields []string
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		tag := field.Tag.Get("json")

		var valStr string

		switch value.Kind() {
		case reflect.String:
			valStr = fmt.Sprintf(`"%s"`, value.String())
		case reflect.Int:
			valStr = fmt.Sprintf("%d", value.Int())
		case reflect.Bool:
			valStr = fmt.Sprintf("%t", value.Bool())
		case reflect.Slice:
			var elements []string
			for j := 0; j < value.Len(); j++ {
				elements = append(elements, fmt.Sprintf(`"%s"`, value.Index(j).String()))
			}
			valStr = "[\n\t\t" + strings.Join(elements, ",\n\t\t") + "\n\t]"
		}

		fields = append(fields, fmt.Sprintf(`	"%s": %s`, tag, valStr))
	}

	return "{\n" + strings.Join(fields, ",\n") + "\n}", nil
}

func main() {
	srv := Server{
		Host:       "localhost",
		Port:       8080,
		Debug:      true,
		AllowedIPs: []string{"192.168.1.1", "10.0.0.1"},
	}

	result, err := ToJSON(srv)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

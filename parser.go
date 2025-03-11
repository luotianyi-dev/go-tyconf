package tyconf

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Field struct {
	CLIName      string
	EnvName      string
	Description  string
	Kind         reflect.Kind
	DefaultValue any
}

func Parse(conf any) any {
	confType := reflect.TypeOf(conf)
	confValue := reflect.ValueOf(conf)
	fields := make([]Field, confType.NumField())
	result := reflect.New(confType).Elem()
	flagsCount := 0

	for i := range confType.NumField() {
		ft := confType.Field(i)
		fv := confValue.Field(i)
		fields[i] = Field{
			CLIName:      ft.Tag.Get("cli"),
			EnvName:      ft.Tag.Get("env"),
			Description:  ft.Tag.Get("description"),
			Kind:         fv.Kind(),
			DefaultValue: fv.Interface(),
		}

		result.Field(i).Set(fv)
		ptr := result.Field(i).Addr().Interface()
		if strings.TrimSpace(fields[i].CLIName) != "" {
			flagsCount++
			fieldFromCLI(fields[i], ptr)
		}
		if strings.TrimSpace(fields[i].EnvName) != "" {
			fieldFromEnv(fields[i], ptr)
		}
	}

	if flagsCount > 0 {
		flag.Parse()
	}
	return result.Interface()
}

func fieldFromEnv(f Field, p any) {
	strValue, ok := os.LookupEnv(f.EnvName)
	if ok {
		if f.Kind == reflect.String {
			*(p.(*string)) = strValue
			return
		}
		if f.Kind == reflect.Bool {
			boolValue, err := strconv.ParseBool(strValue)
			if err == nil {
				*(p.(*bool)) = boolValue
				return
			}
		}
		if f.Kind == reflect.Int {
			intValue, err := strconv.Atoi(strValue)
			if err == nil {
				*(p.(*int)) = intValue
				return
			}
		}
		if f.Kind == reflect.Int64 {
			int64Value, err := strconv.ParseInt(strValue, 10, 64)
			if err == nil {
				*(p.(*int64)) = int64Value
				return
			}
		}
		if f.Kind == reflect.Uint {
			uintValue, err := strconv.ParseUint(strValue, 10, 0)
			if err == nil {
				*(p.(*uint)) = uint(uintValue)
				return
			}
		}
		if f.Kind == reflect.Uint64 {
			uint64Value, err := strconv.ParseUint(strValue, 10, 64)
			if err == nil {
				*(p.(*uint64)) = uint64Value
				return
			}
		}
		if f.Kind == reflect.Float64 {
			float64Value, err := strconv.ParseFloat(strValue, 64)
			if err == nil {
				*(p.(*float64)) = float64Value
				return
			}
		}
	}
}

func fieldFromCLI(f Field, p any) {
	cliDescription := fmt.Sprintf("%s [env: $%s]", f.Description, f.EnvName)
	if f.Kind == reflect.String {
		flag.StringVar(p.(*string), f.CLIName, f.DefaultValue.(string), cliDescription)
		return
	}
	if f.Kind == reflect.Bool {
		flag.BoolVar(p.(*bool), f.CLIName, f.DefaultValue.(bool), cliDescription)
		return
	}
	if f.Kind == reflect.Int {
		flag.IntVar(p.(*int), f.CLIName, f.DefaultValue.(int), cliDescription)
		return
	}
	if f.Kind == reflect.Int64 {
		flag.Int64Var(p.(*int64), f.CLIName, f.DefaultValue.(int64), cliDescription)
		return
	}
	if f.Kind == reflect.Uint {
		flag.UintVar(p.(*uint), f.CLIName, f.DefaultValue.(uint), cliDescription)
		return
	}
	if f.Kind == reflect.Uint64 {
		flag.Uint64Var(p.(*uint64), f.CLIName, f.DefaultValue.(uint64), cliDescription)
		return
	}
	if f.Kind == reflect.Float64 {
		flag.Float64Var(p.(*float64), f.CLIName, f.DefaultValue.(float64), cliDescription)
		return
	}
}

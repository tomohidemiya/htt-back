package main

import (
    user "app/backend/models"
    "fmt"
    "reflect"
)

func ToReflect(itf interface{}) {
    rv := reflect.ValueOf(itf)
    rt := rv.Type()

    for i := 0; i < rt.NumField(); i++ {
        field := rt.Field(i)
        kind := field.Type.Kind()
        value := rv.FieldByName(field.Name)
        switch kind {
        case reflect.Struct:
            fmt.Printf("struct name: %v, value %v, kind: %v\n", field.Name, value, kind)
            ToReflect(field)
            break
        //case reflect.Ptr:
        //   return ToReflect(reflect.Indirect(field))
        default:
            fmt.Printf("name: %v, value %v, kind: %v\n", field.Name, value, kind)
            break
        }
    }
}

func main() {
    var u user.User
    ToReflect(u)
}

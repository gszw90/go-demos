package main

import (
	"fmt"
	"reflect"

	"github.com/huandu/go-sqlbuilder"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func main1() {
	// 创建一个Struct实例
	u := User{ID: 1, Name: "test"}
	st := sqlbuilder.NewStruct(u)

	// 打印字段信息
	fmt.Println("Fields:", st.Columns())

	// 打印结构体标签格式
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field: %s, DB Tag: %s\n", field.Name, field.Tag.Get("db"))
	}

	// 创建一个简单的查询
	sb := st.SelectFrom("users")
	fmt.Println("SQL:", sb.String())
}

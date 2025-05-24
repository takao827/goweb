package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

// OpenAPIスキーマから自動生成した構造体
type AddressForm struct {
	Pref  int    `json:"pref"`
	Addr1 string `json:"addr1"`
	Addr2 string `json:"addr2"`
	Addr3 string `json:"addr3,omitempty" validate:"required"`
}

// 指定したフィールドのOpenAPIスキーマ名を取得する
func (f AddressForm) GetSchemaName(getter func(*AddressForm) any) string {
	typ := reflect.TypeOf(f)
	val := reflect.ValueOf(&f).Elem()
	tgt := getter(&f)

	for i := 0; i < typ.NumField(); i++ {
		field := val.Field(i)
		if field.Addr().Interface() == tgt {
			structField := typ.Field(i)
			return strings.Split(structField.Tag.Get("json"), ",")[0]
		}
	}

	return ""
}

// エンティティ
type Address struct {
	PrefectureCode int
	City           string
	Address        string
	Building       string
}

// 指定したフィールドのフィールド名を取得する
func (a Address) GetFieldName(getter func(*Address) any) string {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(&a).Elem()
	tgt := getter(&a)

	for i := 0; i < typ.NumField(); i++ {
		field := val.Field(i)
		if field.Addr().Interface() == tgt {
			structField := typ.Field(i)

			return structField.Name
		}
	}

	return ""
}

/**
 * 自動生成された構造体と、自分で作成した構造体のフィールド名をマッピングする方法を調べていた
 *
 * エンティティのバリデーションエラーを、OpenAPIスキーマのフィールド名に変換してレスポンスしたかった
 * 自動生成された構造体に変更があった場合、エラーを出して変更が必要なことを把握したかった
 *
 * reflectを使うしかないのか…？
 */
func learn_reflect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, reflect!")

	f := AddressForm{}
	e := Address{}

	entitySchemaMap := map[string]string{
		e.GetFieldName(func(a *Address) any { return &a.PrefectureCode }): f.GetSchemaName(func(f *AddressForm) any { return &f.Pref }),
		e.GetFieldName(func(a *Address) any { return &a.City }):           f.GetSchemaName(func(f *AddressForm) any { return &f.Addr1 }),
		e.GetFieldName(func(a *Address) any { return &a.Address }):        f.GetSchemaName(func(f *AddressForm) any { return &f.Addr2 }),
		e.GetFieldName(func(a *Address) any { return &a.Building }):       f.GetSchemaName(func(f *AddressForm) any { return &f.Addr3 }),
	}

	fmt.Fprintf(w, "PrefectureCode: %s\n", entitySchemaMap["PrefectureCode"])
	fmt.Fprintf(w, "City: %s\n", entitySchemaMap["City"])
	fmt.Fprintf(w, "Address: %s\n", entitySchemaMap["Address"])
	fmt.Fprintf(w, "Building: %s\n", entitySchemaMap["Building"])
}

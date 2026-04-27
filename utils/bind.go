package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

func Bind[T any](c echo.Context, maxBytes int64) (*T, error) {
	switch {
	case c.Request().Method == http.MethodGet, c.Request().Method == http.MethodDelete:
		v := new(T)
		fillFrom(reflect.ValueOf(v).Elem(), c.QueryParam)
		return v, nil
	case IsType(c, "multipart/form-model"):
		v := new(T)
		fillFrom(reflect.ValueOf(v).Elem(), c.FormValue)
		return v, nil
	default:
		v := new(T)
		c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, maxBytes)
		if err := json.NewDecoder(c.Request().Body).Decode(v); err != nil {
			return nil, fmt.Errorf("invalid JSON: %w", err)
		}
		return v, nil
	}
}

func fillFrom(rv reflect.Value, lookup func(string) string) {
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		if !f.IsExported() {
			continue
		}
		tag := f.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		name := strings.SplitN(tag, ",", 2)[0]
		val := lookup(name)
		if val == "" {
			continue
		}
		fv := rv.Field(i)
		if fv.CanSet() && fv.Kind() == reflect.String {
			fv.SetString(val)
		}
	}
}

package repository

import (
	"reflect"

	"github.com/gauas/config-service/utils"
	"gorm.io/gorm"
)

func applyArgs(db *gorm.DB, args ...interface{}) *gorm.DB {
	if len(args) == 0 {
		return db
	}

	first := args[0]

	switch q := first.(type) {

	case string:
		if q != "" {
			return db.Where(q, args[1:]...)
		}
		return db

	default:
		return buildFilter(db, q)
	}
}

func buildFilter(db *gorm.DB, filter interface{}) *gorm.DB {
	v := reflect.ValueOf(filter)
	t := reflect.TypeOf(filter)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		fieldVal := v.Field(i)
		fieldType := t.Field(i)

		if !fieldVal.CanInterface() {
			continue
		}

		if fieldType.Tag.Get("filter") == "-" {
			continue
		}

		if fieldVal.Kind() == reflect.Ptr {
			if fieldVal.IsNil() {
				continue
			}
			fieldVal = fieldVal.Elem()
		} else if fieldVal.IsZero() {
			continue
		}

		column := SnakeCase(fieldType.Name)

		db = db.Where(column+" = ?", fieldVal.Interface())
	}

	return db
}

func SnakeCase(s string) string {
	var result []rune
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if utils.IsUpper(r) {
			if shouldAddUnderscore(runes, i) {
				result = append(result, '_')
			}
			r = utils.ToLower(r)
		}

		result = append(result, r)
	}

	return string(result)
}

func shouldAddUnderscore(runes []rune, i int) bool {
	if i == 0 {
		return false
	}

	prev := runes[i-1]

	if utils.IsLower(prev) || utils.IsDigit(prev) {
		return true
	}

	if i+1 < len(runes) && utils.IsLower(runes[i+1]) {
		return true
	}

	return false
}

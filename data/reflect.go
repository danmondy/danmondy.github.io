package data

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
)

// DELETE this doesn't use reflection, it's uber simple
func Delete(table string, id string) error {
	_, err := db.Exec("delete from "+table+" where id = ?", id)
	return err
}

func Insert(o interface{}) error {
	t := strings.ToLower(reflect.TypeOf(o).Elem().Name())
	ts := strings.Split(t, ".")
	t = ts[len(ts)-1]
	sql := "INSERT INTO %s (%s) VALUES (%s)"

	var cols []string
	var vals []string

	val := reflect.ValueOf(o).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		tag := typeField.Tag.Get("db")
		tags := strings.Split(tag, ",")

		skip := false

		//set column name to the name of the field by default
		col := typeField.Name
		for _, x := range tags {
			if x == "autoinsert" || x == "omit" {
				skip = true
			} else if x != "primarykey" {
				col = x //if there is another flag set this value to its value
			}
		}

		if skip {
			continue
		}
		cols = append(cols, col)

		switch valueField.Interface().(type) {
		case string:
			vals = append(vals, fmt.Sprintf("'%s'", valueField.Interface()))
		case time.Time:
			vals = append(vals, fmt.Sprintf("'%s'", TimeToString(valueField.Interface().(time.Time))))
		case int:
			vals = append(vals, fmt.Sprintf("'%d'", valueField.Interface().(int)))
		default:
			vals = append(vals, fmt.Sprintf("%v", valueField.Interface()))

		}
	}

	fmt.Println("table name: ", t)
	sql = fmt.Sprintf(sql, t, strings.Join(cols, ", "), strings.Join(vals, ", "))

	fmt.Println(sql)

	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	//works with mysql to and auto id increment
	//id, err := result.LastInsertId()

	return nil
}

// take anything, return a strongly typed object from the db
func GetById[T any](id string) (T, error) {
	var item T

	t := strings.ToLower(reflect.TypeOf([0]T{}).Elem().Name())
	ts := strings.Split(t, ".")
	t = ts[len(ts)-1]

	sql := fmt.Sprintf("SELECT * FROM %s where id = ?", t)

	err := db.Get(&item, sql, id)
	if err != nil {
		log.Println(err)
		log.Println(sql, "(id: )", id)
		return item, err
	}

	return item, nil
}

func GetAll[T any]() ([]T, error) {
	items := make([]T, 2)
	zero := [0]T{}
	t := strings.ToLower(reflect.TypeOf(zero).Elem().Name())
	ts := strings.Split(t, ".")
	t = ts[len(ts)-1]

	sql := fmt.Sprintf("SELECT * FROM %s", t)

	err := db.Select(&items, sql)
	if err != nil {
		log.Println(err)
		return items, err
	}

	return items, nil
}

func GetAllFor[T any](id string, colName string) ([]T, error) {
	var items []T
	zero := [0]T{}
	t := strings.ToLower(reflect.TypeOf(zero).Elem().Name())
	ts := strings.Split(t, ".")
	t = ts[len(ts)-1]

	sql := fmt.Sprintf("SELECT * FROM %s where %s = ?", t, colName)

	log.Println(sql)
	err := db.Select(&items, sql, id)
	if err != nil {
		log.Println(err)
		return items, err
	}

	return items, nil
}

func GetAllForSorted[T any](id string, colName string, sortCol string) ([]T, error) {
	var items []T
	zero := [0]T{}
	t := strings.ToLower(reflect.TypeOf(zero).Elem().Name())
	ts := strings.Split(t, ".")
	t = ts[len(ts)-1]

	sql := fmt.Sprintf("SELECT * FROM %s where %s = ? order by %s", t, colName, sortCol)

	log.Println(sql)
	err := db.Select(&items, sql, id)
	if err != nil {
		log.Println(err)
		return items, err
	}

	return items, nil
}

// this one is not tested yet. Needs some work.
func Update(o interface{}) error {
	t := strings.ToLower(reflect.TypeOf(o).Elem().Name())
	ts := strings.Split(t, ".")
	t = ts[len(ts)-1]
	sql := "UPDATE %s set %s WHERE %s"
	setTemplate := "%s = %s"
	where := "%s = '%s'"

	pk := false
	var sets []string

	val := reflect.ValueOf(o).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		tag := typeField.Tag.Get("db")
		tags := strings.Split(tag, ",")

		skip := false

		//set column name to the name of the field by default
		col := typeField.Name
		for _, x := range tags {
			if x == "omit" {
				skip = true
			} else if x == "primarykey" {
				pk = true
			} else {
				col = x //if there is another flag set this value to its value
			}
		}

		if pk {
			where = fmt.Sprintf(where, col, fmt.Sprintf("%v", valueField.Interface()))
			pk = false
			continue
		}

		if skip {
			continue
		}

		switch valueField.Interface().(type) {
		case string:
			sets = append(sets, fmt.Sprintf(setTemplate, col, fmt.Sprintf("'%s'", valueField.Interface())))
		case time.Time:
			sets = append(sets, fmt.Sprintf(setTemplate, col, fmt.Sprintf("'%s'", TimeToString(valueField.Interface().(time.Time)))))
		default:
			sets = append(sets, fmt.Sprintf(setTemplate, col, fmt.Sprintf("%v", valueField.Interface())))
		}
	}

	sql = fmt.Sprintf(sql, t, strings.Join(sets, ", "), where)
	fmt.Println("table name: ", t)
	fmt.Println("sql statement:", sql)

	//TODO: update the ID with on the pointer after it is made
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

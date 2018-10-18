package model

import (
	"fmt"
)


type Record struct {
	Type string
	rec map[string]interface{}
}



func (r *Record) LoadMap(record map[string]interface{}) error {
	r.rec = record
	fmt.Printf("read all %v\n", r.rec)
	return nil
}




//
//func (r *Record) Load(values []bigquery.Value, schema bigquery.Schema) error {
//	if len(r.rec) == 0 {
//		r.rec = make(map[string]bigquery.Value)
//	}
//	loadMap(r.rec, values, schema)
//	return nil
//}
//
//func (r *Record) Set(field string, val interface{})  {
//	if len(r.rec) == 0 {
//		r.rec = make(map[string]bigquery.Value)
//	}
//	r.rec[field] = val
//}
//
//func (r *Record) Get(field string) (interface{}, bool) {
//	value, ok := r.rec[field]
//	return value, ok
//}
//
//
//
//
//
//func loadMap(aMap map[string]bigquery.Value, vals []bigquery.Value, schema bigquery.Schema) {
//	for i, field := range schema {
//		val := vals[i]
//		var v interface{}
//		switch {
//		case val == nil:
//			v = val
//		case field.Schema == nil:
//			v = val
//		case !field.Repeated:
//			m2 := map[string]bigquery.Value{}
//			loadMap(m2, val.([]bigquery.Value), field.Schema)
//			v = m2
//		default: // repeated and nested
//			sval := val.([]bigquery.Value)
//			vs := make([]bigquery.Value, len(sval))
//			for j, e := range sval {
//				m2 := map[string]bigquery.Value{}
//				loadMap(m2, e.([]bigquery.Value), field.Schema)
//				vs[j] = m2
//			}
//			v = vs
//		}
//		aMap[field.Name] = v
//	}
//}
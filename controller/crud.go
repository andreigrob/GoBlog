package controller

import (
	lt "container/list"
	ct "context"
	ft "fmt"
	"log"
	ht "net/http"

	ml "github.com/andreigrob/web_quiz_andrei/model"
	ut "github.com/andreigrob/web_quiz_andrei/utils"
)

type Rows = ut.Rows

func save[T ml.EntityI](c ControllerI, wr ResWr, item T) (id int64, e error) {
	Class := item.Class()
	conn := c.GetConn()
	// insert into TABLE_NAME (ID, FIELD1, FIELD2) values ($1, $2, $3) returning ID
	query := ft.Sprintf("insert into %s (%s) values (%s) returning Id", Class.GetTableName(), Class.GetFieldNames(), Class.GetFieldStr())
	log.Printf("Query: %s", query)
	name := ut.Name(item)
	fields := item.GetFields()
	log.Printf("Fields: %v", fields)
	rows, e := conn.Query(ct.Background(), query, fields...)
	if e != nil {
		ht.Error(wr, "Failed to save item", ht.StatusInternalServerError)
		log.Printf("Failed to save %s: %v", name, e)
		return
	}
	log.Printf("rows: %v\n", rows)
	defer rows.Close()
	if !rows.Next() {
		ht.Error(wr, "No rows returned", ht.StatusInternalServerError)
		log.Printf("No rows returned for %s: %v", name, e)
	}
	if e = rows.Scan(&id); e != nil {
		ht.Error(wr, "Failed to scan item", ht.StatusInternalServerError)
		log.Printf("Failed to scan %s: %v", name, e)
		return
	}
	log.Printf("Saved %s: %v", name, id)
	return id, e
}

func find[T any, P ml.EntityPI[T]](c ControllerI, wr ResWr, entities *[]P) (e error) {
	var entity P
	//var a = ut.Star[ml.IEntity](entity)
	Class := entity.Class()
	conn := c.GetConn()
	query := ft.Sprintf("select %s from %s", Class.GetAllFieldNames(), Class.GetTableName())
	log.Printf("Query: %s", query)
	name := Class.GetName()
	rows, e := conn.Query(ct.Background(), query)
	if e != nil {
		ht.Error(wr, "Failed to retrieve items", ht.StatusInternalServerError)
		log.Printf("Failed to find %s: %v", name, e)
		return
	}
	defer rows.Close()
	// create a linked list of items
	items := lt.New()
	log.Printf("Retrieving %s", name)

	var val P
	for rows.Next() {
		val = new(T)
		if e = val.Scan(rows); e != nil {
			log.Printf("Failed to scan %s: %v", name, e)
			return
		}
		items.PushBack(val)
		log.Printf("Found %s: %v", name, val.GetId())
	}

	Len := items.Len()
	*entities = make([]P, Len)
	var i int
	v := items.Front()
	for ; v != nil; v = v.Next() {
		(*entities)[i] = v.Value.(P)
		i++
	}
	log.Printf("Found %d %s", Len, name)
	return
}

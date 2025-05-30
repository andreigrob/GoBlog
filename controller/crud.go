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

func save[T ml.IEntity](c IController, w RW, item T) (id int64, err error) {
	Class := item.Class()
	conn := c.GetConn()
	query := ft.Sprintf("insert into %s (%s) values (%s) returning Id", Class.GetTableName(), Class.GetFieldNames(), Class.GetFieldStr())
	log.Printf("Query: %s", query)
	name := ut.Name(item)
	fields := item.GetFields()
	log.Printf("Fields: %v", fields)
	var rows Rows
	if rows, err = conn.Query(ct.Background(), query, fields...); err != nil {
		ht.Error(w, "Failed to save item", ht.StatusInternalServerError)
		log.Printf("Failed to save %s: %v", name, err)
		return
	}
	log.Printf("rows: %v\n", rows)
	defer rows.Close()
	if !rows.Next() {
		ht.Error(w, "No rows returned", ht.StatusInternalServerError)
		log.Printf("No rows returned for %s: %v", name, err)
	}
	if err = rows.Scan(&id); err != nil {
		ht.Error(w, "Failed to scan item", ht.StatusInternalServerError)
		log.Printf("Failed to scan %s: %v", name, err)
		return
	}
	log.Printf("Saved %s: %v", name, id)
	return id, err
}

func find[T ml.IEntity](c IController, w RW, entities *[]T) (err error) {
	var entity T
	/var a = ut.Star[ml.IEntity](entity)
	Class := entity.Class()
	conn := c.GetConn()
	query := ft.Sprintf("select %s from %s", Class.GetAllFieldNames(), Class.GetTableName())
	log.Printf("Query: %s", query)
	name := Class.GetName()
	var rows Rows
	if rows, err = conn.Query(ct.Background(), query); err != nil {
		ht.Error(w, "Failed to retrieve items", ht.StatusInternalServerError)
		log.Printf("Failed to find %s: %v", name, err)
		return
	}
	defer rows.Close()
	// create a linked list of items
	items := lt.New()
	log.Printf("Retrieving %s", name)
	for rows.Next() {
		val := entity.NewAny().(T)
		if err = val.Scan(rows); err != nil {
			log.Printf("Failed to scan %s: %v", name, err)
			return
		}
		items.PushBack(val)
		log.Printf("Found %s: %v", name, val.GetId())
	}

	Len := items.Len()
	*entities = make([]T, Len)
	var i int
	v := items.Front()
	for ; v != nil; v = v.Next() {
		(*entities)[i] = v.Value.(T)
		i++
	}
	log.Printf("Found %d %s", Len, name)
	return
}

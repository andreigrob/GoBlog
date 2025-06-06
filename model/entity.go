package model

import (
	"log"
	sc "strconv"
	st "strings"

	ut "github.com/andreigrob/web_quiz_andrei/utils"
)

// A Class has an associated EntityObject.
type ClassI interface {
	Class() (_ *EntityObject)
}

// An Entity is a Class that has an Id and Fields.
type EntityI interface {
	ClassI

	GetId() (_ int64)
	SetId(id int64)
	GetFields() (_ []any)
	NewAny() (_ any)
	InitAny(args ...any)
	Scan(rows Rows) (_ error)
}

type EntityPI[T any] interface {
	EntityI
	*T
}

type ArticlePI = EntityPI[Article]

type AnswerPI = EntityPI[Answer]

type CommentPI = EntityPI[Comment]

func Hello[T any, P EntityPI[T]]() {
	var a T
	log.Printf("Hello %v", a)
	var b P
	log.Printf("Hello %v", b.Class().GetName())
	var c P = &a
	log.Printf("Hello %v", c.Class().GetName())
}

func Hello2() {
	Hello[Article]()
}

// An EntityObject stores the Name, Table Name, and Field Names of a Class, and attributes derived from them.
type EntityObject struct {
	Name          string
	NameLower     string
	TableName     string
	FieldNames    string
	AllFieldNames string
	NFields       int
	FieldStr      string
	AllFieldStr   string
}

// fieldsString creates a field string for a given number of fields.
func fieldsString(nFields int) (_ string) {
	var sb st.Builder
	_, _ = sb.WriteString("$1")
	var i int = 2
	for ; i <= nFields; i++ {
		_, _ = sb.WriteString(", $" + sc.Itoa(i))
	}
	return sb.String()
}

// Init initializes an EntityObject.
func (e *EntityObject) Init(Name, TableName, FieldNames string) {
	e.Name = Name
	e.NameLower = st.ToLower(Name)
	e.TableName = TableName
	e.FieldNames = FieldNames
	e.AllFieldNames = "Id, " + FieldNames
	e.NFields = ut.CountChars(FieldNames, ',') + 1
	e.FieldStr = fieldsString(e.NFields)
	e.AllFieldStr = e.FieldStr + ", $" + sc.Itoa(e.NFields+1)
}

// Name
func (e *EntityObject) GetName() (_ string) {
	return e.Name
}

// NameLower
func (e *EntityObject) GetNameLower() (_ string) {
	return e.NameLower
}

// Table Name
func (e *EntityObject) GetTableName() (_ string) {
	return e.TableName
}

// Field Names
func (e *EntityObject) GetFieldNames() (_ string) {
	return e.FieldNames
}

// All Field Names
func (e *EntityObject) GetAllFieldNames() (_ string) {
	return e.AllFieldNames
}

// Field String
func (e *EntityObject) GetFieldStr() (_ string) {
	return e.FieldStr
}

// Field String for all fields
func (e *EntityObject) GetAllFieldStr() (_ string) {
	return e.AllFieldStr
}

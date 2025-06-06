package controller

import (
	tp "html/template"

	sq "github.com/jackc/pgx/v5"
)

// A controller has an SQL connection and a template object.
type ControllerI interface {
	GetConn() *sq.Conn
	GetTmpl() *tp.Template
	Init(conn *sq.Conn, tmpl *tp.Template)
}

// GetConn returns the SQL connection.
func (c *ControllerT) GetConn() *sq.Conn {
	return c.conn
}

// GetTmpl returns the template object.
func (c *ControllerT) GetTmpl() *tp.Template {
	return c.tmpl
}

// Init initializes the controller with an SQL connection and a template object.
func (c *ControllerT) Init(conn *sq.Conn, tmpl *tp.Template) {
	c.conn = conn
	c.tmpl = tmpl
}

// ControllerT implements IController
type ControllerT struct {
	conn *sq.Conn
	tmpl *tp.Template
}

// FormControler is the controller for most of the forms.
type FormCrT struct {
	ControllerT
}

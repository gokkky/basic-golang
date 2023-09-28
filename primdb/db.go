package primdb

import "log"

type DatabaseNaja struct {
	storage []interface{}
}

func NewDatabase() *DatabaseNaja {
	log.Println("Connect database...")
	defer log.Println("Connection successfully...")
	return &DatabaseNaja{}
}

func (d *DatabaseNaja) Insert(data interface{}) bool {
	d.storage = append(d.storage, data)
	return true
}

func (d *DatabaseNaja) Get() interface{} {
	return d.storage
}

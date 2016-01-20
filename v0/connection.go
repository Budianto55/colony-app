package colonycore

import (
	"github.com/eaciit/orm/v1"
	"github.com/eaciit/toolkit"
)

type Connection struct {
	orm.ModelBase
	ID                               string `json:"_id",bson:"_id"`
	Driver, Host, UserName, Password string
	Settings                         toolkit.M
}

func (c *Connection) TableName() string {
	return "connections"
}

func (c *Connection) RecordID() interface{} {
	return c.ID
}

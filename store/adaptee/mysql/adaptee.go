package mysql

import (
	"github.com/danishk121/go-frame/store/adapter"
	"gorm.io/gorm"
)

// NewAdapter returns store mongodb adapter(*Client)
func NewAdapter(dbConn *gorm.DB) adapter.Store {
	c := &Client{
		DB: dbConn,
	}

	return mySqlAdapter(c)
}

func mySqlAdapter(c *Client) *Client {
	c.LOConn = NewLoStore(c)
	return c
}

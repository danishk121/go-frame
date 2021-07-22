package mysql

import (
	"github.com/danishk121/go-frame/store/adapter"
	"gorm.io/gorm"
)

type Client struct {
	DB     *gorm.DB
	LOConn adapter.LO
}

func (c *Client) LO() adapter.LO {
	return c.LOConn
}

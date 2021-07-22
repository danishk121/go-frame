package adapter

import "github.com/danishk121/go-frame/model"

type Store interface {
	LO() LO
}

type LO interface {
	Create(lo *model.LO) (*model.LO, error)
	Get() (*model.LO, error)
}

package service

import (
	"log"

	"github.com/danishk121/go-frame/model"
	"github.com/danishk121/go-frame/store"
	"github.com/danishk121/go-frame/store/adapter"
)

type LOService struct {
	l     *log.Logger
	store adapter.Store
}

func NewLOService(l *log.Logger) *LOService {
	return &LOService{
		l:     l,
		store: store.Store,
	}

}

func (p *LOService) CreateLOEntry(data model.LO) (*model.LO, error) {
	lo, err := p.store.LO().Create(&data)
	if err != nil {
		return nil, err
	}
	return lo, nil
}

func (p *LOService) GetAllData() (*model.LO, error) {

	lo, err := p.store.LO().Get()
	if err != nil {
		return nil, err
	}
	return lo, nil
}

package service

import (
	"log"
	"strings"

	"github.com/danishk121/go-frame/model"
	"github.com/danishk121/go-frame/repository"
	"gorm.io/gorm"
)

type LOService struct {
	l *log.Logger
	d *gorm.DB
}

func NewLOService(l *log.Logger, d *gorm.DB) *LOService {
	repository.Setup(d)
	return &LOService{l, d}
}

func (p *LOService) CreateLOEntry(data model.LO) {
	model := &repository.LOData{
		Name:        data.Name,
		Description: data.Description,
		Code:        data.Code,
		Applicable:  strings.Join(data.ApplicableClasses, ","),
	}

	repository.Add(model, p.d)

}

func (p *LOService) GetAllData() []repository.LOData {
	var response []repository.LOData
	repository.GetAll(p.d, &response)
	return response
}

package mysql

import "github.com/danishk121/go-frame/model"

type LO struct {
	*Client
}

// LO ...
func NewLoStore(client *Client) *LO {
	return &LO{client}
}

func (l *LO) Create(lo *model.LO) (*model.LO, error) {
	res := l.DB.Create(lo)
	if res.Error != nil {
		return nil, res.Error
	}
	return lo, nil
}

func (l *LO) Get() (*model.LO, error) {
	var lo model.LO
	res := l.DB.Find(&lo)
	if res.Error != nil {
		return nil, res.Error
	}
	return &lo, nil
}

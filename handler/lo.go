package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/danishk121/go-frame/model"
	"github.com/danishk121/go-frame/service"
)

type LO struct {
	l *log.Logger
	s *service.LOService
}

func NewLO(l *log.Logger, s *service.LOService) *LO {
	return &LO{l, s}
}

func (p *LO) AddLO(rw http.ResponseWriter, r *http.Request) {

	data := r.Context().Value(KeyLO{}).(model.LO)
	fmt.Print(data)
	p.s.CreateLOEntry(data)
	p.l.Println("Handle POST Product")

}

func (p *LO) GetLO(rw http.ResponseWriter, r *http.Request) {
	lp, _ := p.s.GetAllData()
	e := json.NewEncoder(rw)
	err := e.Encode(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

}

type KeyLO struct{}

func (p LO) MiddlewareValidateProduct(next http.Handler) http.Handler {
	fmt.Print("inside router")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := model.LO{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading data", http.StatusBadRequest)
			return
		}

		// validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyLO{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}

package model

import (
	"encoding/json"
	"io"
	"log"
	"regexp"

	"github.com/go-playground/validator"
)

type LO struct {
	Code              string   `json:"code" validate:"required"`
	Name              string   `json:"name" validate:"required"`
	Description       string   `json:"description" validate:"required"`
	ApplicableClasses []string `json:"applicableClasses" validate:"required"`
}

func (p *LO) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *LO) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("sku", validateSKU)
	if err != nil {
		//ignore
		log.Print("Error while register")
	}

	return validate.Struct(p)
}

//dummy regex custom validator
func validateSKU(fl validator.FieldLevel) bool {
	// sku is of format abc-absd-dfsdf
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1

}

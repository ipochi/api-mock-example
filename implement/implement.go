package implement

import (
	"github.com/ipochi/api-mock-example/db"
	"github.com/ipochi/api-mock-example/model"
)

type Implementor struct{}

func (impl *Implementor) GetCompanies() ([]model.Company, error) {

	companies := db.FindAll()

	return companies, nil
}

package model

type Functions interface {
	GetCompanies() ([]Company, error)
}

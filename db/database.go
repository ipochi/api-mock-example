package db

import "github.com/ipochi/api-mock-example/model"

var database = make(map[string]model.Company)

func FindAll() []model.Company {
	items := make([]model.Company, 0, len(database))
	for _, v := range database {
		items = append(items, v)
	}

	return items
}

func FindBy(key string) (model.Company, bool) {
	com, ok := database[key]

	return com, ok
}

func Save(key string, item model.Company) {
	database[key] = item
}

func Remove(key string) {
	delete(database, key)
}

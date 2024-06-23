package repositories

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalStock/domain/queries_category"
)

func UpdateStockRepositories(body string, user string, pathParams int) (int, string) {
	status := 200
	response := "Vacio Update Stock Repositories"

	status, response = queries_category.UpdateStockQuery(body, user, pathParams)

	return status, response
}
package use_cases

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalStock/domain/repositories"
)

func UpdateStockUC(body string, user string, pathParams int) (int, string) {
	status := 200
	response := "Vacio Update Stock Use Case"

	status, response = repositories.UpdateStockRepositories(body, user, pathParams)

	return status, response
}
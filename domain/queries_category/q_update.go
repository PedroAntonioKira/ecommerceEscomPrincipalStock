package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
	"strconv"
	"fmt"

	//importaciones externas (descargadas)
//	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalStock/domain/entities"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalStock/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalStock/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func UpdateStockQuery(body string, User string, pathParams int) (int, string) {
	
	var t entities.Product

	//Decodificamos el json que nos mandan en el endpoint en la estructura del producto para poder guardarla.
	err := json.Unmarshal([]byte(body), &t)

	//Verificamos que no tengamos un error al decodificar la información en la estructura.
	if err != nil {
		return 400, "Error en los datos recibidos con el error: " + err.Error()
	}

	//Verificamos si User Is Admin
	isAdmin, msg := secundary.UserIsAdmin(User)

	//Verificamos si efectivamente no es admin
	if !isAdmin {
		return 400, msg
	}

	// el id del producto lo asignamos que nos pasan como parametro
	t.ProdId = pathParams

	fmt.Println("Imprimimos el valor de pathparams antes de entrar a BD")
	fmt.Println(t.ProdId)

	fmt.Println("Imprimimos el valor de stock antes de entrar a BD")
	fmt.Println(t.ProdStock)

	//Actualizamos el Producto.
	err2 := database.UpdateStockDatabase(t)

	//Verificamos no exista un error al momento en que actualizamos el producto.
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el UPDATE del stock" + strconv.Itoa(pathParams) + " > " + err2.Error()
	}

	return 200, "Update OK Producto"

}

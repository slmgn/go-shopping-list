# go-shopping-list

## package

*	github.com/gorilla/mux v1.7.4
*	github.com/mattn/go-sqlite3 v1.14.0

### Descripcion

Se realiza un CRUD para productos y los endpoints expuestos en el puerto 8080 son:

* GET /api/produtcs
* GET /api/product/{id}
* POST /api/product
* PUT /api/product/{id}
* DELETE /api/product/{id}

Se deja la collecion para realizar la prueba de los endoints. 
Para hacer uso de la collecion solo se debe importar la collecion en Postman.

### Despliegue

Desde la raiz del proyecto se debe crear la imagen con el siguiente comando

    docker build -t go-shopping-list .
    
  y luego ejecutar

    docker run -dp 8080:8080 go-shopping-list

 Luego de esto ya se puede hacer uso de la API.
 

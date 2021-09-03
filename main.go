package main

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

func main() {

    //Conecta no DB
    db, err := sql.Open("mysql", "root:@tcp(0.0.0.0:3306)/pizzaria")

    //Pega erros
    if err != nil {
        panic(err)
    }

    //Executa os comandos
    result, err := db.Query("SELECT * FROM pizzas")

    //Pega erros
    if err != nil {
        panic(err)
    }

    //Retorna os resultados
    for result.Next() {

        var id int
        var nome string
        var preco int

        err = result.Scan(&id, &nome, &preco)

        if err != nil {
            panic(err)
        }

        fmt.Printf("Id: %d | Nome: %s | Preço: %d\n", id, nome, preco)
    }

    defer db.Close()
}

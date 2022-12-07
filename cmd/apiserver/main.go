package main

import (
	//	"fmt"

	"flag"

	"github.com/AisanaOlpan/GoProject/internal/app/apiserver"
	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}
func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		//log.Fatal(err)
	}
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		//log.Fatal(err)
	}

	//Lab_7_8()
	////<--------------->
	// connStr := "host=localhost dbname=postgresAis sslmode=disable"
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// result, err := db.Exec("insert into Products (model, company) values ('iPhone 11', $1)",
	// 	"Apple")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result.LastInsertId()) // не поддерживается
	// fmt.Println(result.RowsAffected()) // количество добавленных строк
	//<--------------->
	// var (
	// 	company string
	// 	model   string
	// )
	// connStr := "host=localhost dbname=postgresAis sslmode=disable"
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	panic(err)
	// }

	// result, err := db.Query("Select company, model from products")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer result.Close()

	// for result.Next() {
	// 	err := result.Scan(&company, &model)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(company, model)
	// }

	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	//<--------------->
}

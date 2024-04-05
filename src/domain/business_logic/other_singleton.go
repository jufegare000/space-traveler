package business_logic

import "fmt"

type Dao struct {
}

var daoInstance *Dao

func GetDao() *Dao {

	if daoInstance == nil {
		daoInstance = &Dao{}
		fmt.Println("Creando instancia")
	} else {
		fmt.Println("Creando instancia")
	}

	return daoInstance
}

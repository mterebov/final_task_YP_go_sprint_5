package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// На случай пустого датасета была добавлена проверка, иначе не проходили тесты
	if len(dataset) == 0{
		return
	}
	for _, v := range dataset {
		errPars := dp.Parse(v)
		if errPars != nil {
			log.Println(errPars)
		}
	}
	outStr, errActionInfo := dp.ActionInfo()
	if errActionInfo != nil {
		log.Println(errActionInfo)
		return
	}
	fmt.Println(outStr)
}

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

	for _, v := range dataset {
		errPars := dp.Parse(v)
		if errPars != nil {
			log.Println(errPars)
			continue
		}
		outStr, errActionInfo := dp.ActionInfo()
		if errActionInfo != nil {
			log.Println(errActionInfo)
			continue
		}
		fmt.Println(outStr)
	}
}

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
	var info string
	var err error
	for _, v := range dataset {
		err = dp.Parse(v)
		if err != nil {
			log.Println(err)
			err = nil
		}

		info, err = dp.ActionInfo()
		if err != nil {
			log.Println(err)
			err = nil
		} else {
			fmt.Print(info)
		}
	}
}

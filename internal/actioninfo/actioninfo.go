package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Error parsing data: %v", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Error getting action info: %v", err)
			continue
		}
		fmt.Print(info)
	}
}

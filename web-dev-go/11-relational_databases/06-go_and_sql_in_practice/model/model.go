package model

import "log"

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}

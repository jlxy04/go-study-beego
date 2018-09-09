package common

import (
	"github.com/satori/go.uuid"
	"log"
	"strings"
	)

func GetUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	uuidStr := strings.Replace(uuid.String(), "-", "",len(uuid.String()))
	return uuidStr
}
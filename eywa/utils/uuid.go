package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func UUID2Byte() []byte {
	uuidValue := uuid.Must(uuid.NewRandom())
	fmt.Printf("uuid:%v\n", uuidValue)
	return uuidValue[:]
}

func UUID() uuid.UUID {
	return uuid.Must(uuid.NewRandom())
}

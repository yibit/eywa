package utils

import (
	"github.com/google/uuid"
)

func UUID2Byte() []byte {
	uuidValue := uuid.Must(uuid.NewRandom())
	return uuidValue[:]
}

func UUID() uuid.UUID {
	return uuid.Must(uuid.NewRandom())
}

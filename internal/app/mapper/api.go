package mapper

import (
	"github.com/thirteenths/test-bmstu23/internal/domain/responses"
)

func CreateEcho() *responses.GetEcho {
	return &responses.GetEcho{
		Message: "echo",
	}
}

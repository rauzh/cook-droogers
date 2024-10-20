package common

import (
	modelsDTO "cookdroogers/internal/server/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

func ErrorResponse(status int, err string) middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(status)
		_ = p.Produce(rw, modelsDTO.LeErrorMessage{
			Error: err,
		})
	})
}

package swagger

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupSwaggerRouter() {
	http.Handle("/swagger/", httpSwagger.WrapHandler)
}

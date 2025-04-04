package middlewares

import (
	"github.com/AndresFWilT/afwt-clean-go-crud-test/internal/usecase/authorize/tokens"
	"net/http"

	"github.com/AndresFWilT/afwt-clean-go-crud-test/internal/adapters/handlers"
)

func Authorization(f handlers.FunctionType) handlers.FunctionType {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := tokens.ValidateToken(token)
		if err != nil {
			handlers.Forbidden(w, r)
			return
		}

		f(w, r)
	}
}

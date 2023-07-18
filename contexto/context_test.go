package contexto

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	data := "olá, mundo"

	t.Run("retorna dados da store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("resultado '%s', esperado '%s'", response.Body.String(), data)
		}
	})

	// t.Run("avisa a store para cancelar o trabalho se a requisição for cancelada", func(t *testing.T) {
	// 	store := &SpyStore{response: data, t: t}
	// 	svr := Server(store)

	// 	request := httptest.NewRequest(http.MethodGet, "/", nil)

	// 	cancellingCtx, cancel := context.WithCancel(request.Context())
	// 	time.AfterFunc(5*time.Millisecond, cancel)
	// 	request = request.WithContext(cancellingCtx)

	// 	response := httptest.NewRecorder()

	// 	svr.ServeHTTP(response, request)

	// 	store.assertWasCancelled()
	// })
}

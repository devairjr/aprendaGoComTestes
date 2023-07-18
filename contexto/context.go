package contexto

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := store.Fetch(r.Context())

		// if err != nil {
		// 	return // todo: registre o erro como vc quiser
		// }

		fmt.Fprint(w, data)
	}
}

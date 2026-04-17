package main

import (
	"net/http"

	httpHandler "go-crud/internal/infra/http"
	"go-crud/internal/infra/memory"
	"go-crud/internal/usecase"
)

func main() {
	repo := memory.NewUserMemoryRepo()
	userUsecase := usecase.NewUserUsecase(repo)
	handler := httpHandler.NewUserHandler(userUsecase)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API rodando")) // ✅ escreve resposta
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.Create(w, r)
			return
		}
		if r.Method == http.MethodGet {
			handler.List(w, r)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	http.ListenAndServe(":8080", nil)
}
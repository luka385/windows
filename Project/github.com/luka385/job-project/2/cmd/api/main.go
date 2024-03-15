package main

import (
	"log"

	"github.com/luka385/job-project/2/internal/adapters/handler"
	"github.com/luka385/job-project/2/internal/adapters/repository"
	"github.com/luka385/job-project/2/internal/platform/web"
	"github.com/luka385/job-project/2/internal/usecase"
)

func main() {
	r := repository.NewRepository()
	u := usecase.NewItemUsecase(r)
	h := handler.NewHandler(u)

	err := web.NewHTTPServer(h)
	if err != nil {
		log.Fatalln(err)
	}
}

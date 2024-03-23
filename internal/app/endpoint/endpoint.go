package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Service interface {
	DaysLeft() int
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {

	d := e.s.DaysLeft()
	//Формируем строку ответа
	s := fmt.Sprintf("Days left - %d", d)
	//Отправляем строковый ответ, код состояни
	err := ctx.String(http.StatusOK, s)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

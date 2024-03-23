package service

import "time"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int {
	//Создаем объект времени (1 января 2025 г.)
	d := time.Date(2025, time.January, 1, 1, 1, 1, 1, time.UTC)
	//Вычисляем продолжительность между текущим временем и указанным в "d"
	dur := time.Until(d)

	return int(dur.Hours()) / 24
}

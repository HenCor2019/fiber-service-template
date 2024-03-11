package services

type HealthCheckService interface {
	CheckHealth() string
}

type Service struct {
}

func New() HealthCheckService {
	return &Service{}
}

func (s *Service) CheckHealth() string {
	return "ok"
}

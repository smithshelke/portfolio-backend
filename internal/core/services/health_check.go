package services

type HealthCheckService struct{}

func NewHealthCheckService() *HealthCheckService {
	return &HealthCheckService{}
}

func (s *HealthCheckService) CheckHealth() string {
	return "OK"
}

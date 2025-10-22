package ports

type HealthCheckService interface {
	CheckHealth() string
}

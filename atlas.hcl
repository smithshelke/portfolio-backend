env "dev" {
  url = "postgres://user:password@db:5432/shelke_dev_api?sslmode=disable"
  migration {
    dir = "file://db/migrations"
  }
}

env "prod" {
  migration {
    dir = "file://db/migrations"
  }
}
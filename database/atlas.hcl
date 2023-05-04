env "local" {
  # the real database
  url = postgres://admin:admin@localhost:5434/atlas_db?sslmode=disable
  # a temporary database for atlas compute diff in states
  dev = postgres://admin:admin@localhost:5434/dev?sslmode=disable
}

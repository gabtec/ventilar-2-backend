# Extract db schema from existing local dev database

```sh
$ atlas schema inspect --env local --format '{{ sql . }}' > migrations/20230503125301_init.sql
# or
$ atlas schema inspect -u "postgres://admin:admin@localhost:5432/atlas_db?sslmode=disable" --format '{{ sql . }}' > migrations/20230503125301_init.sql

# create a checksum file
$ atlas migrate hash [--dir "file://migrations"]

# check status (we need a sum file)
$ atlas migrate status -u "postgres://admin:admin@localhost:5432/atlas_db?sslmode=disable" --dir file://automations/atlas/migrations

# apply migrations
$ atlas migrate apply -u "postgres://admin:admin@localhost:5432/atlas_db?sslmode=disable" --dir file://automations/atlas/migrations [N]
$ atlas migrate apply --env local [N]

# check diff between migrations/files and desired state (--to, which can be a db url or a schema file)
# A) against a real db
$ atlas migrate diff --dev-url "postgres://admin:admin@localhost:5432/temp?sslmode=disable" --to "postgres://admin:admin@localhost:5432/atlas_db?sslmode=disable" --dir file://automations/atlas/migrations --schema public --format '{{ sql . "  " }}' new_mig_file_name2

# B) against a schema.hcl file
$ atlas migrate diff --dev-url "postgres://admin:admin@localhost:5432/temp?sslmode=disable" --to file://automations/atlas/schema.hcl --dir file://automations/atlas/migrations  --format '{{ sql . "  " }}' new_mig_file_name

```

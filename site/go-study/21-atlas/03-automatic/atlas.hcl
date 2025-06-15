env "local" {
    name = "local"
    src = "./schema.sql"
    url = "sqlite3://my-automatic-migration.db"
    dev = "sqlite3://my-automatic-migration-dev.db"
}
env "local" {
    name = "local"
    src = "./schema.sql"
    url = "sqlite3://my-tf.db"
    dev = "sqlite3://my-tf-dev.db"
}
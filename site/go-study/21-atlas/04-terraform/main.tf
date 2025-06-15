provider "atlas" {
  dev_url = "sqlite3://my-tf-dev.db"
}

// Load (and normalize) the desired schema from an HCL file.
data "atlas_schema" "myapp" {
  src = "file://${path.module}/schema.sql"
  dev_url = "sqlite3://my-tf-dev.db"
}

// Sync the state of the target database with the hcl file.
resource "atlas_schema" "myapp" {
  url = "sqlite3://my-tf.db"
  hcl = data.atlas_schema.myapp.hcl
}
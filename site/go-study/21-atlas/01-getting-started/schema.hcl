table "customers" {
  schema = schema.main
  column "id" {
    null = true
    type = integer
  }
  column "first_name" {
    null = false
    type = text
  }
  column "last_name" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
}
schema "main" {
}

table "orders" {
  schema = schema.main
  column "id" {
    null = true
    type = integer
  }
  column "customer_id" {
    null = false
    type = integer
  }
  column "order_date" {
    null = false
    type = date
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk1" {
    columns     = [column.customer_id]
    ref_columns = [table.customers.column.id]
  }
}
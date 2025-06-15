package main

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

// Examples:
// 1. build SELECT query
// 2. build SELECT query with a WHERE clause
// 3. build INSERT query
// Print out the SQL and arguments for each query.
func main() {
	users := sq.Select("*").From("users").Join("emails USING (email_id)")
	sql, _, _ := users.ToSql()
	fmt.Printf("SQL: %s\n", sql)

	active := users.Where(sq.Eq{"deleted_at": nil})
	sql, args, err := active.ToSql()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("ARGS: %s\n", args)
	fmt.Printf("SQL: %s\n", sql)
	// ARGS: []
	// SQL: SELECT * FROM users JOIN emails USING (email_id) WHERE deleted_at IS NULL

	sql, args, err = sq.
		Insert("users").Columns("name", "age").
		Values("moe", 13).Values("larry", sq.Expr("? + 5", 12)).
		ToSql()

	fmt.Printf("ARGS: %s\n", args)
	fmt.Printf("SQL: %s\n", sql)
	// ARGS: [moe %!s(int=13) larry %!s(int=12)]
	// SQL: INSERT INTO users (name,age) VALUES (?,?),(?,? + 5)

	//stooges := users.Where(sq.Eq{"username": []string{"moe", "larry", "curly", "shemp"}})
	//three_stooges := stooges.Limit(3)
	//rows, err := three_stooges.RunWith(db).Query()

}

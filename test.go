package main

import "fmt"
type DBQuery struct {
    SQL string
    Retries int
}
func executeQuery(q string) {
    fmt.Println("Executing: " + q)
}
func main() {
    // /* sql */
    // Triggered: short declaration
    /* sql */
	query1 := `SELECT * FROM users WHERE id = 1`
	fmt.Println(query1)

    // Not Triggered: No comment
	query2 := `SELECT id, name FROM products`
	fmt.Println(query2)

    // Triggered: var declaration
    /* sql */
    var query3 string = "INSERT INTO logs (data) VALUES ('test')"
    fmt.Println(query3)

    config := DBQuery{
            /* sql */
            SQL: `UPDATE settings SET value = 'true' WHERE key = 'enabled'`, // SQL should be highlighted here
            Retries: 3,
        }

    fmt.Println(config)


    executeQuery(
            `DELETE FROM cache WHERE expiry < NOW()`, 
        )
}

package main

type DBQuery struct {
    SQL string
}

func executeQuery(q string) string {
    return q
}

func main() {
    // const assignment
    const _ =  /* sql */ "SELECT * FROM users"
    const _ = /* sql */ `SELECT id, name FROM products`

	// := assignment
	test := /* sql */ "SELECT * FROM users"
	println(test)

	// = assignment
    _ =  /* sql */ "SELECT * FROM users WHERE id = 1"
	_ =  /* sql */ `SELECT * FROM users WHERE id = 1`

	// literal elements
	_ = DBQuery{ SQL: /* sql */ "SELECT * FROM users"}
    _ = DBQuery{ SQL: /* sql */ `SELECT * FROM users`}

    executeQuery(/* sql */ "SELECT * FROM users")
    executeQuery(/* sql */ `SELECT * FROM users`)
    }

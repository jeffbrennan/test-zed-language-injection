func test_sql() {
    // const assignment
    const _ =  /* sql */ "SELECT * FROM users"
    const _ = /* sql */ `SELECT id, name FROM products`

    // var assignment
    var _ = /* sql */ `SELECT id, name FROM products`
    var _ = /* sql */ "SELECT id, name FROM products"

	// := assignment
	test := /* sql */ "SELECT * FROM users"
	test2 := /* sql */ `SELECT * FROM users`
	println(test)
	println(test2)

	// = assignment
    _ =  /* sql */ "SELECT * FROM users WHERE id = 1"
	_ =  /* sql */ `SELECT * FROM users WHERE id = 1`

	// literal elements
	_ = testStruct{ Field: /* sql */ "SELECT * FROM users"}
    _ = testStruct{ Field: /* sql */ `SELECT * FROM users`}

    testFunc(/* sql */ "SELECT * FROM users")
    testFunc(/* sql */ `SELECT * FROM users`)

    const backtickString = /* sql */ `SELECT * FROM users;` // it's working with backticks
    const quotedString = /* sql */ "SELECT * FROM users;"  // it's not working with quotes

    const backtickStringNoHighlight = `SELECT * FROM users;`
    const quotedStringNoHighlight = "SELECT * FROM users;"
}

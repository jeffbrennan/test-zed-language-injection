package main


type testStruct struct { Field string }
func testFunc(s string) {}

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

func test_json() {
    // const assignment
    const _ = /* json */ "{\"test\": 1234}"
    const _ = /* json */ `{
        "id": 123,
        "name": "Product Name"
    }`

    // := assignment
    test := /* json */ "{\"test\": 1234}"
    test2 := /* json */ `{"test": 1234}`
    println(test)
    println(test2)

    // = assignment
    _ = /* json */ "{\"data\": [1, 2, 3]}"
    _ = /* json */ `[
        {
            "user_id": 42
        }
    ]`

    // literal elements in a struct
    _ = testStruct{ Field: /* json */ "{\"result\": true}"}
    _ = testStruct{ Field: /* json */ `{
        "errors": []
    }`}

    // function argument
    testFunc(/* json */ "{\"test\": 1234}")
    testFunc(/* json */ `{"test": 1234}`)
}


func test_yaml() {
    // const assignment
    const _ = /* yaml */ `
        settings:
            enabled: true
            port: 8080
    `

    // := assignment
    test := /* yaml */ `
        settings:
            enabled: true
            port: 8080
    `
    println(test)

    // = assignment
    _ = /* yaml */ `
        settings:
            enabled: true
            port: 8080
    `

    // literal elements in a struct
    _ = testStruct{ Field: /* yaml */ `
        settings:
            test: 1234
            port: 8080
    `}

    // function argument
    testFunc(/* yaml */ `
        settings:
            enabled: true
            port: 8080
    `)
}

func test_xml() {
    //const assignment
    const _ = /* xml */ `
        <settings>
            <enabled>true</enabled>
            <port>8080</port>
        </settings>
    `

    // := assignment
    test := /* xml */ `
        <settings>
            <enabled>true</enabled>
            <port>8080</port>
        </settings>
    `
    println(test)

    // = assignment
    _ = /* xml */ `
    <settings>
        <enabled>true</enabled>
        <port>8080</port>
    </settings>
    `


    // literal elements in a struct
    _ = testStruct{
        Field: /* xml */ `
        <settings>
            <enabled>true</enabled>
            <port>8080</port>
        </settings>
    `}

    _ = testStruct{
        Field: /* xml */ "<item><price>9.99</price></item>",
    }


    testFunc(/* xml */ `
        <request>
            <action>update</action>
        </request>
    `)
}

func test_html() {
    //const assignment
    const _ = /* html */ `
        <html>
            <head><title>Test</title></head>
            <body><h1>Hello</h1></body>
        </html>
    `

    // := assignment
    test := /* html */ `
        <div>
            <p>This is a paragraph.</p>
            <span>Note</span>
        </div>
    `
    println(test)

    // = assignment
    _ = /* html */ `
    <article>
        <h2>Title</h2>
        <p>Content goes here.</p>
    </article>
    `


    // literal elements in a struct
    _ = testStruct{
        Field: /* html */ `
        <form action="/submit">
            <input type="text" name="data">
        </form>
    `}

    _ = testStruct{
        Field: /* html */ "<button disabled>Click me</button>",
    }


    // function argument
    // (call_expression, raw_string_literal)
    testFunc(/* html */ `
        <nav>
            <a href="/home">Home</a>
            <a href="/about">About</a>
        </nav>
    `)
}

func test_javascript() {
    // const assignment
    const _ =  /* js */ "console.log('hello world')"
    const _ = /* js */ `console.log('hello world')`

	// := assignment
	test := /* js */ "console.log('hello world')"
	println(test)

	// = assignment
    _ =  /* js */ "console.log('hello world')"
	_ =  /* js */ `console.log('hello world')`

	// literal elements
	_ = testStruct{ Field: /* js */ "console.log('hello world')"}
    _ = testStruct{ Field: /* js */ `console.log('hello world')`}

    testFunc(/* js */ "console.log('hello world')")
    testFunc(/* js */ `console.log('hello world')`)

    const backtickString = /* js */ `console.log('hello world')` // it's working with backticks
    const quotedString = /* js */ "console.log('hello world')"  // it's not working with quotes

    const backtickStringNoHighlight = `console.log('hello world')`
    const quotedStringNoHighlight = "console.log('hello world')"

}

func test_css() {
    // const assignment
    const _ =  /* css */ "body { margin: 0; }"
    const _ = /* css */ `body { margin: 0; }`

	// := assignment
	test := /* css */ "body { margin: 0; }"
	println(test)

	// = assignment
    _ =  /* css */ "body { margin: 0; }"
	_ =  /* css */ `body { margin: 0; }`

	// literal elements
	_ = testStruct{ Field: /* css */ "body { margin: 0; }"}
    _ = testStruct{ Field: /* css */ `body { margin: 0; }`}

    testFunc(/* css */ "body { margin: 0; }")
    testFunc(/* css */ `body { margin: 0; }`)

    const backtickString = /* css */ `body { margin: 0; }` // it's working with backticks
    const quotedString = /* css */ "body { margin: 0; }"  // it's not working with quotes

    const backtickStringNoHighlight = `body { margin: 0; }`
    const quotedStringNoHighlight = "body { margin: 0; }"

}

func test_lua() {
    // const assignment
    const _ =  /* lua */ "local x = 42"
    const _ = /* lua */ `local x = 42`

	// := assignment
	test := /* lua */ "local x = 42"
	println(test)

	// = assignment
    _ =  /* lua */ "local x = 42"
	_ =  /* lua */ `local x = 42`

	// literal elements
	_ = testStruct{ Field: /* lua */ "local x = 42"}
    _ = testStruct{ Field: /* lua */ `local x = 42`}

    testFunc(/* lua */ "local x = 42")
    testFunc(/* lua */ `local x = 42`)

    const backtickString = /* lua */ `local x = 42` // it's working with backticks
    const quotedString = /* lua */ "local x = 42"  // it's not working with quotes

    const backtickStringNoHighlight = `local x = 42`
    const quotedStringNoHighlight = "local x = 42"

}

func test_bash() {
    // const assignment
    const _ =  /* bash */ "echo 'hello world'"
    const _ = /* bash */ `echo 'hello world'`

	// := assignment
	test := /* bash */ "echo 'hello world'"
	println(test)

	// = assignment
    _ =  /* bash */ "echo 'hello world'"
	_ =  /* bash */ `echo 'hello world'`

	// literal elements
	_ = testStruct{ Field: /* bash */ "echo 'hello world'"}
    _ = testStruct{ Field: /* bash */ `echo 'hello world'`}

    testFunc(/* bash */ "echo 'hello world'")
    testFunc(/* bash */ `echo 'hello world'`)

    const backtickString = /* bash */ `echo 'hello world'` // it's working with backticks
    const quotedString = /* bash */ "echo 'hello world'"  // it's not working with quotes

    const backtickStringNoHighlight = `echo 'hello world'`
    const quotedStringNoHighlight = "echo 'hello world'"

}


func test_csv() {
    // const assignment
    const _ =  /* csv */ "1,hello,2,world"
    const _ = /* csv */ `
        col1,col2,col3,col4
        1,hello,2,world
    `

	// := assignment
	test := /* csv */ "1,hello,2,world"
	println(test)

	// = assignment
    _ =  /* csv */ "1,hello,2,world"
	_ =  /* csv */ `1,hello,2,world`

	// literal elements
	_ = testStruct{ Field: /* csv */ "1,hello,2,world"}
    _ = testStruct{ Field: /* csv */ `1,hello,2,world`}

    testFunc(/* csv */ "1,hello,2,world")
    testFunc(/* csv */ `1,hello,2,world`)

    const backtickString = /* csv */ `1,hello,2,world` // it's working with backticks
    const quotedString = /* csv */ "1,hello,2,world"  // it's not working with quotes

    const backtickStringNoHighlight = `1,hello,2,world`
    const quotedStringNoHighlight = "1,hello,2,world"

}

func main() {
    test_sql()
    test_json()
    test_yaml()
    test_xml()
    test_html()
    test_javascript()
    test_css()
    test_lua()
    test_bash()
    test_csv()
}

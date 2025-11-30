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

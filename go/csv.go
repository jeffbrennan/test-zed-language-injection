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

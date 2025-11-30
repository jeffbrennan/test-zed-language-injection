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

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

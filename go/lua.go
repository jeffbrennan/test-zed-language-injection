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

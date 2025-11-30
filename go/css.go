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

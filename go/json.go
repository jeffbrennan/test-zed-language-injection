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

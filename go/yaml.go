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

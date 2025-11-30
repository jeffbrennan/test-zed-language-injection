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

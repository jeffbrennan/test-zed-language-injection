package main

func test() {
    var _ = /* sql */ "SELECT id, name FROM products"
    var _ = /* sql */ "SELECT id, name FROM products"

    var _ = /* sql */ `SELECT id, name FROM products`

    var _ = /* sql */ "SELECT id, name FROM products"
}

package main

type Record struct {
	ContactID string
	Name      string
	Name1     string
	Email     string
	PostalZip string
	Address   string
}

type Response struct {
	ContactIDSource string
	ContactIDMatch  string
	Acurracy        int
}

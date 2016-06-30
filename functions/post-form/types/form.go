package types

type Form struct {
	ID          string // sort/range key
	Email       string // secoundary sort/range key
	Origin      string // primary key
	Secret      string // the user and I kows this secret
	Verifyed    bool
	Subscribing bool
}
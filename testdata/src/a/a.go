package a

type Url string // want "Url should be URL"

type T struct {
	Url string // want "Url should be URL"
}

func f() {
	// The pattern can be written in regular expression.
	var Url string // want "Url should be URL"
	print(Url)     // want "Url should be URL"
}

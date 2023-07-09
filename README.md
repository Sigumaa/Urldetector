# Urldetector


## Install

You can install this tool like this...There may be other ways to do it...

```bash
go install github.com/Sigumaa/Urldetector/cmd/Urldetector@latest
```

## Usage

```bash
go vet -vettool=`which Urldetector` your_package_name
```

## Example

Assume you are in testdata/src/a of this repository.

### Code

```go
package a

type Url string

type T struct {
	Url string
}

func f() {
	var Url string
	print(Url)
}
```

### Command

```bash
go vet -vettool=`which Urldetector` a
```

### Output

```bash
a.go:3:6: Url should be URL
a.go:6:2: Url should be URL
a.go:10:6: Url should be URL
a.go:11:8: Url should be URL
```

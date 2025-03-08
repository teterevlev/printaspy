
# printAsPy

`printAsPy` is a library for printing text with customizable separators and line endings, as you used in python

## Installation

To use this library, add it to your project using `go get`:

```bash
go get github.com/teterevlev/printaspy
```

## Example Usage

Here is an example of how to use the library for printing text with different settings:

```go
package main

import (
	"github.com/teterevlev/printaspy"
)

func main() {
	// Simple example
	printaspy.Print(42, "Hello", 3.14, true, printaspy.WithSep(" | "), printaspy.WithEnd(" !!!\n"))
	// Output:
	// 42 | Hello | 3.14 | true !!!
    
    // To shorten
    print := printaspy.Print
    sep := printaspy.WithSep
    end := printaspy.WithEnd
	print(42, "Hello", 3.14, true, sep(" | "), end(" !!!\n"))
	// Output:
	// 42 | Hello | 3.14 | true !!!
}
```

## Functions

- **WithSep(sep string):** Sets the separator between elements.
- **WithEnd(end string):** Sets the line ending.

## License

MIT


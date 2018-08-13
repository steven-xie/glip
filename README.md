# glip

_A clipboard interface for Go, compatible with Windows, Mac OS X, and Linux._

[![godoc: reference][godoc-img]][godoc]

## Usage

Install `glip` like you would any other Go package:

```bash
go get github.com/steven-xie/glip
```

Since a `glip.Board` (a clipboard interface) implements `io.Writer`,
`io.Reader`, `io.WriterTo`, it can be used just about anywhere:

```go
import (
  "fmt"
  "github.com/steven-xie/glip"
)

func main() {
  // Save "snip snip" into the system clipboard.
  glip.WriteString("snip snip")
  // And we're done!

  // Read clipboard contents into a string (ignoring the error).
  out, _ := glip.ReadString()
  fmt.Println(out)
  // Output: snip snip
}
```

## Compatibility

### Windows:

`glip` uses the `paste` and `copy` commands on Windows to manipulate the
clipboard. This is available _starting from Windows 7, and onwards_.

### Mac OS X:

`glip` uses `pbcopy` and `pbpaste` commands on OS X; this has been available
since 2005, so no compatibility worries here.

### Linux:

`glip` requires the _installation of either `xclip` or `xsel`_ to function on
Linux (since there's no built-in clipboard interface). `glip` will choose
one of those two programs automatically, unless you build a custom board with
the `NewLinuxBoard` function.

[godoc]: https://godoc.org/github.com/steven-xie/glip
[godoc-img]: https://godoc.org/github.com/steven-xie/glip?status.svg

package portal_test

import (
	"github.com/stevenxie/glip/portal"
	"os"
)

const TestPhrase = "Hello Portal!"

func Example() {
	p := portal.New("echo", TestPhrase)
	p.WriteTo(os.Stdout)
	// Output: Hello Portal!
}

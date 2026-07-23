// Command vanity-proof is a trivial proof that chatwright.dev vanity import
// paths resolve via `go install`/`go get`. See
// chatwright/chatwright spec/plans/code-split-restructuring.md (Task 0).
package main

import "fmt"

func main() {
	fmt.Println("chatwright.dev/vanity-proof: vanity import path resolved and ran")
}

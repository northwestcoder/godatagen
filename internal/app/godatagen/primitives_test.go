// tests for primitive fake data types

package primitives

import (
	"testing"
)

func TestNameBundle(t *testing.T) {

	testAddress := NameBundle{"John", "Smith", "jsmith@example.com", "somewhere"}
	print("ok")
	print(testAddress.Email)
	return

}

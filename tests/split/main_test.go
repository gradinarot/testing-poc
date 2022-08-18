package compress_test

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("testing split")
	os.Exit(m.Run())
}

package tester

import "testing"

type TestCase interface {
	validate(t *testing.T)
}

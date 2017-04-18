package thrift

import (
	"testing"
)

func TestThrift(t *testing.T) {
	go ServerMain("127.0.0.1", "1090")
	go ClientMain("127.0.0.1", "1090")
}

package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	r := time.Now().Add(time.Second)

	fmt.Println(time.Now().Before(r))
}

package dpos

import (
	"fmt"
	"testing"
)

func TestBasicHost(t *testing.T) {

	h, err := MakeBasicHost(1000, false, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(h.ID().Pretty())
}

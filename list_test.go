package list

import (
  "testing"
  "github.com/stretchr/testify/assert"
  )

func TestShouldAddItemToList(t *testing.T) {
  myList := NewList()
  myList.Add("Buy milk")

  assert.Equal(t, "Buy milk", myList.Get(0), "They Should be equal")

}

package list

import (
  "testing"
  "github.com/stretchr/testify/assert"
  )

func TestShouldAddItemToList(t *testing.T) {
  myList := NewList()

  myList.Add("Buy milk")

  item, _ := myList.Get(0)

  assert.Equal(t, "Buy milk", item , "They Should be equal")

}

func TestShouldReturnListSize(t *testing.T) {
  myList := NewList()

  myList.Add("Remember to finish this package")
  myList.Add("Learn more about Vagrant/ansible/docker")
  myList.Add("finish integration with Travis CI")

  assert.Equal(t, 3, myList.Size(), "There should be only 3 item")
}

func TestShouldRaiseErrorWhenIndexIsEmpty(t *testing.T) {
  myList    := NewList()
  item, err := myList.Get(0)

  assert.Equal(t, "Index out of bound", err.Error(), "Error should have been raised")
  assert.Equal(t, nil, item, "No value found")
}

func TestShouldRaiseErrorWhenIndexBiggerThenSize(t *testing.T) {
  myList := NewList()

  myList.Add("Write about RubyConf")
  myList.Add("Study more about Tests")
  myList.Add("Learn more GoLang")

  firstItem, _    := myList.Get(0)
  secondItem, _   := myList.Get(1)
  thirdItem, _    := myList.Get(2)
  fourthItem, err := myList.Get(22)

  assert.Equal(t, "Write about RubyConf", firstItem , "Values should be Equal")
  assert.Equal(t, "Study more about Tests", secondItem , "Values should be Equal")
  assert.Equal(t, "Learn more GoLang", thirdItem , "Values should be Equal")

  assert.Nil(t, fourthItem)
  assert.Equal(t, "Index out of bound", err.Error(), "Error should have been raised")
}

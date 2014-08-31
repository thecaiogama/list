package list

import("errors")

type List struct {
  itens []interface{}
  size int
}

func (list *List) Add(item interface{}) {
  list.itens[list.size] = item
  list.size++
}

func (list *List) Get(e int) (item interface{}, err error) {
  if e > len(list.itens) || list.itens[e] == nil {
    return nil, errors.New("Index out of bound")
  }

  return list.itens[e], nil
}

func NewList() *List {
  return &List{itens: make([]interface{}, 10), size:0}
}

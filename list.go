package list

type List struct {
  itens []interface{}
  size int
}

func (list *List) Add(item interface{}) {
  list.itens[list.size] = item
  list.size++
}

func (list *List) Get(e int) interface{} {
  return list.itens[e]
}

func NewList() *List {
  return &List{itens: make([]interface{}, 10), size:0}
}

func (list *List) cara(i int) int {
  return i + 5
}

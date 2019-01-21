package algorithms

import (
	"errors"
	"fmt"
)

const MaxHashTable = 3

type employee struct {
	Id   int
	Name string
	Next *employee
}
type LinkTable struct {
	Head *employee
}

func (l *LinkTable) insert(id int, name string) error {
	em := &employee{
		id,
		name,
		nil,
	}
	if l.Head == nil {
		l.Head = em
		return nil
	}
	last := l.Head
	tmp := l.Head.Next
	for tmp != nil {
		if tmp.Id >= id {
			break
		}
		last = tmp
		tmp = tmp.Next
	}
	em.Next = tmp
	last.Next = em
	return nil
}
func (l *LinkTable) find(id int) (em *employee, err error) {
	tmp := l.Head.Next
	for tmp != nil {
		if tmp.Id == id {
			em = tmp
			return
		}
		tmp = tmp.Next
	}
	return nil, errors.New("not exist!!!")
}
func (l *LinkTable) delete(id int) {
	var last *employee
	tmp := l.Head
	for tmp != nil {
		if tmp.Id == id {
			break
		}
		last = tmp
		tmp = tmp.Next
	}
	//如果删除的是头结点，则让头节点指向头结点的下一个节点
	if last == nil {
		l.Head = tmp.Next
		return
	}
	last.Next = tmp.Next

}
func (l *LinkTable) list() {
	tmp := l.Head
	for tmp != nil {
		fmt.Printf("-->id:%d,name:%s\n", tmp.Id, tmp.Name)
		tmp = tmp.Next
	}
	fmt.Println("-------------------------------------")
}

type hashTable struct {
	table [MaxHashTable]*LinkTable
}

func NewHashTab() *hashTable {
	hashtable := &hashTable{}
	for i := 0; i < MaxHashTable; i++ {
		hashtable.table[i] = &LinkTable{}
	}
	return hashtable
}
func (h *hashTable) insert(id int, name string) error {
	lindIndex := id % MaxHashTable
	linkHead := h.table[lindIndex]
	return linkHead.insert(id, name)
}

func (h *hashTable) find(id int) (em *employee, err error) {
	lindIndex := id % MaxHashTable
	linkHead := h.table[lindIndex]
	return linkHead.find(id)
}

func (h *hashTable) delete(id int) {
	lindIndex := id % MaxHashTable
	linkHead := h.table[lindIndex]
	linkHead.delete(id)
}

func (h *hashTable) list() {
	for _, v := range h.table {
		if v != nil {
			v.list()
		}

	}
}

func hash() {
	tab := NewHashTab()

	tab.insert(2, "jack")
	tab.insert(5, "aaa")
	tab.insert(1, "bbb")
	tab.insert(6, "ccc")
	tab.insert(3, "ddd")
	tab.insert(4, "eee")
	tab.insert(7, "fff")
	tab.insert(8, "ggg")

	tab.list()

	em, err := tab.find(9)
	if err != nil {
		fmt.Println("find err:", err)
	} else {
		fmt.Println(*em)
	}

	tab.delete(1)
	tab.list()

}

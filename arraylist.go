package list

import (
	"errors"
	"fmt"
)

type ArrayList struct {
	values   []int
	inserted int // Esse daqui é a quantidade de valores da lista
}

func (list *ArrayList) Init(len int) error {
	if len > 0 {
		list.values = make([]int, len)
		return nil
	} else {
		return errors.New("size must be greater than 0")
	}

}

func (list *ArrayList) doubleCapacity() {

	doubleArray := make([]int, len(list.values)*2)
	//lint:ignore S1001 because i want
	for i := 0; i < len(list.values); i++ {
		doubleArray[i] = list.values[i]
	}

	list.values = doubleArray

}

func (list *ArrayList) Add(value int) {
	if list.inserted >= len(list.values) {
		list.doubleCapacity()
	}
	list.values[list.inserted] = value
	list.inserted++
}

func (list *ArrayList) AddOnIndex(value int, index int) error {
	if index >= 0 && index <= list.inserted {
		if list.inserted >= len(list.values) {
			list.doubleCapacity()
		}
		for i := list.inserted; i > index; i-- {
			list.values[i] = list.values[i-1]
		}
		list.values[index] = value
		list.inserted++
		return nil
	} else {
		return errors.New("index not accessible")
	}
}

func (list *ArrayList) RemoveLast() {
	if list.inserted > 0 {
		list.inserted--
	}
}

func (list *ArrayList) RemoveOnIndex(index int) error {
	if index >= 0 && index <= list.inserted {
		for i := index; i < (list.inserted - 1); i++ {
			list.values[i] = list.values[i+1]
		}
		list.inserted--
		return nil
	} else {
		return errors.New("index not accessible")
	}
}

func (list *ArrayList) Get(index int) (int, error) {
	if index >= 0 && index <= list.inserted {
		return list.values[index], nil
	} else {
		return index, errors.New("index not accessible")
	}
}

func (list *ArrayList) Update(value int, index int) error {
	if index >= 0 && index <= list.inserted {
		list.values[index] = value
		return nil
	} else {
		return errors.New("index not accessible")
	}
}

func (list *ArrayList) Size() int {
	return list.inserted
}

func (list *ArrayList) GetAll() {
	fmt.Println("Size: ", list.Size())
	for i := 0; i < list.inserted; i++ {
		fmt.Println(list.Get(i))
	}
}

func (list *ArrayList) Reverse() {
	aux := 0
	for i := 0; i < list.inserted/2; i++ {
		aux = list.values[i]
		list.values[i] = list.values[list.inserted-i-1]
		list.values[list.inserted-i-1] = aux
	}
}

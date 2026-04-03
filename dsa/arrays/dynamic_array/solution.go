// Package dynamicarray contains the solution for the Dynamic Array problem.
package dynamicarray

type DynamicArray struct {
	data []int
	size int
	cap  int
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		data: make([]int, capacity),
		size: 0,
		cap:  capacity,
	}

}

func (da *DynamicArray) Get(i int) int {
	return da.data[i]

}

func (da *DynamicArray) Set(i int, n int) {
	da.data[i] = n
}

func (da *DynamicArray) Pushback(n int) {
	if da.size == da.cap {
		da.resize()
	}
	da.data[da.size] = n
	da.size++

}

func (da *DynamicArray) Popback() int {
	if da.size == 0 {
		return -1
	}

	da.size--
	return da.data[da.size]
}

func (da *DynamicArray) resize() {
	newcap := da.cap * 2
	newdata := make([]int, newcap)
	copy(newdata, da.data)
	da.data = newdata
	da.cap = newcap

}

func (da *DynamicArray) GetSize() int {
	return da.size
}

func (da *DynamicArray) GetCapacity() int {
	return da.cap
}

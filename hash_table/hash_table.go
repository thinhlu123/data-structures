package hash_table

import (
	"data_struct/tree"
)

type HashTable struct {
	table map[uint]interface{}
}

// get key
func hashKeyMap(key float64) uint {

	return uint(key)
}

func (ht HashTable) IsExist(data interface{}) (bool, uint) {
	if ht.table == nil {
		return false, 0
	}

	// use hash to convert interface to int
	key, err := tree.HashKey(data)
	if err != nil {
		return false, 0
	}

	index := hashKeyMap(key)

	if _, ok := ht.table[index]; ok {
		return true, index
	}
	return false, 0
}

func (ht *HashTable) Insert(data interface{}) uint {

	isExist, index := ht.IsExist(data)
	if isExist {
		return index
	}

	if ht.table == nil {
		ht.table = make(map[uint]interface{})
	}

	ht.table[index] = data
	return index
}

func (ht HashTable) GetDataByIndex(index uint) interface{} {

	if val, ok := ht.table[index]; ok {
		return val
	}

	return nil
}

func (ht HashTable) Size() int {
	return len(ht.table)
}

func (ht HashTable) Remove(data interface{}) bool {
	isExist, index := ht.IsExist(data)
	if !isExist {
		return false
	}

	delete(ht.table, index)
	return true
}

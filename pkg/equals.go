package pkg

import "reflect"

//Equals uses reflec.DeepEqual to check equality of two slices. Only development over reflect.DeepEquals by this function is syntax
func (list *Array) Equals(compareTo interface{}) bool {
	if reflect.DeepEqual(list.Value, compareTo) {
		return true
	}
	return false
}

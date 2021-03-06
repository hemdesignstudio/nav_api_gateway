// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package salesinvoice

// removeField is used specifically to remove custom fields created for
// from the list of fields being sent to Microsft Navision
func removeField(name string, fields interface{}) interface{} {

	//find index of element to be deleted
	// using its name
	_fields := fields.([]string)

	for index, value := range _fields {

		if name == value {
			// Deleting elements is handled by the following couple of lines of code.
			// If the type of the element is a pointer or a struct with pointer fields,
			// which need to be garbage collected, the implementations of deleting an element
			// have a potential memory leak problem:
			// some elements with values are still referenced by slice a and
			// thus can not be collected.
			// although we are dealing with a []string,
			// It is still better to be oin the safe side and implement it this way
			copy(_fields[index:], _fields[index+1:])
			_fields[len(_fields)-1] = ""
			_fields = _fields[:len(_fields)-1]
		}
	}

	return _fields
}

func addFieldIfNotExist(name string, fields interface{}) interface{} {
	_fields := fields.([]string)
	for _, _name := range _fields {
		if name == _name {
			return fields
		}
	}

	return append(_fields, name)
}

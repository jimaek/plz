package reflect2

import (
	"unsafe"
	"reflect"
)

type eface struct {
	rtype unsafe.Pointer
	data  unsafe.Pointer
}

func unpackEFace(obj interface{}) *eface {
	return (*eface)(unsafe.Pointer(&obj))
}

func packEFace(rtype unsafe.Pointer, data unsafe.Pointer) interface{} {
	var i interface{}
	e := (*eface)(unsafe.Pointer(&i))
	e.rtype = rtype
	e.data = data
	return i
}

type UnsafeEFaceType struct {
	unsafeType
}

func newUnsafeEFaceType(cfg *frozenConfig, type1 reflect.Type) *UnsafeEFaceType {
	return &UnsafeEFaceType{
		unsafeType: *newUnsafeType(cfg, type1),
	}
}

func (type2 *UnsafeEFaceType) Indirect(obj interface{}) interface{} {
	objEFace := unpackEFace(obj)
	assertType("Type.Indirect argument 1", type2.ptrRType, objEFace.rtype)
	return type2.UnsafeIndirect(objEFace.data)
}

func (type2 *UnsafeEFaceType) UnsafeIndirect(ptr unsafe.Pointer) interface{} {
	return *(*interface{})(ptr)
}
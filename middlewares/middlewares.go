package middlewares

import (
	"fmt"

	"github.com/DataDog/go-python3"
)

func dump_to_python(s string) string {
	defer python3.Py_Finalize()
	python3.Py_Initialize()

	sysModule := python3.PyImport_ImportModule("sys")
	path := sysModule.GetAttrString("path")
	//pathStr, _ := pythonRepr2(path)
	//log.Println("before add path is " + pathStr)
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(""))
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString("./py"))

	concurrencyFile := python3.PyImport_ImportModule("test_concurrency")
	oDict := python3.PyModule_GetDict(concurrencyFile)

	genTestdata := python3.PyDict_GetItemString(oDict, "Test_func")
	if !(genTestdata != nil && python3.PyCallable_Check(genTestdata)) {
		// raise error
	}
	//a := python3.PyFloat_FromDouble(1)
	a := python3.PyUnicode_FromString(s)

	fmt.Println(a)

	testdataPy := genTestdata.CallFunctionObjArgs(a) //retval: New reference
	//...

	str, _ := pythonRepr(testdataPy)

	fmt.Printf("%T", s)

	defer testdataPy.DecRef()

	return str
}

func pythonRepr(o *python3.PyObject) (string, error) {
	if o == nil {
		return "", fmt.Errorf("object is nil")
	}

	s := o.Repr()
	if s == nil {
		python3.PyErr_Clear()
		return "", fmt.Errorf("failed to call Repr object method")
	}
	defer s.DecRef()

	return python3.PyUnicode_AsUTF8(s), nil
}

package middlewares

import (
	"fmt"
	"os"
	"shorten/base"

	"github.com/DataDog/go-python3"
)

func Dump_to_python(s string) string {
	defer python3.Py_Finalize()
	python3.Py_Initialize()

	sysModule := python3.PyImport_ImportModule("sys")
	path := sysModule.GetAttrString("path")

	//pathStr, _ := pythonRepr2(path)
	//log.Println("before add path is " + pathStr)
	//python3.PyList_Insert(path, 0, python3.PyUnicode_FromString("//middlewares//"))
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(" "))
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString("/.py"))
	dir, _ := os.Getwd()
	//dir = dir + "\\middlewares\\"
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(base.Python_dir))
	//fmt.Println(dir)
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(dir))
	//pathStr, _ := PythonRepr(path)
	//fmt.Println(pathStr)
	concurrencyFile := python3.PyImport_ImportModule("cat_dog_classifier")
	if concurrencyFile ==nil{

	fmt.Println(PythonRepr(path))
	panic("not found concurrency\n\n\n\n")
	
	}
	oDict := python3.PyModule_GetDict(concurrencyFile)
	genTestdata := python3.PyDict_GetItemString(oDict, "Train")

	
	if !(genTestdata != nil && python3.PyCallable_Check(genTestdata)) {
		// raise error
	}
	//a := python3.PyFloat_FromDouble(1)
	a := python3.PyUnicode_FromString(s)

	//fmt.Println(a)

	testdataPy := genTestdata.CallFunctionObjArgs(a) //retval: New reference
	//...

	str, _ := PythonRepr(testdataPy)

	fmt.Printf("the result :" + str)

	defer testdataPy.DecRef()

	return str
}

func PythonRepr(o *python3.PyObject) (string, error) {
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

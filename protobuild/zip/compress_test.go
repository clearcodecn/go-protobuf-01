package zip

import (
	"io/ioutil"
	"testing"
)

func TestUnZip(t *testing.T) {
	//{
	//	data, err := Zip("./compress.go")
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	ioutil.WriteFile("compress.go.zip", data, 0777)
	//}

	data, err := Zip("../../proto")
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("pb.zip", data, 0777)
}

func TestZip(t *testing.T) {
	err := UnZip("pb.zip", "./dst")
	if err != nil {
		t.Fatal(err)
	}
}

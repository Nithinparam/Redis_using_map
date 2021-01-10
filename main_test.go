package main

import "testing"

func TestGetData(t *testing.T) {
	key := "a"
	val := GetData(key)
	if val != "2" {
		t.Error("Expected", 2, "got", val)
	}

}
func TestSetData(t *testing.T) {
	key := "z"
	value := "1"
	expectedval1 := "variable succesfully added"
	res := SetData(key, value)
	if res != expectedval1 {
		t.Error("expecetd variable got added got", res)
	}
}

func TestUpdateData(t *testing.T) {
	key := "b"
	cmd := []string{"inc", "dec"}
	for _, val := range cmd {
		res := UpdateData(key, val)
		if res != "Updated" {
			t.Error("Expected value will be updated got", res)

		}
	}
}

func TestDeleteData(t *testing.T) {
	key := "z"
	res := DeleteData(key)
	if res != "Data Deleted" {
		t.Error("Expected to delete data but got", res)
	}
}

func TestIsDataExists(t *testing.T) {
	key := "a"
	_, yn := IsDataExists(key)
	if yn != true {
		t.Error("Expected Data exists")
	}
}

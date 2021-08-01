package utils_test

import (
	"testing"

	"github.io/taserbeat/line-family-bot/modules/utils"
)

func TestContaines(t *testing.T) {
	stringSlice := []string{"a", "b", "c"}
	if !utils.Contains(stringSlice, "a") {
		t.Errorf("stringSlice contains `a`, but it could not find `a`.")
	}

	if utils.Contains(stringSlice, "d") {
		t.Errorf("stringSlice does not contain `d`, but it finds `d`.")
	}

	intSlice := []int{10, 20, 30}
	if !utils.Contains(intSlice, 20) {
		t.Errorf("intSlice contains `20`, but it could not find `20`.")
	}

	if utils.Contains(intSlice, 40) {
		t.Errorf("intSlice does not contain `40`, but it finds `40`.")
	}

	type User struct {
		Name string
		Age  int
	}

	objectSlice := []User{}

	tom := User{Name: "Tom", Age: 20}
	objectSlice = append(objectSlice, tom)

	john := User{Name: "John", Age: 30}
	objectSlice = append(objectSlice, john)

	if !utils.Contains(objectSlice, tom) {
		t.Errorf("objectSlice contains %v, but it could not find %v", tom, tom)
	}

	emma := User{Name: "Emma", Age: 25}
	if utils.Contains(objectSlice, emma) {
		t.Errorf("objectSlice does not contain %v, but it finds %v", emma, emma)
	}
}

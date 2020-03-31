package kata

import "testing"

func TestUnique(t *testing.T) {
	if isUnique("josuo") {
		t.Error("'unique' is not a unique string")
	}
	if !isUnique("uniqke") {
		t.Error("'uniqke' is a unique string")
	}
}

func TestIsUniqueVanilla(t *testing.T) {
	if isUniqueVanilla("unique") {
		t.Error("'unique' is a not unique string")
	}
	if !isUniqueVanilla("josue") {
		t.Error("'uniqke' is a unique string")
	}
}

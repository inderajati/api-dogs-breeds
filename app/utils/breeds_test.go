package utils

import "testing"

func TestGetAllBreeds(t *testing.T) {
	got, _ := GetListBreeds()
	if got == nil {
		t.Errorf("GetListBreeds() = %v", got)
	}
}

func TestGetListSubBreeds(t *testing.T) {
	got, _ := GetListSubBreeds("hound")
	if got == nil {
		t.Errorf("GetListSubBreeds() = %v", got)
	}
}

func TestGetListBreedImages(t *testing.T) {
	got, _ := GetListBreedImages("hound")
	if got == nil {
		t.Errorf("GetListBreedImages() = %v", got)
	}
}

func TestGetListSubBreedImages(t *testing.T) {
	got, _ := GetListSubBreedImages("hound", "afghan")
	if got == nil {
		t.Errorf("GetListSubBreeds() = %v", got)
	}
}


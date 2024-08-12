package content_reader

import "testing"

func TestGetExtensionWithoutTrailingQuestionMark(t *testing.T) {
	if ext := getFileExtension("test.pdf"); ext == ".pdf" {
		t.Logf("Success: %v", ext)
	} else {
		t.Errorf("Failed: %v", ext)
	}
}

func TestGetExtensionWithTrailingQuestionMark(t *testing.T) {
	if ext := getFileExtension("test.pdf?a123"); ext == ".pdf" {
		t.Logf("Success: %v", ext)
	} else {
		t.Errorf("Failed: %v", ext)
	}
}

func TestGetExtensionWithoutExtension(t *testing.T) {
	if ext := getFileExtension("test"); ext == "" {
		t.Logf("Success: %v", ext)
	} else {
		t.Errorf("Failed: %v", ext)
	}
}

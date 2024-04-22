package main

import (
	"bytes"
	"os"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathname := CASPathTransformFunc(key)
	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	if pathname != expectedPathName {
		t.Errorf("have %s want %s", pathname, expectedPathName)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeStream("myspecialpicture", data); err != nil {
		t.Error(err)
	}
}

func TestStoreWriteStream(t *testing.T) {
	tests := []struct {
		key             string
		expectedFolders int
	}{
		{"momsbestpicture", 7},       // Expected number of folders created for the given key
		{"anotherspecialpicture", 7}, // Another example
	}

	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc, // Using CASPathTransformFunc for testing
	}
	s := NewStore(opts)

	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			defer os.RemoveAll(test.key) // Cleanup after each test case

			data := bytes.NewReader([]byte("some jpg bytes")) // Mocking JPG bytes
			err := s.writeStream(test.key, data)
			if err != nil {
				t.Errorf("Error writing stream: %v", err)
			}

			// Check if the expected number of folders were created
			// Here, we'll just check if the top-level folder was created
			// You might want to expand this to check for all subfolders
			_, err = os.Stat(test.key)
			if os.IsNotExist(err) {
				t.Errorf("Expected folder %s was not created", test.key)
			}
		})
	}
}

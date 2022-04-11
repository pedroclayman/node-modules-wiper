package modulesearch

import (
	"os"
	"testing"
)

func TestGetNodeModuleDirectories(t *testing.T) {
	testDataRoot := "__testData"

	os.RemoveAll(testDataRoot)

	os.MkdirAll(testDataRoot+"/foo/node_modules", 0777)
	os.MkdirAll(testDataRoot+"/foo/bar", 0777)
	os.Create(testDataRoot + "/foo/bar/node_modules")
	os.MkdirAll(testDataRoot+"/node_modules", 0777)

	var dirs []string
	GetNodeModuleDirectories(testDataRoot, &dirs)

	{
		want := 2
		got := len(dirs)

		if want != got {
			t.Errorf("Expected %v directories, got %v", want, got)
		}
	}

	{
		wants := []string{testDataRoot + "/foo/node_modules", testDataRoot + "/node_modules"}
		got := dirs

		for i := 0; i < 2; i++ {
			if wants[i] != got[i] {
				t.Errorf("Expected %v directories, got %v", wants[i], got[i])
			}
		}

	}

	os.RemoveAll(testDataRoot)
}

package go_embed_data

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	// This test must run from the workspace root since it accesses files using
	// relative paths. We look at parent directories until we find AUTHORS,
	// then change to that directory.
	for {
		if _, err := os.Stat("AUTHORS"); err == nil {
			break
		}
		if err := os.Chdir(".."); err != nil {
			log.Fatal(err)
		}
		if wd, err := os.Getwd(); err != nil {
			log.Fatal(err)
		} else if wd == "/" {
			log.Fatal("could not locate workspace root")
		}
	}
	os.Exit(m.Run())
}

func TestEmpty(t *testing.T) {
	if len(empty) != 0 {
		t.Fatalf("empty is not empty")
	}
}

func TestSingle(t *testing.T) {
	checkFile(t, "AUTHORS", single)
}

func TestLocal(t *testing.T) {
	for path, data := range local {
		checkFile(t, path, data)
	}
}

func TestExternal(t *testing.T) {
	for path, data := range ext {
		checkFile(t, path, data)
	}
}

func TestFlat(t *testing.T) {
	for key := range flat {
		if filepath.Base(key) != key {
			t.Errorf("filename %q is not flat", key)
		}
	}
}

func TestString(t *testing.T) {
	for _, data := range str {
		var _ string = data // just check the type; contents covered by other tests.
	}
}

func checkFile(t *testing.T, path string, data []byte) {
	f, err := os.Open(path)
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	count := 0
	buffer := make([]byte, 8192)
	for {
		n, err := f.Read(buffer)
		if err != nil && err != io.EOF {
			t.Error(err)
			return
		}
		if n == 0 {
			return
		}
		if n > len(data) {
			t.Errorf("%q: file on disk is longer than embedded data", path)
			return
		}

		for i := 0; i < n; i++ {
			if buffer[i] != data[i] {
				t.Errorf("data mismatch on file %q at offset %d", path, count+i)
				return
			}
		}
		count += n
		data = data[n:]
	}
}

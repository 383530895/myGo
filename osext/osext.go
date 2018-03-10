package osext

import (
	"path/filepath"
	"os"
)

var cx, ce = executableClean()

func executableClean() (string, error) {
	p, err := executable()
	return filepath.Clean(p), err
}

// Executable returns an absolute path that can be used to
// re-invoke the current program.
// It may not be valid after the current program exits.
func Executable() (string, error) {
	return cx, ce
}

// Returns same path as Executable, returns just the folder
// path. Excludes the executable name and any trailing slash.
func ExecutableFolder() (string, error) {
	p, err := Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(p), nil
}

//Check file exits
func Exits(filename string) (bool)  {
	if _, err := os.Stat(filename) ; err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//check Read Permission
func HaveReadPermission(filename string) (bool) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)

	if err != nil {
		return false
	}
	defer file.Close()

	return !os.IsPermission(err)
}

//
func HaveWritePermission(filename string) (bool)  {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)

	if err != nil {
		return false
	}
	defer file.Close()

	return !os.IsPermission(err)
}

func HaveRWPermission(filename string) (bool)  {
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)

	if err != nil {
		return false
	}
	defer file.Close()

	return !os.IsPermission(err)
}

//Search dir file
func SearchDir(dir string)  (fi []string)  {
	d, err := os.Open(dir)

	if err != nil {
		return nil
	}
	defer d.Close()

	f, err := d.Readdir(-1)
	if err != nil {
		return nil
	}

	fi = make([]string, len(f))
	for _, finfo := range f {
		if finfo.Mode().IsRegular() {
			fi = append(fi, finfo.Name())
		}
	}

	return fi
}
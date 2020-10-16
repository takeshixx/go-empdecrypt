package embedded

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/markbates/pkger"
)

var loadedExe bool
var execPath string

func ExecResource(resource string, args []string) (encrypted []byte, err error) {
	if !loadedExe {
		tmpFile := WriteResourceTempfile(resource)
		os.Chmod(tmpFile.Name(), 0755)
		execPath = tmpFile.Name()
		loadedExe = true
	}
	cmd := exec.Command(execPath, args...)
	encrypted, err = cmd.CombinedOutput()
	return
}

func WriteResourceTempfile(resource string) *os.File {
	tmpFile, err := os.OpenFile(os.TempDir()+string(os.PathSeparator)+resource, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	resourceFile, err := pkger.Open("/resources/" + resource)
	if err != nil {
		log.Fatal(err)
	}
	defer resourceFile.Close()
	_, err = io.Copy(tmpFile, resourceFile)
	if err != nil {
		log.Fatal(err)
	}
	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}
	return tmpFile
}

func CheckResources() {
	if _, err := os.Stat("resources"); err != nil {
		log.Fatal("resources folder does not exist")
	}
	if _, err := os.Stat("resources/EmpCrypt.exe"); err != nil {
		log.Fatal("resources/EmpCrypt.exe does not exist")
	}
	if _, err := os.Stat("resources/Matrix42.Common.AppVerificator.dll"); err != nil {
		log.Fatal("resources/Matrix42.Common.AppVerificator.dll does not exist")
	}
}

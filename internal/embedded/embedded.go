// +build embedded

package embedded

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/markbates/pkger"
)

func ExecResource(resource string, args []string) (encrypted []byte, err error) {
	tmpFile := WriteResourceTempfile(resource)
	defer os.Remove(tmpFile.Name())
	os.Chmod(tmpFile.Name(), 0755)
	log.Printf("Copied resource to %s\n", tmpFile.Name())
	cmd := exec.Command(tmpFile.Name(), args...)
	encrypted, err = cmd.CombinedOutput()
	return
}

func WriteResourceTempfile(resource string) *os.File {
	tmpFile, err := ioutil.TempFile("", "*"+filepath.Ext(resource))
	if err != nil {
		log.Fatal(err)
	}
	resourceFile, err := pkger.Open("/resources/" + resource)
	if err != nil {
		log.Fatal(err)
	}
	defer resourceFile.Close()
	n, err := io.Copy(tmpFile, resourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes", n)

	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}
	return tmpFile
}

func checkResources() {
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

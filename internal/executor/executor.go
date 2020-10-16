// +build !embedded

package executor

import "os/exec"

var execPath string

func InitExec(path string) {
	execPath = path
}

func ExecuteEmpCrypt(plain string) (encrypted []byte, err error) {
	cmd := exec.Command(execPath, "/S", "/EIS", plain)
	encrypted, err = cmd.CombinedOutput()
	return
}

// +build embedded

package executor

import (
	"fmt"

	"github.com/takeshixx/go-empdecrypt/internal/embedded"
)

func ExecuteEmpCrypt(plain string) (encrypted []byte, err error) {
	return embedded.ExecResource("EmpCrypt.exe", []string{"/S", "/EIS", plain})
}

func InitExec(path string) {
	fmt.Println("Writing DLL")
	embedded.WriteResourceTempfile("Matrix42.Common.AppVerificator.dll")
}

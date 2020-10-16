// +build embedded

package executor

import (
	"github.com/takeshixx/go-empdecrypt/internal/embedded"
)

var loadedDLL bool

func ExecuteEmpCrypt(plain string) (encrypted []byte, err error) {
	return embedded.ExecResource("EmpCrypt.exe", []string{"/S", "/EIS", plain})
}

func InitExec(path string) {
	embedded.CheckResources()
	if !loadedDLL {
		embedded.WriteResourceTempfile("Matrix42.Common.AppVerificator.dll")
		loadedDLL = true
	}
}

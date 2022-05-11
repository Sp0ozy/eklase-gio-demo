package explorer

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"syscall"
)

func Drives() []string {

	kernel32, _ := syscall.LoadLibrary("kernel32.dll")
	getLogicalDrivesHandle, _ := syscall.GetProcAddress(kernel32, "GetLogicalDrives")

	var drives []string

	if ret, _, callErr := syscall.Syscall(uintptr(getLogicalDrivesHandle), 0, 0, 0, 0); callErr != 0 {

	} else {
		drives = bitsToDrives(uint32(ret))
	}
	return drives

}

func bitsToDrives(bitMap uint32) (drives []string) {
	availableDrives := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	for i := range availableDrives {
		if bitMap&1 == 1 {
			drives = append(drives, availableDrives[i])
		}
		bitMap >>= 1
	}
	return
}

func listAll() (files []fs.FileInfo) {
	l := Drives()
	for j := 0; j != len(l); j++ {
		l[j] = l[j] + ":\\"
	}
	log.Println(l)
	for i := 0; i != len(l); i++ {
		f, err := os.Open(l[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		files, err := f.Readdir(0)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		for _, v := range files {
			fmt.Println(v.Name(), v.IsDir())
		}
	}
	return files
}

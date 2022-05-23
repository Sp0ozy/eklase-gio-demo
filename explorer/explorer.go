package explorer

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/shirou/gopsutil/disk"
)

func Drives() []string {

	// kernel32, _ := syscall.LoadLibrary("kernel32.dll")
	// getLogicalDrivesHandle, _ := syscall.GetProcAddress(kernel32, "GetLogicalDrives")

	// var drives []string

	// if ret, _, callErr := syscall.Syscall(uintptr(getLogicalDrivesHandle), 0, 0, 0, 0); callErr != 0 {

	// } else {
	// 	drives = bitsToDrives(uint32(ret))
	// }
	// for j := 0; j != len(drives); j++ {
	// 	drives[j] = drives[j] + ":"
	// }
	// fmt.Println(drives)
	p, err := disk.Partitions( /*all*/ true)
	if err != nil {
		log.Fatal(err)
	}
	drives := make([]string, len(p))
	for i := range p {

		drives[i] = p[i].Mountpoint
		fmt.Println(p[i].Mountpoint)
	}
	fmt.Println(drives)
	return drives

}

// func bitsToDrives(bitMap uint32) (drives []string) {
// 	availableDrives := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

// 	for i := range availableDrives {
// 		if bitMap&1 == 1 {
// 			drives = append(drives, availableDrives[i])
// 		}
// 		bitMap >>= 1
// 	}
// 	return
// }

func ListAll() (files []fs.FileInfo) {
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
func Root(r []string) (root string) {
	return filepath.Join(r...)
}
func List(r []string) (names []string) {
	root := Root(r)
	f, err := os.Open(root)
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
		fmt.Println(v.Name())
		names = append(names, v.Name())
	}
	return names
}

func SaveRoot(root []string) {
	var err = os.Remove("screen/root.txt")
	if err != nil {
		return
	}
	f, e := os.Create("screen/root.txt")
	if e != nil {
		log.Fatal(e)
	}
	f.Close()
	f, _ = os.OpenFile("screen/root.txt", os.O_APPEND, 0666)
	for _, v := range root {
		f.WriteString(v + "\n")
	}
	f.Close()
}
func GetRoot() (root []string) {
	f, _ := os.OpenFile("screen/root.txt", os.O_RDONLY, 0666)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		root = append(root, scanner.Text())
	}
	defer f.Close()
	return root
}

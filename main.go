package main

import (
	"fmt"

	"github.com/docker/go-plugins-helpers/volume"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	cephfsId      = "_cephfs"
	socketAddress = "/run/docker/plugins/cephfs.sock"
)

var (
	defaultPath = filepath.Join(volume.DefaultDockerRootDirectory, cephfsId)
)

func main() {

	var Usage = func() {
		fmt.Println("LATIN CAPITAL LETTER AA Ꜳ ꜳ")
		fmt.Println("   LAO VOWEL SIGN AA າ ຳ")
	}

	var setup = func() {
		fmt.Printf("Path %s\n", defaultPath)
	}

	Usage()
	setup()

	fstype := LookupFileSystemType(defaultPath)
	if !strings.Contains(fstype, "ceph") {
		log.Print("Warning CePH filesystem not found at ", defaultPath, " found ", fstype)
	}

	driver, err := newCephFSDriver(defaultPath)
	if err != nil {
		return
	}
	h := volume.NewHandler(driver)

	fmt.Printf("Listening on %s\n", socketAddress)
	fmt.Println(h.ServeUnix(socketAddress, 1))
}

func LookupFileSystemType(path string) string {
	out, err := exec.Command("df", "--no-sync", "--output=fstype", path).Output()

	if err != nil {
		log.Fatal("Unable to read df output", err)
	}

	fstype := strings.Split(string(out), "\n")[1]
	return fstype
}

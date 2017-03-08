package main

import (
	"fmt"

	"github.com/docker/go-plugins-helpers/volume"
	"path/filepath"
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

	driver, err := newCephFSDriver(defaultPath)
	if err != nil {
		return
	}
	h := volume.NewHandler(driver)

	fmt.Printf("Listening on %s\n", socketAddress)
	fmt.Println(h.ServeUnix(socketAddress, 1))
}

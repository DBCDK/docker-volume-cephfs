package main

import (
	"fmt"

	"github.com/docker/go-plugins-helpers/volume"
)

const (
	cephfsId = "_cephfs"
	socketAddress = "/run/docker/plugins/cephfs.sock"
)

func main() {

	var Usage = func() {
		fmt.Println("Hello, 世界")
	}

	Usage()
	

	//driver := newCephFSDriver( "/tmp/cephfs")
	driver := newCephFSDriver( "/mnt/cephfs_root/")
	h := volume.NewHandler( driver );

	fmt.Printf("Listening on %s\n", socketAddress)
	fmt.Println(h.ServeUnix(socketAddress, 1))
}


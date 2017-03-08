package main


import (
	"log"
	"github.com/docker/go-plugins-helpers/volume"
	"fmt"
	"os"
)


type cephFSDriver struct {
	root string
}

/**

 */
func newCephFSDriver( root  string  ) (cephFSDriver, error) {
	if  !IsDirectory( root ) {
		err := os.MkdirAll( root , os.ModePerm)
		if( err != nil ) {
			log.Print(fmt.Sprintf(" Directory %s Dont Exists, and unable to Create IT ", root ))
			log.Print(fmt.Sprintf(" mkdir failed with:  %s", err ))
			return cephFSDriver{}, err
		}
	}
	
	return cephFSDriver{
		root: root,
	}, nil
}

func (d cephFSDriver ) Create( r volume.Request ) volume.Response {
	log.Print("Create Called ", r.Name, " ", r.Options)
	return volume.Response{}
}

func( d cephFSDriver ) List( r volume.Request ) volume.Response {
	log.Print("List Called ", r.Name, " ", r.Options)
	return volume.Response{}
}


func( d cephFSDriver ) Get( r volume.Request ) volume.Response {
	log.Print("Get Called ", r.Name," ", r.Options)
	log.Print(r)
	log.Print("Get End")
	
	m := d.getAbsolutePathForVolume( r.Name )
	return volume.Response{ Volume: &volume.Volume{
		Name: r.Name,
		Mountpoint: m,
	}}
}

func( d cephFSDriver ) Remove( r volume.Request ) volume.Response {
	log.Print("Remove Called ", r.Name, " ", r.Options)
	return volume.Response{Err: "error Remove NIJ"}
}

func( d cephFSDriver ) Path( r volume.Request ) volume.Response {
	log.Print("Path Called ", r.Name, " ", r.Options)
	
	return volume.Response{ Mountpoint: d.getAbsolutePathForVolume( r.Name ) };
}


func (d cephFSDriver ) Mount( r volume.MountRequest ) volume.Response {
	log.Print("Mount Called ",r.ID," ", r.Name)
	log.Print(r)
	log.Print("Mount End")

	m := fmt.Sprintf("%s/%s",d.root, r.Name)
	if( ! IsDirectory(m) ) {
		return volume.Response{Err: fmt.Sprintf(" %s is not a directory ", m)}
	};
	return volume.Response{ Mountpoint: m};
}

func (d cephFSDriver ) Unmount( r volume.UnmountRequest ) volume.Response {
	log.Print("Unmount Called ", r.ID, " ", r.Name)
	return volume.Response{Err: "error NIJ"}
}
func (d cephFSDriver ) Capabilities( r volume.Request ) volume.Response {
	log.Print("Capabilities Called ", r, " ", r.Name)
	return volume.Response{ }
}


func ( d cephFSDriver) getAbsolutePathForVolume( v string ) string {
	return fmt.Sprintf("%s/%s",d.root, v)
}


func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path);
	if  err != nil {
		return false
	}
	return fileInfo.IsDir()
}
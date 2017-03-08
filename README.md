# docker-volume-cephfs

Volume plugin to use CephFS as distributed data storage

# Use Case 

This Plugin is used to accessing Self Service shared persistent storage for 
docker volumes run on a Mesos/Marathon cluster. 

Teams create gets a directory for use by multiple docker instances. 

# CephFS setup 

Only one CephFS file system is shared by all teams, with quota on the directory level.
 
# Example 
 
Team Dataio requested a shared volume called fileStore, for a staging. 

```$bash
docker run --volume-driver cephfs -v staging/dataio/fileStore:/data --rm -ti docker.dbc.dk/dbc-jessie /bin/bash
```


# Limits 

Only Debian/ubuntu linux with systemd is tested
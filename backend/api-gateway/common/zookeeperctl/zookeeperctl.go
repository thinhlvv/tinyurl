package zookeeperctl

import (
	"errors"
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

type Zookeeper struct {
	client *zk.Conn
}

func New(zookeepers []string) (Zookeeper, error) {
	conn, _, err := zk.Connect(zookeepers, time.Second)
	if err != nil {
		return Zookeeper{}, err
	}
	return Zookeeper{client: conn}, nil
}

func (z *Zookeeper) Read(path string) ([]byte, error) {
	val, state, err := z.client.Get(path)
	logZKState(state)
	return val, err
}

func (z *Zookeeper) Write(path string, data []byte) error {
	// exist
	exist, s, err := z.client.Exists(path)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("Path is not existed.")
	}

	_, err = z.client.Set(path, data, s.Cversion)

	logZKState(s)

	return err
}

func (z *Zookeeper) Create(path string, data []byte) error {
	var flags int32 = 0
	var acls = zk.WorldACL(zk.PermAll) // permission

	// create
	p, err := z.client.Create(path, data, flags, acls)

	fmt.Println("created:", p)
	return err
}

func (z *Zookeeper) Delete(path string) error {
	// exist
	exist, s, err := z.client.Exists(path)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("Path is not existed.")
	}

	err = z.client.Delete(path, s.Version)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func logZKState(s *zk.Stat) {
	fmt.Sprintf("Czxid:%d\nMzxid: %d\nCtime: %d\nMtime: %d\nVersion: %d\nCversion: %d\nAversion: %d\nEphemeralOwner: %d\nDataLength: %d\nNumChildren: %d\nPzxid: %d\n",
		s.Czxid, s.Mzxid, s.Ctime, s.Mtime, s.Version, s.Cversion, s.Aversion, s.EphemeralOwner, s.DataLength, s.NumChildren, s.Pzxid)
}

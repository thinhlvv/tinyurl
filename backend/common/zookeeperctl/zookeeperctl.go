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

func New(zookeepers []strings) (*zk.Conn, error) {
	conn, _, err := zk.Connect(zookeepers, time.Second)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (z *Zookeeper) Get(path string) ([]byte, error) {
	val, state, err := z.client.Get(path)
	logZKState(s)
	return val, err
}

func (z *Zookeeper) Write(path string, data []byte) error {
	// exist
	exist, s, err := z.client.Exists(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !exist {
		return errors.New("Path is not existed.")
	}
	if err != nil {
		return err
	}

	_, err := z.client.Set(path, data, s.Cversion)

	logZKState(s)

	return err
}

func logZKState(s *zk.Stat) string {
	fmt.Sprintf("Czxid:%d\nMzxid: %d\nCtime: %d\nMtime: %d\nVersion: %d\nCversion: %d\nAversion: %d\nEphemeralOwner: %d\nDataLength: %d\nNumChildren: %d\nPzxid: %d\n",
		s.Czxid, s.Mzxid, s.Ctime, s.Mtime, s.Version, s.Cversion, s.Aversion, s.EphemeralOwner, s.DataLength, s.NumChildren, s.Pzxid)
}

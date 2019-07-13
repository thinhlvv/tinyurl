package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	conn, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	// defer conn.Close()

	var path = "/zk_test_go/child1"
	var data = []byte("child1")
	var flags int32 = 0
	// permission
	var acls = zk.WorldACL(zk.PermAll)

	// create
	p, err_create := conn.Create(path, data, flags, acls)
	if err_create != nil {
		fmt.Println(err_create)
		return
	}
	fmt.Println("created:", p)

	// // get
	// v, s, err := conn.Get(path)
	// if err != nil {
	// 	fmt.Println("Getting data from root path", err)
	// 	return
	// }

	// // set
	// _, err = conn.Set(path, data, s.Cversion)
	// if err != nil {
	// 	fmt.Println("Getting data from root path", err)
	// 	return
	// }

	// // get
	// v, s, err = conn.Get(path)
	// if err != nil {
	// 	fmt.Println("Getting data from root path", err)
	// 	return
	// }

	// fmt.Printf("value of path[%s]=[%s].\n", path, v)
	// fmt.Printf("state:\n")
	// fmt.Printf("%s\n", ZkStateStringFormat(s))

	// children, stat, ch, err := c.ChildrenW("/")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("stats: %+v %+v\n", children, stat)
	// e := <-ch
	// fmt.Printf("channel: %+v\n", e)
}

func ZkStateStringFormat(s *zk.Stat) string {
	return fmt.Sprintf("Czxid:%d\nMzxid: %d\nCtime: %d\nMtime: %d\nVersion: %d\nCversion: %d\nAversion: %d\nEphemeralOwner: %d\nDataLength: %d\nNumChildren: %d\nPzxid: %d\n",
		s.Czxid, s.Mzxid, s.Ctime, s.Mtime, s.Version, s.Cversion, s.Aversion, s.EphemeralOwner, s.DataLength, s.NumChildren, s.Pzxid)
}

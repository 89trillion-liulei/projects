package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"grpc_study/protos"
	"log"
	"net/http"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	trans := transport.GetDefaultWebsocketTransport()
	header := http.Header{}
	header.Add("userId", "1")
	trans.RequestHeader = header
	//tk:= "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczpcL1wvZ2l0aHViLmNvbVwvYXV0aDBcL3BocC1qd3QtZXhhbXBsZSIsInN1YiI6IjEyMzQ1NiIsIm5hbWUiOiJKb2huIERvZSIsImVtYWlsIjoiaW5mb0BhdXRoMC5jb20iLCJ1aWQiOiJteUlkIn0.mKXEQvC0zdf271_eD_w6iG1b4mNe2I4UNh9DEE0P260"
	//gd:= "195c08c9-49ee-487b-9922-7b06a910351a"
	c, err := gosocketio.Dial(
		gosocketio.GetUrl("127.0.0.1", 7999, false),
		trans)
	if err != nil {
		log.Fatal(err)
	}
	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Fatal("Disconnected")
	})
	if err != nil {
		log.Fatal(err)
	}
	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("Connected11111")
	})
	if err != nil {
		log.Fatal(err)
	}

	message := &protos.Param{
		Type: 5,
		Data: []string{"1", "1", "2", "1.22123:1.54323", "90001", "10203", "14", "90001", "10301", "2"},
		Ext:  "",
	}

	marshal, _ := proto.Marshal(message)
	fmt.Println(marshal)

	//fmt.Println("building_info start!")

	//err = c.Emit("building_upgrade", marshal)

	//fmt.Println("building_info on start!")
	//
	//c.On("building_info", func(h *gosocketio.Client, data string) {
	//	fmt.Println(data)
	//	fmt.Println("building_info return:", h)
	//})

	//c.Emit("building_upgrade_complete_imdiately",marshal)

	//c.Emit("building_move",marshal)

	var a string

	for {
		fmt.Println("请任意输入：")

		//fmt.Scanf("%s", &a)
		fmt.Scan(&a)

		fmt.Println("您输入的是：", a)

		switch a {
		case "1":
			err = c.Emit("building.info", nil)
		case "2":
			err = c.Emit("building.upgrade", marshal)
		case "3":
			err = c.Emit("building.upgradeCompleteImediately", marshal)
		case "4":
			err = c.Emit("building.move", marshal)
		case "5":
			err = c.Emit("bounty.download", marshal)
		case "6":
			err = c.Emit("bounty.upload", nil)
		case "7":
			err = c.Emit("trophy.status", nil)
		case "8":
			err = c.Emit("collection.reward", marshal)
		case "9":
			err = c.Emit("collection.coin", nil)
		case "10":
			err = c.Emit("change.money", marshal)
		case "11":
			err = c.Emit("attack.upload", marshal)
		}

		//c.On("building_info", func(h *gosocketio.Channel, msg string) {
		//	fmt.Println("building_info return msg:",msg)
		//})

		//c.On("building_upgrade", func(h *gosocketio.Client,msg []byte) {
		//	fmt.Println(string(msg))
		//})
		//
		//
		//
		//_ = c.Emit("building_upgrade", marshal)

		//c.Emit("building_upgrade_complete_imdiately", marshal)
		//c.Emit("building_move", marshal)
		time.Sleep(60 * time.Second)

	}

	//func aaa(a int) error {
	//	fmt.Println(a)
	//	return errors.New("xxx")
	//}

	//func Retry(attempts int, sleep time.Duration, f func() error) (err error) {
	//	for i := 0; ; i++ {
	//		err = f()
	//		if err == nil {
	//			return
	//		}
	//
	//		if i >= (attempts - 1) {
	//			break
	//		}
	//
	//		time.Sleep(sleep)
	//
	//		//log.App().Error("retrying after error:" + err.Error())
	//	}
	//	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
	//
	//}

	//func retry(f func() error) {
	//	fmt.Println(f())
	//}
	//func main() {
	//	//count := math.Abs(float64(time.Now().Unix()) -float64(1597645440))
	//	//fmt.Println(count) // second
	//	try.MaxRetries = 10
	//	try.Do(func(attempt int) (retry bool, err error) {
	//		err = aaa(3)
	//		return attempt < 10,err
	//	})
	//}
}

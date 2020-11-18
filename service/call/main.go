package main

import (
	"log"
	"os"

	"github.com/kardianos/service"
)

var serviceConfig = &service.Config{
	Name: "serviceName",
	// DisplayName: "service Display Name",
	// Description: "service description",
}

func main() {
	// 構建服務物件
	prog := &Program{}
	s, err := service.New(prog, serviceConfig)
	if err != nil {
		log.Fatal(err)
	}
	// 用於記錄系統日誌
	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) < 2 {
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
		return
	}
	// cmd := os.Args[1]
	// if cmd == "install" {
	// 	err = s.Install()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("安裝成功")
	// }
	// if cmd == "uninstall" {
	// 	err = s.Uninstall()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("解除安裝成功")
	// }
	// install, uninstall, start, stop 的另一種實現方式
	err = service.Control(s, os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}

type Program struct{}

func (p *Program) Start(s service.Service) error {
	log.Println("開始服務")
	go p.run()
	return nil
}
func (p *Program) Stop(s service.Service) error {
	log.Println("停止服務")
	return nil
}
func (p *Program) run() {
	text := "this is ..."
	f, err := os.OpenFile("D:/git/golang-medium-example/service/call/file", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

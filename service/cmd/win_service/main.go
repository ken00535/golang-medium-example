package main

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

func main() {
	serviceName := "Hello"
	isWindowsService, _ := svc.IsWindowsService()
	if isWindowsService {
		run(serviceName)
		return
	}
	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			err := install(serviceName)
			if err != nil {
				fmt.Println(err)
			}
		} else if os.Args[1] == "uninstall" {
			err := uninstall(serviceName)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func install(name string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err == nil {
		return fmt.Errorf("service already exist")
	}
	path, err := os.Executable()
	if err != nil {
		return err
	}
	s, err = m.CreateService(name, path, mgr.Config{
		DisplayName: name,
		Description: "This is hello service",
		StartType:   mgr.StartAutomatic,
	}, "")
	if err != nil {
		return err
	}
	defer s.Close()
	return nil
}

func uninstall(name string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		return fmt.Errorf("service %s is not installed", name)
	}
	defer s.Close()
	err = s.Delete()
	if err != nil {
		return err
	}
	return nil
}

func run(name string) error {
	path, err := os.Executable()
	filename := filepath.Join(filepath.Dir(path), "file.log")
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString("hello\n"); err != nil {
		return err
	}
	ws := windowsService{}
	return svc.Run(name, &ws)
}

type windowsService struct{}

func (ws *windowsService) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (svcSpecificEC bool, exitCode uint32) {
	var endCh chan bool
	changes <- svc.Status{State: svc.StartPending}
	changes <- svc.Status{State: svc.Running, Accepts: svc.AcceptStop | svc.AcceptShutdown}
	go func() {
		for {
			request := <-r
			switch request.Cmd {
			case svc.Interrogate:
				changes <- request.CurrentStatus
			case svc.Stop, svc.Shutdown:
				changes <- svc.Status{State: svc.StopPending}
				changes <- svc.Status{State: svc.Stopped}
				endCh <- true
				return
			}
		}
	}()
	<-endCh
	return false, 0
}

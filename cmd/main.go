package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type config struct {
	Filename string `json:"filename"`
}

var (
	help     bool
	filename string
)

func init() {
	flag.BoolVar(&help, "h", false, "this is help")
	flag.StringVar(&filename, "r", "", "select your file")
	flag.Usage = usage
}

func main() {
	data, _ := ioutil.ReadFile("configs/config.json")
	var fileConfig config
	json.Unmarshal(data, &fileConfig)
	fmt.Println("Hello, world")
	noteCmd := exec.Command("cmd", "/c", "type "+fileConfig.Filename)
	buf := make([]byte, 1024)
	stdout, _ := noteCmd.StdoutPipe()
	noteCmd.Start()
	n, _ := stdout.Read(buf)
	os.Stdout.Write(buf[:n])
}

func usage() {
	fmt.Println("Usage: micro-cli [-h] [-r filename]")
	flag.PrintDefaults()
}

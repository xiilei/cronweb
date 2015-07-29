package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type Cron struct {
}

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

// just test code
func readCrontab() string {
	cmd := exec.Command("crontab", "-l")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("pipe error:%s\n", err.Error())
		return ""
	}
	if err := cmd.Start(); err != nil {
		log.Fatalf("run error:%s\n", err.Error())
		return ""
	}
	tasks := make([]string, 2)
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		s := scanner.Text()
		if strings.HasPrefix(s, "#") {
			continue
		}
		tasks = append(tasks, s)
		// parseCron(s)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scan error:%s\n", err.Error())
	}
	return strings.Join(tasks, "\n")
}

func parseCron(taskdesc string) {
	for _, v := range strings.Fields(taskdesc) {
		log.Println(v)
	}
}

func main() {
	addr := "localhost:3881"
	http.Handle("/", String(readCrontab()))
	fmt.Println("listen at", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

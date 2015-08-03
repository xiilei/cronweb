package core

import (
	// "github.com/boltdb/bolt"
	"bufio"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const (
	CAny       = "*"
	CDash      = "-"
	CBackslash = "/"
	CComma     = ","
)

//reading and create tasks from linux crontab
//crontab -l command
func fromCrontab(c int) (tasks []Task, err error) {
	tasks = make([]Task, 0, c)
	cmd := exec.Command("crontab", "-l")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		s := scanner.Text()
		//ignore comments
		if strings.HasPrefix(s, "#") {
			continue
		}
		task := ResolveTask(s)
		tasks = append(tasks, *task)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return
}

//checking times contains by crontab times
func checkInCrontabTime(s_times []string, t_times []int, days int) bool {
	for i, t := range s_times {
		if t == "*" {
			continue
		}
		si, err := strconv.Atoi(t)
		if err != nil {
			log.Printf("convert task.time(%s) to int failed:%s", t, err.Error())
			continue
		}
		if si != t_times[i] {
			return false
		}
	}
	return true
}

//some crontab usage follow
//min   hour    day     mon week
//30    21      *       *   *
//45    4       1,10,22 *   *
//0,30  18-23   *       *   *
//*     */2     *       *   *
//*     23-7/1  *       *   *
//0     4       1       jan *
func resolveCrontabTime(td TDate, time_desc string, days int) {
	//todo
}

//resolve ',' '-'
func resolveCrontabTimeAtom(td TDate, time_desc string, days int) (re_times []int, err error) {
	var st_times []string

	if strings.Contains(time_desc, CComma) {
		st_times = strings.Split(time_desc, CComma)
	} else if strings.Contains(time_desc, CBackslash) {
		//@todo
	}

	//convert []string to []int
	re_times = make([]int, len(st_times))
	for i, v := range st_times {
		re_times[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}
	return re_times, nil
}

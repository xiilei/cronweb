package core

import (
	// "github.com/boltdb/bolt"
	"bufio"
	"errors"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	CAny       = "*"
	CDash      = "-"
	CBackslash = "/"
	CComma     = ","
)

var badTime = errors.New("bad crontab syntax")

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
func checkInCrontabTime(s_times []string, t_times []int, t time.Time) bool {
	for i, t := range s_times {
		if t == CAny {
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
func resolveCrontabTime(time_desc string, t time.Time) {
	//todo
}

//resolve ',' '-'
func resolveCrontabTimeAtom(dt TDate, dst int, atom_desc string, t time.Time) (re bool, err error) {
	//comma
	if strings.Contains(atom_desc, CComma) {
		dstr := strconv.Itoa(dst)
		for _, v := range strings.Split(atom_desc, CComma) {
			if dstr == v {
				return true, nil
			}
		}
	}

	//dash
	if strings.Contains(atom_desc, CDash) {
		re_times, err := asTois(strings.Split(atom_desc, CDash))
		if err != nil {
			return false, err
		}
		if len(re_times) != 2 {
			return false, badTime
		}
		if re_times[0] < re_times[1] {
			return re_times[0] <= dst && dst <= re_times[1], nil
		}
		tmax := TDateMax(dt, t)
		return (re_times[0] <= dst && dst <= tmax) || (0 <= dst && dst <= re_times[1]), nil
	}

	return false, nil
}

//convert []string to []int
func asTois(strs []string) (re_ints []int, err error) {
	re_ints = make([]int, len(strs))
	for i, v := range strs {
		re_ints[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}
	return re_ints, nil
}

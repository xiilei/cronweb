package core

import (
	// "github.com/boltdb/bolt"
	"bufio"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"time"
)

const (
	cany       = "*"
	cdash      = "-"
	cbackslash = "/"
	ccomma     = ","
)

var badTime = errors.New("bad crontab syntax")

var cuser *user.User

func init() {
	var err error
	cuser, err = user.Current()
	if err != nil {
		log.Fatalf("get current user failed:%s\n", err.Error())
	}
}

//reading and create tasks from linux crontab
//crontab -l command
func fromCrontab(c int) (tasks []Task, err error) {
	var last_cline string
	var cline string
	tasks = make([]Task, 0, c)
	cmd := exec.Command("crontab", "-l")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		s := scanner.Text()
		cline = strings.TrimSpace(s)
		if cline == "" {
			continue
		}
		//ignore comments
		if strings.HasPrefix(s, "#") {
			last_cline = cline[1:]
			continue
		}
		task := ResolveTask(s)
		if last_cline != "" {
			task.SetTitle(last_cline)
		}
		tasks = append(tasks, *task)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	err = cmd.Wait()
	return
}

//checking times contains by crontab times
//some crontab usage follow
//min   hour    day     mon week
//30    21      *       *   *
//45    4       1,10,22 *   *
//0,30  18-23   *       *   *
//*     */2     *       *   *
//*     23-7/1  *       *   *
//0     4       1       jan *
func checkInCrontabTime(s_times []string, t_times []int, tm time.Time) (ok bool, err error) {
	for i, t := range s_times {
		ok, err = resolveCrontabTimeAtom(TDate(i), t_times[i], t, tm)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return ok, nil
}

//resolve ',' '-','/'
func resolveCrontabTimeAtom(dt TDate, dst int, atom_desc string, t time.Time) (ok bool, err error) {
	if atom_desc == cany {
		return true, nil
	}

	//month word
	if dt == TMonth {
		atom_desc = strings.ToLower(atom_desc)
		for i, wm := range Months {
			atom_desc = strings.Replace(atom_desc, wm, strconv.Itoa(i+1), -1)
		}
	}

	//backslash
	if strings.Contains(atom_desc, cbackslash) {
		re_times := strings.Split(atom_desc, cbackslash)
		if len(re_times) != 2 {
			return false, badTime
		}
		//*/n (n>1) not support  yet
		if re_times[1] != "1" {
			return false, errors.New("*/n (n>1) not support yet")
		}
		return resolveCrontabTimeAtom(dt, dst, re_times[0], t)
	}

	//comma
	if strings.Contains(atom_desc, ccomma) {
		dstr := strconv.Itoa(dst)
		for _, v := range strings.Split(atom_desc, ccomma) {
			if dstr == v {
				return true, nil
			}
		}
		return false, nil
	}

	//dash
	if strings.Contains(atom_desc, cdash) {
		re_times, err := asTois(strings.Split(atom_desc, cdash))
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

	return strconv.Itoa(dst) == atom_desc, nil
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

//writing raw tasks to crontab
//@todo keep the comments
func writeCrontab(task_desc string) error {
	tmpfile, err := ioutil.TempFile("/tmp", "cronweb-")
	if err != nil {
		return err
	}
	defer tmpfile.Close()
	_, err = tmpfile.Write([]byte(task_desc + "\n"))
	if err != nil {
		return err
	}
	name := tmpfile.Name()
	tmpfile.Close()
	cmd := exec.Command("crontab", name, "-u", cuser.Username)
	err = cmd.Run()
	if err != nil {
		return err
	}
	os.Remove(name)
	return nil
}

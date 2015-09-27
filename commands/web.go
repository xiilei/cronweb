package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/xiilei/cronweb/cron"
	"log"
	"net/http"
)

var WebCmd = cli.Command{
	Name:   "web",
	Usage:  "start cronweb web ui",
	Action: runWeb,
	Flags: []cli.Flag{
		cli.StringFlag{"addr, a", "localhost:3881", "cronweb web service address (e.g., ':3881')", ""},
	},
}

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func runWeb(c *cli.Context) {
	addr := c.String("addr")
	ts := core.NewTaskStore(1)
	http.Handle("/", String(ts.Raw()))
	fmt.Println("listen at", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

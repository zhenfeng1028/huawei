package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cmds = []string{
	"reset board",
	"board add",
	"board delete",
	"reboot backplane",
	"backplane abort",
}

var hashmap = map[string]string{
	"reset board":      "board fault",
	"board add":        "where to add",
	"board delete":     "no board at all",
	"reboot backplane": "impossible",
	"backplane abort":  "install first",
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	res := []string{}
	for input.Scan() {
		strs := strings.Fields(input.Text())
		if len(strs) < 1 || len(strs) > 2 {
			res = append(res, "unknown command")
			continue
		}
		if len(strs) == 1 {
			if strings.HasPrefix("reset", strs[0]) {
				res = append(res, "reset what")
			} else {
				res = append(res, "unknown command")
			}
		} else {
			var found bool
			var count int
			var command string
			for _, cmd := range cmds {
				subCmd := strings.Fields(cmd)
				if strings.HasPrefix(subCmd[0], strs[0]) && strings.HasPrefix(subCmd[1], strs[1]) {
					found = true
					count++
					command = cmd
				}
			}
			if !found || count > 1 {
				res = append(res, "unknown command")
			} else {
				res = append(res, hashmap[command])
			}
		}
	}
	for _, v := range res {
		fmt.Println(v)
	}
}

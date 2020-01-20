package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	for {
		// just pass the file name
		b, err := ioutil.ReadFile("vocabulary.txt")
		if err != nil {
			fmt.Print(err)
		}

		// convert content to a 'string'
		content := string(b)

		groups := strings.Split(content, "\n\n")

		for _, group := range groups {
			index := strings.Index(group, "\n")
			var content string
			var heading string
			if index != -1 {
				heading = group[:index]
				content = group[index+1 : len(group)]
			} else {
				heading = group
				content = ""
			}

			err := notify(heading, content)
			if err != nil {
				panic(err)
			}

			time.Sleep(90 * time.Second)

			err = notify(heading, content)
			if err != nil {
				panic(err)
			}
			time.Sleep(90 * time.Second)
		}
	}
}

func notify(heading string, content string) error {
	return beeep.Notify(heading, content, "assets/information.png")
	//cmd := exec.Command("bash", "-c", " notify-send \""+heading+"\" \""+content+"\"")
	//return cmd.Run()
}

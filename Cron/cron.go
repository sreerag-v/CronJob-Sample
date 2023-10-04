package Cron

import (
	"fmt"
	"io"
	"log"
	"net/http"

	cron "gopkg.in/robfig/cron.v2"
)

func RunCron() {
	cron := cron.New()

	//@every 00h00m00s
	cron.AddFunc("@every 00h00m10s", SendMsg)
	cron.Start()

}

func SendMsg() {
	res, err := http.Get("http://localhost:8080/message")

	if err != nil {
		log.Println("err from cron", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(body))
}

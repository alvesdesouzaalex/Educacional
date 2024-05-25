package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	fmt.Println("Start")
	c := cron.New(cron.WithSeconds())

	_, err := c.AddFunc("*/1 * * * * *", func() {
		fmt.Println("Executando a cada 10 segundos:", time.Now().Format("2006-01-02 15:04:05"))
		task()
	})
	if err != nil {
		fmt.Println("Erro ao adicionar função ao cron:", err)
		return
	}

	c.Start()

	// Impede o programa de sair imediatamente
	select {}

}

func task() {
	fmt.Println("Executou a task")
}

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/InstaySystem/is_v2-be/internal/container"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/background/scheduler"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/background/scheduler/job"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	ctn := container.NewContainer(cfg)
	if err := ctn.InitScheduler(); err != nil {
		log.Fatal(err)
	}
	defer ctn.Cleanup()

	sched := scheduler.NewScheduler(ctn.Log.Logger())

	job := job.NewCleanTokenJob(ctn.Log.Logger(), ctn.TokenRepo)

	if err := sched.AddJob("12 22 * * *", job); err != nil {
		log.Println(err)
		return
	}

	sched.Start()
	log.Println("Scheduler is running")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	sched.Stop()
	log.Println("Scheduler stopped successfully")
}

package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/gogits/cron"
	"github.com/iuriikogan/k8s-file-churner/config"
	"github.com/iuriikogan/k8s-file-churner/utils"
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"
)

var rootCmd = &cobra.Command{
	Use:   "chrnr-cli",
	Short: "chrnr-cli",
	Run:   runCommand,
}

func runCommand(cmd *cobra.Command, args []string) {
	// start the timer
	start := time.Now()

	// // Load the config
	// cfg, err := config.LoadConfig()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	cfg := config.Config{
		SizeOfFileMB:         1,
		SizeOfPVCGB:          2,
		ChurnPercentage:      0.1,
		ChurnIntervalMinutes: time.Duration(1 * time.Minute),
		ChurnDuration:        time.Duration(10 * time.Minute),
	}

	// Create the app/testfiles directory
	err := os.MkdirAll("app/testfiles", 0755)
	if err != nil {
		log.Fatal(err.Error())
	}

	// typecast ChurnIntervalMinutes to time.Duration
	printChurnInterval := time.Duration.Minutes(cfg.ChurnIntervalMinutes)

	log.Printf("Churn interval in minutes: %v\n", printChurnInterval)

	// log stuff
	log.Printf("************************************\nK8s File Churner was made by Iurii Kogan - koganiurii@gmail.com \n************************************\n")
	log.Println("Starting K8s File Churner...\nAll testfiles will be written to app/testfiles directory")
	log.Printf("Size of each file in Mb: %d\n", cfg.SizeOfFileMB)
	log.Printf("Size of PVC in Gb: %d\n", cfg.SizeOfPVCGB)
	log.Printf("Churn percentage: %v\n", (cfg.ChurnPercentage * 100))
	log.Printf("Churn interval in minutes: %v\n", printChurnInterval)
	log.Printf("Churn duration in minutes: %v\n", time.Duration.Minutes(cfg.ChurnDuration))

	// calculate number of files to create
	sizeOfPVCMB := int(cfg.SizeOfPVCGB * 999)
	numberOfFiles := ((sizeOfPVCMB) / (cfg.SizeOfFileMB))
	log.Printf("Number of files to create: %d\n", numberOfFiles)
	fileSizeBytes := int(cfg.SizeOfFileMB * 1024 * 1024)

	// start creating the files
	var wg sync.WaitGroup
	wg.Add(numberOfFiles)
	c := cron.New()
	c.AddFunc("@every", "2m", func() {
		log.Println("waiting for files to be created")
	})
	c.Start()

	// Launch a goroutine for each file creation
	for i := 0; i < numberOfFiles; i++ {
		go utils.CreateFileWithRandomData(fileSizeBytes, i, &wg)
	}

	// Wait for all the goroutines to finish
	wg.Wait()
	c.Stop()

	// once all files are created, set the live probe
	utils.Live()

	// log the number of files created, their size and the time it took
	log.Printf("Created %v files of size %vMb\nTook %s\n", numberOfFiles, cfg.SizeOfFileMB, time.Since(start))

	// start churning the files
	churnTicker := time.NewTicker(cfg.ChurnIntervalMinutes)
	go func() {
		log.Printf("Churning %v percent of files every %v minutes", (cfg.ChurnPercentage * 100), printChurnInterval)

		// Set the churn end time
		churnEndTime := time.Now().Add(cfg.ChurnDuration)

		for {
			select {
			case <-churnTicker.C:
				utils.ChurnFiles(cfg.ChurnPercentage, fileSizeBytes, &wg)
			case <-time.After(120 * time.Second):
				log.Println("Waiting to churn files")
			}
			// Check if churn duration has elapsed
			if time.Now().After(churnEndTime) {
				break
			}
		}
	}()
	<-make(chan struct{})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

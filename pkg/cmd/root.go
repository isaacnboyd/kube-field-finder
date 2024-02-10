package cmd

import (
	"errors"
	"fmt"
	"os"
)

func Execute() {
	fmt.Println("You did it!")
	if err := rootCmd.Execute(); err != nil {
		logrus.Errorf("An error has occurred: %v", err)
		os.Exit(1)
	}
}

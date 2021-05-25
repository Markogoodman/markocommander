package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Markogoodman/markocommander/internal/timer"
	"github.com/spf13/cobra"
)

func NewTimeCmd() *cobra.Command {
	var timeCmd = &cobra.Command{
		Use:   "time",
		Short: "Handle time format",
		Long:  "Handle time format",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	timeCmd.AddCommand(NewNowTimeCmd(), NewCalculateTimeCmd())
	return timeCmd
}

func NewNowTimeCmd() *cobra.Command {
	var nowTimeCmd = &cobra.Command{
		Use:   "now",
		Short: "Get current time",
		Long:  "Get current time",
		Run: func(cmd *cobra.Command, args []string) {
			now := timer.GetNowTime()
			log.Printf("Output: %s, %d", now.Format("2006-01-02 15:04:05"), now.Unix())
		},
	}

	return nowTimeCmd
}

func NewCalculateTimeCmd() *cobra.Command {
	var (
		calculateTime string
		duration      string
	)
	var calculateTimeCmd = &cobra.Command{
		Use:   "calc",
		Short: "Calculate time",
		Long:  "Calculate time",
		Run: func(cmd *cobra.Command, args []string) {
			var currentTimer time.Time
			var layout = time.RFC3339
			if calculateTime == "" {
				currentTimer = timer.GetNowTime()
			} else {
				var err error
				space := strings.Count(calculateTime, " ")
				switch space {
				case 0:
					layout = "2006-01-02"
				case 1:
					layout = "2006-01-02 15:04:05"
				}
				currentTimer, err = time.Parse(layout, calculateTime)
				if err != nil {
					t, _ := strconv.Atoi(calculateTime)
					currentTimer = time.Unix(int64(t), 0)
				}
			}
			t, err := timer.GetCalculateTime(currentTimer, duration)
			if err != nil {
				log.Fatalf("timer.GetCalculateTime err: %v", err)
				return
			}
			log.Printf("Output: %s, %d", t.Format(layout), t.Unix())
		},
	}
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "Time need to be calculated")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `Duration, "ns", "us", "ms", "s", "m", "h"`)
	return calculateTimeCmd
}

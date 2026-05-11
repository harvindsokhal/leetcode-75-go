package tracker

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/harvindsokhal/leetcode-75-go/internal/problems"
)

const ProgressFile = "progress.json"

type ProblemProgress struct {
	Slug            string     `json:"slug"`
	Status          string     `json:"status"`
	StartedAt       *time.Time `json:"started_at,omitempty"`
	FinishedAt      *time.Time `json:"finished_at,omitempty"`
	DurationMinutes int        `json:"duration_minutes,omitempty"`
}

type Progress map[string]ProblemProgress

func Load() (Progress, error) {
	progress := Progress{}

	if _, err := os.Stat(ProgressFile); os.IsNotExist(err) {
		return progress, nil
	}

	data, err := os.ReadFile(ProgressFile)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return progress, nil
	}

	if err := json.Unmarshal(data, &progress); err != nil {
		return nil, err
	}

	return progress, nil
}

func Save(progress Progress) error {
	data, err := json.MarshalIndent(progress, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(ProgressFile, append(data, '\n'), 0o644)
}

func Start(slug string) error {
	progress, err := Load()
	if err != nil {
		return err
	}

	now := time.Now()

	current, exists := progress[slug]
	if exists && current.StartedAt != nil && current.FinishedAt == nil {
		return fmt.Errorf("problem %q is already in progress", slug)
	}

	progress[slug] = ProblemProgress{
		Slug:      slug,
		Status:    "in_progress",
		StartedAt: &now,
	}

	return Save(progress)
}

func Finish(slug string) error {
	progress, err := Load()
	if err != nil {
		return err
	}

	current, exists := progress[slug]
	if !exists || current.StartedAt == nil {
		return fmt.Errorf("problem %q has not been started", slug)
	}

	now := time.Now()
	duration := int(now.Sub(*current.StartedAt).Minutes())

	current.Status = "completed"
	current.FinishedAt = &now
	current.DurationMinutes = duration

	progress[slug] = current

	return Save(progress)
}

func In_progress(slug string) error {
	progress, err := Load()
	if err != nil {
		return err
	}

	current, exists := progress[slug]
	if !exists || current.StartedAt == nil {
		return fmt.Errorf("problem %q has not been started", slug)
	}

	now := time.Now()
	duration := int(now.Sub(*current.StartedAt).Minutes())

	current.Status = "in_progress"
	current.FinishedAt = nil
	current.DurationMinutes = duration

	progress[slug] = current

	return Save(progress)
}

func Status() error {
	progress, err := Load()
	if err != nil {
		return err
	}

	completed := 0
	inProgress := 0

	for _, item := range progress {
		switch item.Status {
		case "completed":
			completed++
		case "in_progress":
			inProgress++
		}
	}

	total := len(problems.All())

	fmt.Println("LeetCode 75 Progress")
	fmt.Println("--------------------")

	fmt.Printf("Completed: %d/%d\n", completed, total)
	fmt.Printf("In Progress: %d\n", inProgress)
	fmt.Printf("Remaining: %d\n", total-completed)

	return nil
}

func Stats() error {
	progress, err := Load()
	if err != nil {
		return err
	}

	completed := 0
	inProgress := 0
	totalMinutes := 0
	fastest := 0
	slowest := 0

	difficultyTotals := map[string]int{}
	difficultyCompleted := map[string]int{}

	for _, problem := range problems.All() {
		difficultyTotals[problem.Difficulty]++
	}

	for _, item := range progress {
		switch item.Status {
		case "completed":
			completed++

			problem, ok := problems.FindBySlug(item.Slug)
			if ok {
				difficultyCompleted[problem.Difficulty]++
			}

			if item.DurationMinutes > 0 {
				totalMinutes += item.DurationMinutes

				if fastest == 0 || item.DurationMinutes < fastest {
					fastest = item.DurationMinutes
				}

				if item.DurationMinutes > slowest {
					slowest = item.DurationMinutes
				}
			}

		case "in_progress":
			inProgress++
		}
	}

	total := len(problems.All())
	remaining := total - completed

	fmt.Println("LeetCode 75 Stats")
	fmt.Println("-----------------")
	fmt.Printf("Completed: %d/%d\n", completed, total)
	fmt.Printf("In Progress: %d\n", inProgress)
	fmt.Printf("Remaining: %d\n", remaining)

	fmt.Println()
	fmt.Println("By Difficulty")
	fmt.Println("-------------")
	fmt.Printf("Easy: %d/%d\n", difficultyCompleted["Easy"], difficultyTotals["Easy"])
	fmt.Printf("Medium: %d/%d\n", difficultyCompleted["Medium"], difficultyTotals["Medium"])
	fmt.Printf("Hard: %d/%d\n", difficultyCompleted["Hard"], difficultyTotals["Hard"])

	fmt.Println()

	if completed == 0 || totalMinutes == 0 {
		fmt.Println("No completed solve times yet.")
		return nil
	}

	average := totalMinutes / completed

	fmt.Printf("Average solve time: %s\n", formatMinutes(average))
	fmt.Printf("Fastest solve: %s\n", formatMinutes(fastest))
	fmt.Printf("Slowest solve: %s\n", formatMinutes(slowest))

	return nil
}

func formatMinutes(minutes int) string {
	if minutes < 60 {
		return fmt.Sprintf("%dm", minutes)
	}

	hours := minutes / 60
	remainingMinutes := minutes % 60

	if remainingMinutes == 0 {
		return fmt.Sprintf("%dh", hours)
	}

	return fmt.Sprintf("%dh %dm", hours, remainingMinutes)
}

func Current() error {
	progress, err := Load()
	if err != nil {
		return err
	}

	for _, item := range progress {
		if item.Status == "in_progress" {
			problem, ok := problems.FindBySlug(item.Slug)
			if !ok {
				fmt.Println(item.Slug)
				return nil
			}

			fmt.Println("Current problem")
			fmt.Println("---------------")
			fmt.Printf("%03d. %s [%s] - %s\n",
				problem.Number,
				problem.Title,
				problem.Difficulty,
				problem.Category,
			)
			fmt.Printf("Slug: %s\n", problem.Slug)
			fmt.Printf("Link: https://leetcode.com/problems/%s/\n", problem.Slug)

			if item.StartedAt != nil {
				elapsed := int(time.Since(*item.StartedAt).Minutes())
				fmt.Printf("Elapsed: %s\n", formatMinutes(elapsed))
			}

			return nil
		}
	}

	fmt.Println("No problem currently in progress.")
	return nil
}

func Next() error {
	progress, err := Load()
	if err != nil {
		return err
	}

	for _, problem := range problems.All() {
		item, exists := progress[problem.Slug]

		if !exists || item.Status != "completed" {
			fmt.Println("Next problem")
			fmt.Println("------------")
			fmt.Printf("%03d. %s [%s] - %s\n",
				problem.Number,
				problem.Title,
				problem.Difficulty,
				problem.Category,
			)
			fmt.Printf("Slug: %s\n", problem.Slug)
			fmt.Printf("Link: https://leetcode.com/problems/%s/\n", problem.Slug)
			return nil
		}
	}

	fmt.Println("All problems completed.")
	return nil
}

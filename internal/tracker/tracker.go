package tracker

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
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

	fmt.Println("LeetCode 75 Progress")
	fmt.Println("--------------------")
	fmt.Printf("Completed: %d/75\n", completed)
	fmt.Printf("In Progress: %d\n", inProgress)
	fmt.Printf("Remaining: %d\n", 75-completed)

	return nil
}

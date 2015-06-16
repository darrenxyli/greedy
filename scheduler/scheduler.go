package scheduler

import (
	"fmt"
)

const (
	defaultLoopInterval = 0.1
)

// Scheduler to schedule
type Scheduler struct {
	projects     []string
	loopInterval float32
}

// NewScheduler create a Scheduler
func NewScheduler(projects []string) *Scheduler {
	return &Scheduler{projects: projects, loopInterval: defaultLoopInterval}
}

// loadTasks
func (scheduler *Scheduler) loadTasks() {
	fmt.Println("loadTask")
}

// Run a schedule
func (scheduler *Scheduler) Run() {
	fmt.Println("Run")
}

// UpdateTask
func (scheduler *Scheduler) updateTasks() {
     fmt.Println("god test")
}
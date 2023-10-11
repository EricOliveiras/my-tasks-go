package builders

import "github/ericoliveiras/basic-crud-go/internal/models"

type TaskBuilder struct {
	ID          string
	UserId      string
	Title       string
	Description string
	Finished    bool
}

func NewTaskBuilder() *TaskBuilder {
	return &TaskBuilder{}
}

func (taskBuilder *TaskBuilder) SetID(id string) *TaskBuilder {
	taskBuilder.ID = id
	return taskBuilder
}

func (taskBuilder *TaskBuilder) SetUserId(userId string) *TaskBuilder {
	taskBuilder.UserId = userId
	return taskBuilder
}

func (taskBuilder *TaskBuilder) SetTitle(title string) *TaskBuilder {
	taskBuilder.Title = title
	return taskBuilder
}

func (taskBuilder *TaskBuilder) SetDescription(description string) *TaskBuilder {
	taskBuilder.Description = description
	return taskBuilder
}

func (taskBuilder *TaskBuilder) SetFinished(finished bool) *TaskBuilder {
	taskBuilder.Finished = finished
	return taskBuilder
}

func (taskBuilder *TaskBuilder) Build() models.Task {
	task := models.Task{
		ID:          taskBuilder.ID,
		UserId:      taskBuilder.UserId,
		Title:       taskBuilder.Title,
		Description: taskBuilder.Description,
		Finished:    taskBuilder.Finished,
	}

	return task
}

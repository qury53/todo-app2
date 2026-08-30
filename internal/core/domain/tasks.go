package domain

import (
	"fmt"
	"time"
	"unicode/utf8"

	core_errors "github.com/qury53/todo-app2/internal/core/errors"
)

type Task struct {
	ID          int
	Version     int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	AuthorID    int
}

type TaskPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Completed   Nullable[bool]
}

func NewTaskPatch(
	title Nullable[string],
	description Nullable[string],
	completed Nullable[bool],
) TaskPatch {
	return TaskPatch{
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}

func NewTaskUninitialized(title string,
	description string,
	completed bool,
	authorID int,
) Task {
	return Task{
		ID:          taskIDUninitialized,
		Version:     taskVersionUninitialized,
		Title:       title,
		Description: description,
		Completed:   completed,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
		AuthorID:    authorID,
	}
}

func (t *Task) Validate() error {
	titleLen := utf8.RuneCountInString(t.Title)
	if titleLen < 1 || titleLen > 100 {
		return fmt.Errorf(
			"title length should be between 1 and 100: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (p *TaskPatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf(
			"title cannot be patched to null: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if p.Completed.Set && p.Completed.Value == nil {
		return fmt.Errorf(
			"completed cannot be patched null: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (t *Task) ApplyPatch(patch TaskPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("failed to validate patch: %w", err)
	}

	tmp := *t

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}

	if patch.Description.Set {
		tmp.Description = *patch.Description.Value
	}

	if patch.Completed.Set {
		tmp.Completed = *patch.Completed.Value
		if tmp.Completed {
			completedAt := time.Now()
			tmp.CompletedAt = &completedAt
		} else {
			tmp.CompletedAt = nil
		}
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("failed to validate patched task: %w", core_errors.ErrInvalidArgument)
	}

	*t = tmp

	return nil
}

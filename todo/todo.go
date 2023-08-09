package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item
type Stringer interface {
	String() string
}

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

func (l *List) Update(i int, task string) error {
	ls := *l
	if i < 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	ls[i-1].Task = task

	return nil
}

func (l *List) Search(task string) bool {
	ls := *l

	for i := 1; i <= len(ls); i++ {
		if ls[i-1].Task == task {
			return true
		}
	}
	return false
}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}

func (l *List) String() string {
	formatted := ""

	for k, t := range *l {
		prefix := " "
		if t.Done {
			prefix = "X"
		}

		formatted += fmt.Sprintf("%s %d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}

func (l List) PrintTodos() List {
	ls := l
	for i, v := range ls {
		if v.Done {
			if i < len(ls) {
				ls = append(ls[:i], ls[i+1:]...)
			} else {
				ls = ls[:i-1]
			}
		}
	}
	return ls
}

package task

import "github.com/johnnyeven/libtools/task/constants"

type Backend struct {

}

func NewBackend() *Backend {
	return &Backend{}
}

func (b *Backend) Feedback(f *constants.TaskFeedback) error {
	return nil
}

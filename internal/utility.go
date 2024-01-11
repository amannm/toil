package internal

import (
	"errors"
	"sync"
)

type Task func() error

func ExecuteAll(tasks ...Task) error {
	numTasks := len(tasks)
	wg := new(sync.WaitGroup)
	wg.Add(numTasks)
	errc := make(chan error, numTasks)
	for i := 0; i < numTasks; i++ {
		go func(t Task) {
			defer wg.Done()
			errc <- t()
		}(tasks[i])
	}
	wg.Wait()
	close(errc)
	errs := make([]error, 0)
	for err := range errc {
		errs = append(errs, err)
	}
	return errors.Join(errs...)
}

package listener

import "errors"

type Qt struct {
	exit    chan interface{}
	updates chan *int
}

var ErrWatcherStopped = errors.New("watcher stopped")

func (w *Qt) Next() (*int, error) {
	select {
	case <-w.exit:
		return nil, ErrWatcherStopped
	case v := <-w.updates:
		return v, nil
	}
}

func (w *Qt) Stop() error {
	select {
	case <-w.exit:
	default:
		close(w.exit)
	}
	return nil
}

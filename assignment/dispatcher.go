package main

var WorkerQueue chan chan WorkRequest //creates a worker queue to keep track of all the workers

func StartDispatcher(nworkers int) {

	WorkerQueue = make(chan chan WorkRequest, *workersno)

	for i := 0; i < nworkers; i++ {
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	//an infinite loop that keeps on assigning work to worker
	go func() {
		for {
			select {
			case work := <-WorkQueue:
				go func() {
					worker := <-WorkerQueue
					worker <- work
				}()
			}
		}
	}()
}

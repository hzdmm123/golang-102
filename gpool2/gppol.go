package gpool2

type Worker struct {
	WorkerPool chan *Worker
	JobQueue   chan Job
	stop       chan struct{}
}

func newWorker(pool chan *Worker) *Worker {

	return &Worker{
		WorkerPool: pool,
		JobQueue:   make(chan Job),
		stop:       make(chan struct{}),
	}
}

func (w *Worker) Start() {
	go func() {
		var job Job
		for {
			w.WorkerPool <- w
			select {
			case job = <-w.JobQueue:
				job()
			case <-w.stop:
				w.stop <- struct{}{}
				return
			}
		}
	}()
}

type Job func()

type Pool struct {
	JobQueue   chan Job
	WorkerPool chan *Worker
	Stop       chan struct{}
}

func NewPool(workerNums int, JobNums int) {
	wp := make(chan *Worker, workerNums)
	jn := make(chan Job, JobNums)

	pool := &Pool{
		JobQueue:   jn,
		WorkerPool: wp,
		Stop:       make(chan struct{}),
	}

	pool.Start()
	return
}

func (p *Pool) Start() {
	for i := 0; i < cap(p.WorkerPool); i++ {
		worker := newWorker(p.WorkerPool)
		worker.Start()
	}

	go p.dispatch()
}

func (p *Pool) dispatch() {

	for {
		select {
		case job := <-p.JobQueue:
			worker := <-p.WorkerPool
			worker.JobQueue <- job
		case <-p.Stop:
			for i := 0; i < cap(p.WorkerPool); i++ {
				worker := <-p.WorkerPool
				worker.stop <- struct{}{}
				<-worker.stop
			}

			p.Stop <- struct{}{}
			return

		}
	}
}

func (p *Pool) Release() {
	p.Stop <- struct{}{}
	<-p.Stop
}

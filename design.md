# design

* queue all download jobs
* workers get started up, number of workers are configure
* workers check job queue for work, if there's work they do it
* when all the jobs have been cleared off the job queue, then the
worker can end, this is because we pre-loaded all jobs before workers
started 
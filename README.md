# WorkerPool / ThreadPool

This step (step1) is like setting up the basic infrastructure always do this at first.

step 1 -> Prepare pool add/spawn X number of workers/go routines and give them access to channel to continuously fetch jobs/work.

precaution : use waitgroups to manage go routines wg.Add() and wg.Done()

step 2 -> prepare jobs as per the problem and keep adding them to the channel.

step 3 -> Dont Forget to Close the channel after (all the jobs are pushed) to avoid blocking/deadlocks and to wait for all go routine to finish using wg.Wait().

## Insights

### Tuning size of pool
-> spawing n goroutines will chok the CPU and ending up using too much of resources.

-> pool size should not be too large or too small (not optimal usage of resources).

-> optimal pool size depends on hardware, task load (I/P bound, CPU load etc) on the basis of that make intelligent guess/ iterative approach.

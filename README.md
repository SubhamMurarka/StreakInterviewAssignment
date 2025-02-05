# WorkerPool / ThreadPool

step 1 -> Prepare pool add/spawn X number of workers/go routines and give them access to channel to continuously fetch jobs/work.

precaution : use waitgroups to manage go routines wg.Add() and wg.Done()

step 2 -> prepare jobs as per the problem and keep adding them to the channel.

step 3 -> Dont Forget to Close the channel to avoid blocking/deadlocks and to wait for all go routine to finish using wg.Wait().

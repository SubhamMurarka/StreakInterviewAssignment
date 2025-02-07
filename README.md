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

# FanIn
// TODO

# FanOut
// TODO

# Pipeline

In pipeline pattern, we do synchronisation with help of channels. Waitgroups, time.Sleep() etc not required.

step 1 -> First decide and break your task into multiple sequential steps (example : Processing of some data).

step 2 -> Number of channels = Number of pipeline connections. Previous stage writes and next stage reads.

step 3 -> Next stage keeps reading until previous channel is writing or not closed. this is followed in the complete chain/pipeline.

step 4 -> Main routine finally reads from last channel and hence no pre-exit and proper synchronisation of routines.

## Insights

### Increasing performance of a stage

-> performance of a particular stage can be increased by spawing multiple routine/ or using other patterns discussed above, hence increasing performance of that step.

-> no need to store intermidiate results hence can be helpful for large datasets processing.

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

-> performance of a particular stage can be increased by spawing multiple routine/ or using other patterns discussed here, hence increasing performance of that step.

-> no need to store intermidiate results hence can be helpful for large datasets processing.

# Semaphore

Mutex allows single go routine to acquire a resource at a time, what if N go routines want to acquire resource at a time here semaphores come.

### Implementation 1 (From Scratch)

step 1 -> A semaphore is just a buffered channel so create a buffered channel, size of it defines how many go routines can require a particular resource.

step 2 -> Whenever a go routine is fired try to book a slot in channel ow keep waiting.

step 3 -> when work is done release the slot.

step 4 -> Use waitgroups to avoid pre-exit of main go routine.

### Implementation 2 (Using Sync Package)

step 1 -> Declare an object of semaphore struct with defining size as above.

step 2 -> Acquire a slot, context is passed to maintain how much time to weight for booking a slot.

step 3 -> can also use tryaquire() which acquires slot if possible otherwise return false immediately without waiting.

step 4 -> release the slot after work is done.

step 5 -> Use waitgroups to avoid pre-exit of main go routine.

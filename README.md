# worker-pool
Golang simple worker pool

## Task Requirements:
 
Create a Golang Worker Function:
 
Create a worker function that accepts the following:
- Some data (e.g., an integer or string).
- A function to process the data (e.g., a function that performs a mathematical operation or string manipulation).

Spawn 10 Workers:
- Create 10 worker instances running in separate Go routines.
- Each worker should process a piece of data using the given function.
- Wait for All Workers to Complete:
 
Use Go's sync.WaitGroup to wait for all workers to finish their tasks before the main function exits. Try to understand how it works.

Test the Program:
Provide an example of data (e.g., a list of integers) and a function (e.g., square each number or append some text to a string) to demonstrate the solution.


example output of finished program:

Worker 3 processed data 4 -> result 16
Worker 7 processed data 8 -> result 64
Worker 8 processed data 9 -> result 81
Worker 5 processed data 6 -> result 36
Worker 4 processed data 5 -> result 25
Worker 1 processed data 2 -> result 4
Worker 0 processed data 1 -> result 1
Worker 9 processed data 10 -> result 100
Worker 2 processed data 3 -> result 9
Worker 6 processed data 7 -> result 49
All workers have finished processing.

order of completion is not guaranteed with go routines so different workers can output result at different time
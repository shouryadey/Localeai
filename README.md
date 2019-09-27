#### PROBLEM: To insert data in postgreSQL database and generating suitable response while managing burst requests

#### PROBLEM: To insert data in postgres SQL database and generating suitable response while managing burst requests

#### SOLUTION: The Idea is to queue requests(work) , use workers to process requests , dispatchers to assign work to workers,provide response using shared cache(to keep track of the request processed) 

#### Key-Terms:
###### Collector:
        Gets request and pushes the required information to the queue(WorkQueue)
###### Work Queue:
        Use a buffered channel so the collector is not blocked
###### Dispatcher:
        Dipatcher distributes work among workers from WorkQueue
###### Workers:
        Responsible for storing data in database  and updating shared cache
###### Shared cache:
        Used to keep track of the request stored in the database ,once the response is generated the ID is removed from the cache


#### RESULTS :Load tested with Vegeta, Capable of processing 20-24 concurrent requests/s with timeouts of 60s
    

#### RESULTS :Load tested with Vegeta, Capable of processing 20-24 concurrent requests/s with timeouts of 30s
    


#### PROBLEM: To insert data in postgreSQL database and generating suitable response while managing burst requests

#### PROBLEM: To insert data in postgres SQL database and generating suitable response while managing burst requests

#### SOLUTION: The Idea is to queue requests(work) , use workers to process requests , dispatchers to assign work to workers,provide response using shared cache(to keep track of the request processed) 

#### Key-Terms:
###### Collector:
        Gets request and pushes the required information to the queue(WorkQueue)
###### Work Queue:
        Use a buffered channel so the collector is not blocked
###### Dispatcher:
        Dispatcher distributes work among workers from WorkQueue
###### Workers:
        Responsible for storing data in database  and updating shared cache
###### Shared cache:
        Used to keep track of the request stored in the database ,once the response is generated the ID is removed from the cache


#### RESULTS :Load tested with Vegeta, Capable of processing 20-24 concurrent requests/s with timeouts of 60s
#### IMPROVEMENTS:
Further improvements can be made by having more servers with 
        
 1)Load balancing

 2) Instead of waiting for the response and showing timeout after a time interval ,we can also use an event triggered (event being a successful insertion in database) solution,where we can generate a SMS response to the user whenever a data is inserted in the database.

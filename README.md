### THE AIM IS TO INSERT DATA IN POSTGRES SQL AND PROVIDE SUITABLE RESPONSE WHILE MANAGING BURST OF REQUESTS

## The Idea is to queue requests(work) , use workers to process requests , dispatchers to assign work to workers,provide response using shared cache(to keep track of the request processed) 

###### Collector:
        Gets request and pushes the required information to the queue(WorkQueue)
###### Work Queue:
        Use a buffered channel so the colector is not blocked
###### Dispatcher:
        Dipatcher distributes work among workers from WorkQueue
###### Workers:
        Responsible for storing data in database  and updating shared cache
###### Shared cache:
        Used to keep track of the request stored in the database ,once the response is generated the ID is removed from the cache
    

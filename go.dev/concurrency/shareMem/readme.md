----------- ## URL POLLER --------------
is a program or a component of a program designed to regularly check("poll") a set of URLs(web addresses) to retreive data or monitor their status. This process involves sending requests to these URLs and handling the responses.
They r used for:
### monitor website availability:
### checking service health:
in microservices architectures or distributed systems, a URL poller might regularly check the health endpoints of various services to ensure they are functioning correctly

### data retrieval:
polling APIs or web services at regular intervals to fetch updated data. this is common in scenarios where real-time data streaming isnt available or necessary

### content change detection:

-- key characteristics of URL pollers:
### automated requests:
 http req(like get or head)
 ### handling responses
 ### scheduling, concurrency, error handling and retry logic

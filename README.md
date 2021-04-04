# Simple Web Tracker

__Disclaimer__: This is not a production ready tracking solution. It was created is being maintained with the commitment of an exploratory side project. That being said feel free to play around with it and get in touch if you'd like to collaborate.

## Why
I wanted to build a tracking solution for small sites that want to learn more about their users behavior without disrespecting their privacy or sending their data to Google.

## How
* Tracking service using Golang and Gin
* Small JS snippet to send events
* Postgres for storage (I don't think this is going to scale, I am thinking about adding a redis layer and only dumping aggregates into pg)

## Setup
I'll make this more straightforward and config file based once I have a handle on the core functionality.  
* Add allowed origins in `tracker/server/router.go`
* Update DB credentials in `tracker/db/db.go`
* Run postgres on same machine, create `tracker` DB
* Run migration scripts `db/init.sql`
* Build golang project for target platform: `make build`
* Run `./dist/tracker`
* Use js snippet in `js-lib` according to example

## Next steps
[ ] Create simple database views as analytics PoC  
[ ] Create simple visualization based on DB  
[ ] Load test solution with Postgres to check for viability   
[ ] Redis layer  
[ ] Config, deployability, documentation

## Contributing
I am open for collaboration - if you find this project useful/ inspiring/ fun to work on please get in touch via Twitter or Email. 

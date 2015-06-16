Part of the SPA 2015 Distributed Databases session. This is a
single node of a diy distributed database. It's a RESTful API
server that wraps up a simple integer to string map.

## Installation
Executables have been pre-built for most platforms.
You can find them on the 
[downloads page](https://github.com/seeemilyplay/diy-dist-db-node/blob/master/downloads/snapshot/downloads.md). From this page find the distribution relevant to you, follow it's link, and then click on GitHub "Raw" button to download. Then unzip to get
the executable.

## Running
Call the program passing it a free local PORT.

    diy-dist-db-node <PORT>

## Adding things
Add things to the map with HTTP POST. For example:

    curl -i -H 'Content-Type: application/json' \
         -d '{"Id": 3, "Value": "foo"}' \
         http://localhost:<PORT>/things

## Querying things
To query for a particular thing in the map use an HTTP GET
request with the thing's id number.

    curl -i http://localhost:<PORT>/things/3 

## Listing all things
To list everything in the map, use a single HTTP GET request.

    curl -i http://localhost:<PORT>/things

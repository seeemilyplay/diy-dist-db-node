Part of the SPA 2015 Distributed Databases session. This is a
single node of a diy distributed database. It's a RESTful API
server that wraps up a simple integer to string map.

## Installation
This program was written in Go, especially to take advantage of
it's easy cross compilation. Please download and install the
suitable version for yourself from the downloads page.

## Running
Once installed, call the program to run, passing it a local PORT.

    diy-dist-db-node <PORT>

## Adding things
Add things to the map with a HTTP POST. For example:

    curl -i -H 'Content-Type: application/json' \
         -d '{"Id": 3, "Value": "foo"}' \
         http://localhost:<PORT>/things

## Querying things
To query for a particular thing in the map, use a HTTP GET
request with the thing's id number.

    curl -i http://localhost:<PORT>/things/3 

## Listing all things
To list everything in the map, use a single HTTP GET request.

    curl -i http://localhost:<PORT>/things

/**
 * Restful API server that wraps up a simple int to string map.
 * 
 * Usage: diy-dist-db-node 8080
 *
 * To put a record in the map:
 *   curl -i -H 'Content-Type: application/json' \
 *        -d '{"Id": 3, "Value": "foo"}' \
 *        http://localhost:8080/things
 * 
 * To get a specific record from the map:
 *   curl -i http://localhost:8080/things/3
 *
 * To list everything in the map:
 *   curl -i http://localhost:8080/things
 */
package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"os"
	"net/http"
	"strconv"
	"time"
)

type Thing struct {
	Id int
	Value string
	Timestamp int64
}

var things map[int]Thing

func main() {

	things = make(map[int]Thing)
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/things", getAllThings),
		rest.Post("/things", postThing),
		rest.Get("/things/:id", getThing),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	port := "8080" 
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	log.Fatal(http.ListenAndServe(":" + port, api.MakeHandler()))
}

func getAllThings(w rest.ResponseWriter, r *rest.Request) {
	thingList := make([]Thing, len(things))
	i := 0
	for _, thing := range things {
		thingList[i] = thing
		i++
	}
	w.WriteJson(thingList)
}

func getThing(w rest.ResponseWriter, r *rest.Request) {
	id, err := strconv.Atoi(r.PathParam("id"))
	if err != nil {
		rest.NotFound(w, r)
		return
	}
	thing, ok := things[id]
	if !ok {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(thing)
}

func postThing(w rest.ResponseWriter, r *rest.Request) {
	thing := Thing{}
	thing.Timestamp = time.Now().UnixNano() / 1000000 
	err := r.DecodeJsonPayload(&thing)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	things[thing.Id] = thing 
	w.WriteJson(thing)
}

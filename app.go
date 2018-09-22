package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/taskqueue"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create_task", createTaskHandler)
	appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		return nil
	}, nil)

	n := r.URL.Query().Get("n")
	num, err := strconv.Atoi(n)
	if err != nil {
		num = 100
	}

	t := taskqueue.NewPOSTTask("/", url.Values{})
	for i := 0; i < num; i++ {
		_, err := taskqueue.Add(ctx, t, "task-queue-experiment")
		if err != nil {
			log.Printf("%#v", err)
			fmt.Fprintln(w, "NG")
			return
		}
	}
	fmt.Fprintf(w, "Created %d tasks", num)
}

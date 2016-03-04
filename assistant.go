package main

import (
	"fmt"
	"net/http"

	"github.com/gartht/assistant/taskService"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	router.HandleFunc("/get/{listId}", getListById)
	router.HandleFunc("/", indexHandler)
	http.ListenAndServe("0.0.0.0:3000", router)

}

func getListById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	taskList := taskService.GetTaskListById(vars["listId"])
	fmt.Fprintln(response, taskList.Title)
}

func indexHandler(response http.ResponseWriter, request *http.Request) {
	taskLists := taskService.GetLists()
	responseTxt := "No lists found"

	if len(taskLists.Items) > 0 {
		title := taskLists.Items[0].Title
		id := taskLists.Items[0].Id
		responseTxt = title + " " + id
	}
	fmt.Fprintln(response, responseTxt)

}

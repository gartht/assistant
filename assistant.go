package main

import (
	"encoding/json"
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

	svc := router.PathPrefix("/svc").Subrouter()

	router.PathPrefix("/styles/").Handler(http.FileServer(http.Dir("./static/")))
	router.Handle("/", http.FileServer(http.Dir("./static/")))

	appDir := http.StripPrefix("/app/", http.FileServer(http.Dir("./app/")))
	router.PathPrefix("/app/").Handler(appDir)

	nodeDir := http.StripPrefix("/node_modules/", http.FileServer(http.Dir("./node_modules/")))
	router.PathPrefix("/node_modules/").Handler(nodeDir)

	svc.HandleFunc("/", indexHandler)
	svc.HandleFunc("/get/{listId}", getTasksByListId)

	http.ListenAndServe("0.0.0.0:3000", router)
}

func getTasksByListId(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	tasks := taskService.GetTasksFromList(vars["listId"])

	if err := json.NewEncoder(response).Encode(tasks.Items); err != nil {
		panic(err)
	}
}

func getListById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	taskList := taskService.GetTaskListById(vars["listId"])
	fmt.Fprintln(response, taskList.Title)
}

func indexHandler(response http.ResponseWriter, request *http.Request) {
	taskLists := taskService.GetLists()

	if err := json.NewEncoder(response).Encode(taskLists); err != nil {
		panic(err)
	}
}

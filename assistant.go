package main

import (
	"fmt"
	"net/http"

	"github.com/gartht/assistant/taskService"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe("0.0.0.0:3000", nil)

}

func indexHandler(response http.ResponseWriter, request *http.Request) {
	taskLists := taskService.GetLists()
	responseTxt := "No lists found"
	if len(taskLists.Items) > 0 {
		responseTxt = taskLists.Items[0].Title

	}
	fmt.Fprintln(response, responseTxt)

}

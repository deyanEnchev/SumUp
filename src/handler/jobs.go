package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deyanEnchev/src/model"
	"github.com/deyanEnchev/src/service"
)

// HandleJobs handles incoming HTTP POST requests for processing a job. It decodes the JSON payload in the request
// body into a job model, sorts the tasks in the job using topological sort and generates a bash script from the sorted
// tasks. The generated bash script is then written to the response writer in plain text format. In case of any errors,
// such as invalid request method, failed parsing of the request body, or error during topological sort or script
// generation, the function returns an appropriate HTTP error response.
func HandleJobs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var job model.Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse request body: %v", err), http.StatusBadRequest)
		return
	}

	sortedTasks, err := service.TopologicalSort(job)
	if err != nil {
		http.Error(w, fmt.Sprintf("An error occured while sorting: %v", err), http.StatusBadRequest)
		return
	}
	bashScript := service.GenerateBashScript(sortedTasks)

	w.Header().Set("Content-Type", "text/plain")

	_, err = w.Write([]byte(bashScript))
	if err != nil {
		http.Error(w, "Could not encode response as JSON", http.StatusInternalServerError)
		return
	}
}

package addtask

type request struct {
	name string `json:"name"`
}

type response struct {
	task_id int `json:"task_id"`
}

package todo

type Event string
const (
	TodoCreate Event = "TODO_CREATE"
	TodoUpdate Event = "TODO_UPDATE"
	TodoDelete Event = "TODO_DELETE"
	TodoRead   Event = "TODO_READ"
)

type TodoEvent struct {
Todo      Todo   `json:"todo"`
RequestId string `json:"requestId"`
}

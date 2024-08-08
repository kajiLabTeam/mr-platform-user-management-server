package common

type UserId struct {
	UserId string `json:"userId"`
}

type ContentIds struct {
	ContentIds []string `json:"contentIds"`
}

type RequestSetContents struct {
	UserId     string   `json:"userId"`
	ContentIds []string `json:"contentIds"`
}

type ResponseSetContents struct {
	ContentIds []string `json:"contentIds"`
}

type RequestGetContents struct {
	UserId string `json:"userId"`
}

type ResponseGetContents struct {
	ContentIds []string `json:"contentIds"`
}

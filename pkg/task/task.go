package task

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/smallpdf/go-worker"
	"github.com/smallpdf/testing-poc/pkg/config"
)

type Client struct {
	endpoint   string
	localToken string
	httpClient *http.Client
}

type Task struct {
	*worker.Task
}

type Result struct {
	*worker.Result
}

// func (r *Result) ge

type ResultResponse struct {
	Success bool          `json:"success"`
	Data    worker.Result `json:"data"`
}

func NewClient(endpoint string) *Client {
	cfg := config.Load()

	client := http.DefaultClient

	return &Client{
		httpClient: client,
		endpoint:   endpoint,
		localToken: cfg.Tasks.LocalToken,
	}
}

func (c *Client) GetResult(ctx context.Context, taskID string) (*Result, error) {
	var result *Result
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(result)

	return result, err
}

func (c *Client) CreateTask(task *Task) (*ResultResponse, error) {
	bs, err := json.Marshal(&task)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(bs)
	res, err := http.Post(c.endpoint, "application/json", body)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	r := &ResultResponse{}
	err = json.Unmarshal(b, &r)

	return r, err
}

func NewTaskBody(tool string, files []string) *Task {
	id := uuid.New().String()

	return &Task{
		Task: &worker.Task{
			ID:          id,
			Tool:        tool,
			Version:     0,
			UserType:    "user_type",
			InputTokens: files,
		},
	}
}

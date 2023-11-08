package openai

import (
	"context"
	"fmt"
	"net/http"
)

const (
	runs              = "/runs"
	submitToolOutputs = "submit_tool_outputs"
)

type Run struct {
	// The identifier, which can be referenced in API endpoints.
	ID string `json:"id"`
	// The object type, which is always `assistant.run`.
	Object string `json:"object"`
	// The Unix timestamp (in seconds) for when the run was created.
	CreatedAt int32 `json:"created_at"`
	// The ID of the [thread](/docs/api-reference/threads) that was executed on as a part of this run.
	ThreadID string `json:"thread_id"`
	// The ID of the [assistant](/docs/api-reference/assistants) used for execution of this run.
	AssistantID string `json:"assistant_id"`
	// The status of the run, which can be either `queued`, `in_progress`, `requires_action`, `cancelling`, `cancelled`, `failed`, `completed`, or `expired`.
	Status         string         `json:"status"`
	RequiredAction RequiredAction `json:"required_action,omitempty"`
	LastError      LastError      `json:"last_error,omitempty"`
	// The Unix timestamp (in seconds) for when the run will expire.
	ExpiresAt int32 `json:"expires_at"`
	// The Unix timestamp (in seconds) for when the run was started.
	StartedAt int64 `json:"started_at,omitempty"`
	// The Unix timestamp (in seconds) for when the run was cancelled.
	CancelledAt int64 `json:"cancelled_at,omitempty"`
	// The Unix timestamp (in seconds) for when the run failed.
	FailedAt int64 `json:"failed_at,omitempty"`
	// The Unix timestamp (in seconds) for when the run was completed.
	CompletedAt int64 `json:"completed_at,omitempty"`
	// The model that the [assistant](/docs/api-reference/assistants) used for this run.
	Model string `json:"model"`
	// The instructions that the [assistant](/docs/api-reference/assistants) used for this run.
	Instructions string `json:"instructions"`
	// The list of tools that the [assistant](/docs/api-reference/assistants) used for this run.
	Tools []RunTool `json:"tools"`
	// The list of [File](/docs/api-reference/files) IDs the [assistant](/docs/api-reference/assistants) used for this run.
	FileIds []string `json:"file_ids"`
	// metadata_description
	Metadata map[string]interface{} `json:"metadata"`

	httpHeader
}

type RequiredAction struct {
	Type              string            `json:"type"`
	SubmitToolOutputs SubmitToolOutputs `json:"submit_tool_outputs"`
}

type RunTool struct {
	Type string `json:"type"`
}

type SubmitToolOutputs struct {
	ToolCalls []ToolCall `json:"tool_calls"`
}

type RunToolCall struct {
	ID       string   `json:"id"`
	Type     string   `json:"type"`
	Function Function `json:"function"`
}

type Function struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type LastError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type RunRequest struct {
	Run Run `json:"run"`
}

type CreateRunRequest struct {
	AssistantID string `json:"assistant_id"`
}

type ToolOutputs struct {
	// The ID of the tool call in the `required_action` object within the run object the output is being submitted for.
	ToolCallID string `json:"tool_call_id,omitempty"`
	// The output of the tool call to be submitted to continue the run.
	Output string `json:"output,omitempty"`
}

type ToolOutputsRequest struct {
	ToolOutputs []ToolOutputs `json:"tool_outputs"`
}

// CreateRun creates a new run on the provided thread.
func (c *Client) CreateRun(ctx context.Context, threadID string, request CreateRunRequest) (response Run, err error) {
	path := fmt.Sprintf("%s/%s/%s", threadsSuffix, threadID, runs)
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(path), withBody(request), func(args *requestOptions) {
		args.header.Set("OpenAI-Beta", "assistants=v1")
	})
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

// RetrieveRun retrieves a run for the provided thread.
func (c *Client) RetrieveRun(ctx context.Context, threadID, runID string) (response Run, err error) {
	path := fmt.Sprintf("%s/%s/%s/%s", threadsSuffix, threadID, runs, runID)
	req, err := c.newRequest(ctx, http.MethodGet, c.fullURL(path), func(args *requestOptions) {
		args.header.Set("OpenAI-Beta", "assistants=v1")
	})
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

// SubmitToolOutputs adds tool outputs to the run for a thread.
func (c *Client) SubmitToolOutputs(ctx context.Context, threadID, runID string, request ToolOutputsRequest) (response Run, err error) {
	path := fmt.Sprintf("%s/%s/%s/%s/%s", threadsSuffix, threadID, runs, runID, submitToolOutputs)
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(path), withBody(request), func(args *requestOptions) {
		args.header.Set("OpenAI-Beta", "assistants=v1")
	})
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

//
//// ModifyThread modifies a thread.
//func (c *Client) ModifyThread(
//	ctx context.Context,
//	threadID string,
//	request ModifyThreadRequest,
//) (response Thread, err error) {
//	urlSuffix := threadsSuffix + "/" + threadID
//	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(urlSuffix), withBody(request))
//	if err != nil {
//		return
//	}
//
//	err = c.sendRequest(req, &response)
//	return
//}
//
//// DeleteThread deletes a thread.
//func (c *Client) DeleteThread(ctx context.Context, threadID string) (err error) {
//	urlSuffix := threadsSuffix + "/" + threadID
//	req, err := c.newRequest(ctx, http.MethodDelete, c.fullURL(urlSuffix))
//	if err != nil {
//		return
//	}
//
//	err = c.sendRequest(req, nil)
//	return
//}

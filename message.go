package openai

import (
	"context"
	"fmt"
	"net/http"
)

const (
	messages = "/messages"
)

type Message struct {
	// The identifier, which can be referenced in API endpoints.
	ID string `json:"id"`
	// The object type, which is always `thread.message`.
	Object string `json:"object"`
	// The Unix timestamp (in seconds) for when the message was created.
	Created int64 `json:"created_at"`
	// The [thread](/docs/api-reference/threads) ID that this message belongs to.
	ThreadID string `json:"thread_id"`
	// The entity that produced the message. One of `user` or `assistant`.
	Role string `json:"role"`
	// The content of the message in array of text and/or images.
	// TODO
	//Content []OneOfMessageContentImageFileObjectMessageContentTextObject `json:"content"`
	// If applicable, the ID of the [assistant](/docs/api-reference/assistants) that authored this message.
	AssistantID string `json:"string,omitempty"`
	// If applicable, the ID of the [run](/docs/api-reference/runs) associated with the authoring of this message.
	RunID string `json:"run_id,omitempty"`
	// A list of [file](/docs/api-reference/files) IDs that the assistant should use. Useful for tools like retrieval and code_interpreter that can access files. A maximum of 10 files can be attached to a message.
	FileIds []string `json:"file_ids"`
	// metadata_description
	Metadata map[string]any `json:"metadata"`

	httpHeader
}

type MessageRequest struct {
	Message ThreadMessage `json:"message"`
}

// CreateMessage creates a new message on the provided thread.
func (c *Client) CreateMessage(ctx context.Context, threadID string, request MessageRequest) (response Thread, err error) {
	postPath := fmt.Sprintf("%s/%s/%s", threadsSuffix, threadID, messages)
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(postPath), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

//
//// RetrieveThread retrieves a thread.
//func (c *Client) RetrieveThread(ctx context.Context, threadID string) (response Thread, err error) {
//	urlSuffix := threadsSuffix + "/" + threadID
//	req, err := c.newRequest(ctx, http.MethodGet, c.fullURL(urlSuffix))
//	if err != nil {
//		return
//	}
//
//	err = c.sendRequest(req, &response)
//	return
//}
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

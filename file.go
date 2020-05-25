package backlog

import (
	"context"
	"net/url"
)

// FileUploadResponse : response of uploading file
type FileUploadResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Size int    `json:"size"`
}

// UploadFile uploads a file
func (api *Client) UploadFile(fpath string) (*FileUploadResponse, error) {
	return api.UploadFileContext(context.Background(), fpath)
}

// UploadFileContext uploads a file and setting a custom context
func (api *Client) UploadFileContext(ctx context.Context, fpath string) (*FileUploadResponse, error) {
	fileUploadResponse := FileUploadResponse{}
	if err := postLocalWithMultipartResponse(ctx, api.httpclient, api.endpoint+"/api/v2/space/attachment?apiKey="+api.apiKey, fpath, "file", url.Values{}, &fileUploadResponse, api); err != nil {
		return nil, err
	}
	return &fileUploadResponse, nil
}
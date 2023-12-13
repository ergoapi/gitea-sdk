// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CreateOrUpdateSecretOption struct {
	Description string `json:"description,omitempty"`
	Data        string `json:"data"`
}

// PushMirrorResponse returns a git push mirror
type PushMirrorResponse struct {
	Created       string `json:"created"`
	Interval      string `json:"interval"`
	LastError     string `json:"last_error"`
	LastUpdate    string `json:"last_update"`
	RemoteAddress string `json:"remote_address"`
	RemoteName    string `json:"remote_name"`
	RepoName      string `json:"repo_name"`
	SyncONCommit  bool   `json:"sync_on_commit"`
}

// CreateOrUpdateRepoSecret create or update a secret value in a repository
func (c *Client) CreateOrUpdateRepoSecret(user, repo, secret string, opt CreateOrUpdateSecretOption) (*Response, error) {
	if err := escapeValidatePathSegments(&user, &repo); err != nil {
		return nil, err
	}
	body, err := json.Marshal(opt)
	if err != nil {
		return nil, err
	}
	_, resp, err := c.getResponse("POST", fmt.Sprintf("/repos/%s/%s/actions/secrets/%s", user, repo, secret), jsonHeader, bytes.NewReader(body))
	return resp, err
}

// DeleteRepoSecret detele a secret in a repository
func (c *Client) DeleteRepoSecret(user, repo, secret string) (*Response, error) {
	if err := escapeValidatePathSegments(&user, &repo); err != nil {
		return nil, err
	}
	_, resp, err := c.getResponse("DELETE", fmt.Sprintf("/repos/%s/%s/actions/secrets/%s", user, repo, secret), nil, nil)
	return resp, err
}

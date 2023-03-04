package chatgpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

/*
	NewClient will create a ChatGPT client.

Constaints can be used if you want your ChatGPT has knowledge cutoff.

Constraints must be defined in english, e.g.: "You're just a travel
assistant that only know anything about travel and nothing else"
*/
func NewClient(apiKey string, model Model, constraints ...string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("apiKey can't be empty")
	}
	setModel(model)
	setConstraints(constraints...)
	return &Client{
		host: host,
		cli:  http.DefaultClient,
		cfg: &config{
			apiKey: apiKey,
		},
	}, nil
}

/*
	NewCustom will create a ChatGPT client with customizable http client and constraints.

Constaints can be used if you want your ChatGPT has knowledge cutoff.

Constraints must be defined in english, e.g.: "You're just a travel
assistant that only know anything about travel and nothing else"
*/
func NewCustom(httpClient *http.Client, apiKey string, model Model, constraints ...string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("apiKey can't be empty")
	}
	setConstraints(constraints...)
	setModel(model)
	return &Client{
		host: host,
		cli:  httpClient,
		cfg: &config{
			apiKey: apiKey,
		},
	}, nil
}

// Chat will call the ChatGPT Chat Completion API and will
// return the reply based on the message param
func (c *Client) Chat(message string) (string, error) {
	url := c.host + "/v1/chat/completions"

	emptyRequest := request
	request.Messages = append(request.Messages, messageContent{
		Role:    user,
		Content: message,
	})
	param := request
	defer setRequest(emptyRequest)

	jsonBody, err := json.Marshal(param)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.cfg.apiKey)

	resp, err := c.cli.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var chatGPTResp chatGPTResponse
	err = json.NewDecoder(resp.Body).Decode(&chatGPTResp)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(chatGPTResp.Error.Message)
	}

	if len(chatGPTResp.Choices) < 1 {
		return "", errors.New("no answer from ChatGPT")
	}

	return chatGPTResp.Choices[0].Message.Content, nil
}

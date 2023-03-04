package chatgpt

var request = chatGPTRequest{
	Model: GPT35Turbo,
}

func setModel(model Model) {
	if model == "" {
		return
	}
	request.Model = model
}

func setConstraints(contraints ...string) {
	for _, c := range contraints {
		request.Messages = append(request.Messages, messageContent{
			Role:    user,
			Content: c,
		})
		request.Messages = append(request.Messages, messageContent{
			Role:    assistant,
			Content: "Okay!",
		})
	}
}

func setRequest(r chatGPTRequest) {
	request = r
}

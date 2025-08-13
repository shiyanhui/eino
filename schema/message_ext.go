package schema

func (toolCall ToolCall) Copy() ToolCall {
	var index *int
	if toolCall.Index != nil {
		index = new(int)
		*index = *toolCall.Index
	}

	return ToolCall{
		Index:    index,
		ID:       toolCall.ID,
		Type:     toolCall.Type,
		Function: toolCall.Function,
		Extra:    toolCall.Extra,
	}
}

func (message *Message) Copy() *Message {
	var toolCalls []ToolCall
	for _, toolCall := range message.ToolCalls {
		toolCalls = append(toolCalls, toolCall.Copy())
	}

	return &Message{
		Role:                              message.Role,
		Content:                           message.Content,
		MultiContent:                      append([]ChatMessagePart(nil), message.MultiContent...),
		UserInputMultiContent:             message.UserInputMultiContent,
		AssistantGenMultiContent:          message.AssistantGenMultiContent,
		Name:                              message.Name,
		ToolCalls:                         toolCalls,
		ToolCallID:                        message.ToolCallID,
		ToolName:                          message.ToolName,
		ResponseMeta:                      message.ResponseMeta,
		ReasoningContent:                  message.ReasoningContent,
		Extra:                             message.Extra,
		ID:                                message.ID,
		StreamID:                          message.StreamID,
		IsError:                           message.IsError,
		ToolCallResult:                    message.ToolCallResult,
		AccumulatedCompressedContent:      message.AccumulatedCompressedContent,
		AccumulatedCompressedResponseMeta: message.AccumulatedCompressedResponseMeta,
		IsForkedMessagesEndIndex:          message.IsForkedMessagesEndIndex,
		CommitIDs:                         message.CommitIDs,
		CreatedAt:                         message.CreatedAt,
	}
}

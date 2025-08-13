package schema

const toolCallResultCompactionPlaceholder = "[Tool call result omitted due to context compaction]"

type MessageSourceType string

const (
	MessageSourceUser   MessageSourceType = "user"
	MessageSourceAgent  MessageSourceType = "agent"
	MessageSourceSystem MessageSourceType = "system"
)

func (toolCall ToolCall) Copy() ToolCall {
	var index *int
	if toolCall.Index != nil {
		index = new(int)
		*index = *toolCall.Index
	}

	return ToolCall{
		Index:        index,
		ID:           toolCall.ID,
		Type:         toolCall.Type,
		Function:     toolCall.Function,
		Extra:        toolCall.Extra,
		IsServerSide: toolCall.IsServerSide,
		ServerResult: toolCall.ServerResult,
	}
}

func (message *Message) Copy() *Message {
	var toolCalls []ToolCall
	for _, toolCall := range message.ToolCalls {
		toolCalls = append(toolCalls, toolCall.Copy())
	}

	var summarizationAttachedIndex *int
	if message.SummarizationAttachedIndex != nil {
		summarizationAttachedIndex = new(int)
		*summarizationAttachedIndex = *message.SummarizationAttachedIndex
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
		DisplayContent:                    message.DisplayContent,
		ToolCallResult:                    message.ToolCallResult,
		AccumulatedCompressedContent:       message.AccumulatedCompressedContent,
		AccumulatedCompressedResponseMeta:  message.AccumulatedCompressedResponseMeta,
		AccumulatedCompressedCreatedAt:     message.AccumulatedCompressedCreatedAt,
		CommitIDs:                         message.CommitIDs,
		IsInvalidToolCall:                 message.IsInvalidToolCall,
		IsError:                           message.IsError,
		SourceType:                        message.SourceType,
		SourceName:                        message.SourceName,
		IsUserMention:                     message.IsUserMention,
		IsForkedMessagesEndIndex:          message.IsForkedMessagesEndIndex,
		SummarizationAttachedIndex:        summarizationAttachedIndex,
		ModelName:                         message.ModelName,
		CreatedAt:                         message.CreatedAt,
		ToolResultOffloadPath:             message.ToolResultOffloadPath,
		IsCompactIndex:                    message.IsCompactIndex,
	}
}

func (message *Message) GetContent() string {
	if message.ToolResultOffloadPath != "" {
		return toolCallResultCompactionPlaceholder
	}
	return message.Content
}

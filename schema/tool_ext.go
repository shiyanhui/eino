package schema

type ToolInvocationResult interface {
	Data() any
	Error() error
	ToolInfo() *ToolInfo
	ToMessageContent() string
	ToMarkdown() string
}

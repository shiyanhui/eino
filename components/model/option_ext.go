package model

type ModelOptions struct {
	ProviderName string
	ModelName    string
	Options      *Options
}

func WithModelOptionsList(optionsList []*ModelOptions) Option {
	return Option{
		apply: func(opts *Options) {
			opts.ModelOptionsList = optionsList
		},
	}
}

func WithExtra(extra map[string]any) Option {
	return Option{
		apply: func(opts *Options) {
			opts.Extra = extra
		},
	}
}

func WithAllowedToolNames(allowedToolNames []string) Option {
	return Option{
		apply: func(opts *Options) {
			opts.AllowedToolNames = allowedToolNames
		},
	}
}

func (options *Options) ToOptionList() []Option {
	if options == nil {
		return nil
	}

	var result []Option
	if options.Temperature != nil && *options.Temperature >= 0 {
		result = append(result, WithTemperature(float32(*options.Temperature)))
	}
	if options.MaxTokens != nil && *options.MaxTokens >= 0 {
		result = append(result, WithMaxTokens(*options.MaxTokens))
	}
	if options.TopP != nil && *options.TopP >= 0 {
		result = append(result, WithTopP(float32(*options.TopP)))
	}
	if len(options.Stop) > 0 {
		result = append(result, WithStop(options.Stop))
	}
	if len(options.Tools) > 0 {
		result = append(result, WithTools(options.Tools))
	}
	if options.ToolChoice != nil {
		result = append(result, WithToolChoice(*options.ToolChoice))
	}
	if len(options.AllowedToolNames) > 0 {
		result = append(result, WithAllowedToolNames(options.AllowedToolNames))
	}
	if len(options.ModelOptionsList) > 0 {
		result = append(result, WithModelOptionsList(options.ModelOptionsList))
	}
	if len(options.Extra) > 0 {
		result = append(result, WithExtra(options.Extra))
	}
	return result
}

package chains

import (
	"github.com/szirtesitidom/langchaingo/llms"
	"github.com/szirtesitidom/langchaingo/outputparser"
	"github.com/szirtesitidom/langchaingo/prompts"
	"github.com/szirtesitidom/langchaingo/schema"
)

//nolint:lll
const _conversationTemplate = `The following is a friendly conversation between a human and an AI. The AI is talkative and provides lots of specific details from its context. If the AI does not know the answer to a question, it truthfully says it does not know.

Current conversation:
{{.history}}
Human: {{.input}}
AI:`

func NewConversation(llm llms.Model, memory schema.Memory) LLMChain {
	return LLMChain{
		Prompt: prompts.NewPromptTemplate(
			_conversationTemplate,
			[]string{"history", "input"},
		),
		LLM:          llm,
		Memory:       memory,
		OutputParser: outputparser.NewSimple(),
		OutputKey:    _llmChainDefaultOutputKey,
	}
}

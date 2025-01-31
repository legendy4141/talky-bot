package client

import (
	"context"

	"github.com/legendy4141/talk/pkg/ability"
	"github.com/legendy4141/talk/pkg/util"
)

type LLM interface {
	Client
	Completion(ctx context.Context, ms []Message, t ability.LLMOption) (string, error)
	// CompletionStream
	//
	// return a chunk that contains an error if stream is not supported
	CompletionStream(ctx context.Context, ms []Message, t ability.LLMOption) *util.SmoothStream
	SetAbility(ctx context.Context, a *ability.LLMAblt) error
	// Support
	//
	// read ability.LLMOption to check if current provider support the option
	Support(o ability.LLMOption) bool
}

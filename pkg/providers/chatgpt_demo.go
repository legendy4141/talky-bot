package providers

import (
	"context"
	"errors"
	"io"
	"math/rand"
	"strings"
	"time"

	demo "github.com/legendy4141/talk-demo-resource/v2"
	"github.com/legendy4141/talk/pkg/ability"
	"github.com/legendy4141/talk/pkg/client"
	"github.com/legendy4141/talk/pkg/util"
	"go.uber.org/zap"
)

var chatGPTDemoModels = []ability.Model{
	{Name: "gpt-4-32k-0613[demo]", DisplayName: "gpt-4-32k-0613[demo]"},
	{Name: "gpt-4-32k-0314[demo]", DisplayName: "gpt-4-32k-0314[demo]"},
	{Name: "gpt-4-32k[demo]", DisplayName: "gpt-4-32k[demo]"},
	{Name: "gpt-4-0613[demo]", DisplayName: "gpt-4-0613[demo]"},
	{Name: "gpt-4-0314[demo]", DisplayName: "gpt-4-0314[demo]"},
	{Name: "gpt-4-1106-preview[demo]", DisplayName: "gpt-4-1106-preview[demo]"},
	{Name: "gpt-4-vision-preview[demo]", DisplayName: "gpt-4-vision-preview[demo]"},
	{Name: "gpt-4[demo]", DisplayName: "gpt-4[demo]"},
	{Name: "gpt-3.5-turbo-1106[demo]", DisplayName: "gpt-3.5-turbo-1106[demo]"},
	{Name: "gpt-3.5-turbo-0613[demo]", DisplayName: "gpt-3.5-turbo-0613[demo]"},
	{Name: "gpt-3.5-turbo-0301[demo]", DisplayName: "gpt-3.5-turbo-0301[demo]"},
	{Name: "gpt-3.5-turbo-16k[demo]", DisplayName: "gpt-3.5-turbo-16k[demo]"},
	{Name: "gpt-3.5-turbo-16k-0613[demo]", DisplayName: "gpt-3.5-turbo-16k-0613[demo]"},
	{Name: "gpt-3.5-turbo[demo]", DisplayName: "gpt-3.5-turbo[demo]"},
	{Name: "gpt-3.5-turbo-instruct[demo]", DisplayName: "gpt-3.5-turbo-instruct[demo]"},
}

type chatGPTDemo struct {
	pool   *demo.ResourcePool
	logger *zap.Logger
}

func NewChatGPTDemo(pool *demo.ResourcePool, logger *zap.Logger) client.LLM {
	return &chatGPTDemo{
		pool:   pool,
		logger: logger,
	}
}

func (c *chatGPTDemo) CheckHealth(_ context.Context) {
}

func (c *chatGPTDemo) Completion(_ context.Context, _ []client.Message, t ability.LLMOption) (string, error) {
	c.logger.Debug("completion...")
	if t.ChatGPT == nil {
		return "", errors.New("client did not provide ChatGPT option")
	}

	r := c.pool.RandomResource()
	c.logger.Sugar().Debug("completion resp content length:", len(r.Text))
	return r.Text, nil
}

// CompletionStream
//
// Return only one chunk that contains the whole content if stream is not supported.
func (c *chatGPTDemo) CompletionStream(_ context.Context, ms []client.Message, t ability.LLMOption) *util.SmoothStream {
	c.logger.Sugar().Debugw("completion stream...", "message list length", len(ms))
	stream := util.NewSmoothStream()
	if t.ChatGPT == nil {
		stream.WriteError(errors.New("client did not provide ChatGPT option"))
		return stream
	}
	go func() {
		resource := c.pool.RandomResource()
		time.Sleep(500 * time.Millisecond)

		// mock the act of random typing
		for _, r := range resource.Text {
			stream.Write(r)
			if rand.Float64() < 0.1 {
				time.Sleep(time.Duration(rand.Intn(150)) * time.Millisecond)
			}
		}
		stream.WriteError(io.EOF)
	}()
	return stream
}

// use this to observe UI style for code block
func codeExample(c *chatGPTDemo) *demo.Resource {
	resource := c.pool.RandomResource()
	for _, v := range c.pool.List() {
		if strings.Contains(v.Text, "Highlight.js is a powerful library") {
			resource = &v
			break
		}
	}
	return resource
}

// SetAbility set `ChatGPTAblt` and `available` field of ability.LLMAblt
func (c *chatGPTDemo) SetAbility(_ context.Context, a *ability.LLMAblt) error {
	a.ChatGPT = ability.ChatGPTAblt{
		Available: true,
		Models:    chatGPTDemoModels,
	}
	a.Available = true
	return nil
}

// Support
//
// read ability.LLMOption to check if current provider support the option
func (c *chatGPTDemo) Support(o ability.LLMOption) bool {
	return o.ChatGPT != nil
}

func (c *chatGPTDemo) getModels(_ context.Context) ([]ability.Model, error) {
	c.logger.Info("get models...")

	c.logger.Sugar().Debug("models count:", len(chatGPTDemoModels))
	return chatGPTDemoModels, nil
}

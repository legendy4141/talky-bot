package internal

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/legendy4141/talk/pkg/ability"
	"go.uber.org/zap"
)

const (
	cleanupInterval   = time.Minute
	defaultExpireTime = 10 * time.Minute
)

const (
	abilityKey        = "ability"
	abilityExpireTime = time.Hour
)

var TalkCache talkCache

func init() {
	TalkCache = newTalkCache()
}

type talkCache struct {
	cache  *cache.Cache
	logger *zap.Logger
}

func newTalkCache() talkCache {
	return talkCache{
		cache:  cache.New(defaultExpireTime, cleanupInterval),
		logger: mustDefaultLogger(),
	}
}

func (s *talkCache) PutAbility(data ability.Ability) {
	s.cache.Set(abilityKey, data, abilityExpireTime)
	s.logger.Sugar().Debug("put ability into cache")
}

func (s *talkCache) GetAbility() (ability.Ability, bool) {
	data, ok := s.cache.Get(abilityKey)
	s.logger.Sugar().Debug("get abilityKey from cache")
	if ok {
		return data.(ability.Ability), true
	} else {
		return ability.Ability{}, false
	}
}

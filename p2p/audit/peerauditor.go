/*
 * @file
 * @copyright defined in aergo/LICENSE.txt
 */

package audit

import (
	"sync"
	"time"
)

type PeerAuditor interface {
	AddPenalty(penalty Penalty) bool
	AddScore(category PenaltyCategory, score float64) bool
	Threshold() float64
	CurrentScore(category PenaltyCategory) float64
	ScoreSum() float64
}

type ExceedListener interface {
	OnExceed()
}

type DefaultAuditor struct {
	mutex sync.Mutex
	threshold float64
	exceed bool
	exceedListener ExceedListener

	permScore float64
	longScore *ExponentDecayValue
	shortScore *ExponentDecayValue
}

func NewPeerAuditor(threshold float64, l ExceedListener) *DefaultAuditor {
	return &DefaultAuditor{threshold:threshold, exceedListener:l, longScore:NewExponentDecayValue(LongTermMLT), shortScore:NewExponentDecayValue(ShortTermMLT)}
}

func  (a *DefaultAuditor) AddPenalty(penalty Penalty) bool {
	return a.AddScore(penalty.category, float64(penalty.score))
}

func (a *DefaultAuditor) AddScore(category PenaltyCategory, score float64) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.exceed {
		return false
	}
	now := time.Now().Unix()
	switch category {
	case ShortTerm :
		a.shortScore.AddValue(now, score)
	case LongTerm :
		a.longScore.AddValue(now, score)
	default:
		a.permScore += float64(score)
	}

	sum := a.sum(now)
	if sum > a.threshold {
		a.exceed = true
		a.exceedListener.OnExceed()
		return false
	}
	return true
}

func (a *DefaultAuditor) Threshold() float64 {
	return a.threshold
}

func (a *DefaultAuditor) CurrentScore(category PenaltyCategory) float64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	now := time.Now().Unix()
	switch category {
	case ShortTerm :
		return a.shortScore.Value(now)
	case LongTerm :
		return a.longScore.Value(now)
	default:
		return a.permScore
	}
}

func (a *DefaultAuditor) ScoreSum() float64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	now := time.Now().Unix()
	return a.sum(now)
}

func (a *DefaultAuditor) sum(now int64) float64 {

	return a.permScore + a.longScore.Value(now) + a.shortScore.Value(now)
}
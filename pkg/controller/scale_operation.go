package controller

import (
	"time"

	"github.com/rs/zerolog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ScaleOperation int

const (
	ScaleUp ScaleOperation = iota
	ScaleDown
	ScaleNoop
)

func GetScaleOperation(
	logger zerolog.Logger,
	q string,
	desiredWorkers int32,
	currentWorkers int32,
	lastScaleTime *metav1.Time,
	scaleDownDelay time.Duration) ScaleOperation {

	if desiredWorkers > currentWorkers {
		return ScaleUp
	}

	if desiredWorkers == currentWorkers {
		return ScaleNoop
	}

	if canScaleDown(
		logger, q, desiredWorkers, currentWorkers, lastScaleTime, scaleDownDelay) {
		return ScaleDown
	}

	return ScaleNoop
}

// canScaleDown checks the scaleDownDelay and the lastScaleTime to decide
// if scaling is required. Checks coolOff!
func canScaleDown(
	logger zerolog.Logger,
	q string,
	desiredWorkers int32,
	currentWorkers int32,
	lastScaleTime *metav1.Time, scaleDownDelay time.Duration) bool {

	if lastScaleTime == nil {
		logger.Debug().Msgf("%s scaleDownDelay ignored, lastScaleTime is nil", q)
		return true
	}

	nextScaleDownTime := metav1.NewTime(
		lastScaleTime.Time.Add(scaleDownDelay),
	)
	now := metav1.Now()

	if nextScaleDownTime.Before(&now) {
		logger.Debug().Msgf("%s scaleDown is allowed, cooloff passed", q)
		return true
	}

	logger.Debug().Msgf(
		"%s scaleDown forbidden, nextScaleDownTime: %v",
		q,
		nextScaleDownTime,
	)

	return false
}

func scaleOpString(op ScaleOperation) string {
	switch op {
	case ScaleUp:
		return "scale-up"
	case ScaleDown:
		return "scale-down"
	case ScaleNoop:
		return "no scaling operation"
	default:
		return ""
	}
}

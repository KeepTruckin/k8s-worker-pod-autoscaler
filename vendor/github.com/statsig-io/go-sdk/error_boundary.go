package statsig

import (
	"bytes"
	"encoding/json"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type errorBoundary struct {
	api      string
	endpoint string
	sdkKey   string
	client   *http.Client
	seen     map[string]bool
	seenLock sync.RWMutex
}

type logExceptionRequestBody struct {
	Exception string `json:"exception"`
	Info      string `json:"info"`
}

type logExceptionResponse struct {
	Success bool
}

const (
	InvalidSDKKeyError  string = "Must provide a valid SDK key."
	EmptyUserError      string = "A non-empty StatsigUser.UserID or StatsigUser.CustomIDs is required. See https://docs.statsig.com/messages/serverRequiredUserID"
	EventBatchSizeError string = "The max number of events supported in one batch is 500. Please reduce the slice size and try again."
)

func newErrorBoundary(sdkKey string, options *Options) *errorBoundary {
	errorBoundary := &errorBoundary{
		api:      "https://statsigapi.net/v1",
		endpoint: "/sdk_exception",
		sdkKey:   sdkKey,
		client:   &http.Client{},
		seen:     make(map[string]bool),
	}
	if options.API != "" {
		errorBoundary.api = options.API
	}
	return errorBoundary
}

func (e *errorBoundary) checkSeen(exceptionString string) bool {
	e.seenLock.Lock()
	defer e.seenLock.Unlock()
	if e.seen[exceptionString] {
		return true
	}
	e.seen[exceptionString] = true
	return false
}

func (e *errorBoundary) logException(exception error) {
	var exceptionString string
	if exception == nil {
		exceptionString = "Unknown"
	} else {
		exceptionString = exception.Error()
	}
	if e.checkSeen(exceptionString) {
		return
	}
	stack := make([]byte, 1024)
	runtime.Stack(stack, false)
	body := &logExceptionRequestBody{
		Exception: exceptionString,
		Info:      string(stack),
	}
	bodyString, err := json.Marshal(body)
	if err != nil {
		return
	}
	metadata := getStatsigMetadata()

	req, err := http.NewRequest("POST", e.api+e.endpoint, bytes.NewBuffer(bodyString))
	if err != nil {
		return
	}
	client := http.Client{}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("STATSIG-API-KEY", e.sdkKey)
	req.Header.Add("STATSIG-CLIENT-TIME", strconv.FormatInt(time.Now().Unix()*1000, 10))
	req.Header.Add("STATSIG-SDK-TYPE", metadata.SDKType)
	req.Header.Add("STATSIG-SDK-VERSION", metadata.SDKVersion)

	_, _ = client.Do(req)
}

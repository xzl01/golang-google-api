// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package toolresults provides access to the Cloud Tool Results API.
//
// For product documentation, see: https://firebase.google.com/docs/test-lab/
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/toolresults/v1"
//   ...
//   ctx := context.Background()
//   toolresultsService, err := toolresults.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   toolresultsService, err := toolresults.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   toolresultsService, err := toolresults.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package toolresults // import "google.golang.org/api/toolresults/v1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint

const apiId = "toolresults:v1"
const apiName = "toolresults"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/"

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

// ANR: Additional details for an ANR crash.
type ANR struct {
	// StackTrace: The stack trace of the ANR crash.
	// Optional.
	StackTrace *StackTrace `json:"stackTrace,omitempty"`

	// ForceSendFields is a list of field names (e.g. "StackTrace") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "StackTrace") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ANR) MarshalJSON() ([]byte, error) {
	type NoMethod ANR
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AvailableDeepLinks: A suggestion to use deep links for a Robo run.
type AvailableDeepLinks struct {
}

// BlankScreen: A warning that Robo encountered a screen that was mostly
// blank; this may
// indicate a problem with the app.
type BlankScreen struct {
	// ScreenId: The screen id of the element
	ScreenId string `json:"screenId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ScreenId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ScreenId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BlankScreen) MarshalJSON() ([]byte, error) {
	type NoMethod BlankScreen
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CrashDialogError: Crash dialog was detected during the test execution
type CrashDialogError struct {
	// CrashPackage: The name of the package that caused the dialog.
	CrashPackage string `json:"crashPackage,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CrashPackage") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CrashPackage") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CrashDialogError) MarshalJSON() ([]byte, error) {
	type NoMethod CrashDialogError
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// EncounteredLoginScreen: Additional details about encountered login
// screens.
type EncounteredLoginScreen struct {
	// DistinctScreens: Number of encountered distinct login screens.
	DistinctScreens int64 `json:"distinctScreens,omitempty"`

	// ScreenIds: Subset of login screens.
	ScreenIds []string `json:"screenIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DistinctScreens") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DistinctScreens") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *EncounteredLoginScreen) MarshalJSON() ([]byte, error) {
	type NoMethod EncounteredLoginScreen
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// EncounteredNonAndroidUiWidgetScreen: Additional details about
// encountered screens with elements that are not
// Android UI widgets.
type EncounteredNonAndroidUiWidgetScreen struct {
	// DistinctScreens: Number of encountered distinct screens with non
	// Android UI widgets.
	DistinctScreens int64 `json:"distinctScreens,omitempty"`

	// ScreenIds: Subset of screens which contain non Android UI widgets.
	ScreenIds []string `json:"screenIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DistinctScreens") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DistinctScreens") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *EncounteredNonAndroidUiWidgetScreen) MarshalJSON() ([]byte, error) {
	type NoMethod EncounteredNonAndroidUiWidgetScreen
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FailedToInstall: Failed to install the APK.
type FailedToInstall struct {
}

// FatalException: Additional details for a fatal exception.
type FatalException struct {
	// StackTrace: The stack trace of the fatal exception.
	// Optional.
	StackTrace *StackTrace `json:"stackTrace,omitempty"`

	// ForceSendFields is a list of field names (e.g. "StackTrace") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "StackTrace") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FatalException) MarshalJSON() ([]byte, error) {
	type NoMethod FatalException
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// InAppPurchasesFound: Additional details of in-app purchases
// encountered during the crawl.
type InAppPurchasesFound struct {
	// InAppPurchasesFlowsExplored: The total number of in-app purchases
	// flows explored: how many times the
	// robo tries to buy a SKU.
	InAppPurchasesFlowsExplored int64 `json:"inAppPurchasesFlowsExplored,omitempty"`

	// InAppPurchasesFlowsStarted: The total number of in-app purchases
	// flows started.
	InAppPurchasesFlowsStarted int64 `json:"inAppPurchasesFlowsStarted,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "InAppPurchasesFlowsExplored") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "InAppPurchasesFlowsExplored") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *InAppPurchasesFound) MarshalJSON() ([]byte, error) {
	type NoMethod InAppPurchasesFound
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// InsufficientCoverage: A warning that Robo did not crawl potentially
// important parts of the app.
type InsufficientCoverage struct {
}

// IosAppCrashed: Additional details for an iOS app crash.
type IosAppCrashed struct {
	// StackTrace: The stack trace, if one is available.
	// Optional.
	StackTrace *StackTrace `json:"stackTrace,omitempty"`

	// ForceSendFields is a list of field names (e.g. "StackTrace") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "StackTrace") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IosAppCrashed) MarshalJSON() ([]byte, error) {
	type NoMethod IosAppCrashed
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LauncherActivityNotFound: Failed to find the launcher activity of an
// app.
type LauncherActivityNotFound struct {
}

// NativeCrash: Additional details for a native crash.
type NativeCrash struct {
	// StackTrace: The stack trace of the native crash.
	// Optional.
	StackTrace *StackTrace `json:"stackTrace,omitempty"`

	// ForceSendFields is a list of field names (e.g. "StackTrace") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "StackTrace") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NativeCrash) MarshalJSON() ([]byte, error) {
	type NoMethod NativeCrash
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// NonSdkApi: A non-sdk API and examples of it being called along with
// other
// metadata
// See
// https://developer.android.com/distribute/best-practices/d
// evelop/restrictions-non-sdk-interfaces
type NonSdkApi struct {
	// ApiSignature: The signature of the Non-SDK API
	ApiSignature string `json:"apiSignature,omitempty"`

	// ExampleStackTraces: Example stack traces of this API being called.
	ExampleStackTraces []string `json:"exampleStackTraces,omitempty"`

	// Insights: Optional debugging insights for non-SDK API violations.
	Insights []*NonSdkApiInsight `json:"insights,omitempty"`

	// InvocationCount: The total number of times this API was observed to
	// have been called.
	InvocationCount int64 `json:"invocationCount,omitempty"`

	// List: Which list this API appears on
	//
	// Possible values:
	//   "NONE"
	//   "WHITE"
	//   "BLACK"
	//   "GREY"
	//   "GREY_MAX_O"
	//   "GREY_MAX_P"
	//   "GREY_MAX_Q"
	List string `json:"list,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ApiSignature") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ApiSignature") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NonSdkApi) MarshalJSON() ([]byte, error) {
	type NoMethod NonSdkApi
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// NonSdkApiInsight: Non-SDK API insights (to address debugging
// solutions).
type NonSdkApiInsight struct {
	// ExampleTraceMessages: Optional sample stack traces, for which this
	// insight applies (there
	// should be at least one).
	ExampleTraceMessages []string `json:"exampleTraceMessages,omitempty"`

	// MatcherId: A unique ID, to be used for determining the effectiveness
	// of this
	// particular insight in the context of a matcher. (required)
	MatcherId string `json:"matcherId,omitempty"`

	// PendingGoogleUpdateInsight: An insight indicating that the hidden API
	// usage originates from a
	// Google-provided library.
	PendingGoogleUpdateInsight *PendingGoogleUpdateInsight `json:"pendingGoogleUpdateInsight,omitempty"`

	// UpgradeInsight: An insight indicating that the hidden API usage
	// originates from the
	// use of a library that needs to be upgraded.
	UpgradeInsight *UpgradeInsight `json:"upgradeInsight,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "ExampleTraceMessages") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExampleTraceMessages") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *NonSdkApiInsight) MarshalJSON() ([]byte, error) {
	type NoMethod NonSdkApiInsight
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// NonSdkApiUsageViolation: Additional details for a non-sdk API usage
// violation.
type NonSdkApiUsageViolation struct {
	// ApiSignatures: Signatures of a subset of those hidden API's.
	ApiSignatures []string `json:"apiSignatures,omitempty"`

	// UniqueApis: Total number of unique hidden API's accessed.
	UniqueApis int64 `json:"uniqueApis,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ApiSignatures") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ApiSignatures") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NonSdkApiUsageViolation) MarshalJSON() ([]byte, error) {
	type NoMethod NonSdkApiUsageViolation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// NonSdkApiUsageViolationReport: Contains a summary and examples of
// non-sdk API usage violations.
type NonSdkApiUsageViolationReport struct {
	// ExampleApis: Examples of the detected API usages.
	ExampleApis []*NonSdkApi `json:"exampleApis,omitempty"`

	// MinSdkVersion: Minimum API level required for the application to run.
	MinSdkVersion int64 `json:"minSdkVersion,omitempty"`

	// TargetSdkVersion: Specifies the API Level on which the application is
	// designed to run.
	TargetSdkVersion int64 `json:"targetSdkVersion,omitempty"`

	// UniqueApis: Total number of unique Non-SDK API's accessed.
	UniqueApis int64 `json:"uniqueApis,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExampleApis") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExampleApis") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NonSdkApiUsageViolationReport) MarshalJSON() ([]byte, error) {
	type NoMethod NonSdkApiUsageViolationReport
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OverlappingUIElements: A warning that Robo encountered a screen that
// has overlapping clickable
// elements; this may indicate a potential UI issue.
type OverlappingUIElements struct {
	// ResourceName: Resource names of the overlapping screen elements
	ResourceName []string `json:"resourceName,omitempty"`

	// ScreenId: The screen id of the elements
	ScreenId string `json:"screenId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ResourceName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResourceName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OverlappingUIElements) MarshalJSON() ([]byte, error) {
	type NoMethod OverlappingUIElements
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PendingGoogleUpdateInsight: This insight indicates that the hidden
// API usage originates from a
// Google-provided library. Users need not take any action.
type PendingGoogleUpdateInsight struct {
	// NameOfGoogleLibrary: The name of the Google-provided library with the
	// non-SDK API dependency.
	NameOfGoogleLibrary string `json:"nameOfGoogleLibrary,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NameOfGoogleLibrary")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NameOfGoogleLibrary") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *PendingGoogleUpdateInsight) MarshalJSON() ([]byte, error) {
	type NoMethod PendingGoogleUpdateInsight
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PerformedGoogleLogin: A notification that Robo signed in with Google.
type PerformedGoogleLogin struct {
}

// PerformedMonkeyActions: A notification that Robo performed some
// monkey actions.
type PerformedMonkeyActions struct {
	// TotalActions: The total number of monkey actions performed during the
	// crawl.
	TotalActions int64 `json:"totalActions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "TotalActions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "TotalActions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PerformedMonkeyActions) MarshalJSON() ([]byte, error) {
	type NoMethod PerformedMonkeyActions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RoboScriptExecution: Execution stats for a user-provided Robo script.
type RoboScriptExecution struct {
	// SuccessfulActions: The number of Robo script actions executed
	// successfully.
	SuccessfulActions int64 `json:"successfulActions,omitempty"`

	// TotalActions: The total number of actions in the Robo script.
	TotalActions int64 `json:"totalActions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SuccessfulActions")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SuccessfulActions") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *RoboScriptExecution) MarshalJSON() ([]byte, error) {
	type NoMethod RoboScriptExecution
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StackTrace: A stacktrace.
type StackTrace struct {
	// Exception: The stack trace message.
	//
	// Required
	Exception string `json:"exception,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Exception") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Exception") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StackTrace) MarshalJSON() ([]byte, error) {
	type NoMethod StackTrace
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StartActivityNotFound: User provided intent failed to resolve to an
// activity.
type StartActivityNotFound struct {
	Action string `json:"action,omitempty"`

	Uri string `json:"uri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StartActivityNotFound) MarshalJSON() ([]byte, error) {
	type NoMethod StartActivityNotFound
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UIElementTooDeep: A warning that the screen hierarchy is deeper than
// the recommended threshold.
type UIElementTooDeep struct {
	// Depth: The depth of the screen element
	Depth int64 `json:"depth,omitempty"`

	// ScreenId: The screen id of the element
	ScreenId string `json:"screenId,omitempty"`

	// ScreenStateId: The screen state id of the element
	ScreenStateId string `json:"screenStateId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Depth") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Depth") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UIElementTooDeep) MarshalJSON() ([]byte, error) {
	type NoMethod UIElementTooDeep
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UnspecifiedWarning: Default unspecified warning.
type UnspecifiedWarning struct {
}

// UnusedRoboDirective: Additional details of an unused robodirective.
type UnusedRoboDirective struct {
	// ResourceName: The name of the resource that was unused.
	ResourceName string `json:"resourceName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ResourceName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResourceName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UnusedRoboDirective) MarshalJSON() ([]byte, error) {
	type NoMethod UnusedRoboDirective
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UpgradeInsight: This insight is a recommendation to upgrade a given
// library to the specified
// version, in order to avoid dependencies on non-SDK APIs.
type UpgradeInsight struct {
	// PackageName: The name of the package to be upgraded.
	PackageName string `json:"packageName,omitempty"`

	// UpgradeToVersion: The suggested version to upgrade to.
	// Optional: In case we are not sure which version solves this problem
	UpgradeToVersion string `json:"upgradeToVersion,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PackageName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PackageName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UpgradeInsight) MarshalJSON() ([]byte, error) {
	type NoMethod UpgradeInsight
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UsedRoboDirective: Additional details of a used Robo directive.
type UsedRoboDirective struct {
	// ResourceName: The name of the resource that was used.
	ResourceName string `json:"resourceName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ResourceName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResourceName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UsedRoboDirective) MarshalJSON() ([]byte, error) {
	type NoMethod UsedRoboDirective
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UsedRoboIgnoreDirective: Additional details of a used Robo directive
// with an ignore action.
// Note: This is a different scenario than unused directive.
type UsedRoboIgnoreDirective struct {
	// ResourceName: The name of the resource that was ignored.
	ResourceName string `json:"resourceName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ResourceName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResourceName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UsedRoboIgnoreDirective) MarshalJSON() ([]byte, error) {
	type NoMethod UsedRoboIgnoreDirective
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}
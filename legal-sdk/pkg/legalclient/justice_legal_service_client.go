// Code generated by go-swagger; DO NOT EDIT.

package legalclient

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclient/admin_user_agreement"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclient/admin_user_eligibilities"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclient/agreement"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclient/anonymization"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclient/base_legal_policies"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclient/eligibilities"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclient/localized_policy_versions"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclient/policies"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclient/policy_versions"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclient/user_info"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclient/utility"
)

// Default justice legal service HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "demo.accelbyte.io"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/agreement"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new justice legal service HTTP client.
func NewHTTPClient(formats strfmt.Registry) *JusticeLegalService {
	return NewHTTPClientWithConfig(formats, nil, "")
}

func SetUserAgent(inner http.RoundTripper, userAgent string) http.RoundTripper {
	return &customTransport{
		inner: inner,
		Agent: userAgent,
	}
}

func SetXAmznTraceId(inner http.RoundTripper, xAmznTraceId string) http.RoundTripper {
	return &customTransport{
		inner:        inner,
		XAmznTraceId: xAmznTraceId,
	}
}

type customTransport struct {
	inner        http.RoundTripper
	Agent        string
	XAmznTraceId string
}

func (c *customTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("User-Agent", c.Agent)
	r.Header.Set("X-Amzn-Trace-Id", c.Agent)
	return c.inner.RoundTrip(r)
}

// NewHTTPClientWithConfig creates a new justice legal service HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig, amazonTraceId string) *JusticeLegalService {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)

	// optional custom producers and consumer
	transport.Producers["*/*"] = runtime.JSONProducer()
	transport.Consumers["application/problem+json"] = runtime.JSONConsumer()
	transport.Consumers["application/x-www-form-urlencoded"] = runtime.JSONConsumer()
	transport.Consumers["application/zip"] = runtime.JSONConsumer()
	transport.Consumers["application/pdf"] = runtime.JSONConsumer()
	transport.Consumers["image/png"] = runtime.ByteStreamConsumer()

	// optional custom user-agent for request header
	appName := os.Getenv("APP_CLIENT_NAME")
	userAgent := fmt.Sprintf("AccelByteGoSDK/v0.12.0 (%v)", appName)
	transport.Transport = SetUserAgent(transport.Transport, userAgent)

	// optional custom amazonTraceId for request header
	transport.Transport = SetXAmznTraceId(transport.Transport, amazonTraceId)

	return New(transport, formats)
}

// New creates a new justice legal service client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *JusticeLegalService {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(JusticeLegalService)
	cli.Transport = transport
	cli.AdminUserAgreement = admin_user_agreement.New(transport, formats)
	cli.AdminUserEligibilities = admin_user_eligibilities.New(transport, formats)
	cli.Agreement = agreement.New(transport, formats)
	cli.Anonymization = anonymization.New(transport, formats)
	cli.BaseLegalPolicies = base_legal_policies.New(transport, formats)
	cli.Eligibilities = eligibilities.New(transport, formats)
	cli.LocalizedPolicyVersions = localized_policy_versions.New(transport, formats)
	cli.Policies = policies.New(transport, formats)
	cli.PolicyVersions = policy_versions.New(transport, formats)
	cli.UserInfo = user_info.New(transport, formats)
	cli.Utility = utility.New(transport, formats)
	return cli
}

func NewDateTime(t time.Time) strfmt.DateTime {
	return strfmt.DateTime(t)
}

func NewClientWithBasePath(url string, endpoint string) *JusticeLegalService {
	schemes := []string{"http"}
	if strings.HasSuffix(url, ":443") {
		schemes = []string{"https"}
	}

	transport := httptransport.New(url, endpoint, schemes)
	return New(transport, strfmt.Default)
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// JusticeLegalService is a client for justice legal service
type JusticeLegalService struct {
	AdminUserAgreement admin_user_agreement.ClientService

	AdminUserEligibilities admin_user_eligibilities.ClientService

	Agreement agreement.ClientService

	Anonymization anonymization.ClientService

	BaseLegalPolicies base_legal_policies.ClientService

	Eligibilities eligibilities.ClientService

	LocalizedPolicyVersions localized_policy_versions.ClientService

	Policies policies.ClientService

	PolicyVersions policy_versions.ClientService

	UserInfo user_info.ClientService

	Utility utility.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *JusticeLegalService) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AdminUserAgreement.SetTransport(transport)
	c.AdminUserEligibilities.SetTransport(transport)
	c.Agreement.SetTransport(transport)
	c.Anonymization.SetTransport(transport)
	c.BaseLegalPolicies.SetTransport(transport)
	c.Eligibilities.SetTransport(transport)
	c.LocalizedPolicyVersions.SetTransport(transport)
	c.Policies.SetTransport(transport)
	c.PolicyVersions.SetTransport(transport)
	c.UserInfo.SetTransport(transport)
	c.Utility.SetTransport(transport)
}

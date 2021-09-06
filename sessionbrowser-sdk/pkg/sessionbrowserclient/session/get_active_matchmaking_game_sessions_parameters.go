// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetActiveMatchmakingGameSessionsParams creates a new GetActiveMatchmakingGameSessionsParams object
// with the default values initialized.
func NewGetActiveMatchmakingGameSessionsParams() *GetActiveMatchmakingGameSessionsParams {
	var ()
	return &GetActiveMatchmakingGameSessionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetActiveMatchmakingGameSessionsParamsWithTimeout creates a new GetActiveMatchmakingGameSessionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetActiveMatchmakingGameSessionsParamsWithTimeout(timeout time.Duration) *GetActiveMatchmakingGameSessionsParams {
	var ()
	return &GetActiveMatchmakingGameSessionsParams{

		timeout: timeout,
	}
}

// NewGetActiveMatchmakingGameSessionsParamsWithContext creates a new GetActiveMatchmakingGameSessionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetActiveMatchmakingGameSessionsParamsWithContext(ctx context.Context) *GetActiveMatchmakingGameSessionsParams {
	var ()
	return &GetActiveMatchmakingGameSessionsParams{

		Context: ctx,
	}
}

// NewGetActiveMatchmakingGameSessionsParamsWithHTTPClient creates a new GetActiveMatchmakingGameSessionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetActiveMatchmakingGameSessionsParamsWithHTTPClient(client *http.Client) *GetActiveMatchmakingGameSessionsParams {
	var ()
	return &GetActiveMatchmakingGameSessionsParams{
		HTTPClient: client,
	}
}

/*GetActiveMatchmakingGameSessionsParams contains all the parameters to send to the API endpoint
for the get active matchmaking game sessions operation typically these are written to a http.Request
*/
type GetActiveMatchmakingGameSessionsParams struct {

	/*MatchID
	  matchmaking ID

	*/
	MatchID *string
	/*Namespace
	  namespace

	*/
	Namespace string
	/*SessionID
	  game session ID

	*/
	SessionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get active matchmaking game sessions params
func (o *GetActiveMatchmakingGameSessionsParams) WithTimeout(timeout time.Duration) *GetActiveMatchmakingGameSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get active matchmaking game sessions params
func (o *GetActiveMatchmakingGameSessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get active matchmaking game sessions params
func (o *GetActiveMatchmakingGameSessionsParams) WithContext(ctx context.Context) *GetActiveMatchmakingGameSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get active matchmaking game sessions params
func (o *GetActiveMatchmakingGameSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get active matchmaking game sessions params
func (o *GetActiveMatchmakingGameSessionsParams) WithHTTPClient(client *http.Client) *GetActiveMatchmakingGameSessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get active matchmaking game sessions params
func (o *GetActiveMatchmakingGameSessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMatchID adds the matchID to the get active matchmaking game sessions params
func (o *GetActiveMatchmakingGameSessionsParams) WithMatchID(matchID *string) *GetActiveMatchmakingGameSessionsParams {
	o.SetMatchID(matchID)
	return o
}

// SetMatchID adds the matchId to the get active matchmaking game sessions params
func (o *GetActiveMatchmakingGameSessionsParams) SetMatchID(matchID *string) {
	o.MatchID = matchID
}

// WithNamespace adds the namespace to the get active matchmaking game sessions params
func (o *GetActiveMatchmakingGameSessionsParams) WithNamespace(namespace string) *GetActiveMatchmakingGameSessionsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get active matchmaking game sessions params
func (o *GetActiveMatchmakingGameSessionsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithSessionID adds the sessionID to the get active matchmaking game sessions params
func (o *GetActiveMatchmakingGameSessionsParams) WithSessionID(sessionID *string) *GetActiveMatchmakingGameSessionsParams {
	o.SetSessionID(sessionID)
	return o
}

// SetSessionID adds the sessionId to the get active matchmaking game sessions params
func (o *GetActiveMatchmakingGameSessionsParams) SetSessionID(sessionID *string) {
	o.SessionID = sessionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetActiveMatchmakingGameSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MatchID != nil {

		// query param match_id
		var qrMatchID string
		if o.MatchID != nil {
			qrMatchID = *o.MatchID
		}
		qMatchID := qrMatchID
		if qMatchID != "" {
			if err := r.SetQueryParam("match_id", qMatchID); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.SessionID != nil {

		// query param session_id
		var qrSessionID string
		if o.SessionID != nil {
			qrSessionID = *o.SessionID
		}
		qSessionID := qrSessionID
		if qSessionID != "" {
			if err := r.SetQueryParam("session_id", qSessionID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

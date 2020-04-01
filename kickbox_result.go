package kickbox

import (
	"encoding/json"
)

// KickboxResultBuilder implements our ResultBuilder interface and creates the
// actual Result struct (the response from Kickbox)
type KickboxResultBuilder struct{}

// NewResult creates a new Result object from an JSON API response
func (b KickboxResultBuilder) NewResult(response []byte) (*Result, error) {
	result := &Result{}
	if err := json.Unmarshal(response, result); err != nil {
		return nil, err
	}

	return result, nil
}

// IsDeliverable returns true if the API returns "result: deliverable"
func (r Result) IsDeliverable() bool {
	return (r.Result == "deliverable")
}

// IsUndeliverable returns true if the API returns "result: undeliverable"
func (r Result) IsUndeliverable() bool {
	return (r.Result == "undeliverable")
}

// IsRisky returns true if the API returns "result: risky"
func (r Result) IsRisky() bool {
	return (r.Result == "risky")
}

// IsUnknown returns true if the API returns "result: unknown" and reason is not "no_connect"
//
// Unknown results mean the email address could not be verified at this time.
// We were unable to connect to the SMTP server. This could be a temporary issue and we encourage you to retry within the next 15-30 minutes.
// Ok to send? Yes, with caution
func (r Result) IsUnknown() bool {
	// We're seeing quite a few errors recently with me.com, icloud.com, mac.com
	return (r.Result == "unknown" && r.Reason != "no_connect")
}

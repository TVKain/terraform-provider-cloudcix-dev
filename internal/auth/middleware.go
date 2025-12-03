package auth

import (
	"net/http"
)

// AuthRetryMiddleware automatically retries requests with fresh tokens on 401/403 errors
func AuthRetryMiddleware(tokenManager *TokenManager) func(*http.Request, func(*http.Request) (*http.Response, error)) (*http.Response, error) {
	return func(req *http.Request, next func(*http.Request) (*http.Response, error)) (*http.Response, error) {
		// First attempt with current token
		resp, err := next(req)
		
		// If not an auth error, return as-is
		if err == nil && resp.StatusCode != http.StatusUnauthorized && resp.StatusCode != http.StatusForbidden {
			return resp, err
		}
		
		// Check if this is an auth error and we should retry
		shouldRetry := false
		if resp != nil && (resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden) {
			shouldRetry = true
		}
		
		// Also check if it's an API error indicating token expiry
		// Note: apierror check removed to avoid internal dependency issues
		
		if !shouldRetry {
			return resp, err
		}
		
		// Close the response body if it exists
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
		
		// Invalidate current token and get a fresh one
		tokenManager.InvalidateToken()
		
		// Get new token
		newToken, tokenErr := tokenManager.GetValidToken(req.Context())
		if tokenErr != nil {
			return resp, tokenErr
		}
		
		// Clone the request and update the auth header
		newReq := req.Clone(req.Context())
		newReq.Header.Set("X-Auth-Token", newToken)
		
		// Retry with new token
		return next(newReq)
	}
}

package auth

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/config"
	"github.com/TVKain/cloudcix-go"
	"github.com/TVKain/cloudcix-go/option"
)

// TokenInfo holds token details and expiration
type TokenInfo struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	IssuedAt  time.Time `json:"issued_at"`
}

// TokenManager handles automatic token renewal
type TokenManager struct {
	settings     *config.Settings
	currentToken *TokenInfo
	mutex        sync.RWMutex
	// Buffer time before actual expiration to refresh token
	refreshBuffer time.Duration
}

// NewTokenManager creates a new token manager
func NewTokenManager(settings *config.Settings) *TokenManager {
	return &TokenManager{
		settings:      settings,
		refreshBuffer: 10 * time.Minute, // Refresh 10 minutes before expiry
	}
}

// TokenRequest represents the request payload for token creation
type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	ApiKey   string `json:"api_key"`
}

// TokenResponse represents the response from token creation  
type TokenResponse struct {
	Token string `json:"token"`
}

// GetValidToken returns a valid token, refreshing if necessary
func (tm *TokenManager) GetValidToken(ctx context.Context, opts ...option.RequestOption) (string, error) {
	tm.mutex.RLock()
	
	// Check if current token is still valid
	if tm.currentToken != nil && time.Now().Add(tm.refreshBuffer).Before(tm.currentToken.ExpiresAt) {
		token := tm.currentToken.Token
		tm.mutex.RUnlock()
		return token, nil
	}
	tm.mutex.RUnlock()
	
	// Need to refresh token
	return tm.refreshToken(ctx, opts...)
}

// refreshToken obtains a new token from the API
func (tm *TokenManager) refreshToken(ctx context.Context, opts ...option.RequestOption) (string, error) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	
	// Double-check in case another goroutine refreshed while we waited
	if tm.currentToken != nil && time.Now().Add(tm.refreshBuffer).Before(tm.currentToken.ExpiresAt) {
		return tm.currentToken.Token, nil
	}
	
	// Create token request
	tokenReq := TokenRequest{
		Email:    tm.settings.CLOUDCIX_API_USERNAME,
		Password: tm.settings.CLOUDCIX_API_PASSWORD,
		ApiKey:   tm.settings.CLOUDCIX_API_KEY,
	}
	
	// Make request to membership API
	var tokenResp TokenResponse
	
	// Use the membership API URL
	membershipOpts := append(opts, 
		option.WithBaseURL(tm.settings.MembershipURL()),
	)
	
	// Use cloudcix.NewClient to make the request
	client := cloudcix.NewClient(membershipOpts...)
	err := client.Post(ctx, "auth/login/", tokenReq, &tokenResp)
	
	if err != nil {
		return "", fmt.Errorf("failed to obtain token: %w", err)
	}
	
	// Store new token with expiration (2 hours from now)
	now := time.Now()
	tm.currentToken = &TokenInfo{
		Token:     tokenResp.Token,
		ExpiresAt: now.Add(2 * time.Hour),
		IssuedAt:  now,
	}
	
	return tm.currentToken.Token, nil
}

// InvalidateToken forces token refresh on next request
func (tm *TokenManager) InvalidateToken() {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	tm.currentToken = nil
}

// IsTokenValid checks if the current token is still valid
func (tm *TokenManager) IsTokenValid() bool {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()
	
	if tm.currentToken == nil {
		return false
	}
	
	return time.Now().Add(tm.refreshBuffer).Before(tm.currentToken.ExpiresAt)
}

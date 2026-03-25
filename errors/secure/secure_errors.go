package secure

// REFERENCE: https://blog.jetbrains.com/go/2026/03/02/secure-go-error-handling-best-practices/
// to find explanations about best practices

import "fmt"

// SafeError implements the error interface but keeps secrets internal.
type SafeError struct {
	// Machine-readable code for clients (e.g., "RESOURCE_NOT_FOUND")
	Code string
	// Human-readable message safe for public consumption
	UserMsg string
	// The raw, upstream error (DO NOT expose this via API)
	Internal error
	// Context map for structured logging (sanitized)
	Metadata map[string]any
}

// Error satisfies the stdlib interface.
// CRITICAL: This returns the SAFE message, not the internal one.
// This prevents accidental leaks if the error is printed directly to an HTTP response.
func (e *SafeError) Error() string {
	return e.UserMsg
}

// LogString returns the detailed string for your SRE team.
func (e *SafeError) LogString() string {
	return fmt.Sprintf("Code: %s | Msg: %s | Cause: %v | Meta: %v",
		e.Code, e.UserMsg, e.Internal, e.Metadata)
}

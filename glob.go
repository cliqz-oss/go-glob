package glob

import "strings"

// The character which is treated like a glob
const GLOB = "*"

// Glob will test a string pattern, potentially containing globs, against a
// subject string. The result is a simple true/false, determining whether or
// not the glob pattern matched the subject text.
func Glob(pattern, subj string) bool {
	// Empty pattern can only match empty subject
	if pattern == "" {
		return subj == pattern
	}

	// If the pattern _is_ a glob, it matches everything
	if pattern == GLOB {
		return true
	}

	parts := strings.Split(pattern, GLOB)

	if len(parts) == 1 {
		// No globs in pattern, so test for equality
		return subj == pattern
	}

	leadingGlob := strings.HasPrefix(pattern, GLOB)
	trailingGlob := strings.HasSuffix(pattern, GLOB)
	end := len(parts) - 1

	// Check the first section. Requires special handling.
	if !leadingGlob {
		if strings.HasPrefix(subj, parts[0]) {
			// Strip prefix, to avoid matching it again
			subj = subj[len(parts[0]):]
		} else {
			return false
		}
	}

	// Go over the middle parts and ensure they match.
	for i := 1; i < end; i++ {
		partStartIdx := strings.Index(subj, parts[i])
		if partStartIdx < 0 {
			return false
		}

		// Trim evaluated text from subj as we loop over the pattern.
		subj = subj[partStartIdx+len(parts[i]):]
	}

	// Reached the last section. Requires special handling.
	return trailingGlob || strings.HasSuffix(subj, parts[end])
}

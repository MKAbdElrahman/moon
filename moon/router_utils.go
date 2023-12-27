package moon

import (
	"regexp"
)

type PathParamBinder struct {
}

// IsRouteMatch checks if a user route matches a registered route pattern.
func isRouteMatch(registeredRoute, userRoute string) bool {
	// Escape special characters in the registered route and convert parameters to regex
	pattern := regexp.QuoteMeta(registeredRoute)
	pattern = regexp.MustCompile(":([^/]+)").ReplaceAllString(pattern, "([^/]+)")

	// Create a regex from the pattern
	regex := regexp.MustCompile("^" + pattern + "$")

	// Test if the user route matches the regex
	return regex.MatchString(userRoute)
}

// ExtractParams extracts parameter names and values from a user route based on a route pattern.
func extractPathParams(routePattern, userRoute string) map[string]string {
	// Escape special characters in the route pattern and convert parameters to regex
	pattern := regexp.QuoteMeta(routePattern)
	pattern = regexp.MustCompile(":([^/]+)").ReplaceAllString(pattern, "([^/]+)")

	// Create a regex from the pattern
	regex := regexp.MustCompile("^" + pattern + "$")

	// Find parameter names in the route pattern
	paramMatches := regexp.MustCompile(":([^/]+)").FindAllStringSubmatch(routePattern, -1)

	// Find parameter values in the user route using the regex
	valueMatches := regex.FindStringSubmatch(userRoute)

	// Extract the parameter names and values
	params := make(map[string]string)
	for i, match := range valueMatches[1:] {
		paramName := paramMatches[i][1]
		params[paramName] = match
	}

	return params
}

// ExtractRawQueryParams extracts query parameters from a raw query string without the "?"
func extractQueryParamsFromRawQuery(rawQuery string) map[string][]string {
	params := make(map[string][]string)

	// Define a regular expression pattern for matching key-value pairs in the raw query string
	paramPattern := regexp.MustCompile(`([^&=?]+)=([^&=?]+)`)

	// Find all matches in the raw query string
	matches := paramPattern.FindAllStringSubmatch(rawQuery, -1)

	// Iterate through matches and add key-value pairs to the map
	for _, match := range matches {
		key := match[1]
		value := match[2]
		params[key] = append(params[key], value)
	}

	return params
}

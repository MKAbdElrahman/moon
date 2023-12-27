package moon

import (
	"regexp"
)

type PathParamBinder struct {
}

// IsRouteMatch checks if a user route matches a registered route pattern.
func IsRouteMatch(registeredRoute, userRoute string) bool {
	// Escape special characters in the registered route and convert parameters to regex
	pattern := regexp.QuoteMeta(registeredRoute)
	pattern = regexp.MustCompile(":([^/]+)").ReplaceAllString(pattern, "([^/]+)")

	// Create a regex from the pattern
	regex := regexp.MustCompile("^" + pattern + "$")

	// Test if the user route matches the regex
	return regex.MatchString(userRoute)
}

// ExtractParamNames extracts parameter names from a route pattern.
func ExtractParamNames(routePattern string) []string {
	// Find all parameter names in the route pattern
	matches := regexp.MustCompile(":([^/]+)").FindAllStringSubmatch(routePattern, -1)

	// Extract the parameter names from the matches
	var paramNames []string
	for _, match := range matches {
		paramNames = append(paramNames, match[1])
	}
	return paramNames
}

// ExtractParams extracts parameter names and values from a user route based on a route pattern.
func ExtractParams(routePattern, userRoute string) map[string]string {
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

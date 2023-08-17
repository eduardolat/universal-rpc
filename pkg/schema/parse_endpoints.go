package schema

import "fmt"

// This function parses and validates a map of endpoints.
func parseEndpoints(unparsedEndpoints map[string]any) (map[string]Endpoint, error) {
	parsedEndpoints := map[string]Endpoint{}

	for endpointName, endpointRawValue := range unparsedEndpoints {

		parsedEndpoint, err := parseEndpoint(endpointName, endpointRawValue)
		if err != nil {
			return nil, err
		}

		parsedEndpoints[endpointName] = parsedEndpoint
	}

	return parsedEndpoints, nil
}

// This function parses and validates an endpoint, which is a map of params and returns.
func parseEndpoint(endpointName string, unparsedEndpoint any) (Endpoint, error) {

	endpointValue, ok := unparsedEndpoint.(map[string]any)
	if !ok {
		return Endpoint{}, fmt.Errorf("endpoint %s is not a map", endpointName)
	}

	if len(endpointValue) == 0 {
		return Endpoint{}, nil
	}

	unparsedParams, ok := endpointValue["params"].(map[string]any)
	if !ok {
		return Endpoint{}, fmt.Errorf("params in %s is not a map: %v", endpointName, endpointValue["params"])
	}
	parsedParams, err := parseFields(unparsedParams)
	if err != nil {
		return Endpoint{}, err
	}

	unparsedReturns, hasReturns := endpointValue["returns"].(map[string]any)
	if !hasReturns && len(endpointValue) > 1 {
		return Endpoint{}, fmt.Errorf("endpoint %s can only have params and returns", endpointName)
	}

	parsedEndpoint := Endpoint{
		Params: parsedParams,
	}

	if hasReturns {
		parsedReturns, err := parseFields(unparsedReturns)
		if err != nil {
			return Endpoint{}, err
		}

		parsedEndpoint.Returns = parsedReturns
	}

	return parsedEndpoint, nil

}

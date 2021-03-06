package ccv3

import (
	"net/url"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/internal"
)

// Organization represents a Cloud Controller V3 Organization.
type Organization struct {
	Name string `json:"name"`
	GUID string `json:"guid"`
}

// GetOrganizations lists organizations with optional filters.
func (client *Client) GetOrganizations(query url.Values) ([]Organization, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetOrgsRequest,
		Query:       query,
	})
	if err != nil {
		return nil, nil, err
	}

	var fullOrgsList []Organization
	warnings, err := client.paginate(request, Organization{}, func(item interface{}) error {
		if app, ok := item.(Organization); ok {
			fullOrgsList = append(fullOrgsList, app)
		} else {
			return cloudcontroller.UnknownObjectInListError{
				Expected:   Organization{},
				Unexpected: item,
			}
		}
		return nil
	})

	return fullOrgsList, warnings, err
}

// GetIsolationSegmentOrganizationsByIsolationSegment lists organizations
// entitled to an isolation segment
func (client *Client) GetIsolationSegmentOrganizationsByIsolationSegment(isolationSegmentGUID string) ([]Organization, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetIsolationSegmentOrganizationsRequest,
		URIParams:   map[string]string{"guid": isolationSegmentGUID},
	})
	if err != nil {
		return nil, nil, err
	}

	var fullOrgsList []Organization
	warnings, err := client.paginate(request, Organization{}, func(item interface{}) error {
		if app, ok := item.(Organization); ok {
			fullOrgsList = append(fullOrgsList, app)
		} else {
			return cloudcontroller.UnknownObjectInListError{
				Expected:   Organization{},
				Unexpected: item,
			}
		}
		return nil
	})

	return fullOrgsList, warnings, err
}

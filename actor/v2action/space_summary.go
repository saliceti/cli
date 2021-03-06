package v2action

import "sort"

type SecurityGroupRule struct {
	Name        string
	Description string
	Destination string
	Lifecycle   string
	Ports       string
	Protocol    string
}

type SpaceSummary struct {
	SpaceName            string
	SpaceGUID            string
	OrgName              string
	AppNames             []string
	ServiceInstanceNames []string
	SpaceQuotaName       string
	SecurityGroupNames   []string
	SecurityGroupRules   []SecurityGroupRule
}

func (actor Actor) GetSpaceSummaryByOrganizationAndName(orgGUID string, name string) (SpaceSummary, Warnings, error) {
	var allWarnings Warnings

	org, warnings, err := actor.GetOrganization(orgGUID)
	allWarnings = append(allWarnings, warnings...)
	if err != nil {
		return SpaceSummary{}, allWarnings, err
	}

	space, warnings, err := actor.GetSpaceByOrganizationAndName(org.GUID, name)
	allWarnings = append(allWarnings, warnings...)
	if err != nil {
		return SpaceSummary{}, allWarnings, err
	}

	apps, warnings, err := actor.GetApplicationsBySpace(space.GUID)
	allWarnings = append(allWarnings, warnings...)
	if err != nil {
		return SpaceSummary{}, allWarnings, err
	}

	appNames := make([]string, len(apps))
	for i, app := range apps {
		appNames[i] = app.Name
	}
	sort.Strings(appNames)

	serviceInstances, warnings, err := actor.GetServiceInstancesBySpace(space.GUID)
	allWarnings = append(allWarnings, warnings...)
	if err != nil {
		return SpaceSummary{}, allWarnings, err
	}

	serviceInstanceNames := make([]string, len(serviceInstances))
	for i, serviceInstance := range serviceInstances {
		serviceInstanceNames[i] = serviceInstance.Name
	}
	sort.Strings(serviceInstanceNames)

	var spaceQuota SpaceQuota

	if space.SpaceQuotaDefinitionGUID != "" {
		spaceQuota, warnings, err = actor.GetSpaceQuota(space.SpaceQuotaDefinitionGUID)
		allWarnings = append(allWarnings, warnings...)
		if err != nil {
			return SpaceSummary{}, allWarnings, err
		}
	}

	securityGroups, warnings, err := actor.GetSpaceRunningSecurityGroupsBySpace(space.GUID)
	allWarnings = append(allWarnings, warnings...)
	if err != nil {
		return SpaceSummary{}, allWarnings, err
	}

	var securityGroupNames []string
	var securityGroupRules []SecurityGroupRule

	for _, securityGroup := range securityGroups {
		securityGroupNames = append(securityGroupNames, securityGroup.Name)
		securityGroupRules = append(securityGroupRules, extractSecurityGroupRules(securityGroup, "running")...)
	}

	securityGroups, warnings, err = actor.GetSpaceStagingSecurityGroupsBySpace(space.GUID)
	allWarnings = append(allWarnings, warnings...)
	if err != nil {
		return SpaceSummary{}, allWarnings, err
	}
	for _, securityGroup := range securityGroups {
		securityGroupRules = append(securityGroupRules, extractSecurityGroupRules(securityGroup, "staging")...)
	}

	sort.Sort(sortableSecurityGroupRules(securityGroupRules))

	sort.Strings(securityGroupNames)

	spaceSummary := SpaceSummary{
		SpaceName:            name,
		SpaceGUID:            space.GUID,
		OrgName:              org.Name,
		AppNames:             appNames,
		ServiceInstanceNames: serviceInstanceNames,
		SpaceQuotaName:       spaceQuota.Name,
		SecurityGroupNames:   securityGroupNames,
		SecurityGroupRules:   securityGroupRules,
	}

	return spaceSummary, allWarnings, nil
}

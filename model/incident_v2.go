package model

import (
	"github.com/incident-io/singer-tap/client"
	"github.com/samber/lo"
)

type incidentV2 struct{}

var IncidentV2 incidentV2

func (incidentV2) Schema() Property {
	return Property{
		Types: []string{"object"},
		Properties: map[string]Property{
			"id": {
				Types: []string{"string"},
			},
			"name": {
				Types: []string{"string"},
			},
			"call_url": {
				Types: []string{"string", "null"},
			},
			"creator":                   ActorV2.Schema(),
			"custom_field_entries":      ArrayOf(CustomFieldEntryV1.Schema()),
			"external_issue_reference":  Optional(ExternalIssueReferenceV2.Schema()),
			"incident_role_assignments": ArrayOf(IncidentRoleAssignmentV1.Schema()),
			"incident_status":           IncidentStatusV1.Schema(),
			"incident_timestamp_values": Optional(ArrayOf(IncidentTimestampWithValueV2.Schema())),
			"incident_type":             Optional(IncidentTypeV1.Schema()),
			"mode": {
				Types: []string{"string"},
			},
			"permalink": {
				Types: []string{"string", "null"},
			},
			"postmortem_document_url": {
				Types: []string{"string", "null"},
			},
			"reference": {
				Types: []string{"string"},
			},
			"severity": Optional(SeverityV2.Schema()),
			"slack_channel_id": {
				Types: []string{"string"},
			},
			"slack_channel_name": {
				Types: []string{"string", "null"},
			},
			"slack_team_id": {
				Types: []string{"string"},
			},
			"summary": {
				Types: []string{"string", "null"},
			},
			"visibility": {
				Types: []string{"string"},
			},
			"workload_minutes_late": {
				Types: []string{"number", "null"},
			},
			"workload_minutes_sleeping": {
				Types: []string{"number", "null"},
			},
			"workload_minutes_total": {
				Types: []string{"number", "null"},
			},
			"workload_minutes_working": {
				Types: []string{"number", "null"},
			},
			"created_at": DateTime.Schema(),
			"updated_at": DateTime.Schema(),
		},
	}
}

func (incidentV2) Serialize(input client.IncidentV2) map[string]any {
	return map[string]any{
		"id":       input.Id,
		"name":     input.Name,
		"call_url": input.CallUrl,
		"creator":  ActorV2.Serialize(input.Creator),
		"custom_field_entries": lo.Map(input.CustomFieldEntries, func(entry client.CustomFieldEntryV1, _ int) map[string]any {
			return CustomFieldEntryV1.Serialize(entry)
		}),
		"external_issue_reference": ExternalIssueReferenceV2.Serialize(input.ExternalIssueReference),
		"incident_role_assignments": lo.Map(input.IncidentRoleAssignments, func(assignment client.IncidentRoleAssignmentV1, _ int) map[string]any {
			return IncidentRoleAssignmentV1.Serialize(assignment)
		}),
		"incident_status": IncidentStatusV1.Serialize(input.IncidentStatus),
		"incident_timestamp_values": lo.Map(*input.IncidentTimestampValues, func(timestamp client.IncidentTimestampWithValueV2, _ int) map[string]any {
			return IncidentTimestampWithValueV2.Serialize(timestamp)
		}),
		"incident_type":             IncidentTypeV1.Serialize(input.IncidentType),
		"mode":                      input.Mode,
		"permalink":                 input.Permalink,
		"postmortem_document_url":   input.PostmortemDocumentUrl,
		"reference":                 input.Reference,
		"severity":                  SeverityV2.Serialize(input.Severity),
		"slack_channel_id":          input.SlackChannelId,
		"slack_channel_name":        input.SlackChannelName,
		"slack_team_id":             input.SlackTeamId,
		"summary":                   input.Summary,
		"visibility":                input.Visibility,
		"workload_minutes_late":     input.WorkloadMinutesLate,
		"workload_minutes_sleeping": input.WorkloadMinutesSleeping,
		"workload_minutes_total":    input.WorkloadMinutesTotal,
		"workload_minutes_working":  input.WorkloadMinutesWorking,
		"created_at":                input.CreatedAt,
		"updated_at":                input.UpdatedAt,
	}
}

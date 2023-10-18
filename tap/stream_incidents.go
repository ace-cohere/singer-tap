package tap

import (
	"context"

	kitlog "github.com/go-kit/log"
	"github.com/incident-io/singer-tap/client"
	"github.com/incident-io/singer-tap/model"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func init() {
	register(&StreamIncidents{})
}

type StreamIncidents struct {
}

func (s *StreamIncidents) Output() *Output {
	return &Output{
		Type:   OutputTypeSchema,
		Stream: "incidents",
		Schema: &model.Schema{
			HasAdditionalProperties: false,
			Type:                    []string{"object"},
			Properties:              model.IncidentV2.Schema().Properties,
		},
		KeyProperties:      []string{"id"},
		BookmarkProperties: []string{},
	}
}

func (s *StreamIncidents) GetRecords(ctx context.Context, logger kitlog.Logger, cl *client.ClientWithResponses) ([]map[string]any, error) {
	var (
		after    *string
		pageSize = int64(250)
		results  = []map[string]any{}
	)

	for {
		logger.Log("msg", "loading page", "page_size", pageSize, "after", after)
		page, err := cl.IncidentsV2ListWithResponse(ctx, &client.IncidentsV2ListParams{
			PageSize: &pageSize,
			After:    after,
		})
		if err != nil {
			return nil, errors.Wrap(err, "listing incidents")
		}

		for _, element := range page.JSON200.Incidents {
			results = append(results, model.IncidentV2.Serialize(element))
		}
		if count := len(page.JSON200.Incidents); count == 0 {
			return results, nil // end pagination
		} else {
			after = lo.ToPtr(page.JSON200.Incidents[count-1].Id)
		}
	}
}

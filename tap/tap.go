package tap

import (
	"context"
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/incident-io/singer-tap/client"
)

func Sync(ctx context.Context, logger kitlog.Logger, ol *OutputLogger, cl *client.ClientWithResponses, catalog *Catalog) error {
	// If we weren't given a catalog, create a default one and use that
	if catalog == nil {
		catalog = NewDefaultCatalog(streams)
	}

	// We only want to sync enabled streams
	enabledStreams := catalog.GetEnabledStreams()

	for _, catalogEntry := range enabledStreams {
		stream := streams[catalogEntry.Stream]
		logger := kitlog.With(logger, "stream", catalogEntry.Stream)

		logger.Log("msg", "outputting schema")
		if err := ol.Log(stream.Output()); err != nil {
			return err
		}

		timeExtracted := time.Now().Format(time.RFC3339)
		logger.Log("msg", "loading records", "time_extracted", timeExtracted)

		records, err := stream.GetRecords(ctx, logger, cl)
		if err != nil {
			return err
		}

		// Get the enabled fields for this stream
		disabledFields, err := catalog.GetDisabledFields(catalogEntry.Stream)
		if err != nil {
			return err
		}

		// Filter out the disabled fields from each record (ew)
		for _, record := range records {
			for fieldName := range disabledFields {
				delete(record, fieldName)
			}
		}

		logger.Log("msg", "outputting records", "count", len(records))
		for _, record := range records {
			op := &Output{
				Type:          OutputTypeRecord,
				Stream:        catalogEntry.Stream,
				Record:        record,
				TimeExtracted: timeExtracted,
			}
			if err := ol.Log(op); err != nil {
				return err
			}
		}
	}

	return nil
}

func Discover(ctx context.Context, logger kitlog.Logger, ol *OutputLogger) error {
	catalog := NewDefaultCatalog(streams)

	if err := ol.CataLog(catalog); err != nil {
		return err
	}

	return nil
}

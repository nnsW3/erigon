package heimdall

import (
	"context"
	"testing"
	"time"

	"github.com/erigontech/erigon-lib/log/v3"
)

func TestOver50EventBlockFetch(t *testing.T) {
	heimdallClient := NewHeimdallClient("https://heimdall-api.polygon.technology/", log.New())

	// block      := 48077376
	// block time := Sep-28-2023 08:13:58 AM
	events, err := heimdallClient.FetchStateSyncEvents(context.Background(), 2774290, time.Unix(1695888838-128, 0), 0)

	if err != nil {
		t.Fatal(err)
	}

	if len(events) != 113 {
		t.Fatal("Unexpected event count, exptected: 113, got:", len(events))
	}

	// block :=  23893568
	// block time := Jan-19-2022 04:48:04 AM

	events, err = heimdallClient.FetchStateSyncEvents(context.Background(), 1479540, time.Unix(1642567374, 0), 0)

	if err != nil {
		t.Fatal(err)
	}

	if len(events) != 9417 {
		t.Fatal("Unexpected event count, exptected: 9417, got:", len(events))
	}

}

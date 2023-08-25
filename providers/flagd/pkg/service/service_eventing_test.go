package service

import (
	schemaV1 "buf.build/gen/go/open-feature/flagd/protocolbuffers/go/schema/v1"
	"context"
	"errors"
	"github.com/open-feature/go-sdk-contrib/providers/flagd/pkg/cache"
	of "github.com/open-feature/go-sdk/pkg/openfeature"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
	"time"
)

func TestRetries(t *testing.T) {
	// given - stream is always errored
	client := MockClient{
		error: errors.New("streaming error"),
	}

	service := Service{
		maxRetries:     1,
		baseRetryDelay: 100 * time.Millisecond,
		client:         &client,
		cache:          cache.NewCacheService(cache.DisabledValue, 0, log),
		events:         make(chan of.Event),
	}

	// when - start event stream, knowing it will result in error
	go func() {
		service.startEventStream(context.Background())
	}()

	// then - expect an error event after retries
	var event of.Event
	select {
	case event = <-service.EventChannel():
		break
	case <-time.After(1 * time.Second):
		t.Fatal("timed out waiting for event")

	}

	if event.EventType != of.ProviderError {
		t.Errorf("expected event of %s, got %s", of.ProviderError, event.EventType)
	}
}

func TestConfigChane(t *testing.T) {
	data := map[string]interface{}{
		"flags": map[string]interface{}{
			"a": "",
			"b": "",
		},
	}

	stData, err := structpb.NewStruct(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("no cache - do nothing", func(t *testing.T) {
		// given
		service := Service{
			cache:  cache.NewCacheService(cache.DisabledValue, 0, log),
			events: make(chan of.Event),
		}

		// when
		go func() {
			service.handleConfigurationChangeEvent(&schemaV1.EventStreamResponse{
				Data: stData,
			})
		}()

		// then - expect no event
		select {
		case event := <-service.EventChannel():
			t.Fatalf("expected no event, but got with type: %s", event.EventType)
		case <-time.After(100 * time.Millisecond):
			// no events mean pass
			break
		}
	})

	t.Run("with cache - validate config change event", func(t *testing.T) {
		// given
		service := Service{
			cache:  cache.NewCacheService(cache.InMemValue, 1, log),
			events: make(chan of.Event),
		}

		// when
		go func() {
			service.handleConfigurationChangeEvent(&schemaV1.EventStreamResponse{
				Data: stData,
			})
		}()

		// then - expect no event
		select {
		case event := <-service.EventChannel():
			if event.EventType != of.ProviderConfigChange {
				t.Fatalf("expected event %s, got %s", of.ProviderConfigChange, event.EventType)
			}
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("timed out waiting for event")
		}
	})

}

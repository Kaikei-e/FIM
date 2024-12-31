package feed_collector_driver

import (
	"context"
	"encoding/json"
	"errors"
	"federation_orchestrator/protector"
	"federation_orchestrator/slogger"
	"net/http"
	"time"

	feedv1 "buf_schema/gen/feed/v1"
	feedv1connect "buf_schema/gen/feed/v1/feedv1connect"

	"connectrpc.com/connect"
	"github.com/doyensec/safeurl"
	"github.com/gorilla/mux"
)

type BffServer struct {
	apiClient     feedv1connect.FeedServiceClient
	safeurlClient *safeurl.WrappedClient
}

func NewBffServer(apiClient feedv1connect.FeedServiceClient, safeurlClient *safeurl.WrappedClient) *BffServer {
	return &BffServer{
		apiClient:     apiClient,
		safeurlClient: safeurlClient,
	}
}

func (s *BffServer) RegisterSingleFeedCollectorDriver(w http.ResponseWriter, r *http.Request, ctx context.Context) {
	vars := mux.Vars(r)
	feedURL := vars["url"]

	// Validate the URL
	targetLink, err := protector.AntiURLAttack(feedURL, s.safeurlClient)
	if err != nil {
		slogger.Logger.Error("Error validating URL", "error", err)
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	req := connect.NewRequest(&feedv1.GetFeedRequest{
		FeedUrl: targetLink.String(),
	})

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := s.apiClient.GetFeed(ctx, req)
	if err != nil {
		var connectErr *connect.Error
		if errors.As(err, &connectErr) {
			switch connectErr.Code() {
			case connect.CodeNotFound:
				http.Error(w, "Feed not found", http.StatusNotFound)
			case connect.CodeInvalidArgument:
				http.Error(w, "Invalid argument", http.StatusBadRequest)
			default:
				slogger.Logger.Error("GetFeed RPC error", "error", connectErr)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		} else {
			slogger.Logger.Error("GetFeed RPC unexpected error", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp.Msg); err != nil {
		slogger.Logger.Error("Failed to encode response", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

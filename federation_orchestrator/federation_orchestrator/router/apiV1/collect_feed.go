package apiv1

import (
	"context"
	"errors"
	feedv1 "federation_orchestrator/gen/schema/feed/v1"
	"federation_orchestrator/gen/schema/feed/v1/feedv1connect"
	"federation_orchestrator/slogger"
	"fmt"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/doyensec/safeurl"
	"github.com/labstack/echo/v4"
)

type targetURL struct {
	URL string `json:"url"`
}

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

func CollectSingleFeed(e echo.Context, ctx context.Context) error {
	var reqString targetURL

	err := e.Bind(&reqString)
	if err != nil {
		slogger.Logger.Error("Failed to bind the request body.", "Caused by", err)
		return e.JSON(400, "Failed to bind the request body.")
	}

	fmt.Println(reqString)

	s := feedv1connect.NewFeedServiceClient(http.DefaultClient,
		"http://10.0.100.30:8080") // TODO: replace with the dynamic address

	reqRPC := connect.NewRequest(&feedv1.GetFeedRequest{
		FeedUrl: reqString.URL,
	})

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err = s.GetFeed(ctx, reqRPC)
	if err != nil {
		var connectErr *connect.Error
		if errors.As(err, &connectErr) {
			switch connectErr.Code() {
			case connect.CodeNotFound:
				return e.JSON(http.StatusNotFound, map[string]string{"error": "Feed not found."})
			case connect.CodeInvalidArgument:
				return e.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid argument."})
			default:
				slogger.Logger.Error("GetFeed RPC error", "error", connectErr)
				return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error."})
			}
		}
		slogger.Logger.Error("GetFeed RPC unexpected error", "error", err)
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error."})
	}

	return e.JSON(200, "Successfully collected the feed.")
}

package collect_feed_usecase

import (
	"context"
	"database/sql"
	"testing"
)

func TestCollectSingleFeedUsecase(t *testing.T) {
	ctx := context.Background()

	type args struct {
		reqBody string
		db      *sql.DB
		ctx     context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "successfull collect single feed",
			args: args{
				reqBody: "https://lorem-rss.herokuapp.com/feed",
				db:      nil,
				ctx:     ctx,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CollectSingleFeedUsecase(tt.args.reqBody, tt.args.db, tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("CollectSingleFeedUsecase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

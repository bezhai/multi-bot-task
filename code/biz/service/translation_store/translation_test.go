package translation_store

import (
	"context"
	"testing"

	"github.com/bezhai/multi-bot-task/biz/dal"
	"github.com/bezhai/multi-bot-task/biz/model/translation"
	"github.com/bezhai/multi-bot-task/biz/utils/langx/ptr"
	"github.com/bezhai/multi-bot-task/config"
)

func TestListTranslation(t *testing.T) {
	type args struct {
		ctx context.Context
		req *translation.ListTranslationRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				ctx: context.Background(),
				req: &translation.ListTranslationRequest{
					Page:             1,
					PageSize:         10,
					SearchKey:        ptr.Ptr("embarrassed face"),
					OnlyUntranslated: true,
				},
			},
		},
	}
	config.InitEnv()
	dal.Init()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _ = ListTranslation(tt.args.ctx, tt.args.req)
		})
	}
}

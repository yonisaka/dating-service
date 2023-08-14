package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/yonisaka/dating-service/internal/consts"
	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/presentations"
	"go.uber.org/mock/gomock"
)

func TestSubscribe(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	type test struct {
		fields  subscribeFields
		args    args
		want    *presentations.SubscribeResponse
		wantErr error
	}

	tests := map[string]func(t *testing.T, ctrl *gomock.Controller) test{
		"Given valid request of Subscribe, When repository executed successfully, Return no error": func(t *testing.T, ctrl *gomock.Controller) test {

			ctx := context.Background()

			userID := int64(1)
			dateNow := time.Now()

			ctx = context.WithValue(ctx, consts.CtxAuthInfo, jwt.MapClaims{
				"authorized":  true,
				"dtid":        uuid.New().String(),
				"expiry_time": dateNow,
				"user_id":     userID,
			})

			args := args{
				ctx: ctx,
			}

			mockUserRepo := repository.NewGoMockUserRepo(ctrl)

			mockUser := repository.User{
				ID:        1,
				FirstName: "Yoni",
				LastName:  "Saka",
				Email:     "yonisaka0@gmail.com",
				Password:  "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8",
			}
			mockUserRepo.EXPECT().FindByID(args.ctx, userID).Return(&mockUser, nil)

			mockUserSubscriptionRepo := repository.NewGoMockUserSubscriptionRepo(ctrl)

			subscriptionCode := uuid.New().String()
			subscriptionExpiredAt := time.Now().AddDate(0, 1, 0)

			mockUserSubscriptionRepo.EXPECT().Store(args.ctx, gomock.Any()).Return(nil).AnyTimes()

			want := presentations.SubscribeResponse{
				SubscriptionCode: subscriptionCode,
				ExpiredAt:        subscriptionExpiredAt.Format(time.DateTime),
			}

			return test{
				fields: subscribeFields{
					userRepo:             mockUserRepo,
					userSubscriptionRepo: mockUserSubscriptionRepo,
				},
				args:    args,
				want:    &want,
				wantErr: nil,
			}
		},
	}

	for name, testFn := range tests {

		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt := testFn(t, ctrl)

			sut := subscribeSut(tt.fields)

			got, err := sut.Subscribe(tt.args.ctx)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			t.Log(got)
		})
	}
}

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

func TestActionHistory_List(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	type test struct {
		fields  actionHistoryFields
		args    args
		want    []presentations.ActionHistoryResponse
		wantErr error
	}

	tests := map[string]func(t *testing.T, ctrl *gomock.Controller) test{
		"Given valid request of Action History, When repository executed successfully, Return no error": func(t *testing.T, ctrl *gomock.Controller) test {

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

			mockUserActionHistoryRepo := repository.NewGoMockUserActionHistoryRepo(ctrl)

			mockUserActionHistoryRepo.EXPECT().FindByUserID(args.ctx, userID).Return([]repository.UserActionHistory{
				{
					ID:        1,
					UserID:    1,
					ProfileID: 2,
					Action:    "pass",
					UpdatedAt: &dateNow,
				},
				{
					ID:        2,
					UserID:    1,
					ProfileID: 3,
					Action:    "like",
					UpdatedAt: &dateNow,
				},
			}, nil)

			mockUserRepo := repository.NewGoMockUserRepo(ctrl)

			mockUserRepo.EXPECT().FindByID(args.ctx, gomock.Any()).Return(&repository.User{
				ID:        2,
				FirstName: "User",
				LastName:  "2",
				Email:     "user2@gmail.com",
				Dob:       &dateNow,
			}, nil).MaxTimes(1)

			mockUserRepo.EXPECT().FindByID(args.ctx, gomock.Any()).Return(&repository.User{
				ID:        3,
				FirstName: "User",
				LastName:  "3",
				Email:     "user3@gmail.com",
				Dob:       &dateNow,
			}, nil).MaxTimes(1)

			want := []presentations.ActionHistoryResponse{
				{
					FirstName: "User",
					LastName:  "2",
					Age:       0,
					Action:    "pass",
					UpdatedAt: dateNow.Format(time.DateTime),
				},
				{
					FirstName: "User",
					LastName:  "3",
					Age:       0,
					Action:    "like",
					UpdatedAt: dateNow.Format(time.DateTime),
				},
			}

			return test{
				fields: actionHistoryFields{
					userActionHistoryRepo: mockUserActionHistoryRepo,
					userRepo:              mockUserRepo,
				},
				args:    args,
				want:    want,
				wantErr: nil,
			}
		},
	}

	for name, testFn := range tests {

		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt := testFn(t, ctrl)

			sut := actionHistorySut(tt.fields)

			got, err := sut.GetActionHistories(tt.args.ctx)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

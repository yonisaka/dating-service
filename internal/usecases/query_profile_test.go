package usecases_test

import (
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/yonisaka/dating-service/internal/consts"
	"github.com/yonisaka/dating-service/internal/di"
	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/pkg/kvs"
	"github.com/yonisaka/dating-service/pkg/storage"
	"go.uber.org/mock/gomock"
)

func TestQueryProfile(t *testing.T) {
	type args struct {
		ctx context.Context
		req presentations.QueryProfileRequest
	}

	type test struct {
		fields  queryProfileFields
		args    args
		want    *presentations.QueryProfileResponse
		wantErr error
	}

	tests := map[string]func(t *testing.T, ctrl *gomock.Controller) test{
		"Given valid request of Query Profile, When repository executed successfully, Return no error": func(t *testing.T, ctrl *gomock.Controller) test {

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

			cfg := di.GetConfig()

			mockUserRepo := repository.NewGoMockUserRepo(ctrl)

			mockUser := repository.User{
				ID:        1,
				FirstName: "Yoni",
				LastName:  "Saka",
				Email:     "yonisaka0@gmail.com",
				Phone:     "08123456789",
				Password:  "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8",
				Dob:       &dateNow,
				Gender:    consts.Man,
			}

			mockUserRepo.EXPECT().FindByID(args.ctx, userID).Return(&mockUser, nil)

			mockProfiles := []repository.User{
				{
					ID:        2,
					FirstName: "User",
					LastName:  "1",
					Email:     "user1@gmail.com",
					Phone:     "08123456789",
					Dob:       &dateNow,
					Gender:    consts.Woman,
				},
				{
					ID:        3,
					FirstName: "User",
					LastName:  "2",
					Email:     "user2@gmail.com",
					Phone:     "08123456789",
					Dob:       &dateNow,
					Gender:    consts.Woman,
				},
				{
					ID:        4,
					FirstName: "User",
					LastName:  "3",
					Email:     "user4@gmail.com",
					Phone:     "08123456789",
					Dob:       &dateNow,
					Gender:    consts.Woman,
				},
			}

			mockUserRepo.EXPECT().Find(args.ctx, gomock.Any()).Return(mockProfiles, nil)

			mockUserPreferenceRepo := repository.NewGoMockUserPreferenceRepo(ctrl)

			mockUserPreference := repository.UserPreference{
				ID:        1,
				UserID:    1,
				MinAge:    18,
				MaxAge:    30,
				Gender:    consts.Woman,
				UseIntend: false,
			}

			mockUserPreferenceRepo.EXPECT().FindByUserID(args.ctx, userID).Return(&mockUserPreference, nil)

			mockUserSubscriptionRepo := repository.NewGoMockUserSubscriptionRepo(ctrl)

			expiredAt := time.Now().AddDate(1, 0, 0)
			mockUserSubscription := repository.UserSubscription{
				ID:               1,
				UserID:           1,
				SubscriptionCode: "77aec485-5196-4974-90a6-d0831eb72c8c",
				ExpiredAt:        &expiredAt,
			}

			mockUserSubscriptionRepo.EXPECT().FindByUserID(args.ctx, userID).Return(&mockUserSubscription, nil)

			mockUserImageRepo := repository.NewGoMockUserImageRepo(ctrl)

			mockUserImages := []repository.UserImage{
				{
					ID:       1,
					UserID:   1,
					ImageURL: "77aec485-5196-4974-90a6-d0831eb72c8c",
				},
			}

			mockUserImageRepo.EXPECT().FindByUserID(args.ctx, gomock.Any()).Return(mockUserImages, nil).AnyTimes()

			mockKvs := kvs.NewGoMockClient(ctrl)

			keyUserQuery := fmt.Sprintf("user:query:%d", userID)
			keyRecentQuery := fmt.Sprintf("user:recent:query:%d", userID)

			mockKvs.EXPECT().Get(ctx, keyUserQuery).Return(nil, nil)

			mockKvs.EXPECT().Get(ctx, keyRecentQuery).Return(nil, nil)

			mockKvs.EXPECT().Set(ctx, keyUserQuery, gomock.Any(), time.Duration(24)*time.Hour).Return(nil, nil)

			mockKvs.EXPECT().Set(ctx, keyRecentQuery, gomock.Any(), time.Duration(24)*time.Hour).Return(nil, nil)

			mockStorage := storage.NewGoMockStorage(ctrl)

			mockURL := url.URL{
				Scheme: "https",
				Host:   "cloudinary.com",
				Path:   "image",
			}

			mockStorage.EXPECT().GetSignedURL(cfg.Cloudinary.Bucket, mockUserImages[0].ImageURL).Return(&mockURL, nil)

			want := presentations.QueryProfileResponse{
				FirstName: mockProfiles[0].FirstName,
				LastName:  mockProfiles[0].LastName,
				Age:       mockProfiles[0].Age(),
				Gender:    mockProfiles[0].Gender,
				Images: []string{
					mockURL.String(),
				},
			}

			return test{
				fields: queryProfileFields{
					cfg:                  cfg,
					userRepo:             mockUserRepo,
					userPreferenceRepo:   mockUserPreferenceRepo,
					userSubscriptionRepo: mockUserSubscriptionRepo,
					userImageRepo:        mockUserImageRepo,
					kvs:                  mockKvs,
					storage:              mockStorage,
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

			sut := queryProfileSut(tt.fields)

			got, err := sut.GetQueryProfile(tt.args.ctx, tt.args.req)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

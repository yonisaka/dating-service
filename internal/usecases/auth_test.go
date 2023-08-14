package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yonisaka/dating-service/internal/consts"
	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/pkg/auth"
	"go.uber.org/mock/gomock"
)

// nolint:all
func TestAuth_Login(t *testing.T) {
	type args struct {
		ctx context.Context
		req presentations.LoginRequest
	}

	type test struct {
		fields  authFields
		args    args
		want    *presentations.LoginResponse
		wantErr error
	}

	tests := map[string]func(t *testing.T, ctrl *gomock.Controller) test{
		"Given valid request of Login, When repository executed successfully, Return no error": func(t *testing.T, ctrl *gomock.Controller) test {

			ctx := context.Background()

			args := args{
				ctx: ctx,
				req: presentations.LoginRequest{
					Email:    "yonisaka0@gmail.com",
					Password: "password",
				},
			}

			mockUserRepo := repository.NewGoMockUserRepo(ctrl)

			mockUser := repository.User{
				ID:        1,
				FirstName: "Yoni",
				LastName:  "Saka",
				Email:     "yonisaka0@gmail.com",
				Password:  "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8",
			}
			mockUserRepo.EXPECT().FindByEmail(args.ctx, args.req.Email).Return(&mockUser, nil)

			mockAuthenticator := auth.NewMockAuthenticator(ctrl)

			mockToken := auth.TokenDetails{
				AccessToken:  "token",
				RefreshToken: "refresh_token",
				AtExpires:    time.Now().Add(time.Minute * 15),
				RtExpires:    time.Now().Add(time.Hour * 24 * 7),
			}
			mockAuthenticator.EXPECT().CreateTokenJWT(mockUser.ID).Return(&mockToken)

			want := presentations.LoginResponse{
				AccessToken:  mockToken.AccessToken,
				RefreshToken: mockToken.RefreshToken,
				ExpiredAt:    mockToken.AtExpires.Format(time.DateTime),
			}

			return test{
				fields: authFields{
					userRepo:      mockUserRepo,
					authenticator: mockAuthenticator,
				},
				args:    args,
				want:    &want,
				wantErr: nil,
			}
		},
		"Given valid request of Login, When repository executed failed, Return error": func(t *testing.T, ctrl *gomock.Controller) test {

			ctx := context.Background()

			args := args{
				ctx: ctx,
				req: presentations.LoginRequest{
					Email:    "yonisaka0@gmail.com",
					Password: "password",
				},
			}

			mockUserRepo := repository.NewGoMockUserRepo(ctrl)

			mockUserRepo.EXPECT().FindByEmail(args.ctx, args.req.Email).Return(nil, errInternal)

			return test{
				fields: authFields{
					userRepo: mockUserRepo,
				},
				args:    args,
				want:    nil,
				wantErr: errInternal,
			}
		},
	}

	for name, testFn := range tests {

		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt := testFn(t, ctrl)

			sut := authSut(tt.fields)

			got, err := sut.Login(tt.args.ctx, tt.args.req)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAuth_Register(t *testing.T) {
	type args struct {
		ctx context.Context
		req presentations.RegisterRequest
	}

	type test struct {
		fields  authFields
		args    args
		want    *presentations.RegisterResponse
		wantErr error
	}

	tests := map[string]func(t *testing.T, ctrl *gomock.Controller) test{
		"Given valid request of Register, When repository executed successfully, Return no error": func(t *testing.T, ctrl *gomock.Controller) test {

			ctx := context.Background()

			args := args{
				ctx: ctx,
				req: presentations.RegisterRequest{
					FirstName: "Yoni",
					LastName:  "Saka",
					Email:     "yonisaka0@gmail.com",
					Password:  "password",
					Phone:     "08123456789",
					Dob:       "1997-01-01",
					Gender:    consts.Man,
					Intend:    "sort-term",
				},
			}

			mockUserRepo := repository.NewGoMockUserRepo(ctrl)

			dob := time.Now()
			mockUser := repository.User{
				ID:        1,
				FirstName: "Yoni",
				LastName:  "Saka",
				Email:     "yonisaka0@gmail.com",
				Phone:     "08123456789",
				Password:  "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8",
				Dob:       &dob,
				Gender:    consts.Man,
				Intend:    "sort-term",
			}

			mockUserRepo.EXPECT().Store(args.ctx, gomock.Any()).Return(nil)
			mockUserRepo.EXPECT().FindByEmail(args.ctx, args.req.Email).Return(&mockUser, nil)

			want := presentations.RegisterResponse{
				FirstName: mockUser.FirstName,
				LastName:  mockUser.LastName,
				Email:     mockUser.Email,
				Phone:     mockUser.Phone,
				Dob:       mockUser.Dob.Format(time.DateTime),
				Gender:    mockUser.Gender,
			}

			return test{
				fields: authFields{
					userRepo: mockUserRepo,
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

			sut := authSut(tt.fields)

			got, err := sut.Register(tt.args.ctx, tt.args.req)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

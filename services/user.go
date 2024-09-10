package services

import (
	"context"

	"github.com/marlonmp/books-app/models"
	"github.com/marlonmp/books-app/payloads"
	"github.com/marlonmp/books-app/repos"
)

type userService struct {
	users repos.UserRepo
	books repos.BookRepo
}

func (us userService) ListUsers(ctx context.Context, uf *repos.UserFilters) ([]payloads.UserList, error) {
	users, err := us.users.FilterMany(ctx, uf)

	if err != nil {
		return nil, err
	}

	usersPayload := payloads.UserListFromModels(users)

	return usersPayload, nil
}

func (us userService) SignUp(ctx context.Context, payload payloads.UserCreate) (payloads.UserList, error) {
	// make the user validation
	// err := payload.Validate()
	// if err != nil {
	// 	return payloads.UserList{}, err
	// }

	user, err := payload.ToModel()

	if err != nil {
		return payloads.UserList{}, err
	}

	// this status will make the users to verify their emails
	user.Status = models.UserStatusUnverified

	user, err = us.users.CreateOne(ctx, user)

	if err != nil {
		return payloads.UserList{}, err
	}

	// send the verification email
	// verification, err := auth.CreateUserVerification(ctx, user)
	// if err != nil {
	// 	return payloads.UserList{}, err
	// }
	// go mails.SendHTMLMail(ctx, user.Email, "subject", templates.VerificationEmail, verification, retry=true)

	usersPayload := payloads.UserListFromModel(user)

	return usersPayload, nil
}

func (us userService) SignIn(ctx context.Context, payload payloads.UserCredentials) (payloads.UserList, error) {
	// validate
	// err := payload.Validate()
	// if err != nil {
	// 	return payloads.UserList{}, err
	// }

	user, err := us.users.GetCredentialsByUsername(ctx, payload.Username, models.UserStatusActive)

	if repos.IsNotFoundError(err) {
		// must return an permissions denied
		return payloads.UserList{}, err
	}

	if err != nil {
		return payloads.UserList{}, err
	}

	if !user.Password.IsEqual(payload.Password) {
		// err = us.IncrementSignInTryes(ctx, user)
		// if err != nil {
		// 	return payloads.UserList{}, err
		// }
		return payloads.UserList{}, err
	}

	// session, err := us.sessions.CreateUserSession(ctx, user)
	// if err != nil {
	// 	return payloads.UserList{}, err
	// }

	usersPayload := payloads.UserListFromModel(user)

	return usersPayload, nil
}

func (us userService) UserProfile(ctx context.Context, username string) (payloads.UserProfile, error) {
	// validate username

	user, err := us.users.GetByUsername(ctx, username, models.UserStatusActive)

	if err != nil {
		return payloads.UserProfile{}, err
	}

	payload := payloads.UserProfileFromModel(user)

	return payload, nil
}

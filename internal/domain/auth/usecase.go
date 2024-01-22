package auth

import (
	"crypto/sha256"

	"cars/pkg/errors"
)

type Usecase struct {
	repo  IRepository
	token IToken
}

func New(repo IRepository, token IToken) *Usecase {
	return &Usecase{
		repo:  repo,
		token: token,
	}
}

// Signup hashes password and saves user data in repository.
func (uc *Usecase) Signup(req *SignupReq) (*SignupRes, error) {
	// Add user to database
	resRepo, err := uc.repo.Signup(&RepoSignupReq{
		Username:     req.Username,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: hash(req.Password),
	})
	if err != nil {
		return nil, errors.Wrap(err, "repository")
	}

	return NewSignupRes(resRepo), nil
}

// Signin compares request data with data in repository.
// If they are equal, then generates and saves a token.
func (uc *Usecase) Signin(req *SigninReq) (*SigninRes, error) {
	// Check credentials
	resRepo, err := uc.repo.Signin(&RepoSigninReq{
		Email:        req.Email,
		PasswordHash: hash(req.Password),
	})
	if err != nil {
		return nil, errors.Wrap(err, "repository")
	}

	// Create token
	token, err := uc.token.Create(resRepo.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "create token")
	}

	// Store token
	if err := uc.token.StoreUserRefresh(resRepo.UserID, token); err != nil {
		return nil, errors.Wrap(err, "store token")
	}

	return NewSigninRes(token), nil
}

// Refresh revokes requested token, generates and saves a new token.
func (uc *Usecase) Refresh(req *RefreshReq) (*RefreshRes, error) {
	// Parse expired token
	claims, err := uc.token.ParseExpired(req.Token)
	if err != nil {
		return nil, errors.Wrap(err, "parse token")
	}

	// Revoke expired token
	if err := uc.token.RevokeUserRefresh(claims.UserID, req.Token); err != nil {
		return nil, errors.Wrap(err, "revoke token")
	}

	// Check if user account is not deleted
	repoRes, err := uc.repo.IsDeleted(&RepoIsDeletedReq{
		UserID: claims.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "check if account deleted")
	}
	if repoRes.Deleted {
		return nil, errors.Wrap(errors.DeletedAccount, "account was deleted")
	}

	// Create new token
	token, err := uc.token.Create(claims.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "create token")
	}

	// Store new token
	if err := uc.token.StoreUserRefresh(claims.UserID, token); err != nil {
		return nil, errors.Wrap(err, "store token")
	}

	return NewRefreshRes(token), nil
}

// Sessions returns all stored user refresh tokens.
func (uc *Usecase) Sessions(req *SessionsReq) (*SessionsRes, error) {
	// Parse token
	claims, err := uc.token.Parse(req.Token)
	if err != nil {
		return nil, errors.Wrap(err, "parse token")
	}

	// Get tokens
	tokens, err := uc.token.ListUserRefresh(claims.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "get all tokens")
	}

	return NewSessionsRes(tokens), nil
}

// Signout parses and revokes requested token.
func (uc *Usecase) Signout(req *SignoutReq) (*SignoutRes, error) {
	// Parse token
	claims, err := uc.token.Parse(req.Token)
	if err != nil {
		return nil, errors.Wrap(err, "parse token")
	}

	// Revoke token
	if err := uc.token.RevokeUserRefresh(claims.UserID, req.Token); err != nil {
		return nil, errors.Wrap(err, "revoke token")
	}

	return NewSignoutRes(req.Token), nil
}

// SignoutAll parses and revokes all user tokens.
func (uc *Usecase) SignoutAll(req *SignoutAllReq) (*SignoutAllRes, error) {
	// Parse token
	claims, err := uc.token.Parse(req.Token)
	if err != nil {
		return nil, errors.Wrap(err, "parse token")
	}

	// Check if token is stored and remove them
	if err := uc.token.RevokeUserRefresh(claims.UserID, req.Token); err != nil {
		return nil, errors.Wrap(err, "revoke token")
	}

	// Revoke all other tokens
	tokens, err := uc.token.RevokeUserRefreshAll(claims.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "revoke all tokens")
	}

	return NewSignoutAllRes(tokens), nil
}

// hash generates SHA-256 hash string.
func hash(text string) string {
	ph := sha256.New()
	ph.Write([]byte(text)) // No errors
	return string(ph.Sum(nil))
}

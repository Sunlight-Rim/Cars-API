package auth

// Signup

type SignupReq struct {
	Username string
	Email    string
	Phone    uint64
	Password string
}

type SignupRes struct {
	UserID uint64
}

type RepoSignupReq struct {
	Username     string
	Email        string
	Phone        uint64
	PasswordHash string
}

type RepoSignupRes struct {
	UserID uint64
}

func NewSignupRes(r *RepoSignupRes) *SignupRes {
	return &SignupRes{UserID: r.UserID}
}

// Signin

type SigninReq struct {
	Email    string
	Password string
}

type SigninRes struct {
	Token string
}

type RepoSigninReq struct {
	Email        string
	PasswordHash string
}

type RepoSigninRes struct {
	UserID uint64
}

func NewSigninRes(token string) *SigninRes {
	return &SigninRes{Token: token}
}

// Is deleted

type RepoIsDeletedReq struct {
	UserID uint64
}

type RepoIsDeletedRes struct {
	Deleted bool
}

// Refresh

type RefreshReq struct {
	Token string
}

type RefreshRes struct {
	Token string
}

func NewRefreshRes(token string) *RefreshRes {
	return &RefreshRes{Token: token}
}

// Sessions

type SessionsReq struct {
	Token string
}

type SessionsRes struct {
	Tokens []string
}

func NewSessionsRes(tokens []string) *SessionsRes {
	return &SessionsRes{Tokens: tokens}
}

// Signout

type SignoutReq struct {
	Token string
}

type SignoutRes struct {
	Token string
}

func NewSignoutRes(token string) *SignoutRes {
	return &SignoutRes{Token: token}
}

// Signout all

type SignoutAllReq struct {
	Token string
}

type SignoutAllRes struct {
	Tokens []string
}

func NewSignoutAllRes(tokens []string) *SignoutAllRes {
	return &SignoutAllRes{Tokens: tokens}
}

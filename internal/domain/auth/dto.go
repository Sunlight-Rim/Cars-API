package auth

// Signup

type SignupReq struct {
	Username string
	Email    string
	Address  string
	Password string
}

type SignupRes struct {
	ID uint64
}

type RepoSignupReq struct {
	Username     string
	Email        string
	Address      string
	PasswordHash string
}

type RepoSignupRes struct {
	ID uint64
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
	Email    string
	Password string
}

type RepoSigninRes struct {
	ID uint64
}

// Signout

type SignoutReq struct {
	Token string
}

type SignoutRes struct {
	Token string
}

// Refresh

type RefreshReq struct {
	Token string
}

type RefreshRes struct {
	Token string
}

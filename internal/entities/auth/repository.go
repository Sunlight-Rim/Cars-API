package authEntity

// Signup

type RepositoryReqSignup struct {
	Username string
	Email    string
	Address  string
}

type RepositoryResSignup struct {
	ID uint64
}

// Signin

type RepositoryReqSignin struct {
	Email    string
	Password string
}

type RepositoryResSignin struct {
	Token string
}

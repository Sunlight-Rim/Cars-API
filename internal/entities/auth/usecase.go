package authEntity

// Signup

type UsecaseReqSignup struct {
	Username string
	Email    string
	Address  string
}

type UsecaseResSignup struct {
	ID uint64
}

// Signin

type UsecaseReqSignin struct {
	Email    string
	Password string
}

type UsecaseResSignin struct {
	Token string
}

// Signout

type UsecaseReqSignout struct {
	Token string
}

type UsecaseResSignout struct {
	Token string
}

// Refresh

type UsecaseReqRefresh struct {
	Token string
}

type UsecaseResRefresh struct {
	Token string
}

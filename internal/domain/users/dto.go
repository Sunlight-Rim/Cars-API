package usersDomain

// Usecase

type UcaseReqSignup struct {
	Username string
	Email    string
	Address  string
}

type UcaseReqSignin struct {
	Email    string
	Password string
}

type UcaseResSignin struct {
	Token string
}

type UcaseReqSignout struct {
	Token string
}

type UcaseReqRefresh struct {
	Token string
}

type UcaseResRefresh struct {
	Token string
}

type UcaseReqGet struct {
	Token string
}

type UcaseResGet struct {
	User *User
}

type UcaseReqUpdateInfo struct {
	Token    string
	Username string
	Address  string
}

type UcaseReqUpdatePassword struct {
	Token       string
	NewPassword string
}

type UcaseReqDelete struct {
	Token string
}

// Repository

type RepoReqSignup struct {
	Username string
	Email    string
	Address  string
}

type RepoReqSignin struct {
	Email    string
	Password string
}

type RepoResSignin struct {
	Token string
}

type RepoReqSignout struct {
	Token string
}

type RepoReqGet struct {
	Token string
}

type RepoResGet struct {
	User *User
}

type RepoReqUpdateInfo struct {
	Token    string
	Username string
	Address  string
}

type RepoReqUpdatePassword struct {
	Token       string
	NewPassword string
}

type RepoReqDelete struct {
	Token string
}

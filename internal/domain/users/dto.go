package usersDomain

// Usecase

type UcaseReqRegister struct {
	Username string
	Email    string
	Address  string
}

type UcaseReqLogin struct {
	Email    string
	Password string
}

type UcaseResLogin struct {
	Token string
}

type UcaseReqLogout struct {
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

type RepoReqRegister struct {
	Username string
	Email    string
	Address  string
}

type RepoReqLogin struct {
	Email    string
	Password string
}

type RepoResLogin struct {
	Token string
}

type RepoReqLogout struct {
	Token string
}

type RepoReqRefresh struct {
	Token string
}

type RepoResRefresh struct {
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

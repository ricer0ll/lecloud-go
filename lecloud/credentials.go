package lecloud

type Credentials struct {
	Username  string
	Password  string
	AccountID string
}

func NewCredentials(username string, password string, accountID string) *Credentials {
	return &Credentials{
		Username:  username,
		Password:  password,
		AccountID: accountID,
	}
}

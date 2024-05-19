package reset_password

import "PopcornMovie/ent"

type Repository interface {
	ResetPasswordQuery() *ent.ResetPasswordQuery
	ResetPasswordCreate() *ent.ResetPasswordCreate
}

type impl struct {
	client *ent.Client
}

func (i impl) ResetPasswordQuery() *ent.ResetPasswordQuery {
	return i.client.ResetPassword.Query()
}

func (i impl) ResetPasswordCreate() *ent.ResetPasswordCreate {
	return i.client.ResetPassword.Create()
}

func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}

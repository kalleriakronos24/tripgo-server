package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/models/master"
)

type CheckExistingEntityCredentialsStruct struct {
	*master.Credentials
}

func (module *module) RetrieveEntityCredentialsByUserID(userId uuid.UUID) (m master.Credentials, err error) {
	if m, err = module.db.credentialModel.GetOneByID(userId); err != nil {
		return m, errors.New("failed to get user information")
	}
	return
}

func (module *module) RetrieveEntityCredentialsByEmail(email string) (m master.Credentials, err error) {
	if m, err = module.db.credentialModel.GetOneByEmail(email); err != nil {
		return m, errors.New("failed to get user information")
	}
	return
}

func (module *module) RetrieveAllEntityCredentials(id uuid.UUID) (m []master.Credentials, err error) {
	return
}

func (module *module) RetrieveAllEntityCredentialsPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	return
}

func (module *module) RetrieveEntityCredentials(id uuid.UUID) (m master.Credentials, err error) {
	return
}

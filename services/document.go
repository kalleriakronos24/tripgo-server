package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
)

type CheckExistingDocumentStruct struct {
	*models.Document
}

func (module *module) RetrieveAllDocument(id uuid.UUID) (m []models.Document, err error) {
	if m, err = module.db.documentModel.GetAllDocument(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveDocument(id uuid.UUID) (m models.Document, err error) {
	if m, err = module.db.documentModel.GetOneDocumentByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertDocument(p *dto.InsertDocument) (err error) {
	if _, err = module.db.documentModel.InsertDocument(models.Document{
		Base64:            p.Base64,
		Path:              p.Path,
		AbsolutePath:      p.AbsolutePath,
		FileName:          p.FileName,
		Extension:         p.Extension,
		Location:          p.Location,
		DocumentCreatedBy: p.CreatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) UpdateDocument(id uuid.UUID, p *dto.UpdateDocument) (err error) {
	if _, err = module.db.documentModel.UpdateDocument(id, models.Document{
		Base64:            p.Base64,
		Path:              p.Path,
		AbsolutePath:      p.AbsolutePath,
		FileName:          p.FileName,
		Extension:         p.Extension,
		Location:          p.Location,
		DocumentUpdatedBy: p.UpdatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) CheckExistingDocument(id string, param CheckExistingDocumentStruct) (err error) {

	if param.FileName != "" {
		if _, dbErr := module.db.documentModel.GetOneDocumentByFileName(param.FileName); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.documentModel.GetOneDocumentByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}

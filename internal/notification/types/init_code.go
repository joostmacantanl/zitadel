package types

import (
	"github.com/caos/zitadel/internal/config/systemdefaults"
	"github.com/caos/zitadel/internal/crypto"
	"github.com/caos/zitadel/internal/notification/templates"
	es_model "github.com/caos/zitadel/internal/user/repository/eventsourcing/model"
	view_model "github.com/caos/zitadel/internal/user/repository/view/model"
)

type InitCodeEmailData struct {
	templates.TemplateData
	FirstName string
	LastName  string
	URL       string
}

type UrlData struct {
	UserID string
	Code   string
}

func SendUserInitCode(user *view_model.NotifyUser, code *es_model.InitUserCode, systemDefaults systemdefaults.SystemDefaults, alg crypto.EncryptionAlgorithm) error {
	codeString, err := crypto.DecryptString(code.Code, alg)
	if err != nil {
		return err
	}
	url, err := templates.ParseTemplateText(systemDefaults.Notifications.Endpoints.InitCode, &UrlData{UserID: user.ID, Code: codeString})
	if err != nil {
		return err
	}
	initCodeData := &InitCodeEmailData{TemplateData: systemDefaults.Notifications.TemplateData.InitCode, FirstName: user.FirstName, LastName: user.LastName, URL: url}

	template, err := templates.GetParsedTemplate(initCodeData)
	if err != nil {
		return err
	}
	return generateEmail(user, systemDefaults.Notifications.TemplateData.InitCode.Subject, template, systemDefaults.Notifications, true)
}
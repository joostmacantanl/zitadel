package view

import (
	global_model "github.com/caos/zitadel/internal/model"
	proj_model "github.com/caos/zitadel/internal/project/model"
	"github.com/caos/zitadel/internal/project/repository/view/model"
	"github.com/caos/zitadel/internal/view"
	"github.com/jinzhu/gorm"
)

func ProjectMemberByIDs(db *gorm.DB, table, projectID, userID string) (*model.ProjectMemberView, error) {
	role := new(model.ProjectMemberView)

	projectIDQuery := model.ProjectMemberSearchQuery{Key: proj_model.PROJECTMEMBERSEARCHKEY_PROJECT_ID, Value: projectID, Method: global_model.SEARCHMETHOD_EQUALS}
	userIDQuery := model.ProjectMemberSearchQuery{Key: proj_model.PROJECTMEMBERSEARCHKEY_USER_ID, Value: userID, Method: global_model.SEARCHMETHOD_EQUALS}
	query := view.PrepareGetByQuery(table, projectIDQuery, userIDQuery)
	err := query(db, role)
	return role, err
}

func SearchProjectMembers(db *gorm.DB, table string, req *proj_model.ProjectMemberSearchRequest) ([]*model.ProjectMemberView, int, error) {
	roles := make([]*model.ProjectMemberView, 0)
	query := view.PrepareSearchQuery(table, model.ProjectMemberSearchRequest{Limit: req.Limit, Offset: req.Offset, Queries: req.Queries})
	count, err := query(db, &roles)
	if err != nil {
		return nil, 0, err
	}
	return roles, count, nil
}
func ProjectMembersByUserID(db *gorm.DB, table string, userID string) ([]*model.ProjectMemberView, error) {
	members := make([]*model.ProjectMemberView, 0)
	queries := []*proj_model.ProjectMemberSearchQuery{
		&proj_model.ProjectMemberSearchQuery{Key: proj_model.PROJECTMEMBERSEARCHKEY_USER_ID, Value: userID, Method: global_model.SEARCHMETHOD_EQUALS},
	}
	query := view.PrepareSearchQuery(table, model.ProjectMemberSearchRequest{Queries: queries})
	_, err := query(db, &members)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func PutProjectMember(db *gorm.DB, table string, role *model.ProjectMemberView) error {
	save := view.PrepareSave(table)
	return save(db, role)
}

func DeleteProjectMember(db *gorm.DB, table, projectID, userID string) error {
	role, err := ProjectMemberByIDs(db, table, projectID, userID)
	if err != nil {
		return err
	}
	delete := view.PrepareDeleteByObject(table, role)
	return delete(db)
}
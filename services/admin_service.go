package services

import (
	"mitchvillanueva.com/pulseid/domain/admin"
)

var (
	AdminService adminServiceInterface = &adminService{}
)

type adminService struct{}

type adminServiceInterface interface {
	GenerateToken(admin.Admin) (*admin.RequestToken, error)
}

func (s *adminService) GenerateToken(request admin.Admin) (*admin.RequestToken, error) {
	requestToken := &admin.RequestToken{}
	dao := &admin.Admin{
		Username: request.Username,
		Password: request.Password,
	}

	if err := dao.FindByUsernameAndPassword(); err != nil {
		return nil, err
	}

	jw := JwtServices{}
	token := jw.GenerateToken(request.Username, true)


	requestToken.Id = dao.Id
	requestToken.Username = dao.Username
	requestToken.Token = token

	return requestToken, nil
}


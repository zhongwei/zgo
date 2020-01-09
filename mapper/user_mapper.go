package mapper

import (
	"zgo/dto"
	"zgo/model"
)

func ToUser(userDTO dto.UserDTO) model.User {
	return model.User{Name: userDTO.Name, Age: userDTO.Age}
}

func ToUserDTO(user model.User) dto.UserDTO {
	return dto.UserDTO{ID: user.ID, Name: user.Name, Age: user.Age}
}

func ToUserDTOs(users []model.User) []dto.UserDTO {
	userdtos := make([]dto.UserDTO, len(users))

	for i, itm := range users {
		userdtos[i] = ToUserDTO(itm)
	}

	return userdtos
}

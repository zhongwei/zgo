package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"zgo/dto"
	"zgo/mapper"
	"zgo/model"
	"zgo/service"
)

type UserAPI struct {
	UserService service.UserService
}

func CreateUserAPI(u service.UserService) UserAPI {
	return UserAPI{UserService: u}
}

func (u *UserAPI) FindAll(c *gin.Context) {
	users := u.UserService.FindAll()

	c.JSON(http.StatusOK, gin.H{"users": mapper.ToUserDTOs(users)})
}

func (u *UserAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := u.UserService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"user": mapper.ToUserDTO(user)})
}

func (u *UserAPI) Create(c *gin.Context) {
	var userDTO dto.UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdUser := u.UserService.Save(mapper.ToUser(userDTO))

	c.JSON(http.StatusOK, gin.H{"user": mapper.ToUserDTO(createdUser)})
}

func (p *UserAPI) Update(c *gin.Context) {
	var userDTO dto.UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user := p.UserService.FindByID(uint(id))
	if user == (model.User{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	user.Name = userDTO.Name
	user.Age = userDTO.Age
	p.UserService.Save(user)

	c.Status(http.StatusOK)
}

func (p *UserAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := p.UserService.FindByID(uint(id))
	if user == (model.User{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.UserService.Delete(user)

	c.Status(http.StatusOK)
}

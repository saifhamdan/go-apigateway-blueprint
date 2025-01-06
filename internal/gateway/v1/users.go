// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saifhamdan/go-apigateway-blueprint/internal/gateway/utils"
	models "github.com/saifhamdan/go-apigateway-blueprint/models/v1"
	_ "github.com/saifhamdan/go-apigateway-blueprint/pkg/http"

	"gorm.io/gorm"
)

// @Id				GetUsers
// @Description	Get Users
// @Tags			Users
// @Accept			json
// @Produce		json
// @Success		200	{object}	[]models.User
// @Failure		500	{object}	http.HttpResponse
// @Param			page	query	int		false	"page number"
// @Param			limit	query	int		false	"limit number"
// @Param			from	query	string	false	"from date"
// @Param			to		query	string	false	"to date"
// @Param			sort_by	query	string	false	"sort by"
// @Router			/api/v1/users [get]
func (g *GatewayV1) GetUsers(c *fiber.Ctx) error {
	queryObj, err := utils.QueryFilter(c)
	if err != nil {
		return g.App.HttpResponseBadQueryParams(c, err)
	}

	users := new([]*models.User)
	err = g.DB.Postgres.
		Offset(queryObj.Page * queryObj.Limit).
		Limit(queryObj.Limit).
		Order(queryObj.SortBy).
		Find(&users).
		Error
	if err != nil {
		return g.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	return g.App.HttpResponseOK(c, users)
}

// @Id				GetUser
// @Description	Get User
// @Tags			Users
// @Accept			json
// @Produce		json
// @Success		200	{object}	models.User
// @Failure		500	{object}	http.HttpResponse
// @Param			username	path	int	true	"User ID"
// @Router			/api/v1/users/{id} [get]
func (g *GatewayV1) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return g.App.HttpResponseBadRequest(c, err)
	}

	user := new(models.User)

	err = g.DB.Postgres.
		Where("id = ?", id).
		First(&user).
		Error
	if err != nil {
		return g.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	return g.App.HttpResponseOK(c, user)
}

// @Id				CreateUser
// @Description	Create User
// @Tags			Users
// @Accept			json
// @Produce		json
// @Success		201	{object}	models.User
// @Failure		400	{object}	http.HttpResponse
// @Param			body	body	models.User	true	"User"
// @Router			/api/v1/users [post]
func (g *GatewayV1) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	err := c.BodyParser(user)
	if err != nil {
		return g.App.HttpResponseBadRequest(c, err)
	}

	err = g.DB.Postgres.Create(user).Error
	if err != nil {
		return g.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	return g.App.HttpResponseCreated(c, user)
}

// @Id				UpdateUser
// @Description	Update User
// @Tags			Users
// @Accept			json
// @Produce		json
// @Success		200	{object}	models.User
// @Failure		400	{object}	http.HttpResponse
// @Param			id		path	int			true	"User ID"
// @Param			body	body	models.User	true	"User"
// @Router			/api/v1/users/{id} [patch]
func (g *GatewayV1) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return g.App.HttpResponseBadRequest(c, err)
	}

	user := new(models.User)
	err = g.DB.Postgres.First(user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return g.App.HttpResponseNotFound(c, err)
		}
		return g.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	err = c.BodyParser(user)
	if err != nil {
		return g.App.HttpResponseBadRequest(c, err)
	}

	err = g.DB.Postgres.Save(user).Error
	if err != nil {
		return g.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	return g.App.HttpResponseOK(c, user)
}

// @Id				DeleteUser
// @Description	Delete User
// @Tags			Users
// @Accept			json
// @Produce		json
// @Success		204	{object}	http.HttpResponse
// @Failure		400	{object}	http.HttpResponse
// @Param			id	path	int	true	"User ID"
// @Router			/api/v1/users/{id} [delete]
func (g *GatewayV1) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return g.App.HttpResponseBadRequest(c, err)
	}

	user := new(models.User)
	err = g.DB.Postgres.First(user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return g.App.HttpResponseNotFound(c, err)
		}
		return g.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	err = g.DB.Postgres.Delete(user).Error
	if err != nil {
		return g.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	return g.App.HttpResponseNoContent(c)
}

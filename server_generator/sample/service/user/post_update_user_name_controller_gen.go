// generated version: 0.3.5

package user

import (
	"context"

	"github.com/labstack/echo/v4"
)

// PostUpdateUserNameController ...
type PostUpdateUserNameController struct {
}

// NewPostUpdateUserNameController ...
func NewPostUpdateUserNameController() *PostUpdateUserNameController {
	p := &PostUpdateUserNameController{}
	return p
}

// PostUpdateUserName ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param Name body string WIP:${isRequire} WIP:${description}
// @Success 200 {object} PostUpdateUserNameResponse
// @Failure 400 {object} WIP
// @Router /service/user/update_user_name [POST]
func (p *PostUpdateUserNameController) PostUpdateUserName(
	ctx context.Context, c echo.Context, req *PostUpdateUserNameRequest,
) (res *PostUpdateUserNameResponse, err error) {
	panic("require implements.")
}

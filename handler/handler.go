package handler

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/yeom-c/admin-template-go-fiber-api/app"
	"github.com/yeom-c/admin-template-go-fiber-api/auth_token"
	"github.com/yeom-c/admin-template-go-fiber-api/database"
	"github.com/yeom-c/admin-template-go-fiber-api/util"
)

type handler struct{}

var once sync.Once
var instance *handler

func Handler() *handler {
	once.Do(func() {
		if instance == nil {
			instance = &handler{}
		}
	})

	return instance
}

func (h *handler) errRes(c *fiber.Ctx, status int, res string) error {
	return c.Status(status).JSON(fiber.Map{
		"error": res,
	})
}

func (h *handler) SignIn(c *fiber.Ctx) error {
	var req signInReq
	if err := c.BodyParser(&req); err != nil {
		return h.errRes(c, fiber.StatusBadRequest, fmt.Sprintf("요청 파라미터 오류: %s", err.Error()))
	}

	user := database.User{}
	if req.AuthProvider == "email" {
		user.Email = req.Email
		has, err := database.Database().Conn.Get(&user)
		if err != nil {
			return h.errRes(c, fiber.StatusInternalServerError, err.Error())
		}

		if !has {
			return h.errRes(c, fiber.StatusNotFound, "유저 정보 없음")
		}

		err = util.CheckPassword(req.Password, user.HashedPassword)
		if err != nil {
			return h.errRes(c, fiber.StatusUnauthorized, "인증 실패")
		}
	} else if req.AuthProvider == "google" {
		authCode := req.AuthCode
		userInfo, err := util.GetGoogleUserInfo(authCode)
		if err != nil {
			return h.errRes(c, fiber.StatusUnauthorized, fmt.Sprintf("구글 인증 오류: %s", err.Error()))
		}

		has, err := database.Database().Conn.Where("email = ?", userInfo.Email).Get(&user)
		if err != nil {
			return h.errRes(c, fiber.StatusInternalServerError, err.Error())
		}

		if !has {
			user.Email = userInfo.Email
			user.Name = userInfo.Name
			_, err = database.Database().Conn.Insert(&user)
			if err != nil {
				return h.errRes(c, fiber.StatusInternalServerError, err.Error())
			}
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "잘못된 로그인 유형",
		})
	}

	accessToken, err := auth_token.TokenMaker().CreateToken(int64(user.Id), user.Name, app.Config().AuthTokenDuration)
	if err != nil {
		return h.errRes(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(signInRes{
		AccessToken: accessToken,
		UserRes: userRes{
			Id:        user.Id,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
		},
	})
}

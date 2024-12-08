package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/config"
	"server/internal/entities"
	"server/internal/log"
	"server/internal/repository/postgres"
	"server/pkg"
	"server/util"
	"strconv"
)

// SignUp
// @Tags         user
// @Summary      User signup
// @Description  Registers a new user and returns access and refresh tokens.
// @Accept       json
// @Produce      json
// @Param        data body entities.CreateUserRequest true "User  signup information"
// @Success      200 {object} entities.CreateUserResponse "User  successfully registered"
// @Failure      400 {object} entities.ErrorResponse "User  already exists or invalid request payload"
// @Failure      500 {object} entities.ErrorResponse "Internal server error"
// @Router       /signup [post]
func (h *Handler) SignUp(c *fiber.Ctx) error {
	var u entities.CreateUserRequest
	if err := c.BodyParser(&u); err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call postgres.DBUserExists")
	exists, err := postgres.DBUserExists(h.db, u.Email)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if exists {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("user already exists")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user already exists"})
	}

	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	user := &entities.User{
		Email:     u.Email,
		Password:  hashedPassword,
		Role:      u.Role,
		Name:      u.Name,
		Surname:   u.Surname,
		ThirdName: u.ThirdName,
	}

	h.logger.Debug().Msg("call postgres.DBUserCreate")
	r, err := postgres.DBUserCreate(h.db, user)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	TokenExpiration, err := strconv.Atoi(config.TokenExpiration)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg("wrong data")
		return c.Status(500).JSON(fiber.Map{"error": "wrong data"})
	}
	h.logger.Debug().Msg("call pkg.GenerateAccessToken")
	accessToken, err := pkg.GenerateAccessToken(user.ID, TokenExpiration,
		config.SigningKey)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call pkg.GenerateRefreshToken")
	refreshToken, err := pkg.GenerateRefreshToken(user.ID, config.SigningKey)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	res := &entities.CreateUserResponse{
		ID:           r.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(res)

}

// Login
// @Tags         user
// @Summary      User login
// @Description  Authenticates a user and returns access and refresh tokens.
// @Accept       json
// @Produce      json
// @Param        data body entities.LoginUserRequest true "User  login credentials"
// @Success      200 {object} entities.LoginUserResponse "User  successfully logged in"
// @Failure      400 {object} entities.ErrorResponse "Invalid email or password"
// @Failure      500 {object} entities.ErrorResponse "Internal server error"
// @Router       /login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	var user entities.LoginUserRequest
	if err := c.BodyParser(&user); err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	h.logger.Debug().Msg("call postgres.DBUserGetByLogin")
	u, err := postgres.DBUserGetByEmail(h.db, user.Email)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("wrong data")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "wrong data"})
	}
	h.logger.Debug().Msg("call util.CheckPassword")
	err = util.CheckPassword(user.Password, u.Password)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("wrong data")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "wrong data"})
	}

	TokenExpiration, err := strconv.Atoi(config.TokenExpiration)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("wrong data")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "wrong data"})
	}
	h.logger.Debug().Msg("call pkg.GenerateAccessToken")
	accessToken, err := pkg.GenerateAccessToken(u.ID, TokenExpiration,
		config.SigningKey)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call pkg.GenerateRefreshToken")
	refreshToken, err := pkg.GenerateRefreshToken(u.ID, config.SigningKey)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res := entities.LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ID:           u.ID,
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(res)
}
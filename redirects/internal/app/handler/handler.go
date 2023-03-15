package handler

import (
	"github.com/alikhanturusbekov/redirects/internal/app/cache"
	"github.com/alikhanturusbekov/redirects/internal/app/database"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/alikhanturusbekov/redirects/internal/app/model"
)

func Redirect(c *fiber.Ctx) error {
	var redirect model.Redirect
	link := c.Query("link")
	if link == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Provide link parameter"})
	}

	database.DB.Db.Where("history_link = ?", link).First(&redirect)

	if redirect.HistoryLink != "" {
		cache.Cr.Add(redirect.HistoryLink, redirect.ActiveLink)

		return c.Status(fiber.StatusMovedPermanently).JSON(fiber.Map{"status": "success", "link": redirect.HistoryLink})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "your link is active"})
}

func CreateRedirect(c *fiber.Ctx) error {
	var payload model.Redirect

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := model.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	now := time.Now()
	newNote := model.Redirect{
		ActiveLink:  payload.ActiveLink,
		HistoryLink: payload.HistoryLink,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	result := database.DB.Db.Create(&newNote)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Title already exist, please use another title"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"redirect": newNote}})
}

func GetRedirects(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var notes []model.Redirect
	results := database.DB.Db.Limit(intLimit).Offset(offset).Find(&notes)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(notes), "notes": notes})
}

func UpdateRedirect(c *fiber.Ctx) error {
	redirectId := c.Params("redirectId")

	var payload model.UpdateRedirect

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var redirect model.Redirect
	result := database.DB.Db.First(&redirect, "id = ?", redirectId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.ActiveLink != "" {
		updates["active_link"] = payload.ActiveLink
	}
	if payload.HistoryLink != "" {
		updates["history_link"] = payload.HistoryLink
	}

	updates["updated_at"] = time.Now()

	database.DB.Db.Model(&redirect).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": redirect}})
}

func GetRedirectById(c *fiber.Ctx) error {
	redirectId := c.Params("redirectId")

	var redirect model.Redirect
	result := database.DB.Db.First(&redirect, "id = ?", redirectId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": redirect}})
}

func DeleteRedirect(c *fiber.Ctx) error {
	redirectId := c.Params("redirectId")

	result := database.DB.Db.Delete(&model.Redirect{}, "id = ?", redirectId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

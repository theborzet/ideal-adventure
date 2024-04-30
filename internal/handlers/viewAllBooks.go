package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/pagination"
)

func (h Handler) ViewAllBooks(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	title := c.FormValue("search")
	genre := c.FormValue("genre")
	authorID, _ := strconv.Atoi(c.FormValue("author"))
	yearFrom, _ := strconv.Atoi(c.FormValue("year_from"))
	yearTo, _ := strconv.Atoi(c.FormValue("year_to"))

	books, err := h.repo.GetAllBooks(genre, title, uint(authorID), yearFrom, yearTo)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	authors, err := h.repo.GetAllAuthors("", "")
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	genres, err := h.repo.GetAllGenres()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	pageSize := 10
	paginatedBooks, paginator := pagination.PaginateBooks(books, page, pageSize)

	return c.Render("book_list", fiber.Map{
		"Books":       paginatedBooks,
		"Paginator":   paginator,
		"Title":       "Список книг",
		"IsPaginated": paginator.TotalPages > 1,
		"Authors":     authors,
		"Genres":      genres,
	})
}

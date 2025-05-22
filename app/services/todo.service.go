package services

import (
	"context"
	"log"
	"strconv"

	"skillsrock-test/app/types"
	"skillsrock-test/config/database"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	var t types.Task
	if err := c.BodyParser(&t); err != nil {
		log.Println("Ошибка заполнения задачи:", err)
		return c.Status(400).SendString(err.Error())
	}

	err := database.DB.QueryRow(context.Background(),
		"INSERT INTO tasks(title, description, status) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		t.Title, t.Description, t.Status, t.CreatedAt, t.UpdatedAt).Scan(&t.ID)
	if err != nil {
		log.Println("Ошибка при создании задачи:", err)
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(t)
}

func GetTasks(c *fiber.Ctx) error {
	rows, err := database.DB.Query(context.Background(),
		"SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil {
		log.Println("Ошибка запроса списка задач:", err)
		return c.Status(500).SendString(err.Error())
	}

	defer rows.Close()

	var tasks []types.Task

	for rows.Next() {
		var t types.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			log.Println("Ошибка при сканировании строки:", err)
			return c.Status(500).SendString(err.Error())
		}

		tasks = append(tasks, t)
	}

	// if len(tasks) == 0 {
	// 	return c.Status(200).JSON([]Task{})
	// }

	return c.JSON(tasks)
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")

	var t types.Task
	if err := c.BodyParser(&t); err != nil {
		log.Println("Ошибка заполнения полей:", err)
		return c.Status(400).SendString(err.Error())
	}

	_, err := database.DB.Exec(context.Background(),
		"UPDATE tasks SET title=$1, description=$2, status=$3, created_at=$4, updated_at=$5 WHERE id=$6",
		t.Title, t.Description, t.Status, t.CreatedAt, t.UpdatedAt, id)
	if err != nil {
		log.Println("Ошибка при обновлении задачи:", err)
		return c.Status(500).SendString(err.Error())
	}

	t.ID, err = strconv.Atoi(id)
	if err != nil {
		log.Println("Ошибка конвертации id:", err)
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(t)
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := database.DB.Exec(context.Background(),
		"DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		log.Println("Ошибка при удалении задачи:", err)
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}

func Preview(c *fiber.Ctx) error {
	endpoints := "POST /tasks – создание задачи.\n\nGET /tasks – получение списка всех задач.\n\nPUT /tasks/:id – обновление задачи.\n\nDELETE /tasks/:id – удаление задачи."
	return c.SendString(endpoints)
}

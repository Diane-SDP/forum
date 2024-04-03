package models

import (
	"strconv"
)

type Category struct {
	Title string
	Bio   string
	Image string
	Id    string
}

func GetCategories() ([]Category, error) {
	rows, err := DB.Query("SELECT id, title, description, image FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []Category
	var title string
	var bio string
	var idcategory string
	var pathImg string
	for rows.Next() {
		err := rows.Scan(&idcategory, &title, &bio, &pathImg)
		if err != nil {
			return nil, err
		}
		var category Category
		category.Title = title
		category.Bio = bio
		category.Id = idcategory
		category.Image = pathImg
		categories = append(categories, category)
	}
	return categories, nil
}

func GetCategory(id int) Category {
	rows, err := DB.Query("SELECT title, description, image FROM categories WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var title string
	var description string
	var pathImg string
	for rows.Next() {
		err := rows.Scan(&title, &description, &pathImg)
		if err != nil {
			panic(err)
		}
	}
	var category Category
	category.Title = title
	category.Bio = description
	category.Image = pathImg
	category.Id = strconv.Itoa(id)
	return category
}

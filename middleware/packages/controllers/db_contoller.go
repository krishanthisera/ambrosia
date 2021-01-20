package controllers

type TopicInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Stories     string `json:"stories" binding:"required"`
}

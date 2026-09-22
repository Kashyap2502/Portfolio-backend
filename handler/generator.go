package handler

import (
	"Portfolio/models"
	"Portfolio/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeneratorHandler struct {
	Repo *repository.Repository
}

func (h *GeneratorHandler) About(c *gin.Context) {
	description, err := h.Repo.GetAbout(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := models.AboutResp{ID: 1, Description: description, Profile: "Generic Profile"}
	c.JSON(http.StatusOK, resp)
}

func (h *GeneratorHandler) Projects(c *gin.Context) {
	projects, err := h.Repo.GetProjects(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := models.ProjectsResp{Projects: projects}
	c.JSON(http.StatusOK, resp)
}

func (h *GeneratorHandler) Education(c *gin.Context) {
	education, err := h.Repo.GetEducation(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := models.EducationResp{Education: education}
	c.JSON(http.StatusOK, resp)
}

func (h *GeneratorHandler) Experience(c *gin.Context) {
	experience, err := h.Repo.GetExperience(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := models.ExperienceResp{Experience: experience}
	c.JSON(http.StatusOK, resp)
}

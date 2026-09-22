package repository

import (
	"Portfolio/models"
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

func (r *Repository) GetAbout(ctx context.Context) (string, error) {
	var description string
	err := r.DB.QueryRow(ctx, "SELECT description FROM about WHERE id = $1", 10).Scan(&description)
	if err != nil {
		return "", err
	}
	return description, nil
}

func (r *Repository) GetProjects(ctx context.Context) ([]models.Project, error) {
	rows, err := r.DB.Query(ctx, "SELECT title, description, techstack, url FROM projects")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []models.Project
	for rows.Next() {
		var project models.Project
		var title, description, techstack, url sql.NullString
		err := rows.Scan(&title, &description, &techstack, &url)
		if err != nil {
			return nil, err
		}
		if title.Valid {
			project.Title = title.String
		}
		if description.Valid {
			project.Description = description.String
		}
		if techstack.Valid {
			project.Techstack = techstack.String
		}
		if url.Valid {
			project.Url = url.String
		}
		projects = append(projects, project)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *Repository) GetEducation(ctx context.Context) ([]models.Education, error) {
	rows, err := r.DB.Query(ctx, "SELECT institution, degree, year FROM education")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var education []models.Education
	for rows.Next() {
		var edu models.Education
		err := rows.Scan(&edu.Institution, &edu.Degree, &edu.Year)
		if err != nil {
			return nil, err
		}
		education = append(education, edu)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return education, nil
}

func (r *Repository) GetExperience(ctx context.Context) ([]models.Experience, error) {
	rows, err := r.DB.Query(ctx, "SELECT Organization, title, duration, description FROM experience")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var experience []models.Experience
	for rows.Next() {
		var exp models.Experience
		err := rows.Scan(&exp.Company, &exp.Role, &exp.Duration, &exp.Description)
		if err != nil {
			return nil, err
		}
		experience = append(experience, exp)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return experience, nil
}

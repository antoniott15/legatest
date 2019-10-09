package main

import (
	"context"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/rs/xid"
)

func (repo *RepositoryDB) CreateNewProject(project *Project) error {
	if project.ID == "" {
		project.ID = xid.New().String()
	}

	project.CreationDate = time.Now().String()

	_, err := repo.Project.InsertOne(context.Background(), project)
	if err != nil {
		return err
	}

	return nil

}

func (repo *RepositoryDB) CreateNewInvestor(invest *Investor) error {

	if invest.ID == "" {
		invest.ID = xid.New().String()
	}

	invest.Creationdate = time.Now().String()

	_, err := repo.Project.InsertOne(context.Background(), invest)
	if err != nil {
		return err
	}

	return nil

}

func (repo *RepositoryDB) GetInvestorByID(id string) (*Investor, error) {
	investor := new(Investor)
	if err := repo.Investor.FindOne(context.Background(), bson.M{"id": id}).Decode(investor); err != nil {
		return nil, err
	}
	return investor, nil
}

func (repo *RepositoryDB) GetProjectByID(id string) (*Project, error) {
	log.Println("operations: ", id)
	project := new(Project)
	if err := repo.Project.FindOne(context.Background(), bson.M{"id": id}).Decode(project); err != nil {
		return nil, err
	}
	return project, nil
}

func (repo *RepositoryDB) UpdateInvestorByID(investor *Investor) error {
	_, err := repo.Project.UpdateOne(context.Background(), bson.M{"id": investor.ID}, bson.M{
		"name":                  investor.Name,
		"last_name":             investor.LastName,
		"gender":                investor.Gender,
		"creation_date":         investor.Creationdate,
		"email":                 investor.Email,
		"phone":                 investor.Phone,
		"total_invest":          investor.TotalInvest,
		"country":               investor.Country,
		"number_of_transaction": investor.NumberTransaction,
		"birthday":              investor.Birthday,
	})
	return err
}

func (repo *RepositoryDB) UpdateProjectByID(project *Project) error {
	_, err := repo.Project.UpdateOne(context.Background(), bson.M{"id": project.ID}, bson.M{
		"_location":      project.Location,
		"_name":          project.Name,
		"_status":        project.Status,
		"_total_actions": project.TotalActions,
	})
	return err
}

func (repo *RepositoryDB) Projects() ([]*Project, error) {
	c, _ := repo.Project.Find(context.Background(), bson.M{})
	allProjects := make([]*Project, 0)

	for c.Next(context.Background()) {
		projects := new(Project)
		if err := c.Decode(projects); err != nil {
			return nil, err
		}
		allProjects = append(allProjects, projects)
	}
	return allProjects, nil
}

func (repo *RepositoryDB) Investors() ([]*Investor, error) {

	c, _ := repo.Project.Find(context.Background(), bson.M{})
	allInvestors := make([]*Investor, 0)

	for c.Next(context.Background()) {
		investors := new(Investor)
		if err := c.Decode(investors); err != nil {
			return nil, err
		}
		allInvestors = append(allInvestors, investors)
	}
	return allInvestors, nil
}

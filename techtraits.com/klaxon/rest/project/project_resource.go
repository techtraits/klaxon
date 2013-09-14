package project

import (
	"techtraits.com/klaxon/router"
	"techtraits.com/log"
)

func init() {
	router.Register("/rest/project", router.GET, nil, nil, getProjects)
	router.Register("/rest/project/{project_id}", router.GET, nil, nil, getProjects)
	router.Register("/rest/project", router.POST, []string{"application/json"}, nil, postProject)
}

//Get all projects
func getProjects(request router.Request) {

	log.Info("Get Project")
}

//Create/Update a project
func postProject(request router.Request) {

	log.Info("Post Project")
}

//Get a specific project
func getProject(request router.Request) {

	log.Info("Get Project")
}

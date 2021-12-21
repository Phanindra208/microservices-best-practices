package core

import (
	"net/http"

	"github.com/iAmPlus/microservice/mongo"
)

// SkillContext is the context for the skill. Useful for sharing common components with Handlers
type SkillContext struct {
	Config Config

	Mongo mongo.Manager
}

// SkillParams ... contains the building blocks for a skill
type SkillParams struct {
	SkillContext  SkillContext
	RoutesHandler http.Handler
}

//Skill ... The skill interface
type Skill interface {
	Initialise(SkillParams) // Should be called first before any other methods are called

	Context() SkillContext
}

//Config ... Interface to the skill config Manager
type Config interface {
	Get(string) interface{}
	SetDefault(string, interface{})
}

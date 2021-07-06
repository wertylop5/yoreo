package util

// Primary data types

type User struct {
	id   int
	name string
}

type Routine struct {
	id                 int
	name               string
	ownerId            int
	collaboratorListId int
	data               string
}

type PermissionLevel struct {
	id          int
	description string
}

// Secondary data types

type CollaboratorList struct {
	id        int
	routineId int
}

type Collaborator struct {
	userId             int
	collaboratorListId int
	permissionLevelId  int
}

// Aggregate data types

type UserCollaborator struct {
	User
	permissionLevelId int
}

type RoutineWithCollaborators struct {
	Routine
	collaborators []UserCollaborator
}

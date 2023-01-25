package graph

import "github.com/Ndav07/GraphQL/internal/database"

type Resolver struct {
	CategoryDB *database.Category
	CourseDB *database.Course
}

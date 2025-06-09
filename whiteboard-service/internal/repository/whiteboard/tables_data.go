package whiteboard

const (
	// TableName specifies the name of the database table for whiteboard records.
	TableName = "whiteboard"

	// IDColumn defines the column name for the whiteboard ID in the database.
	IDColumn = "id"

	// IDColumn defines the column name for the whiteboard title in the database.
	TitleColumn = "title"

	// IDColumn defines the column name for the whiteboard title in the database.
	DescriptionColumn = "description"

	// ReturningID is a SQL clause used to return the ID of a newly inserted record.
	ReturningID = "RETURNING id"
)

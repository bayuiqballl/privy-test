package query

const (
	InsertDataCake = `
	INSERT INTO
		cake(
			title,
			description,
			rating,
			image,
			created_at,
			updated_at		
		)
		VALUES(
			?, ?, ?, ?, now(), now()
		)
	`

	GetAllDataCakes = `
	SELECT
		id,
		title,
		description,
		rating,
		image,
		created_at,
		updated_at	
	FROM cake WHERE deleted_at is null 
	`

	GetCakeDetailByID = `
	SELECT
		id,
		title,
		description,
		rating,
		image,
		created_at,
		updated_at	
	FROM cake 
	WHERE deleted_at is null 
	AND id = ?
	`

	UpdateCakeByID = `
	UPDATE cake SET
			title = ?,
			description = ?,
			rating = ?,
			image = ?,
			updated_at	= now()	
	WHERE id = ?
	`

	DeleteCakeByID = `
	DELETE FROM cake WHERE id = ?
	`

	SoftDeleteCakeByID = `
	UPDATE cake SET deleted_at = now()
	WHERE id = ?
	`
)

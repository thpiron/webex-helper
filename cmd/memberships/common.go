package memberships

var (
	roomID                   string
	personID                 string
	personEmail              string
	isModerator              bool
	membershipsFields        []string
	defaultMembershipsFields = []string{"roomId", "personEmail", "IsModerator", "roomType"}
)

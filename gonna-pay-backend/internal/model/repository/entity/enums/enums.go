package enums

type ContactCategory string

const (
	ContactCategoryFamily ContactCategory = "Family"
	ContactCategoryFriend ContactCategory = "Friend"
	ContactCategoryWork   ContactCategory = "Work"
)

type CostCategory string

const (
	CostCategoryDinner        CostCategory = "Dinner"
	CostCategoryLunch         CostCategory = "Lunch"
	CostCategoryEntertainment CostCategory = "Entertainment"
	CostCategoryTravel        CostCategory = "Travel"
	CostCategoryOthers        CostCategory = "Others"
)

type GroupCategory string

const (
	GroupCategoryDinner        GroupCategory = "Dinner"
	GroupCategoryLunch         GroupCategory = "Lunch"
	GroupCategoryEntertainment GroupCategory = "Entertainment"
	GroupCategoryTravel        GroupCategory = "Travel"
	GroupCategoryOthers        GroupCategory = "Others"
)

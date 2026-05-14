package models

type Company struct {
	ID uint `gorm:"primaryKey"`
	Name string 
	Location string //example: [Amman]
	LocationURL string //Google map link
	Size string // startup, mid-size, locally known[large-l], regional known[large-r], international [large-i]  
	Logo string //card logo path file
}

/*var Companies = []Company{
	{
		Name : "Amazon",
		Location: "Amman",
		Size: "International",
		Logo: "",
	},
	{
		Name : "Google",
		Location: "Amman",
		Size: "International",
		Logo: "",
	},
	{
		Name : "Madar",
		Location: "Irbid",
		Size: "Mid-Size",
		Logo: "",
	},
}*/
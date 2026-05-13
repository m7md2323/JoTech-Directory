package models

type Company struct {
	ID uint `gorm:"primaryKey"`
	Name string 
	Location string 
	LocationURL string
	Size string 
	Logo string //card logo path file
}

var Companies = []Company{
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
}
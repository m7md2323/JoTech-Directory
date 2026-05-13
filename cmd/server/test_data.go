package main


type Company struct {
	Name string
	Location string
	Size string
	cardPhoto string //card logo path file
}

var companies = []Company{
	{
		Name : "Amazon",
		Location: "Amman",
		Size: "International",
		cardPhoto: "",
	},
	{
		Name : "Google",
		Location: "Amman",
		Size: "International",
		cardPhoto: "",
	},
	{
		Name : "Madar",
		Location: "Irbid",
		Size: "Mid-Size",
		cardPhoto: "",
	},
}
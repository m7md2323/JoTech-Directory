package services


import (
	"bufio"
    "context"
    "fmt"
    "log"
    "google.golang.org/genai"
	"os"
	"net/http"
	"io"
	"encoding/json"
	"time"
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/internal/models"

)

type ExtractionResult struct {
    Name        string   `json:"name"`
    Size        string   `json:"size"`
    Description string   `json:"description"`
	ContactInfo string 	 `json:"contactInfo"`
	Type		string   `json:"type"`
    Locations 	[]models.Location `json:"locations"`
    Tags []models.Tag 	 `json:"tags"`
    Links []models.Link  `json:"links"`
}

type GeminiRequest struct {
    Contents []struct {
        Parts []struct {
            Text string `json:"text"`
        } `json:"parts"`
    } `json:"contents"`
    GenerationConfig struct {
        ResponseMimeType string `json:"responseMimeType"`
    } `json:"generationConfig"`
}

func getURLs() []string{
	var result []string 

	file, err := os.Open("E://Jordan-Tech-Companies//TestURLs.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result=append(result, line)
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	return result
}

func callGeminiAPI(client *genai.Client, markdownContent string) (*ExtractionResult, error) {
    ctx := context.Background()
    
    prompt := `You are an expert data extraction assistant. Your task is to extract company profile information from the following Markdown text and return it as a structured JSON object.

	Follow these strict rules:
	1. Return ONLY valid JSON. Do not include markdown formatting blocks (like ` + "```json" + `).
	2. If a specific piece of information is not found in the text, return an empty string "" for strings, or an empty array [] for lists.
	3. The "size" field MUST be exactly one of these strings: "Startup", "Small", "Medium", "Large", "Enterprise", "Multi-National". For Size check linkedIn Employee Count, and based on that decied.
	4. The "type" field MUST be exactly one of these strings: "Private Company", "Governmental Company", "Non-Profit". 
	5. For "locations", extract the city, The Google Map URL starts with (https://www.google.com/maps/), Mostly they are in the contact page.
	6. For "tags", extract the industry or tech stack : {
		"Software Development", "Web Development", "Mobile Development", "AI", 
		"Machine Learning", "Deep Learning", "Data Science", "Cloud Computing", 
		"Blockchain", "Internet Of Things", "Cyber Security", "Robotics", 
		"Embedded Systems", "Augmented Reality", "Virtual Reality", "Game Development", 
		"ECommerce", "Health Tech", "FinTech", "EdTech", "AdTech", "MarTech", "Drones", 
		"Internship Programs", "Data Analysis", "Networking", "Freelance", 
		"Computer Vision", "Human Resources",
	}.
	7. For "links", extract social media or website links, Mostly they are in the contact page.

	You MUST use this exact JSON schema:
	{
	"name": "",
	"size": "",
	"description": "",
	"contactInfo": "",
	"type": "",
	"locations": [
		{
		"City": "",
		"URL": ""
		}
	],
	"tags": [
		{
		"Name": ""
		}
	],
	"links": [
		{
		"Platform": "",
		"URL": ""
		}
	]
	}

	Markdown text to process:
	` + markdownContent

    response, err := client.Models.GenerateContent(ctx, "gemini-3.1-flash-lite", 
        genai.Text(prompt),
        &genai.GenerateContentConfig{
            ResponseMIMEType: "application/json",
        },
    )
    if err != nil {
        return nil, err
    }

    var result ExtractionResult
    err = json.Unmarshal([]byte(response.Text()), &result)
    return &result, err
}



func AIWebScraper(){
    ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }


	URLs := getURLs()
    for _, url := range URLs {

        markdownText := jinaURLHandler(url)
        

        result, err := callGeminiAPI(client, string(markdownText))
        if err != nil {
            fmt.Printf("Error processing %s: %v\n", url, err)
            continue
        }


        saveToDatabase(*result)
        

        time.Sleep(4 * time.Second)
    }
}

func jinaURLHandler(Url string) []byte{

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://r.jina.ai/"+Url, nil)
	jinaAPIKey:=os.Getenv("JINA_API_KEY")
	req.Header.Add("Authorization", "Bearer "+jinaAPIKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error Handling jina req and resp: %s", err)
	}
	defer resp.Body.Close() 

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error while Reading the body into a byte slice inside Jina Handler: %s", err)
	}

	fmt.Println(string(bodyBytes))
	return bodyBytes

}

func saveToDatabase(data ExtractionResult) error {
	/*
	type Company struct {

	ID uint `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
	Size Size //startup, small, medium, large, enterprise, multi-national
	Logo string //card logo path file
	ContactInfo string 
	EmployeeCount string //example: "1-10", "11-50", "51-200", "201-500"
	ProfileImage string //profile page image path
	Description string `gorm:"type:text"`
	Type Type //Private Company, Governmental Company, Non-Profit

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Locations []Location `gorm:"foreignKey:CompanyID"` //example: [{city:"Amman", url:"example.com"}] the first location is the base one.
	Tags []Tag `gorm:"foreignKey:CompanyID"` //example: [AI, Fintech, E-commerce, Healthtech, Gaming, Edtech, etc.], and this will be used for searching and filtering.
	Links []Link `gorm:"serializer:json"` //example: [{name:"Facebook", url:"https://facebook.com/"}]

}	*/
    company := models.Company{
        Name:        	data.Name,
		Size: 		 	models.Size(data.Size),
		Logo:			" ",
		ContactInfo: 	data.ContactInfo,
		EmployeeCount: 	" ",
		ProfileImage: 	"",
		Description: 	data.Description,
		Type:		 	models.Type(data.Type),
		Locations: 		data.Locations,
		Tags:			data.Tags,
		Links:			data.Links,	
    }

	result:=database.DB.Create(&company)
    
	return result.Error
	
}
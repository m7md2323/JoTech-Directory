package services

import (
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/internal/models"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/genai"
)

type ExtractionResult struct {
	Name        string            `json:"name"`
	Size        string            `json:"size"`
	Description string            `json:"description"`
	ContactInfo string            `json:"contactInfo"`
	Type        string            `json:"type"`
	Locations   []models.Location `json:"locations"`
	Tags        []models.Tag      `json:"tags"`
	Links       []models.Link     `json:"links"`
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

func getURLs() []string {
	var result []string

	file, err := os.Open("E://Jordan-Tech-Companies//TestURLs.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	return result
}

func callGeminiAPI(client *genai.Client, markdownContent string) (*ExtractionResult, error) {
	ctx := context.Background()

	prompt := `You are an expert data extraction assistant. Your task is to extract company profile information from the provided Markdown text and return it as a structured JSON object.

		Follow these strict rules:
		1. Output Format: Return ONLY a valid, minified JSON object. Do NOT wrap the response in markdown code blocks (like ` + "```json" + `). Do not include any explanations.
		2. Missing Data: If a specific piece of information is not found in the text, return an empty string "" for strings, or an empty array [] for lists. Do not invent data.
		3. Company Size: You must map the LinkedIn Employee Count (or general employee count) to exactly ONE of the following strings based on these specific ranges:
		- "Startup" (1-10 employees)
		- "Small" (11-50 employees)
		- "Medium" (51-200 employees)
		- "Large" (201-1000 employees)
		- "Enterprise" (1001-5000 employees)
		- "Multi-National" (5000+ employees)
		4. Company Type: MUST be exactly one of: "Private Company", "Governmental Company", or "Non-Profit". Default to "Private Company" if unclear.
		5. Locations: Always return an empty array: []
		6. Tags: Extract the industry, tech stack, or focus areas. You are STRICTLY limited to these exact strings:
		["Software Development", "Web Development", "Mobile Development", "AI", "Machine Learning", "Deep Learning", "Data Science", "Cloud Computing", 
		"Blockchain", "Internet Of Things", "Cyber Security", "Robotics", "Embedded Systems", "Augmented Reality", 
		"Virtual Reality", "Game Development", "ECommerce", "Health Tech", "FinTech", "EdTech", "AdTech", 
		"MarTech", "Drones", "Internship Programs", "Data Analysis", "Network Engineering", "Freelance", "Computer Vision", 
		"Human Resources", "Quality Assurance" , "Open-Source" , "Fiber Infrastructure"]
		7. Links: Extract the company's main website and their LinkedIn URL if present. 
		8. Description: Write a concise, professional summary of the company's main tasks and offerings. Strictly constrain this to 70-100 words.

		You MUST adhere to this exact JSON schema (note the exact lowercase/camelCase key names):
		{
		"name": "",
		"size": "",
		"description": "",
		"contactInfo": "",
		"type": "",
		"locations": [],
		"tags": [
			{
			"name": ""
			}
		],
		"links": [
			{
			"platform": "",
			"url": ""
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

func AIWebScraper() {
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

func jinaURLHandler(Url string) []byte {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://r.jina.ai/"+Url, nil)
	jinaAPIKey := os.Getenv("JINA_API_KEY")
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

	locations := []models.Location{}
	tags := []models.Tag{}
	links := []models.Link{}

	for _, loc := range data.Locations {
		locations = append(locations, models.Location{
			City: loc.City,
			URL:  loc.URL,
		})
	}

	for _, tag := range data.Tags {
		tags = append(tags, models.Tag{
			Name: tag.Name,
		})
	}

	for _, link := range data.Links {
		links = append(links, models.Link{
			Platform: link.Platform,
			URL:      link.URL,
		})
	}

	company := models.Company{
		Name:          data.Name,
		Size:          models.Size(data.Size),
		Logo:          " ",
		ContactInfo:   data.ContactInfo,
		EmployeeCount: " ",
		Description:   data.Description,
		Type:          models.Type(data.Type),
		Locations:     locations,
		Tags:          tags,
		Links:         links,
	}

	result := database.DB.Create(&company)
	if result.Error != nil {
		log.Printf("Failed to save company %s to database: %v\n", data.Name, result.Error)
	} else {
		log.Printf("Successfully saved %s to database.\n", data.Name)
	}
	return result.Error

}

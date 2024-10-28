package main

import (
	"log"
	"os"

	_ "example.com/kaisheng"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	_ "github.com/go-sql-driver/mysql"
)

// const openaiURL = "https://api.openai.com/v1/chat/completions"
// const default_provider_name = "chatgpt"
// const default_model_name = "gpt-4o-mini"
// const default_prompt_message = "This is Data Retrieve from GoogleVisionOCR. PLs help me to find {type: nric | passport | driving-license | Others , number:,name:,country: {code: ,name: }, return result as **JSON (JavaScript Object Notation)** and must in stringify Json format make it machine readable message. dont use ```json. The number should not mixed with alpha.No explaination or further questions needed !!!"

// var db *sql.DB

// type RequestBody struct {
// 	Base64Image string `json:"base64image"`
// }
// type ResponseBody struct {
// 	AIResponse string `json:"aiResponse"`
// }

// func ProcessChatgptAI(formatedText string, modelname string) (string, error) {
// 	ChatgptKey := os.Getenv("CHATGPT_KEY")
// 	promptMessage, err := GetPromptMessage()
// 	if err != nil {
// 		return " ", fmt.Errorf("error retrieving prompt message: %v", err)

// 	}
// 	prompt := fmt.Sprintf("%s, %s", promptMessage, formatedText)

// 	// Create the request body (adjust based on Ollama's API requirements)
// 	requestBody := map[string]interface{}{
// 		"model": modelname,
// 		"messages": []map[string]string{
// 			{
// 				"role":    "user",
// 				"content": prompt,
// 			},
// 		},
// 	}

// 	// Convert the request body to JSON
// 	jsonBody, err := json.Marshal(requestBody)
// 	if err != nil {
// 		return " ", fmt.Errorf("error marshaling JSON: %v", err)

// 	}

// 	// Create the HTTP request
// 	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(jsonBody))
// 	if err != nil {
// 		return " ", fmt.Errorf("error creating request: %v", err)
// 	}

// 	// Set the headers
// 	req.Header.Set("Authorization", "Bearer "+ChatgptKey)
// 	req.Header.Set("Content-Type", "application/json")

// 	// Send the request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return " ", fmt.Errorf("error sending request: %v", err)

// 	}
// 	defer resp.Body.Close()

// 	// Check the response status
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
// 		body, _ := io.ReadAll(resp.Body)
// 		return " ", fmt.Errorf("response body: %v", string(body))

// 	}

// 	// Read the response stream and accumulate the content
// 	var fullResponse string
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return " ", fmt.Errorf("error reading response body: %v", err)

// 	}

// 	// Parse the JSON response
// 	var response map[string]interface{}
// 	err = json.Unmarshal(body, &response)
// 	if err != nil {
// 		return " ", fmt.Errorf("error unmarshalling response body: %v", err)

// 	}

// 	// Extract and print the assistant's response content
// 	if choices, ok := response["choices"].([]interface{}); ok && len(choices) > 0 {
// 		choice := choices[0].(map[string]interface{})
// 		if message, ok := choice["message"].(map[string]interface{}); ok {
// 			if content, ok := message["content"].(string); ok {
// 				fullResponse = content // Assign the content directly
// 				return fullResponse, nil
// 			}
// 		}
// 	}
// 	return "", fmt.Errorf("no valid response from ChatGPT")
// }

// func ProcessGemmaAI(formatedText string, modelname string) (string, error) {
// 	host := os.Getenv("OLLAMA_HOST")
// 	api := os.Getenv("OLLAMA_API")
// 	endpoint := os.Getenv("OLLAMA_ENDPOINT")
// 	OllamaKey := os.Getenv("OLLAMA_KEY")
// 	ollamaURL := fmt.Sprintf("http://%s/%s/%s", host, api, endpoint)
// 	promptMessage, err := GetPromptMessage()
// 	if err != nil {
// 		return " ", fmt.Errorf("error retrieving prompt message: %v", err)
// 	}
// 	prompt := fmt.Sprintf("%s, %s", promptMessage, formatedText)
// 	// Create the request body (adjust based on Ollama's API requirements)
// 	requestBody := map[string]interface{}{
// 		"model": modelname,
// 		"messages": []map[string]string{
// 			{
// 				"role":    "user",
// 				"content": prompt,
// 			},
// 		},
// 	}

// 	// Convert the request body to JSON
// 	jsonBody, err := json.Marshal(requestBody)
// 	if err != nil {
// 		return " ", fmt.Errorf("error marshaling JSON: %v", err)

// 	}

// 	// Create the HTTP request
// 	req, err := http.NewRequest("POST", ollamaURL, bytes.NewBuffer(jsonBody))
// 	if err != nil {
// 		return " ", fmt.Errorf("error creating request: %v", err)

// 	}

// 	// Set the headers
// 	req.Header.Set("Authorization", "Bearer "+OllamaKey)
// 	req.Header.Set("Content-Type", "application/json")

// 	// Send the request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return " ", fmt.Errorf("error sending request: %v", err)

// 	}
// 	defer resp.Body.Close()

// 	// Check the response status
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
// 		body, _ := io.ReadAll(resp.Body)
// 		return " ", fmt.Errorf("response body: %v", string(body))

// 	}

// 	// Read the response stream and accumulate the content
// 	var fullResponse string
// 	decoder := json.NewDecoder(resp.Body)
// 	for {
// 		var chunk map[string]interface{}
// 		if err := decoder.Decode(&chunk); err == io.EOF {
// 			break // End of stream
// 		} else if err != nil {
// 			return " ", fmt.Errorf("error decoding response: %v", err)
// 		}

// 		if message, ok := chunk["message"]; ok {
// 			if content, ok := message.(map[string]interface{})["content"]; ok {
// 				fullResponse += content.(string)
// 			}
// 		}

// 		if done, ok := chunk["done"]; ok && done.(bool) {
// 			break
// 		}
// 	}
// 	return fullResponse, nil
// }

// func GetOCRText(imageData []byte, googleCred option.ClientOption) (string, error) {
// 	ctx := context.Background()
// 	// Create a new Vision API client
// 	client, err := vision.NewImageAnnotatorClient(ctx, googleCred)
// 	if err != nil {
// 		return "", fmt.Errorf("vision.NewImageAnnotatorClient: %v", err)
// 	}
// 	defer client.Close()

// 	// Create an image object
// 	image, err := vision.NewImageFromReader(bytes.NewReader(imageData))

// 	if err != nil {
// 		return "", fmt.Errorf("vision.NewImageFromReader: %v", err)
// 	}

// 	// Perform OCR (text detection)
// 	annotations, err := client.DetectTexts(ctx, image, nil, 1)
// 	if err != nil {
// 		return "", fmt.Errorf("DetectTexts: %v", err)
// 	}

// 	// Check if text was detected
// 	if len(annotations) == 0 {
// 		return "", fmt.Errorf("no text found in image")
// 	}

// 	// Return the detected text
// 	return annotations[0].Description, nil
// }
// func InitSQL() error {
// 	user := os.Getenv("MY_SQL_USER")
// 	password := os.Getenv("MY_SQL_PASSWORD")
// 	host := os.Getenv("MY_SQL_HOST")
// 	dbName := os.Getenv("MY_SQL_DB")
// 	tcp := os.Getenv("MY_SQL_TCP")
// 	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s", user, password, tcp, host, dbName)
// 	var err error
// 	db, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		return fmt.Errorf("error connecting to MySQL: %v", err)
// 	}

// 	// Test the connection
// 	err = db.Ping()
// 	if err != nil {
// 		return fmt.Errorf("error pinging MySQL: %v", err)
// 	}

// 	fmt.Println("Connected to MySQL!")
// 	return nil
// }

// func GetAIConfig() (string, string, error) {

// 	// Query the AI configuration
// 	var providerName, modelName string
// 	err := db.QueryRow("SELECT provider_name, model_name FROM system_config WHERE id = 1").Scan(&providerName, &modelName)
// 	if err != nil {
// 		log.Println("Setting to Default Provider and Model")
// 		return default_provider_name, default_model_name, nil
// 	}
// 	log.Printf("Retrieved AI config: provider=%s, model=%s\n", providerName, modelName)
// 	return providerName, modelName, nil
// }

// func GetPromptMessage() (string, error) {

// 	// Query the AI configuration
// 	var promptMessage string
// 	err := db.QueryRow("SELECT ocr_prompt_message FROM ocr_prompt_message;").Scan(&promptMessage)
// 	if err != nil {
// 		log.Println("Attempting Default Prompt Message...")
// 		return default_prompt_message, nil
// 	}

// 	return promptMessage, nil
// }
// func ProcessAI(formatedText string) (string, error) {
// 	providerName, modelName, err := GetAIConfig() // Fetch provider and model from MySQL
// 	if err != nil {
// 		log.Printf("Error retrieving AI configuration: %v\n", err)
// 		return "", fmt.Errorf("error retrieving AI configuration: %v", err)
// 	}
// 	var aiResponse string
// 	if providerName == "chatgpt" {
// 		aiResponse, err = ProcessChatgptAI(formatedText, modelName) // Call ChatGPT API function
// 		if err != nil {
// 			log.Printf("Error processing chatgpt AI: %v\n", err)
// 			return "", fmt.Errorf("error processing chatgpt AI: %v", err)
// 		}
// 	} else if providerName == "gemma" {
// 		aiResponse, err = ProcessGemmaAI(formatedText, modelName) // Call Gemma API function
// 		if err != nil {
// 			log.Printf("Error processing gemma AI: %v\n", err)
// 			return "", fmt.Errorf("error processing Gemma AI: %v", err)
// 		}
// 	} else {
// 		log.Printf("Unknown provider: %s\n", providerName)
// 		aiResponse, err = ProcessGemmaAI(formatedText, modelName)
// 		if err != nil {
// 			log.Printf("Error processing default AI: %v\n", err)
// 			return "", fmt.Errorf("error processing default AI: %v", err)
// 		}
// 	}
// 	return aiResponse, nil
// }

// func main() {
// 	env := os.Getenv("ENVIRONMENT")
// 	var envFile string
// 	if env == "cloud" {
// 		envFile = "Capp.env"
// 	} else {
// 		envFile = "app.env"
// 	}

// 	err := godotenv.Load(envFile)
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}
// 	err = InitSQL()
// 	if err != nil {
// 		log.Fatalf("Failed to connect to MySQL: %v", err)
// 	}
// 	defer db.Close()

// 	port := os.Getenv("LOCAL_SERVER_PORT")
// 	if port == "" {
// 		port = "5000"
// 	}

//		if env == "cloud" {
//			// For Cloud Functions, use funcframework to start the server
//			if err := funcframework.Start(port); err != nil {
//				log.Fatalf("funcframework.Start: %v\n", err)
//			}
//		}
//		// else {
//		// 	// For local development, use http.ListenAndServe
//		// 	http.HandleFunc("/upload", PostImage)
//		// 	log.Fatal(http.ListenAndServe(":"+port, nil))
//		// }
//	}
func main() {
	env := os.Getenv("ENVIRONMENT")
	port := os.Getenv("LOCAL_SERVER_PORT")
	if port == "" {
		port = "5000"
	}

	if env == "cloud" {
		// For Cloud Functions, use funcframework to start the server
		if err := funcframework.Start(port); err != nil {
			log.Fatalf("funcframework.Start: %v\n", err)
		}
	}
	// else {
	// 	// For local development, use http.ListenAndServe
	// 	http.HandleFunc("/upload", PostImage)
	// 	log.Fatal(http.ListenAndServe(":"+port, nil))
	// }
}
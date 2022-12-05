package shared

type SDBGeneration struct {
	Id            string `json:"id"`
	Status        string `json:"status"`
	FailureReason string `json:"failure_reason"`
	CreatedAt     string `json:"created_at"`
}

type SDBServer struct {
	Id                string `json:"id"`
	Enabled           bool   `json:"enabled"`
	Healthy           bool   `json:"healthy"`
	LastHealthCheckAt string `json:"last_health_check_at"`
	Url               string `json:"url"`
}

type SDiscordWebhookBody struct {
	Content     *string                     `json:"content"`
	Embeds      []SDiscordWebhookEmbed      `json:"embeds"`
	Attachments []SDiscordWebhookAttachment `json:"attachments"`
}

type SDiscordWebhookEmbed struct {
	Title  string                     `json:"title,omitempty"`
	Fields []SDiscordWebhookField     `json:"fields"`
	Color  int                        `json:"color"`
	Footer SDiscordWebhookEmbedFooter `json:"footer"`
}

type SDiscordWebhookEmbedFooter struct {
	Text string `json:"text"`
}

type SDiscordWebhookField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type SDiscordWebhookAttachment struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

type SInternalServerErrorResponse struct {
	Error string `json:"error"`
}

type SCogGenerateRequestBody struct {
	Input SCogGenerateRequestInput `json:"input"`
}

type SCogGenerateResponseBody struct {
	Output []string `json:"output"`
}

type SGenerateRequestBody struct {
	Prompt                string `json:"prompt"`
	NegativePrompt        string `json:"negative_prompt,omitempty"`
	Width                 int    `json:"width"`
	Height                int    `json:"height"`
	NumInferenceSteps     int    `json:"num_inference_steps"`
	GuidanceScale         int    `json:"guidance_scale"`
	ServerUrl             string `json:"server_url"`
	ModelId               string `json:"model_id"`
	SchedulerId           string `json:"scheduler_id"`
	Seed                  int    `json:"seed"`
	ShouldSubmitToGallery bool   `json:"should_submit_to_gallery"`
}

type SCogGenerateRequestInput struct {
	Prompt            string `json:"prompt"`
	NegativePrompt    string `json:"negative_prompt,omitempty"`
	Width             string `json:"width"`
	Height            string `json:"height"`
	OutputImageExt    string `json:"output_image_ext"`
	NumInferenceSteps string `json:"num_inference_steps"`
	GuidanceScale     string `json:"guidance_scale"`
	Model             string `json:"model"`
	Scheduler         string `json:"scheduler"`
	Seed              string `json:"seed"`
}
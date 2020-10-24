package mapResponse

// mpStruct struct for Here map return response
type mpStruct struct {
	Results struct {
		Next  string `json:"next"`
		Items []struct {
			Position      []float64 `json:"position"`
			Distance      int       `json:"distance"`
			Title         string    `json:"title"`
			AverageRating float64   `json:"averageRating"`
			Category      struct {
				ID     string `json:"id"`
				Title  string `json:"title"`
				Href   string `json:"href"`
				Type   string `json:"type"`
				System string `json:"system"`
			} `json:"category"`
			Icon     string        `json:"icon"`
			Vicinity string        `json:"vicinity"`
			Having   []interface{} `json:"having"`
			Type     string        `json:"type"`
			Href     string        `json:"href"`
			Tags     []struct {
				ID    string `json:"id"`
				Title string `json:"title"`
				Group string `json:"group"`
			} `json:"tags,omitempty"`
			ID           string `json:"id"`
			OpeningHours struct {
				Text       string `json:"text"`
				Label      string `json:"label"`
				IsOpen     bool   `json:"isOpen"`
				Structured []struct {
					Start      string `json:"start"`
					Duration   string `json:"duration"`
					Recurrence string `json:"recurrence"`
				} `json:"structured"`
			} `json:"openingHours,omitempty"`
			AlternativeNames []struct {
				Name     string `json:"name"`
				Language string `json:"language"`
			} `json:"alternativeNames,omitempty"`
		} `json:"items"`
	} `json:"results"`
	Search struct {
		Context struct {
			Location struct {
				Position []float64 `json:"position"`
				Address  struct {
					Text        string `json:"text"`
					House       string `json:"house"`
					Street      string `json:"street"`
					PostalCode  string `json:"postalCode"`
					District    string `json:"district"`
					City        string `json:"city"`
					County      string `json:"county"`
					StateCode   string `json:"stateCode"`
					Country     string `json:"country"`
					CountryCode string `json:"countryCode"`
				} `json:"address"`
			} `json:"location"`
			Type string `json:"type"`
			Href string `json:"href"`
		} `json:"context"`
		SupportsPanning bool   `json:"supportsPanning"`
		Ranking         string `json:"ranking"`
	} `json:"search"`
}

// New function used as constructor to return pointer of mpStruct
func New() *mpStruct {
	return &mpStruct{}
}

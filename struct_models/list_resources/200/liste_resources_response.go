package struct_models
import "encoding/json"

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Page       int64   `json:"page"`       
	PerPage    int64   `json:"per_page"`   
	Total      int64   `json:"total"`      
	TotalPages int64   `json:"total_pages"`
	Data       []Datum `json:"data"`       
	Support    Support `json:"support"`    
}

type Datum struct {
	ID        int64  `json:"id"`        
	Email     string `json:"email"`     
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"` 
	Avatar    string `json:"avatar"`    
}

type Support struct {
	URL  string `json:"url"` 
	Text string `json:"text"`
}
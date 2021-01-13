package structs

type ReqData struct {
	Client1 Client `json:"client_1"`
	Client2 Client `json:"client_2"`
	Action  int    `json:"action"`
}

type Client struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	SurName   string `json:"sur_name"`
	CountThx  int    `json:"count_thx"`
	CountStar int    `json:"count_star"`
	CanDo     bool   `json:"can_do"`
}

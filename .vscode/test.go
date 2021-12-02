type Customer struct {
	ID           string                 `json:"id"`
	ReferenceID  string                 `json:"reference_id"`
	MobileNumber string                 `json:"mobile_number,omitempty"`
	Email        string                 `json:"email,omitempty"`
	GivenNames   string                 `json:"given_names"`
	MiddleName   string                 `json:"middle_name"`
	Surname      string                 `json:"surname"`
	Description  string                 `json:"description,omitempty"`
	PhoneNumber  string                 `json:"phone_number"`
	Nationality  string                 `json:"nationality"`
	Addresses    []CustomerAddress      `json:"addresses"`
	DateOfBirth  string                 `json:"date_of_birth"`
	Metadata     map[string]interface{} `json:"metadata"`
}

if !isParamsNil {
	// reqBody, err = json.Marshal(body)
	reqBody, err = json.MarshalIndent(body, "", " ")
	if err != nil {
		return FromGoErr(err)
	}
	fmt.Println(url)
	fmt.Println(method)
	fmt.Println(string(reqBody))
}
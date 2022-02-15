package bigcommerce

type Customer struct {
	ID               int64             `json:"id"`
	Company          string            `json:"company"`
	Firstname        string            `json:"first_name"`
	Lastname         string            `json:"last_name"`
	Email            string            `json:"email"`
	Phone            string            `json:"phone"`
	FormFields       interface{}       `json:"form_fields"`
	DateCreated      string            `json:"date_created"`
	DateModified     string            `json:"date_modified"`
	StoreCredit      string            `json:"store_credit"`
	RegistrationIP   string            `json:"registration_ip_address"`
	CustomerGroup    int64             `json:"customer_group_id"`
	Notes            string            `json:"notes"`
	TaxExempt        string            `json:"tax_exempt_category"`
	ResetPassword    bool              `json:"reset_pass_on_login"`
	AcceptsMarketing bool              `json:"accepts_marketing"`
	Addresses        map[string]string `json:"addresses"`
}

package account

// AccountData contains all the account metadata (ID, organisationID...) and the account attributes
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

// AccountAttributes contains specific account attributes
type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

// AccountCreate contains all attributes we can defined when creating an account
type AccountCreate struct {
	Country                 string                `json:"country"`
	BaseCurrency            string                `json:"base_currency,omitempty"`
	BankID                  string                `json:"bank_id,omitempty"`
	BankIDCode              string                `json:"bank_id_code,omitempty"`
	AccountNumber           string                `json:"account_number,omitempty"`
	Bic                     string                `json:"bic,omitempty"`
	Iban                    string                `json:"iban,omitempty"`
	CustomerID              string                `json:"customer_id,omitempty"`
	Name                    []string              `json:"name,omitempty"`
	AlternativeNames        []string              `json:"alternative_names,omitempty"`
	AccountClassification   AccountClassification `json:"account_classification,omitempty"`
	JointAccount            bool                  `json:"joint_account"`
	SecondaryIdentification string                `json:"secondary_identification,omitempty"`
	NameMatchingStatus      NameMatchingStatus    `json:"name_matching_status,omitempty"`
}

type AccountClassification string

const (
	Personal AccountClassification = "Personal"
	Business AccountClassification = "Business"
)

type NameMatchingStatus string

const (
	Supported    NameMatchingStatus = "supported"
	Switched     NameMatchingStatus = "switched"
	OptedOut     NameMatchingStatus = "opted_out"
	NotSupported NameMatchingStatus = "not_supported"
)

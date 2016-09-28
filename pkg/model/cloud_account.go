package model

type CloudAccountList struct {
	BaseList
	Items []*CloudAccount `json:"items"`
}

type CloudAccount struct {
	BaseModel

	// has_many Kubes
	Kubes []*Kube `json:"kubes,omitempty"`

	Name string `json:"name" validate:"nonzero" gorm:"not null;unique_index"`

	Provider string `json:"provider" validate:"regexp=^(aws|digitalocean|localkube)$" gorm:"not null"`

	// NOTE this is loose map to allow for multiple clouds (eventually)
	Credentials     map[string]string `json:"credentials,omitempty" validate:"nonzero" gorm:"-" sg:"store_as_json_in=CredentialsJSON,private"`
	CredentialsJSON []byte            `json:"-" gorm:"not null"`
}

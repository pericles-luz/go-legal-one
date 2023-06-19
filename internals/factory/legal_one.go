package factory

import (
	"github.com/pericles-luz/go-legal-one/pkg/legal_one"
	"github.com/pericles-luz/go-rest/pkg/rest"
)

func NewLegalOne(file string) (*legal_one.LegalOne, error) {
	config := legal_one.NewConfig()
	err := config.Load(file)
	if err != nil {
		return nil, err
	}
	rest := rest.NewRest(config.GetConfig())
	legalOne := legal_one.NewLegalOne(rest)
	return legalOne, nil
}

package legal_one_test

import (
	"testing"
	"time"

	"github.com/pericles-luz/go-legal-one/internals/factory"
	"github.com/stretchr/testify/require"
)

func TestLegalOneAuthentication(t *testing.T) {
	t.Skip("Test only if necessary")
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	require.NoError(t, legalOne.Autenticate())
}

func TestLegalOneGetContactByCPF(t *testing.T) {
	t.Skip("Test only if necessary")
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	contact, err := legalOne.GetContactByCPF("000.000.001-91")
	require.NoError(t, err)
	require.Equal(t, "000.000.001-91", contact.Value[0].IdentificationNumber)
	t.Log(contact)
}

func TestLegalOneGetLawsuits(t *testing.T) {
	t.Skip("Test only if necessary")
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	lawsuits, err := legalOne.GetLawsuits()
	require.NoError(t, err)
	require.NotEmpty(t, lawsuits.Value)
}

func TestLegalOneIndividualRegistrate(t *testing.T) {
	t.Skip("Test only if necessary")
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	data := map[string]interface{}{
		"DE_Pessoa": "Joaquim de Teste",
		"CO_CPF":    "00000000191",
	}
	individual, err := legalOne.IndividualRegistrate(data)
	require.NoError(t, err)
	require.Equal(t, "000.000.001-91", individual.IdentificationNumber)
	require.Equal(t, "Joaquim de Teste", individual.Name)
	t.Log(individual)
}

func TestLegalOneIndividualDelete(t *testing.T) {
	t.Skip("Test only if necessary")
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	require.NoError(t, legalOne.IndividualDelete(25087))
}

func TestLegalOneGetLawsuitParticipationByContactID(t *testing.T) {
	t.Skip("Test only if necessary")
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	participations, err := legalOne.GetLawsuitParticipationByContactID(1, 3)
	require.NoError(t, err)
	require.NotEmpty(t, participations.Value)
	t.Log(participations)
}

func TestLegalOneGetAppealParticipationByContactID(t *testing.T) {
	t.Skip("Test only if necessary")
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	participations, err := legalOne.GetAppealParticipationByContactID(9184, 3)
	require.NoError(t, err)
	require.NotEmpty(t, participations.Value)
	t.Log(participations)
}

func TestLegalOneGetLawsuitByProcessNumber(t *testing.T) {
	t.Skip("Test only if necessary")
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	lawsuits, err := legalOne.GetLawsuitByProcessNumber("1004915-65.2018.4.01.3400")
	require.NoError(t, err)
	require.NotEmpty(t, lawsuits.Value)
	t.Log(lawsuits)
}

func TestLegalOneGetLawsuitByFolder(t *testing.T) {
	t.Skip("Test only if necessary")
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	lawsuits, err := legalOne.GetLawsuitByFolder("Colet-0295")
	require.NoError(t, err)
	require.NotEmpty(t, lawsuits.Value)
	t.Log(lawsuits)
}

func TestLegalLitigationsByContactID(t *testing.T) {
	t.Skip("Test only if necessary")
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	litigations, err := legalOne.GetLitigationByContactID(3175)
	require.NoError(t, err)
	require.NotEmpty(t, litigations.Value)
	t.Log(litigations)
}

func TestLegalOneParticipationRegistrate(t *testing.T) {
	t.Skip("Test only if necessary")
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	data := map[string]interface{}{
		"ID_Contato": 25089,
		"ID_Acao":    1,
	}
	participation, err := legalOne.ParticipationRegistrate(data)
	t.Log(participation)
	require.NoError(t, err)
	require.NotEmpty(t, participation.ID)
	time.Sleep(1 * time.Second)
	require.NoError(t, legalOne.ParticipationDelete(1, participation.ID))
}

func TestLegalOneGetAppealByFolder(t *testing.T) {
	t.Skip("Test only if necessary")
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	lawsuits, err := legalOne.GetAppealByFolder("Colet-0295/001")
	require.NoError(t, err)
	require.NotEmpty(t, lawsuits.Value)
	t.Log(lawsuits)
}

func TestLegalOneAppealParticipationRegistrate(t *testing.T) {
	t.Skip("Test only if necessary")
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	data := map[string]interface{}{
		"ID_Contato": 25089,
		"ID_Acao":    9184,
	}
	participation, err := legalOne.AppealParticipationRegistrate(data)
	t.Log(participation)
	require.NoError(t, err)
	require.NotEmpty(t, participation.ID)
	time.Sleep(1 * time.Second)
	require.NoError(t, legalOne.AppealParticipationDelete(9184, participation.ID))
}

func TestLegalOneGetLitigationUpdateByID(t *testing.T) {
	t.Skip("Test only if necessary")
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	litigation, err := legalOne.GetLitigationUpdateByID(1, 1)
	require.NoError(t, err)
	require.NotEmpty(t, litigation.Value)
	t.Log(litigation)
}

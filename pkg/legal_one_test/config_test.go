package legal_one_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-legal-one/pkg/legal_one"
	"github.com/stretchr/testify/require"
)

func TestLegalOneMustLoadFile(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skip when running on github")
	}
	cfg := legal_one.NewConfig()
	err := cfg.Load("legalone.dev")
	require.NoError(t, err)
	require.Equal(t, "https://api.test.com/legalone/v1/api/rest", cfg.LinkAPI)
	require.Equal(t, "https://api.test.com/legalone/oauth/oauth?grant_type=client_credentials", cfg.LinkAuth)
	require.Equal(t, "Z5wDvkq2oSG23rP7pw273kNki7h", cfg.User)
	require.Equal(t, "ofnLSaoGst", cfg.Password)
	require.NotEmpty(t, cfg.User)
}

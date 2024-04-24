package integration

import (
	"testing"

	"github.com/juanfont/headscale/integration/hsic"
	"github.com/juanfont/headscale/integration/tsic"
)

func TestWebUI(t *testing.T) {
	IntegrationSkip(t)

	scenario, err := NewScenario(dockertestMaxWait())
	assertNoErr(t, err)
	t.Cleanup(scenario.Shutdown)

	spec := map[string]int{
		"user1": len(MustTestVersions),
	}

	err = scenario.CreateHeadscaleEnv(
		spec, []tsic.Option{}, hsic.WithTestName("webui"),
		hsic.WithHostPortBindings(
			map[string][]string{"8080": {"8080"}},
		),
	)
	assertNoErrHeadscaleEnv(t, err)

	allClients, err := scenario.ListTailscaleClients()
	assertNoErrListClients(t, err)

	_ = allClients
}

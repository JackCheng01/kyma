package extractor

import (
	"encoding/json"

	"github.com/kyma-project/kyma/components/asset-store-controller-manager/pkg/apis/assetstore/v1alpha2"
	"github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
)

type Common struct{}

func (e *Common) Status(status v1alpha2.CommonAssetStatus) gqlschema.AssetStatus {
	return gqlschema.AssetStatus{
		Phase:   e.phase(status.Phase),
		Reason:  string(status.Reason),
		Message: status.Message,
	}
}

func (*Common) Parameters(parameters *runtime.RawExtension) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	if parameters == nil {
		return result, nil
	}

	err := json.Unmarshal(parameters.Raw, &result)
	if err != nil {
		return nil, errors.Wrap(err, "while unmarshalling parameters")
	}

	return result, nil
}

func (e *Common) phase(phase v1alpha2.AssetPhase) gqlschema.AssetPhaseType {
	switch phase {
	case v1alpha2.AssetReady:
		return gqlschema.AssetPhaseTypeReady
	case v1alpha2.AssetPending:
		return gqlschema.AssetPhaseTypePending
	default:
		return gqlschema.AssetPhaseTypeFailed
	}
}

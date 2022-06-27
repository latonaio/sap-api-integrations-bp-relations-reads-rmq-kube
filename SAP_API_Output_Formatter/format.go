package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-bp-relations-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToBusinessPartnerRelationshipCollection(raw []byte, l *logger.Logger) ([]BusinessPartnerRelationshipCollection, error) {
	pm := &responses.BusinessPartnerRelationshipCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to BusinessPartnerRelationshipCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	businessPartnerRelationshipCollection := make([]BusinessPartnerRelationshipCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		businessPartnerRelationshipCollection = append(businessPartnerRelationshipCollection, BusinessPartnerRelationshipCollection{
			ObjectID:                               data.ObjectID,
			RelationshipType:                       data.RelationshipType,
			RelationshipTypeText:                   data.RelationshipTypeText,
			FirstBusinessPartnerID:                 data.FirstBusinessPartnerID,
			SecondBusinessPartnerID:                data.SecondBusinessPartnerID,
			MainIndicator:                          data.MainIndicator,
			SalesOrganisationID:                    data.SalesOrganisationID,
			DistributionChannelCode:                data.DistributionChannelCode,
			DistributionChannelCodeText:            data.DistributionChannelCodeText,
			DivisionCode:                           data.DivisionCode,
			DivisionCodeText:                       data.DivisionCodeText,
			BuyingCenterAttitude:                   data.BuyingCenterAttitude,
			BuyingCenterAttitudeText:               data.BuyingCenterAttitudeText,
			BuyingCenterFrequencyOfInteraction:     data.BuyingCenterFrequencyOfInteraction,
			BuyingCenterFrequencyOfInteractionText: data.BuyingCenterFrequencyOfInteractionText,
			BuyingCenterLevelOfInfluence:           data.BuyingCenterLevelOfInfluence,
			BuyingCenterNotes:                      data.BuyingCenterNotes,
			BuyingCenterStrengthOfInfluence:        data.BuyingCenterStrengthOfInfluence,
			BuyingCenterStrengthOfInfluenceText:    data.BuyingCenterStrengthOfInfluenceText,
			CreationOn:                             data.CreationOn,
			CreatedBy:                              data.CreatedBy,
			CreatedByIdentityUUID:                  data.CreatedByIdentityUUID,
			ChangedOn:                              data.ChangedOn,
			ChangedBy:                              data.ChangedBy,
			ChangedByIdentityUUID:                  data.ChangedByIdentityUUID,
			EntityLastChangedOn:                    data.EntityLastChangedOn,
			ETag:                                   data.ETag,
		})
	}

	return businessPartnerRelationshipCollection, nil
}

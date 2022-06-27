package responses

type BusinessPartnerRelationshipCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                                     string `json:"ObjectID"`
			RelationshipType                             string `json:"RelationshipType"`
			RelationshipTypeText                         string `json:"RelationshipTypeText"`
			FirstBusinessPartnerID                       string `json:"FirstBusinessPartnerID"`
			SecondBusinessPartnerID                      string `json:"SecondBusinessPartnerID"`
			MainIndicator                                bool   `json:"MainIndicator"`
			SalesOrganisationID                          string `json:"SalesOrganisationID"`
			DistributionChannelCode                      string `json:"DistributionChannelCode"`
			DistributionChannelCodeText                  string `json:"DistributionChannelCodeText"`
			DivisionCode                                 string `json:"DivisionCode"`
			DivisionCodeText                             string `json:"DivisionCodeText"`
			BuyingCenterAttitude                         string `json:"BuyingCenterAttitude"`
			BuyingCenterAttitudeText                     string `json:"BuyingCenterAttitudeText"`
			BuyingCenterFrequencyOfInteraction           string `json:"BuyingCenterFrequencyOfInteraction"`
			BuyingCenterFrequencyOfInteractionText       string `json:"BuyingCenterFrequencyOfInteractionText"`
			BuyingCenterLevelOfInfluence                 string `json:"BuyingCenterLevelOfInfluence"`
			BuyingCenterNotes                            string `json:"BuyingCenterNotes"`
			BuyingCenterStrengthOfInfluence              string `json:"BuyingCenterStrengthOfInfluence"`
			BuyingCenterStrengthOfInfluenceText          string `json:"BuyingCenterStrengthOfInfluenceText"`
			CreationOn                                   string `json:"CreationOn"`
			CreatedBy                                    string `json:"CreatedBy"`
			CreatedByIdentityUUID                        string `json:"CreatedByIdentityUUID"`
			ChangedOn                                    string `json:"ChangedOn"`
			ChangedBy                                    string `json:"ChangedBy"`
			ChangedByIdentityUUID                        string `json:"ChangedByIdentityUUID"`
			EntityLastChangedOn                          string `json:"EntityLastChangedOn"`
			ETag                                         string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}

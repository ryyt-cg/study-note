package main

type PortfolioEnvelope struct {
	Data PortfolioData `json:"data"`
}

type PortfolioData struct {
	Value PortfolioValue `json:"value"`
}

type PortfolioValue struct {
	Balance  string `json:"balance"`
	AsOfDate string `json:"AsOfDate"`
	//List<InvestorAccount> investorAccounts;
	//InvestorProfile investorProfile;
}

type PortfolioResponse struct {
	TotalValue string `json:"totalValue"`
	AsOfDate   string `json:"AsOfDate"`
}

func (pr PortfolioResponse) FromPortfolioData(pd PortfolioData) {
	pr.TotalValue = pd.Value.Balance
	pr.AsOfDate = pd.Value.AsOfDate
}

func FromPortfolioData(pds []PortfolioData) []PortfolioResponse {
	portfolioResponses := make([]PortfolioResponse, len(pds))

	for i, v := range pds {
		portfolioResponses[i].FromPortfolioData(v)
	}

	return portfolioResponses
}

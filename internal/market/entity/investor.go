package entity

type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetPosition // slica (array sem tamanho fixo) com posições do investidor
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

// adiciona um novo asset na carteira do investidor
func (i *Investor) addAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

// atualiza a posicao de um asset na carteira do investidor
func (i *Investor) UpdateAssetPosition(assetID string, qtdShares int) {
	assetPosition := i.GetAssetPosition(assetID)
	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetID, qtdShares))
	} else {
		assetPosition.Shares += qtdShares
	}
}

// retorna a quantidade de assets na carteira do investidor
func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}
	return nil
}

// estrutura de dados de um asset position do investidor
type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}

// cria uma nova asset positon
func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID: assetID,
		Shares:  shares,
	}
}

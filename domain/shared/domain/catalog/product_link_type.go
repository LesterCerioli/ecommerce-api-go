package catalog

type ProductLinkType int

const (
	ProductLinkTypeSuper     ProductLinkType = 1
	ProductLinkTypeRelated   ProductLinkType = 2
	ProductLinkTypeCrossSell ProductLinkType = 3
	ProductLinkTypeUpSell    ProductLinkType = 4
)

func (t ProductLinkType) String() string {
	switch t {
	case ProductLinkTypeSuper:
		return "Super"
	case ProductLinkTypeRelated:
		return "Related"
	case ProductLinkTypeCrossSell:
		return "CrossSell"
	case ProductLinkTypeUpSell:
		return "UpSell"
	default:
		return "Unknown"
	}
}

func (t ProductLinkType) IsValid() bool {
	switch t {
	case ProductLinkTypeSuper, ProductLinkTypeRelated, ProductLinkTypeCrossSell, ProductLinkTypeUpSell:
		return true
	default:
		return false
	}
}

func ParseProductLinkType(v int) (ProductLinkType, bool) {
	t := ProductLinkType(v)
	return t, t.IsValid()
}

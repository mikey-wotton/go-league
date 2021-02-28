package go_league


type Region string

const(
	RegionUNKNOWN Region = ""
	RegionEUW Region = "euw"
	RegionBR Region = "br"
	RegionEUNE Region = "EUNE"
	RegionLAN Region = "lan"
	RegionNA Region = "na"
	RegionOCE Region = "oce"
	RegionRU Region = "ru"
	RegionTR Region = "tr"
	RegionJP Region = "jp"
	RegionKR Region = "kr"
)

var regions = []Region{RegionEUW, RegionBR, RegionEUNE, RegionLAN, RegionNA, RegionOCE,
	RegionRU, RegionTR, RegionJP, RegionKR}

func (r Region) Valid() bool {
	for _, region := range regions {
		if region == r {
			return true
		}
	}

	return false
}
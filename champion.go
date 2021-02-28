package go_league

type ChampionName string

const (
	ChampionNameUnknown        ChampionName = ""
	ChampionNameAatrox         ChampionName = "Aatrox"
	ChampionNameAhri           ChampionName = "Ahri"
	ChampionNameAkali          ChampionName = "Akali"
	ChampionNameAlistar        ChampionName = "Alistar"
	ChampionNameAmumu          ChampionName = "Amumu"
	ChampionNameAnivia         ChampionName = "Anivia"
	ChampionNameAnnie          ChampionName = "Annie"
	ChampionNameAphelios       ChampionName = "Aphelios"
	ChampionNameAshe           ChampionName = "Ashe"
	ChampionNameAurelion       ChampionName = "Aurelion Sol"
	ChampionNameAzir           ChampionName = "Azir"
	ChampionNameBard           ChampionName = "Bard"
	ChampionNameBlitzcrank     ChampionName = "Blitzcrank"
	ChampionNameBrand          ChampionName = "Brand"
	ChampionNameBraum          ChampionName = "Braum"
	ChampionNameCaitlyn        ChampionName = "Caitlyn"
	ChampionNameCamille        ChampionName = "Camille"
	ChampionNameCassiopeia     ChampionName = "Cassiopeia"
	ChampionNameChoGath        ChampionName = "Cho'Gath"
	ChampionNameCorki          ChampionName = "Corki"
	ChampionNameDarius         ChampionName = "Darius"
	ChampionNameDiana          ChampionName = "Diana"
	ChampionNameDrMundo        ChampionName = "Dr. Mundo"
	ChampionNameDraven         ChampionName = "Draven"
	ChampionNameEkko           ChampionName = "Ekko"
	ChampionNameElise          ChampionName = "Elise"
	ChampionNameEvelynn        ChampionName = "Evelynn"
	ChampionNameEzreal         ChampionName = "Ezreal"
	ChampionNameFiddlesticks   ChampionName = "Fiddlesticks"
	ChampionNameFiora          ChampionName = "Fiora"
	ChampionNameFizz           ChampionName = "Fizz"
	ChampionNameGalio          ChampionName = "Galio"
	ChampionNameGangplank      ChampionName = "Gangplank"
	ChampionNameGaren          ChampionName = "Garen"
	ChampionNameGnar           ChampionName = "Gnar"
	ChampionNameGragas         ChampionName = "Gragas"
	ChampionNameGraves         ChampionName = "Graves"
	ChampionNameHecarim        ChampionName = "Hecarim"
	ChampionNameHeimerdinger   ChampionName = "Heimerdinger"
	ChampionNameIllaoi         ChampionName = "Illaoi"
	ChampionNameIrelia         ChampionName = "Irelia"
	ChampionNameIvern          ChampionName = "Ivern"
	ChampionNameJanna          ChampionName = "Janna"
	ChampionNameJarvan         ChampionName = "Jarvan IV"
	ChampionNameJax            ChampionName = "Jax"
	ChampionNameJayce          ChampionName = "Jayce"
	ChampionNameJhin           ChampionName = "Jhin"
	ChampionNameJinx           ChampionName = "Jinx"
	ChampionNameKaiSa          ChampionName = "Kai'Sa"
	ChampionNameKalista        ChampionName = "Kalista"
	ChampionNameKarma          ChampionName = "Karma"
	ChampionNameKarthus        ChampionName = "Karthus"
	ChampionNameKassadin       ChampionName = "Kassadin"
	ChampionNameKatarina       ChampionName = "Katarina"
	ChampionNameKayle          ChampionName = "Kayle"
	ChampionNameKayn           ChampionName = "Kayn"
	ChampionNameKennen         ChampionName = "Kennen"
	ChampionNameKhaZix         ChampionName = "Kha'Zix"
	ChampionNameKindred        ChampionName = "Kindred"
	ChampionNameKled           ChampionName = "Kled"
	ChampionNameKogMaw         ChampionName = "Kog'Maw"
	ChampionNameLeBlanc        ChampionName = "LeBlanc"
	ChampionNameLeeSin         ChampionName = "Lee Sin"
	ChampionNameLeona          ChampionName = "Leona"
	ChampionNameLillia         ChampionName = "Lillia"
	ChampionNameLissandra      ChampionName = "Lissandra"
	ChampionNameLucian         ChampionName = "Lucian"
	ChampionNameLulu           ChampionName = "Lulu"
	ChampionNameLux            ChampionName = "Lux"
	ChampionNameMalphite       ChampionName = "Malphite"
	ChampionNameMalzahar       ChampionName = "Malzahar"
	ChampionNameMaokai         ChampionName = "Maokai"
	ChampionNameMasterYi       ChampionName = "Master Yi"
	ChampionNameMissFortune    ChampionName = "Miss Fortune"
	ChampionNameMordekaiser    ChampionName = "Mordekaiser"
	ChampionNameMorgana        ChampionName = "Morgana"
	ChampionNameNami           ChampionName = "Nami"
	ChampionNameNasus          ChampionName = "Nasus"
	ChampionNameNautilus       ChampionName = "Nautilus"
	ChampionNameNeeko          ChampionName = "Neeko"
	ChampionNameNidalee        ChampionName = "Nidalee"
	ChampionNameNocturne       ChampionName = "Nocturne"
	ChampionNameNunuAndWillump ChampionName = "Nunu and Willump"
	ChampionNameOlaf           ChampionName = "Olaf"
	ChampionNameOrianna        ChampionName = "Orianna"
	ChampionNameOrnn           ChampionName = "Ornn"
	ChampionNamePantheon       ChampionName = "Pantheon"
	ChampionNamePoppy          ChampionName = "Poppy"
	ChampionNamePyke           ChampionName = "Pyke"
	ChampionNameQiyana         ChampionName = "Qiyana"
	ChampionNameQuinn          ChampionName = "Quinn"
	ChampionNameRakan          ChampionName = "Rakan"
	ChampionNameRammus         ChampionName = "Rammus"
	ChampionNameRekSai         ChampionName = "Rek'Sai"
	ChampionNameRell           ChampionName = "Rell"
	ChampionNameRenekton       ChampionName = "Renekton"
	ChampionNameRengar         ChampionName = "Rengar"
	ChampionNameRiven          ChampionName = "Riven"
	ChampionNameRumble         ChampionName = "Rumble"
	ChampionNameRyze           ChampionName = "Ryze"
	ChampionNameSamira         ChampionName = "Samira"
	ChampionNameSejuani        ChampionName = "Sejuani"
	ChampionNameSenna          ChampionName = "Senna"
	ChampionNameSeraphine      ChampionName = "Seraphine"
	ChampionNameSett           ChampionName = "Sett"
	ChampionNameShaco          ChampionName = "Shaco"
	ChampionNameShen           ChampionName = "Shen"
	ChampionNameShyvana        ChampionName = "Shyvana"
	ChampionNameSinged         ChampionName = "Singed"
	ChampionNameSion           ChampionName = "Sion"
	ChampionNameSivir          ChampionName = "Sivir"
	ChampionNameSkarner        ChampionName = "Skarner"
	ChampionNameSona           ChampionName = "Sona"
	ChampionNameSoraka         ChampionName = "Soraka"
	ChampionNameSwain          ChampionName = "Swain"
	ChampionNameSylas          ChampionName = "Sylas"
	ChampionNameSyndra         ChampionName = "Syndra"
	ChampionNameTahmKench      ChampionName = "Tahm Kench"
	ChampionNameTaliyah        ChampionName = "Taliyah"
	ChampionNameTalon          ChampionName = "Talon"
	ChampionNameTaric          ChampionName = "Taric"
	ChampionNameTeemo          ChampionName = "Teemo"
	ChampionNameThresh         ChampionName = "Thresh"
	ChampionNameTristana       ChampionName = "Tristana"
	ChampionNameTrundle        ChampionName = "Trundle"
	ChampionNameTryndamere     ChampionName = "Tryndamere"
	ChampionNameTwistedFate    ChampionName = "Twisted Fate"
	ChampionNameTwitch         ChampionName = "Twitch"
	ChampionNameUdyr           ChampionName = "Udyr"
	ChampionNameUrgot          ChampionName = "Urgot"
	ChampionNameVarus          ChampionName = "Varus"
	ChampionNameVayne          ChampionName = "Vayne"
	ChampionNameVeigar         ChampionName = "Veigar"
	ChampionNameVelKoz         ChampionName = "Vel'Koz"
	ChampionNameVi             ChampionName = "Vi"
	ChampionNameViktor         ChampionName = "Viktor"
	ChampionNameVladimir       ChampionName = "Vladimir"
	ChampionNameVolibear       ChampionName = "Volibear"
	ChampionNameWarwick        ChampionName = "Warwick"
	ChampionNameWukong         ChampionName = "Wukong"
	ChampionNameXayah          ChampionName = "Xayah"
	ChampionNameXerath         ChampionName = "Xerath"
	ChampionNameXinZhao        ChampionName = "Xin Zhao"
	ChampionNameYasuo          ChampionName = "Yasuo"
	ChampionNameYone           ChampionName = "Yone"
	ChampionNameYorick         ChampionName = "Yorick"
	ChampionNameYuumi          ChampionName = "Yuumi"
	ChampionNameZac            ChampionName = "Zac"
	ChampionNameZed            ChampionName = "Zed"
	ChampionNameZiggs          ChampionName = "Ziggs"
	ChampionNameZilean         ChampionName = "Zilean"
	ChampionNameZoe            ChampionName = "Zoe"
	ChampionNameZyra           ChampionName = "Zyra"
)

var championNames = []ChampionName{
	ChampionNameUnknown,
	ChampionNameAatrox,
	ChampionNameAhri,
	ChampionNameAkali,
	ChampionNameAlistar,
	ChampionNameAmumu,
	ChampionNameAnivia,
	ChampionNameAnnie,
	ChampionNameAphelios,
	ChampionNameAshe,
	ChampionNameAurelion,
	ChampionNameAzir,
	ChampionNameBard,
	ChampionNameBlitzcrank,
	ChampionNameBrand,
	ChampionNameBraum,
	ChampionNameCaitlyn,
	ChampionNameCamille,
	ChampionNameCassiopeia,
	ChampionNameChoGath,
	ChampionNameCorki,
	ChampionNameDarius,
	ChampionNameDiana,
	ChampionNameDrMundo,
	ChampionNameDraven,
	ChampionNameEkko,
	ChampionNameElise,
	ChampionNameEvelynn,
	ChampionNameEzreal,
	ChampionNameFiddlesticks,
	ChampionNameFiora,
	ChampionNameFizz,
	ChampionNameGalio,
	ChampionNameGangplank,
	ChampionNameGaren,
	ChampionNameGnar,
	ChampionNameGragas,
	ChampionNameGraves,
	ChampionNameHecarim,
	ChampionNameHeimerdinger,
	ChampionNameIllaoi,
	ChampionNameIrelia,
	ChampionNameIvern,
	ChampionNameJanna,
	ChampionNameJarvan,
	ChampionNameJax,
	ChampionNameJayce,
	ChampionNameJhin,
	ChampionNameJinx,
	ChampionNameKaiSa,
	ChampionNameKalista,
	ChampionNameKarma,
	ChampionNameKarthus,
	ChampionNameKassadin,
	ChampionNameKatarina,
	ChampionNameKayle,
	ChampionNameKayn,
	ChampionNameKennen,
	ChampionNameKhaZix,
	ChampionNameKindred,
	ChampionNameKled,
	ChampionNameKogMaw,
	ChampionNameLeBlanc,
	ChampionNameLeeSin,
	ChampionNameLeona,
	ChampionNameLillia,
	ChampionNameLissandra,
	ChampionNameLucian,
	ChampionNameLulu,
	ChampionNameLux,
	ChampionNameMalphite,
	ChampionNameMalzahar,
	ChampionNameMaokai,
	ChampionNameMasterYi,
	ChampionNameMissFortune,
	ChampionNameMordekaiser,
	ChampionNameMorgana,
	ChampionNameNami,
	ChampionNameNasus,
	ChampionNameNautilus,
	ChampionNameNeeko,
	ChampionNameNidalee,
	ChampionNameNocturne,
	ChampionNameNunuAndWillump,
	ChampionNameOlaf,
	ChampionNameOrianna,
	ChampionNameOrnn,
	ChampionNamePantheon,
	ChampionNamePoppy,
	ChampionNamePyke,
	ChampionNameQiyana,
	ChampionNameQuinn,
	ChampionNameRakan,
	ChampionNameRammus,
	ChampionNameRekSai,
	ChampionNameRell,
	ChampionNameRenekton,
	ChampionNameRengar,
	ChampionNameRiven,
	ChampionNameRumble,
	ChampionNameRyze,
	ChampionNameSamira,
	ChampionNameSejuani,
	ChampionNameSenna,
	ChampionNameSeraphine,
	ChampionNameSett,
	ChampionNameShaco,
	ChampionNameShen,
	ChampionNameShyvana,
	ChampionNameSinged,
	ChampionNameSion,
	ChampionNameSivir,
	ChampionNameSkarner,
	ChampionNameSona,
	ChampionNameSoraka,
	ChampionNameSwain,
	ChampionNameSylas,
	ChampionNameSyndra,
	ChampionNameTahmKench,
	ChampionNameTaliyah,
	ChampionNameTalon,
	ChampionNameTaric,
	ChampionNameTeemo,
	ChampionNameThresh,
	ChampionNameTristana,
	ChampionNameTrundle,
	ChampionNameTryndamere,
	ChampionNameTwistedFate,
	ChampionNameTwitch,
	ChampionNameUdyr,
	ChampionNameUrgot,
	ChampionNameVarus,
	ChampionNameVayne,
	ChampionNameVeigar,
	ChampionNameVelKoz,
	ChampionNameVi,
	ChampionNameViktor,
	ChampionNameVladimir,
	ChampionNameVolibear,
	ChampionNameWarwick,
	ChampionNameWukong,
	ChampionNameXayah,
	ChampionNameXerath,
	ChampionNameXinZhao,
	ChampionNameYasuo,
	ChampionNameYone,
	ChampionNameYorick,
	ChampionNameYuumi,
	ChampionNameZac,
	ChampionNameZed,
	ChampionNameZiggs,
	ChampionNameZilean,
	ChampionNameZoe,
	ChampionNameZyra,
}

func (c ChampionName) Valid() bool {
	for _, championName := range championNames {
		if championName == c {
			return true
		}
	}

	return false
}
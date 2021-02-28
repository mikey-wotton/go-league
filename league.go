package go_league

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	allGameDataURL      = "https://127.0.0.1:2999/liveclientdata/allgamedata"
	eventsGameDataURL   = "​https://127.0.0.1:2999/liveclientdata/eventdata"
	activePlayerNameURL = "https://127.0.0.1:2999/liveclientdata/activeplayername"
)

func GetLiveData() (*GameStats, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get(allGameDataURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	gd := GameStats{}
	err = json.Unmarshal(bytes, &gd)
	if err != nil {
		return nil, err
	}

	if len(gd.AllPlayers) > 0 {
		log.Println(gd.AllPlayers)
	}

	if err := gd.SetEvents(); err != nil {
		return nil, err
	}

	return &gd, nil
}

func GetLiveEvents() (*Events, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get(eventsGameDataURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	events := Events{}
	err = json.Unmarshal(bytes, &events)
	if err != nil {
		return nil, err
	}

	return &events, nil
}

func GetActivePlayerName() (string, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get(activePlayerNameURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

//GetDemoData returns stale data but increments game time every time it is called by 1 minute.
func GetDemoData() (*GameStats, error) {
	data := []byte(demoData)

	var stats *GameStats

	err := json.Unmarshal(data, &stats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

const demoData = `{
	"activePlayer": {
		"abilities": {
			"E": {
				"abilityLevel": 5,
				"displayName": "Lucent Singularity",
				"id": "LuxLightStrikeKugel",
				"rawDescription": "GeneratedTip_Spell_LuxLightStrikeKugel_Description",
				"rawDisplayName": "GeneratedTip_Spell_LuxLightStrikeKugel_DisplayName"
			},
			"Passive": {
				"displayName": "Illumination",
				"id": "LuxIlluminationPassive",
				"rawDescription": "GeneratedTip_Passive_LuxIlluminationPassive_Description",
				"rawDisplayName": "GeneratedTip_Passive_LuxIlluminationPassive_DisplayName"
			},
			"Q": {
				"abilityLevel": 2,
				"displayName": "Light Binding",
				"id": "LuxLightBinding",
				"rawDescription": "GeneratedTip_Spell_LuxLightBinding_Description",
				"rawDisplayName": "GeneratedTip_Spell_LuxLightBinding_DisplayName"
			},
			"R": {
				"abilityLevel": 2,
				"displayName": "Final Spark",
				"id": "LuxMaliceCannon",
				"rawDescription": "GeneratedTip_Spell_LuxMaliceCannon_Description",
				"rawDisplayName": "GeneratedTip_Spell_LuxMaliceCannon_DisplayName"
			},
			"W": {
				"abilityLevel": 5,
				"displayName": "Prismatic Barrier",
				"id": "LuxPrismaticWave",
				"rawDescription": "GeneratedTip_Spell_LuxPrismaticWave_Description",
				"rawDisplayName": "GeneratedTip_Spell_LuxPrismaticWave_DisplayName"
			}
		},
		"championStats": {
			"abilityPower": 265.1999816894531,
			"armor": 118.36000061035156,
			"armorPenetrationFlat": 0.0,
			"armorPenetrationPercent": 1.0,
			"attackDamage": 97.65287780761719,
			"attackRange": 550.0,
			"attackSpeed": 0.744562566280365,
			"bonusArmorPenetrationPercent": 1.0,
			"bonusMagicPenetrationPercent": 1.0,
			"cooldownReduction": 0.0,
			"critChance": 0.0,
			"critDamage": 175.0,
			"currentHealth": 1280.851806640625,
			"healthRegenRate": 2.4299001693725588,
			"lifeSteal": 0.0,
			"magicLethality": 0.0,
			"magicPenetrationFlat": 39.0,
			"magicPenetrationPercent": 1.0,
			"magicResist": 36.04500198364258,
			"maxHealth": 2081.38818359375,
			"moveSpeed": 375.0,
			"physicalLethality": 0.0,
			"resourceMax": 1514.114990234375,
			"resourceRegenRate": 7.06879997253418,
			"resourceType": "MANA",
			"resourceValue": 1477.81982421875,
			"spellVamp": 0.0,
			"tenacity": 0.0
		},
		"currentGold": 836.0624389648438,
		"fullRunes": {
			"generalRunes": [
				{
					"displayName": "Glacial Augment",
					"id": 8351,
					"rawDescription": "perk_tooltip_GlacialAugment",
					"rawDisplayName": "perk_displayname_GlacialAugment"
				},
				{
					"displayName": "Perfect Timing",
					"id": 8313,
					"rawDescription": "perk_tooltip_PerfectTiming",
					"rawDisplayName": "perk_displayname_PerfectTiming"
				},
				{
					"displayName": "Biscuit Delivery",
					"id": 8345,
					"rawDescription": "perk_tooltip_BiscuitDelivery",
					"rawDisplayName": "perk_displayname_BiscuitDelivery"
				},
				{
					"displayName": "Cosmic Insight",
					"id": 8347,
					"rawDescription": "perk_tooltip_CosmicInsight",
					"rawDisplayName": "perk_displayname_CosmicInsight"
				},
				{
					"displayName": "Overgrowth",
					"id": 8451,
					"rawDescription": "perk_tooltip_Overgrowth",
					"rawDisplayName": "perk_displayname_Overgrowth"
				},
				{
					"displayName": "Second Wind",
					"id": 8444,
					"rawDescription": "perk_tooltip_SecondWind",
					"rawDisplayName": "perk_displayname_SecondWind"
				}
			],
			"keystone": {
				"displayName": "Glacial Augment",
				"id": 8351,
				"rawDescription": "perk_tooltip_GlacialAugment",
				"rawDisplayName": "perk_displayname_GlacialAugment"
			},
			"primaryRuneTree": {
				"displayName": "Inspiration",
				"id": 8300,
				"rawDescription": "perkstyle_tooltip_7203",
				"rawDisplayName": "perkstyle_displayname_7203"
			},
			"secondaryRuneTree": {
				"displayName": "Resolve",
				"id": 8400,
				"rawDescription": "perkstyle_tooltip_7204",
				"rawDisplayName": "perkstyle_displayname_7204"
			},
			"statRunes": [
				{
					"id": 5007,
					"rawDescription": "perk_tooltip_StatModCooldownReductionScaling"
				},
				{
					"id": 5002,
					"rawDescription": "perk_tooltip_StatModArmor"
				},
				{
					"id": 5001,
					"rawDescription": "perk_tooltip_StatModHealthScaling"
				}
			]
		},
		"level": 14,
		"summonerName": "mikehpoo"
	},
	"allPlayers": [
		{
			"championName": "Brand",
			"isBot": false,
			"isDead": false,
			"items": [
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Demonic Embrace",
					"itemID": 4637,
					"price": 815,
					"rawDescription": "GeneratedTip_Item_4637_Description",
					"rawDisplayName": "Item_4637_Name",
					"slot": 0
				},
				{
					"canUse": true,
					"consumable": true,
					"count": 1,
					"displayName": "Elixir of Sorcery",
					"itemID": 2139,
					"price": 500,
					"rawDescription": "GeneratedTip_Item_2139_Description",
					"rawDisplayName": "Item_2139_Name",
					"slot": 1
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Seeker's Armguard",
					"itemID": 3191,
					"price": 265,
					"rawDescription": "GeneratedTip_Item_3191_Description",
					"rawDisplayName": "Item_3191_Name",
					"slot": 2
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Broken Stopwatch",
					"itemID": 2421,
					"price": 650,
					"rawDescription": "GeneratedTip_Item_2421_Description",
					"rawDisplayName": "Item_2421_Name",
					"slot": 3
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Liandry's Anguish",
					"itemID": 6653,
					"price": 1200,
					"rawDescription": "GeneratedTip_Item_6653_Description",
					"rawDisplayName": "Item_6653_Name",
					"slot": 4
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Sorcerer's Shoes",
					"itemID": 3020,
					"price": 800,
					"rawDescription": "GeneratedTip_Item_3020_Description",
					"rawDisplayName": "Item_3020_Name",
					"slot": 5
				},
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Stealth Ward",
					"itemID": 3340,
					"price": 0,
					"rawDescription": "GeneratedTip_Item_3340_Description",
					"rawDisplayName": "Item_3340_Name",
					"slot": 6
				}
			],
			"level": 16,
			"position": "TOP",
			"rawChampionName": "game_character_displayname_Brand",
			"respawnTimer": 0.0,
			"runes": {
				"keystone": {
					"displayName": "Dark Harvest",
					"id": 8128,
					"rawDescription": "perk_tooltip_DarkHarvest",
					"rawDisplayName": "perk_displayname_DarkHarvest"
				},
				"primaryRuneTree": {
					"displayName": "Domination",
					"id": 8100,
					"rawDescription": "perkstyle_tooltip_7200",
					"rawDisplayName": "perkstyle_displayname_7200"
				},
				"secondaryRuneTree": {
					"displayName": "Inspiration",
					"id": 8300,
					"rawDescription": "perkstyle_tooltip_7203",
					"rawDisplayName": "perkstyle_displayname_7203"
				}
			},
			"scores": {
				"assists": 5,
				"creepScore": 140,
				"deaths": 4,
				"kills": 0,
				"wardScore": 4.8
			},
			"skinID": 0,
			"summonerName": "wezznco",
			"summonerSpells": {
				"summonerSpellOne": {
					"displayName": "Flash",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerFlash_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerFlash_DisplayName"
				},
				"summonerSpellTwo": {
					"displayName": "Ignite",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerDot_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerDot_DisplayName"
				}
			},
			"team": "ORDER"
		},
		{
			"championName": "Shyvana",
			"isBot": false,
			"isDead": false,
			"items": [
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Night Harvester",
					"itemID": 4636,
					"price": 900,
					"rawDescription": "GeneratedTip_Item_4636_Description",
					"rawDisplayName": "Item_4636_Name",
					"slot": 0
				},
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Refillable Potion",
					"itemID": 2031,
					"price": 150,
					"rawDescription": "GeneratedTip_Item_2031_Description",
					"rawDisplayName": "Item_2031_Name",
					"slot": 1
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Morellonomicon",
					"itemID": 3165,
					"price": 450,
					"rawDescription": "GeneratedTip_Item_3165_Description",
					"rawDisplayName": "Item_3165_Name",
					"slot": 2
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Sorcerer's Shoes",
					"itemID": 3020,
					"price": 800,
					"rawDescription": "GeneratedTip_Item_3020_Description",
					"rawDisplayName": "Item_3020_Name",
					"slot": 3
				},
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Zhonya's Hourglass",
					"itemID": 3157,
					"price": 50,
					"rawDescription": "GeneratedTip_Item_3157_Description",
					"rawDisplayName": "Item_3157_Name",
					"slot": 4
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Dark Seal",
					"itemID": 1082,
					"price": 350,
					"rawDescription": "GeneratedTip_Item_1082_Description",
					"rawDisplayName": "Item_1082_Name",
					"slot": 5
				},
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Oracle Lens",
					"itemID": 3364,
					"price": 0,
					"rawDescription": "GeneratedTip_Item_3364_Description",
					"rawDisplayName": "Item_3364_Name",
					"slot": 6
				}
			],
			"level": 15,
			"position": "JUNGLE",
			"rawChampionName": "game_character_displayname_Shyvana",
			"respawnTimer": 0.0,
			"runes": {
				"keystone": {
					"displayName": "Dark Harvest",
					"id": 8128,
					"rawDescription": "perk_tooltip_DarkHarvest",
					"rawDisplayName": "perk_displayname_DarkHarvest"
				},
				"primaryRuneTree": {
					"displayName": "Domination",
					"id": 8100,
					"rawDescription": "perkstyle_tooltip_7200",
					"rawDisplayName": "perkstyle_displayname_7200"
				},
				"secondaryRuneTree": {
					"displayName": "Sorcery",
					"id": 8200,
					"rawDescription": "perkstyle_tooltip_7202",
					"rawDisplayName": "perkstyle_displayname_7202"
				}
			},
			"scores": {
				"assists": 14,
				"creepScore": 150,
				"deaths": 8,
				"kills": 5,
				"wardScore": 1.5
			},
			"skinID": 0,
			"summonerName": "piperx",
			"summonerSpells": {
				"summonerSpellOne": {
					"displayName": "Flash",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerFlash_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerFlash_DisplayName"
				},
				"summonerSpellTwo": {
					"displayName": "Chilling Smite",
					"rawDescription": "GeneratedTip_SummonerSpell_S5_SummonerSmitePlayerGanker_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_S5_SummonerSmitePlayerGanker_DisplayName"
				}
			},
			"team": "ORDER"
		},
		{
			"championName": "Ezreal",
			"isBot": false,
			"isDead": false,
			"items": [
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Muramana",
					"itemID": 3042,
					"price": 2900,
					"rawDescription": "GeneratedTip_Item_3042_Description",
					"rawDisplayName": "Item_3042_Name",
					"slot": 0
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Ionian Boots of Lucidity",
					"itemID": 3158,
					"price": 600,
					"rawDescription": "GeneratedTip_Item_3158_Description",
					"rawDisplayName": "Item_3158_Name",
					"slot": 1
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Divine Sunderer",
					"itemID": 6632,
					"price": 700,
					"rawDescription": "GeneratedTip_Item_6632_Description",
					"rawDisplayName": "Item_6632_Name",
					"slot": 2
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Vampiric Scepter",
					"itemID": 1053,
					"price": 550,
					"rawDescription": "GeneratedTip_Item_1053_Description",
					"rawDisplayName": "Item_1053_Name",
					"slot": 3
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Dagger",
					"itemID": 1042,
					"price": 300,
					"rawDescription": "GeneratedTip_Item_1042_Description",
					"rawDisplayName": "Item_1042_Name",
					"slot": 4
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Dagger",
					"itemID": 1042,
					"price": 300,
					"rawDescription": "GeneratedTip_Item_1042_Description",
					"rawDisplayName": "Item_1042_Name",
					"slot": 5
				},
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Stealth Ward",
					"itemID": 3340,
					"price": 0,
					"rawDescription": "GeneratedTip_Item_3340_Description",
					"rawDisplayName": "Item_3340_Name",
					"slot": 6
				}
			],
			"level": 14,
			"position": "MIDDLE",
			"rawChampionName": "game_character_displayname_Ezreal",
			"rawSkinName": "game_character_skin_displayname_Ezreal_18",
			"respawnTimer": 0.0,
			"runes": {
				"keystone": {
					"displayName": "Conqueror",
					"id": 8010,
					"rawDescription": "perk_tooltip_Conqueror",
					"rawDisplayName": "perk_displayname_Conqueror"
				},
				"primaryRuneTree": {
					"displayName": "Precision",
					"id": 8000,
					"rawDescription": "perkstyle_tooltip_7201",
					"rawDisplayName": "perkstyle_displayname_7201"
				},
				"secondaryRuneTree": {
					"displayName": "Inspiration",
					"id": 8300,
					"rawDescription": "perkstyle_tooltip_7203",
					"rawDisplayName": "perkstyle_displayname_7203"
				}
			},
			"scores": {
				"assists": 5,
				"creepScore": 120,
				"deaths": 10,
				"kills": 7,
				"wardScore": 5.947086334228516
			},
			"skinID": 18,
			"skinName": "Star Guardian Ezreal",
			"summonerName": "Zero Damage ADC",
			"summonerSpells": {
				"summonerSpellOne": {
					"displayName": "Heal",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerHeal_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerHeal_DisplayName"
				},
				"summonerSpellTwo": {
					"displayName": "Flash",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerFlash_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerFlash_DisplayName"
				}
			},
			"team": "ORDER"
		},
		{
			"championName": "Miss Fortune",
			"isBot": false,
			"isDead": false,
			"items": [
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Galeforce",
					"itemID": 6671,
					"price": 625,
					"rawDescription": "GeneratedTip_Item_6671_Description",
					"rawDisplayName": "Item_6671_Name",
					"slot": 0
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Mortal Reminder",
					"itemID": 3033,
					"price": 650,
					"rawDescription": "GeneratedTip_Item_3033_Description",
					"rawDisplayName": "Item_3033_Name",
					"slot": 1
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "The Collector",
					"itemID": 6676,
					"price": 425,
					"rawDescription": "GeneratedTip_Item_6676_Description",
					"rawDisplayName": "Item_6676_Name",
					"slot": 2
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Berserker's Greaves",
					"itemID": 3006,
					"price": 500,
					"rawDescription": "GeneratedTip_Item_3006_Description",
					"rawDisplayName": "Item_3006_Name",
					"slot": 4
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Doran's Blade",
					"itemID": 1055,
					"price": 450,
					"rawDescription": "GeneratedTip_Item_1055_Description",
					"rawDisplayName": "Item_1055_Name",
					"slot": 5
				},
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Stealth Ward",
					"itemID": 3340,
					"price": 0,
					"rawDescription": "GeneratedTip_Item_3340_Description",
					"rawDisplayName": "Item_3340_Name",
					"slot": 6
				}
			],
			"level": 14,
			"position": "BOTTOM",
			"rawChampionName": "game_character_displayname_MissFortune",
			"rawSkinName": "game_character_skin_displayname_MissFortune_17",
			"respawnTimer": 0.0,
			"runes": {
				"keystone": {
					"displayName": "Press the Attack",
					"id": 8005,
					"rawDescription": "perk_tooltip_PressTheAttack",
					"rawDisplayName": "perk_displayname_PressTheAttack"
				},
				"primaryRuneTree": {
					"displayName": "Precision",
					"id": 8000,
					"rawDescription": "perkstyle_tooltip_7201",
					"rawDisplayName": "perkstyle_displayname_7201"
				},
				"secondaryRuneTree": {
					"displayName": "Domination",
					"id": 8100,
					"rawDescription": "perkstyle_tooltip_7200",
					"rawDisplayName": "perkstyle_displayname_7200"
				}
			},
			"scores": {
				"assists": 13,
				"creepScore": 170,
				"deaths": 9,
				"kills": 5,
				"wardScore": 0.0
			},
			"skinID": 17,
			"skinName": "Pajama Guardian Miss Fortune",
			"summonerName": "ntalice",
			"summonerSpells": {
				"summonerSpellOne": {
					"displayName": "Heal",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerHeal_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerHeal_DisplayName"
				},
				"summonerSpellTwo": {
					"displayName": "Flash",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerFlash_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerFlash_DisplayName"
				}
			},
			"team": "ORDER"
		},
		{
			"championName": "Lux",
			"isBot": false,
			"isDead": false,
			"items": [
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Zhonya's Hourglass",
					"itemID": 3157,
					"price": 50,
					"rawDescription": "GeneratedTip_Item_3157_Description",
					"rawDisplayName": "Item_3157_Name",
					"slot": 0
				},
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Shard of True Ice",
					"itemID": 3853,
					"price": 400,
					"rawDescription": "GeneratedTip_Item_3853_Description",
					"rawDisplayName": "Item_3853_Name",
					"slot": 2
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Luden's Tempest",
					"itemID": 6655,
					"price": 1250,
					"rawDescription": "GeneratedTip_Item_6655_Description",
					"rawDisplayName": "Item_6655_Name",
					"slot": 3
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Sorcerer's Shoes",
					"itemID": 3020,
					"price": 800,
					"rawDescription": "GeneratedTip_Item_3020_Description",
					"rawDisplayName": "Item_3020_Name",
					"slot": 4
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Morellonomicon",
					"itemID": 3165,
					"price": 450,
					"rawDescription": "GeneratedTip_Item_3165_Description",
					"rawDisplayName": "Item_3165_Name",
					"slot": 5
				},
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Stealth Ward",
					"itemID": 3340,
					"price": 0,
					"rawDescription": "GeneratedTip_Item_3340_Description",
					"rawDisplayName": "Item_3340_Name",
					"slot": 6
				}
			],
			"level": 14,
			"position": "UTILITY",
			"rawChampionName": "game_character_displayname_Lux",
			"respawnTimer": 0.0,
			"runes": {
				"keystone": {
					"displayName": "Glacial Augment",
					"id": 8351,
					"rawDescription": "perk_tooltip_GlacialAugment",
					"rawDisplayName": "perk_displayname_GlacialAugment"
				},
				"primaryRuneTree": {
					"displayName": "Inspiration",
					"id": 8300,
					"rawDescription": "perkstyle_tooltip_7203",
					"rawDisplayName": "perkstyle_displayname_7203"
				},
				"secondaryRuneTree": {
					"displayName": "Resolve",
					"id": 8400,
					"rawDescription": "perkstyle_tooltip_7204",
					"rawDisplayName": "perkstyle_displayname_7204"
				}
			},
			"scores": {
				"assists": 10,
				"creepScore": 30,
				"deaths": 9,
				"kills": 6,
				"wardScore": 42.56794357299805
			},
			"skinID": 0,
			"summonerName": "mikehpoo",
			"summonerSpells": {
				"summonerSpellOne": {
					"displayName": "Flash",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerFlash_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerFlash_DisplayName"
				},
				"summonerSpellTwo": {
					"displayName": "Exhaust",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerExhaust_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerExhaust_DisplayName"
				}
			},
			"team": "ORDER"
		},
		{
			"championName": "Shen",
			"isBot": false,
			"isDead": false,
			"items": [
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Titanic Hydra",
					"itemID": 3748,
					"price": 800,
					"rawDescription": "GeneratedTip_Item_3748_Description",
					"rawDisplayName": "Item_3748_Name",
					"slot": 0
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Divine Sunderer",
					"itemID": 6632,
					"price": 700,
					"rawDescription": "GeneratedTip_Item_6632_Description",
					"rawDisplayName": "Item_6632_Name",
					"slot": 1
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Mercury's Treads",
					"itemID": 3111,
					"price": 350,
					"rawDescription": "GeneratedTip_Item_3111_Description",
					"rawDisplayName": "Item_3111_Name",
					"slot": 2
				},
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Stealth Ward",
					"itemID": 3340,
					"price": 0,
					"rawDescription": "GeneratedTip_Item_3340_Description",
					"rawDisplayName": "Item_3340_Name",
					"slot": 6
				}
			],
			"level": 14,
			"position": "TOP",
			"rawChampionName": "game_character_displayname_Shen",
			"rawSkinName": "game_character_skin_displayname_Shen_20",
			"respawnTimer": 0.0,
			"runes": {
				"keystone": {
					"displayName": "Grasp of the Undying",
					"id": 8437,
					"rawDescription": "perk_tooltip_GraspOfTheUndying",
					"rawDisplayName": "perk_displayname_GraspOfTheUndying"
				},
				"primaryRuneTree": {
					"displayName": "Resolve",
					"id": 8400,
					"rawDescription": "perkstyle_tooltip_7204",
					"rawDisplayName": "perkstyle_displayname_7204"
				},
				"secondaryRuneTree": {
					"displayName": "Domination",
					"id": 8100,
					"rawDescription": "perkstyle_tooltip_7200",
					"rawDisplayName": "perkstyle_displayname_7200"
				}
			},
			"scores": {
				"assists": 6,
				"creepScore": 110,
				"deaths": 9,
				"kills": 1,
				"wardScore": 17.243696212768556
			},
			"skinID": 20,
			"skinName": "Infernal Shen",
			"summonerName": "GrantReaper",
			"summonerSpells": {
				"summonerSpellOne": {
					"displayName": "Teleport",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerTeleport_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerTeleport_DisplayName"
				},
				"summonerSpellTwo": {
					"displayName": "Flash",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerFlash_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerFlash_DisplayName"
				}
			},
			"team": "CHAOS"
		},
		{
			"championName": "Master Yi",
			"isBot": false,
			"isDead": false,
			"items": [
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Kraken Slayer",
					"itemID": 6672,
					"price": 625,
					"rawDescription": "GeneratedTip_Item_6672_Description",
					"rawDisplayName": "Item_6672_Name",
					"slot": 0
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Guinsoo's Rageblade",
					"itemID": 3124,
					"price": 1100,
					"rawDescription": "GeneratedTip_Item_3124_Description",
					"rawDisplayName": "Item_3124_Name",
					"slot": 1
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Berserker's Greaves",
					"itemID": 3006,
					"price": 500,
					"rawDescription": "GeneratedTip_Item_3006_Description",
					"rawDisplayName": "Item_3006_Name",
					"slot": 2
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Blade of The Ruined King",
					"itemID": 3153,
					"price": 425,
					"rawDescription": "GeneratedTip_Item_3153_Description",
					"rawDisplayName": "Item_3153_Name",
					"slot": 3
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Guardian Angel",
					"itemID": 3026,
					"price": 50,
					"rawDescription": "GeneratedTip_Item_3026_Description",
					"rawDisplayName": "Item_3026_Name",
					"slot": 4
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Caulfield's Warhammer",
					"itemID": 3133,
					"price": 400,
					"rawDescription": "GeneratedTip_Item_3133_Description",
					"rawDisplayName": "Item_3133_Name",
					"slot": 5
				},
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Stealth Ward",
					"itemID": 3340,
					"price": 0,
					"rawDescription": "GeneratedTip_Item_3340_Description",
					"rawDisplayName": "Item_3340_Name",
					"slot": 6
				}
			],
			"level": 16,
			"position": "JUNGLE",
			"rawChampionName": "game_character_displayname_MasterYi",
			"rawSkinName": "game_character_skin_displayname_MasterYi_2",
			"respawnTimer": 0.0,
			"runes": {
				"keystone": {
					"displayName": "Conqueror",
					"id": 8010,
					"rawDescription": "perk_tooltip_Conqueror",
					"rawDisplayName": "perk_displayname_Conqueror"
				},
				"primaryRuneTree": {
					"displayName": "Precision",
					"id": 8000,
					"rawDescription": "perkstyle_tooltip_7201",
					"rawDisplayName": "perkstyle_displayname_7201"
				},
				"secondaryRuneTree": {
					"displayName": "Domination",
					"id": 8100,
					"rawDescription": "perkstyle_tooltip_7200",
					"rawDisplayName": "perkstyle_displayname_7200"
				}
			},
			"scores": {
				"assists": 5,
				"creepScore": 130,
				"deaths": 7,
				"kills": 23,
				"wardScore": 12.207406044006348
			},
			"skinID": 2,
			"skinName": "Chosen Master Yi",
			"summonerName": "MoRice651",
			"summonerSpells": {
				"summonerSpellOne": {
					"displayName": "Ghost",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerHaste_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerHaste_DisplayName"
				},
				"summonerSpellTwo": {
					"displayName": "Chilling Smite",
					"rawDescription": "GeneratedTip_SummonerSpell_S5_SummonerSmitePlayerGanker_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_S5_SummonerSmitePlayerGanker_DisplayName"
				}
			},
			"team": "CHAOS"
		},
		{
			"championName": "Ziggs",
			"isBot": false,
			"isDead": false,
			"items": [
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Zhonya's Hourglass",
					"itemID": 3157,
					"price": 50,
					"rawDescription": "GeneratedTip_Item_3157_Description",
					"rawDisplayName": "Item_3157_Name",
					"slot": 0
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Luden's Tempest",
					"itemID": 6655,
					"price": 1250,
					"rawDescription": "GeneratedTip_Item_6655_Description",
					"rawDisplayName": "Item_6655_Name",
					"slot": 1
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Seraph's Embrace",
					"itemID": 3040,
					"price": 3000,
					"rawDescription": "GeneratedTip_Item_3040_Description",
					"rawDisplayName": "Item_3040_Name",
					"slot": 2
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Sorcerer's Shoes",
					"itemID": 3020,
					"price": 800,
					"rawDescription": "GeneratedTip_Item_3020_Description",
					"rawDisplayName": "Item_3020_Name",
					"slot": 3
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Morellonomicon",
					"itemID": 3165,
					"price": 450,
					"rawDescription": "GeneratedTip_Item_3165_Description",
					"rawDisplayName": "Item_3165_Name",
					"slot": 4
				},
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Stealth Ward",
					"itemID": 3340,
					"price": 0,
					"rawDescription": "GeneratedTip_Item_3340_Description",
					"rawDisplayName": "Item_3340_Name",
					"slot": 6
				}
			],
			"level": 16,
			"position": "MIDDLE",
			"rawChampionName": "game_character_displayname_Ziggs",
			"rawSkinName": "game_character_skin_displayname_Ziggs_5",
			"respawnTimer": 0.0,
			"runes": {
				"keystone": {
					"displayName": "Arcane Comet",
					"id": 8229,
					"rawDescription": "perk_tooltip_ArcaneComet",
					"rawDisplayName": "perk_displayname_ArcaneComet"
				},
				"primaryRuneTree": {
					"displayName": "Sorcery",
					"id": 8200,
					"rawDescription": "perkstyle_tooltip_7202",
					"rawDisplayName": "perkstyle_displayname_7202"
				},
				"secondaryRuneTree": {
					"displayName": "Inspiration",
					"id": 8300,
					"rawDescription": "perkstyle_tooltip_7203",
					"rawDisplayName": "perkstyle_displayname_7203"
				}
			},
			"scores": {
				"assists": 12,
				"creepScore": 170,
				"deaths": 5,
				"kills": 8,
				"wardScore": 16.796411514282228
			},
			"skinID": 5,
			"skinName": "Master Arcanist Ziggs",
			"summonerName": "Japudi",
			"summonerSpells": {
				"summonerSpellOne": {
					"displayName": "Teleport",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerTeleport_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerTeleport_DisplayName"
				},
				"summonerSpellTwo": {
					"displayName": "Flash",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerFlash_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerFlash_DisplayName"
				}
			},
			"team": "CHAOS"
		},
		{
			"championName": "Samira",
			"isBot": false,
			"isDead": false,
			"items": [
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Immortal Shieldbow",
					"itemID": 6673,
					"price": 600,
					"rawDescription": "GeneratedTip_Item_6673_Description",
					"rawDisplayName": "Item_6673_Name",
					"slot": 0
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Berserker's Greaves",
					"itemID": 3006,
					"price": 500,
					"rawDescription": "GeneratedTip_Item_3006_Description",
					"rawDisplayName": "Item_3006_Name",
					"slot": 1
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "The Collector",
					"itemID": 6676,
					"price": 425,
					"rawDescription": "GeneratedTip_Item_6676_Description",
					"rawDisplayName": "Item_6676_Name",
					"slot": 2
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Doran's Blade",
					"itemID": 1055,
					"price": 450,
					"rawDescription": "GeneratedTip_Item_1055_Description",
					"rawDisplayName": "Item_1055_Name",
					"slot": 3
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Cloak of Agility",
					"itemID": 1018,
					"price": 600,
					"rawDescription": "GeneratedTip_Item_1018_Description",
					"rawDisplayName": "Item_1018_Name",
					"slot": 4
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Vampiric Scepter",
					"itemID": 1053,
					"price": 550,
					"rawDescription": "GeneratedTip_Item_1053_Description",
					"rawDisplayName": "Item_1053_Name",
					"slot": 5
				},
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Stealth Ward",
					"itemID": 3340,
					"price": 0,
					"rawDescription": "GeneratedTip_Item_3340_Description",
					"rawDisplayName": "Item_3340_Name",
					"slot": 6
				}
			],
			"level": 14,
			"position": "BOTTOM",
			"rawChampionName": "game_character_displayname_Samira",
			"respawnTimer": 0.0,
			"runes": {
				"keystone": {
					"displayName": "Conqueror",
					"id": 8010,
					"rawDescription": "perk_tooltip_Conqueror",
					"rawDisplayName": "perk_displayname_Conqueror"
				},
				"primaryRuneTree": {
					"displayName": "Precision",
					"id": 8000,
					"rawDescription": "perkstyle_tooltip_7201",
					"rawDisplayName": "perkstyle_displayname_7201"
				},
				"secondaryRuneTree": {
					"displayName": "Domination",
					"id": 8100,
					"rawDescription": "perkstyle_tooltip_7200",
					"rawDisplayName": "perkstyle_displayname_7200"
				}
			},
			"scores": {
				"assists": 7,
				"creepScore": 140,
				"deaths": 4,
				"kills": 9,
				"wardScore": 15.132512092590332
			},
			"skinID": 0,
			"summonerName": "GhostBig",
			"summonerSpells": {
				"summonerSpellOne": {
					"displayName": "Flash",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerFlash_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerFlash_DisplayName"
				},
				"summonerSpellTwo": {
					"displayName": "Heal",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerHeal_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerHeal_DisplayName"
				}
			},
			"team": "CHAOS"
		},
		{
			"championName": "Leona",
			"isBot": false,
			"isDead": false,
			"items": [
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Bulwark of the Mountain",
					"itemID": 3860,
					"price": 400,
					"rawDescription": "GeneratedTip_Item_3860_Description",
					"rawDisplayName": "Item_3860_Name",
					"slot": 0
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Mercury's Treads",
					"itemID": 3111,
					"price": 350,
					"rawDescription": "GeneratedTip_Item_3111_Description",
					"rawDisplayName": "Item_3111_Name",
					"slot": 1
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Sunfire Aegis",
					"itemID": 3068,
					"price": 600,
					"rawDescription": "GeneratedTip_Item_3068_Description",
					"rawDisplayName": "Item_3068_Name",
					"slot": 2
				},
				{
					"canUse": false,
					"consumable": false,
					"count": 1,
					"displayName": "Abyssal Mask",
					"itemID": 3001,
					"price": 900,
					"rawDescription": "GeneratedTip_Item_3001_Description",
					"rawDisplayName": "Item_3001_Name",
					"slot": 3
				},
				{
					"canUse": true,
					"consumable": false,
					"count": 1,
					"displayName": "Stealth Ward",
					"itemID": 3340,
					"price": 0,
					"rawDescription": "GeneratedTip_Item_3340_Description",
					"rawDisplayName": "Item_3340_Name",
					"slot": 6
				}
			],
			"level": 13,
			"position": "UTILITY",
			"rawChampionName": "game_character_displayname_Leona",
			"rawSkinName": "game_character_skin_displayname_Leona_11",
			"respawnTimer": 0.0,
			"runes": {
				"keystone": {
					"displayName": "Aftershock",
					"id": 8439,
					"rawDescription": "perk_tooltip_VeteranAftershock",
					"rawDisplayName": "perk_displayname_VeteranAftershock"
				},
				"primaryRuneTree": {
					"displayName": "Resolve",
					"id": 8400,
					"rawDescription": "perkstyle_tooltip_7204",
					"rawDisplayName": "perkstyle_displayname_7204"
				},
				"secondaryRuneTree": {
					"displayName": "Inspiration",
					"id": 8300,
					"rawDescription": "perkstyle_tooltip_7203",
					"rawDisplayName": "perkstyle_displayname_7203"
				}
			},
			"scores": {
				"assists": 13,
				"creepScore": 30,
				"deaths": 5,
				"kills": 4,
				"wardScore": 35.7579460144043
			},
			"skinID": 11,
			"skinName": "Lunar Eclipse Leona",
			"summonerName": "KubiK651",
			"summonerSpells": {
				"summonerSpellOne": {
					"displayName": "Flash",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerFlash_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerFlash_DisplayName"
				},
				"summonerSpellTwo": {
					"displayName": "Ignite",
					"rawDescription": "GeneratedTip_SummonerSpell_SummonerDot_Description",
					"rawDisplayName": "GeneratedTip_SummonerSpell_SummonerDot_DisplayName"
				}
			},
			"team": "CHAOS"
		}
	],
	"events": {
		"Events": [
			{
				"EventID": 0,
				"EventName": "GameStart",
				"EventTime": 0.0658269003033638
			},
			{
				"EventID": 1,
				"EventName": "MinionsSpawning",
				"EventTime": 65.0288314819336
			},
			{
				"Assisters": [
					"MoRice651"
				],
				"EventID": 2,
				"EventName": "ChampionKill",
				"EventTime": 93.05531311035156,
				"KillerName": "GrantReaper",
				"VictimName": "wezznco"
			},
			{
				"EventID": 3,
				"EventName": "FirstBlood",
				"EventTime": 93.05531311035156,
				"Recipient": "GrantReaper"
			},
			{
				"Assisters": [
					"piperx"
				],
				"EventID": 4,
				"EventName": "ChampionKill",
				"EventTime": 254.929931640625,
				"KillerName": "Zero Damage ADC",
				"VictimName": "Japudi"
			},
			{
				"Assisters": [],
				"EventID": 5,
				"EventName": "ChampionKill",
				"EventTime": 302.123291015625,
				"KillerName": "MoRice651",
				"VictimName": "Zero Damage ADC"
			},
			{
				"Assisters": [
					"Japudi"
				],
				"EventID": 6,
				"EventName": "ChampionKill",
				"EventTime": 310.0793762207031,
				"KillerName": "MoRice651",
				"VictimName": "piperx"
			},
			{
				"EventID": 7,
				"EventName": "Multikill",
				"EventTime": 310.0793762207031,
				"KillStreak": 2,
				"KillerName": "MoRice651"
			},
			{
				"Assisters": [],
				"EventID": 8,
				"EventName": "ChampionKill",
				"EventTime": 371.86322021484377,
				"KillerName": "Zero Damage ADC",
				"VictimName": "Japudi"
			},
			{
				"Assisters": [],
				"DragonType": "Water",
				"EventID": 9,
				"EventName": "DragonKill",
				"EventTime": 387.78826904296877,
				"KillerName": "piperx",
				"Stolen": "False"
			},
			{
				"Assisters": [],
				"EventID": 10,
				"EventName": "ChampionKill",
				"EventTime": 476.97442626953127,
				"KillerName": "wezznco",
				"VictimName": "GrantReaper"
			},
			{
				"Assisters": [
					"Japudi"
				],
				"EventID": 11,
				"EventName": "ChampionKill",
				"EventTime": 500.7935485839844,
				"KillerName": "MoRice651",
				"VictimName": "Zero Damage ADC"
			},
			{
				"Assisters": [
					"GhostBig",
					"KubiK651"
				],
				"EventID": 12,
				"EventName": "ChampionKill",
				"EventTime": 599.2725219726563,
				"KillerName": "MoRice651",
				"VictimName": "ntalice"
			},
			{
				"Assisters": [
					"GhostBig",
					"KubiK651"
				],
				"EventID": 13,
				"EventName": "ChampionKill",
				"EventTime": 603.2362060546875,
				"KillerName": "MoRice651",
				"VictimName": "mikehpoo"
			},
			{
				"EventID": 14,
				"EventName": "Multikill",
				"EventTime": 603.2362060546875,
				"KillStreak": 2,
				"KillerName": "MoRice651"
			},
			{
				"Assisters": [],
				"EventID": 15,
				"EventName": "ChampionKill",
				"EventTime": 638.1248779296875,
				"KillerName": "wezznco",
				"VictimName": "GrantReaper"
			},
			{
				"Assisters": [
					"MoRice651"
				],
				"EventID": 16,
				"EventName": "ChampionKill",
				"EventTime": 640.3074340820313,
				"KillerName": "Japudi",
				"VictimName": "Zero Damage ADC"
			},
			{
				"Assisters": [],
				"EventID": 17,
				"EventName": "HeraldKill",
				"EventTime": 648.8998413085938,
				"KillerName": "piperx",
				"Stolen": "False"
			},
			{
				"Assisters": [
					"GrantReaper"
				],
				"EventID": 18,
				"EventName": "ChampionKill",
				"EventTime": 686.2005615234375,
				"KillerName": "MoRice651",
				"VictimName": "wezznco"
			},
			{
				"Assisters": [
					"ntalice"
				],
				"EventID": 19,
				"EventName": "TurretKilled",
				"EventTime": 721.640380859375,
				"KillerName": "mikehpoo",
				"TurretKilled": "Turret_T2_R_03_A"
			},
			{
				"EventID": 20,
				"EventName": "FirstBrick",
				"EventTime": 721.640380859375,
				"KillerName": "mikehpoo"
			},
			{
				"Assisters": [
					"GhostBig",
					"KubiK651"
				],
				"EventID": 21,
				"EventName": "ChampionKill",
				"EventTime": 722.0499267578125,
				"KillerName": "Japudi",
				"VictimName": "ntalice"
			},
			{
				"Assisters": [
					"Japudi",
					"KubiK651"
				],
				"EventID": 22,
				"EventName": "ChampionKill",
				"EventTime": 723.52587890625,
				"KillerName": "GhostBig",
				"VictimName": "mikehpoo"
			},
			{
				"Assisters": [
					"Japudi",
					"KubiK651"
				],
				"DragonType": "Air",
				"EventID": 23,
				"EventName": "DragonKill",
				"EventTime": 739.8707885742188,
				"KillerName": "MoRice651",
				"Stolen": "False"
			},
			{
				"Assisters": [],
				"EventID": 24,
				"EventName": "ChampionKill",
				"EventTime": 775.3260498046875,
				"KillerName": "wezznco",
				"VictimName": "GrantReaper"
			},
			{
				"Assisters": [
					"piperx",
					"Zero Damage ADC",
					"ntalice"
				],
				"EventID": 25,
				"EventName": "ChampionKill",
				"EventTime": 778.3513793945313,
				"KillerName": "mikehpoo",
				"VictimName": "MoRice651"
			},
			{
				"Assisters": [
					"MoRice651"
				],
				"EventID": 26,
				"EventName": "ChampionKill",
				"EventTime": 782.0662231445313,
				"KillerName": "Japudi",
				"VictimName": "Zero Damage ADC"
			},
			{
				"Assisters": [
					"MoRice651",
					"Japudi"
				],
				"EventID": 27,
				"EventName": "ChampionKill",
				"EventTime": 783.1781616210938,
				"KillerName": "GhostBig",
				"VictimName": "piperx"
			},
			{
				"Assisters": [
					"Japudi"
				],
				"EventID": 28,
				"EventName": "ChampionKill",
				"EventTime": 785.1104736328125,
				"KillerName": "GhostBig",
				"VictimName": "mikehpoo"
			},
			{
				"EventID": 29,
				"EventName": "Multikill",
				"EventTime": 785.1104736328125,
				"KillStreak": 2,
				"KillerName": "GhostBig"
			},
			{
				"Assisters": [
					"Japudi"
				],
				"EventID": 30,
				"EventName": "ChampionKill",
				"EventTime": 788.8355712890625,
				"KillerName": "GhostBig",
				"VictimName": "ntalice"
			},
			{
				"EventID": 31,
				"EventName": "Multikill",
				"EventTime": 788.8355712890625,
				"KillStreak": 3,
				"KillerName": "GhostBig"
			},
			{
				"Assisters": [],
				"EventID": 32,
				"EventName": "ChampionKill",
				"EventTime": 854.5018920898438,
				"KillerName": "MoRice651",
				"VictimName": "wezznco"
			},
			{
				"Assisters": [
					"KubiK651"
				],
				"EventID": 33,
				"EventName": "ChampionKill",
				"EventTime": 858.3759765625,
				"KillerName": "GhostBig",
				"VictimName": "mikehpoo"
			},
			{
				"Assisters": [
					"GhostBig"
				],
				"EventID": 34,
				"EventName": "ChampionKill",
				"EventTime": 862.6747436523438,
				"KillerName": "KubiK651",
				"VictimName": "ntalice"
			},
			{
				"Assisters": [],
				"EventID": 35,
				"EventName": "ChampionKill",
				"EventTime": 873.408203125,
				"KillerName": "MoRice651",
				"VictimName": "piperx"
			},
			{
				"Assisters": [],
				"EventID": 36,
				"EventName": "TurretKilled",
				"EventTime": 882.1104736328125,
				"KillerName": "Japudi",
				"TurretKilled": "Turret_T1_C_05_A"
			},
			{
				"Assisters": [
					"piperx"
				],
				"EventID": 37,
				"EventName": "ChampionKill",
				"EventTime": 883.364501953125,
				"KillerName": "Zero Damage ADC",
				"VictimName": "MoRice651"
			},
			{
				"Assisters": [
					"KubiK651"
				],
				"EventID": 38,
				"EventName": "TurretKilled",
				"EventTime": 902.2656860351563,
				"KillerName": "GhostBig",
				"TurretKilled": "Turret_T1_R_03_A"
			},
			{
				"Assisters": [],
				"EventID": 39,
				"EventName": "ChampionKill",
				"EventTime": 934.3631591796875,
				"KillerName": "Japudi",
				"VictimName": "Zero Damage ADC"
			},
			{
				"Assisters": [
					"piperx",
					"mikehpoo"
				],
				"EventID": 40,
				"EventName": "ChampionKill",
				"EventTime": 950.64892578125,
				"KillerName": "ntalice",
				"VictimName": "GhostBig"
			},
			{
				"Assisters": [
					"piperx",
					"ntalice"
				],
				"EventID": 41,
				"EventName": "ChampionKill",
				"EventTime": 952.1148071289063,
				"KillerName": "mikehpoo",
				"VictimName": "KubiK651"
			},
			{
				"Assisters": [],
				"EventID": 42,
				"EventName": "TurretKilled",
				"EventTime": 953.5919189453125,
				"KillerName": "Minion_T100L2S29N0178",
				"TurretKilled": "Turret_T2_L_03_A"
			},
			{
				"Assisters": [
					"piperx",
					"mikehpoo"
				],
				"EventID": 43,
				"EventName": "ChampionKill",
				"EventTime": 958.2581787109375,
				"KillerName": "ntalice",
				"VictimName": "GrantReaper"
			},
			{
				"EventID": 44,
				"EventName": "Multikill",
				"EventTime": 958.2581787109375,
				"KillStreak": 2,
				"KillerName": "ntalice"
			},
			{
				"Assisters": [
					"GrantReaper",
					"KubiK651"
				],
				"EventID": 45,
				"EventName": "ChampionKill",
				"EventTime": 964.437744140625,
				"KillerName": "MoRice651",
				"VictimName": "piperx"
			},
			{
				"Assisters": [],
				"EventID": 46,
				"EventName": "ChampionKill",
				"EventTime": 968.0065307617188,
				"KillerName": "MoRice651",
				"VictimName": "mikehpoo"
			},
			{
				"EventID": 47,
				"EventName": "Multikill",
				"EventTime": 968.0065307617188,
				"KillStreak": 2,
				"KillerName": "MoRice651"
			},
			{
				"Assisters": [],
				"EventID": 48,
				"EventName": "TurretKilled",
				"EventTime": 968.69921875,
				"KillerName": "Japudi",
				"TurretKilled": "Turret_T1_C_04_A"
			},
			{
				"Assisters": [],
				"EventID": 49,
				"EventName": "ChampionKill",
				"EventTime": 972.037841796875,
				"KillerName": "MoRice651",
				"VictimName": "ntalice"
			},
			{
				"EventID": 50,
				"EventName": "Multikill",
				"EventTime": 972.037841796875,
				"KillStreak": 3,
				"KillerName": "MoRice651"
			},
			{
				"Assisters": [
					"KubiK651"
				],
				"EventID": 51,
				"EventName": "ChampionKill",
				"EventTime": 1023.1942138671875,
				"KillerName": "Japudi",
				"VictimName": "wezznco"
			},
			{
				"Assisters": [],
				"EventID": 52,
				"EventName": "ChampionKill",
				"EventTime": 1026.886962890625,
				"KillerName": "wezznco",
				"VictimName": "KubiK651"
			},
			{
				"Assisters": [],
				"EventID": 53,
				"EventName": "TurretKilled",
				"EventTime": 1050.8404541015626,
				"KillerName": "Japudi",
				"TurretKilled": "Turret_T1_L_03_A"
			},
			{
				"Assisters": [
					"piperx",
					"Zero Damage ADC",
					"ntalice",
					"mikehpoo"
				],
				"DragonType": "Fire",
				"EventID": 54,
				"EventName": "DragonKill",
				"EventTime": 1051.360107421875,
				"KillerName": "MoRice651",
				"Stolen": "True"
			},
			{
				"Assisters": [
					"piperx",
					"Zero Damage ADC",
					"ntalice"
				],
				"EventID": 55,
				"EventName": "ChampionKill",
				"EventTime": 1053.919677734375,
				"KillerName": "mikehpoo",
				"VictimName": "MoRice651"
			},
			{
				"Assisters": [
					"GhostBig"
				],
				"EventID": 56,
				"EventName": "ChampionKill",
				"EventTime": 1094.4326171875,
				"KillerName": "KubiK651",
				"VictimName": "ntalice"
			},
			{
				"Assisters": [
					"KubiK651"
				],
				"EventID": 57,
				"EventName": "ChampionKill",
				"EventTime": 1095.4161376953126,
				"KillerName": "GhostBig",
				"VictimName": "mikehpoo"
			},
			{
				"Assisters": [
					"piperx"
				],
				"EventID": 58,
				"EventName": "ChampionKill",
				"EventTime": 1112.4810791015626,
				"KillerName": "Zero Damage ADC",
				"VictimName": "GrantReaper"
			},
			{
				"Assisters": [
					"Japudi"
				],
				"EventID": 59,
				"EventName": "HeraldKill",
				"EventTime": 1127.34228515625,
				"KillerName": "MoRice651",
				"Stolen": "False"
			},
			{
				"Assisters": [
					"KubiK651"
				],
				"EventID": 60,
				"EventName": "ChampionKill",
				"EventTime": 1141.85400390625,
				"KillerName": "GhostBig",
				"VictimName": "wezznco"
			},
			{
				"Assisters": [
					"KubiK651"
				],
				"EventID": 61,
				"EventName": "ChampionKill",
				"EventTime": 1143.474609375,
				"KillerName": "GhostBig",
				"VictimName": "Zero Damage ADC"
			},
			{
				"EventID": 62,
				"EventName": "Multikill",
				"EventTime": 1143.474609375,
				"KillStreak": 2,
				"KillerName": "GhostBig"
			},
			{
				"Assisters": [
					"wezznco",
					"Zero Damage ADC"
				],
				"EventID": 63,
				"EventName": "ChampionKill",
				"EventTime": 1144.693115234375,
				"KillerName": "piperx",
				"VictimName": "GhostBig"
			},
			{
				"Assisters": [
					"GrantReaper",
					"GhostBig"
				],
				"EventID": 64,
				"EventName": "ChampionKill",
				"EventTime": 1156.3231201171876,
				"KillerName": "KubiK651",
				"VictimName": "piperx"
			},
			{
				"Assisters": [],
				"EventID": 65,
				"EventName": "TurretKilled",
				"EventTime": 1161.8486328125,
				"KillerName": "Japudi",
				"TurretKilled": "Turret_T1_L_02_A"
			},
			{
				"Assisters": [
					"MoRice651"
				],
				"EventID": 66,
				"EventName": "TurretKilled",
				"EventTime": 1176.5400390625,
				"KillerName": "Japudi",
				"TurretKilled": "Turret_T1_C_06_A"
			},
			{
				"Assisters": [
					"Japudi"
				],
				"EventID": 67,
				"EventName": "InhibKilled",
				"EventTime": 1182.5499267578126,
				"InhibKilled": "Barracks_T1_L1",
				"KillerName": "RiftHeraldMercenary"
			},
			{
				"Assisters": [
					"Japudi"
				],
				"EventID": 68,
				"EventName": "ChampionKill",
				"EventTime": 1201.0078125,
				"KillerName": "MoRice651",
				"VictimName": "wezznco"
			},
			{
				"Assisters": [
					"Japudi"
				],
				"EventID": 69,
				"EventName": "ChampionKill",
				"EventTime": 1202.0162353515626,
				"KillerName": "MoRice651",
				"VictimName": "ntalice"
			},
			{
				"EventID": 70,
				"EventName": "Multikill",
				"EventTime": 1202.0162353515626,
				"KillStreak": 2,
				"KillerName": "MoRice651"
			},
			{
				"Assisters": [],
				"EventID": 71,
				"EventName": "ChampionKill",
				"EventTime": 1205.1785888671876,
				"KillerName": "MoRice651",
				"VictimName": "Zero Damage ADC"
			},
			{
				"EventID": 72,
				"EventName": "Multikill",
				"EventTime": 1205.1785888671876,
				"KillStreak": 3,
				"KillerName": "MoRice651"
			},
			{
				"Assisters": [
					"Japudi"
				],
				"EventID": 73,
				"EventName": "ChampionKill",
				"EventTime": 1210.1029052734376,
				"KillerName": "MoRice651",
				"VictimName": "mikehpoo"
			},
			{
				"EventID": 74,
				"EventName": "Multikill",
				"EventTime": 1210.1029052734376,
				"KillStreak": 4,
				"KillerName": "MoRice651"
			},
			{
				"Assisters": [
					"Zero Damage ADC",
					"ntalice",
					"mikehpoo"
				],
				"EventID": 75,
				"EventName": "ChampionKill",
				"EventTime": 1216.2138671875,
				"KillerName": "piperx",
				"VictimName": "MoRice651"
			},
			{
				"Assisters": [
					"GrantReaper",
					"Japudi",
					"KubiK651"
				],
				"EventID": 76,
				"EventName": "TurretKilled",
				"EventTime": 1236.6787109375,
				"KillerName": "GhostBig",
				"TurretKilled": "Turret_T1_C_03_A"
			},
			{
				"Assisters": [
					"GrantReaper",
					"KubiK651"
				],
				"EventID": 77,
				"EventName": "InhibKilled",
				"EventTime": 1247.778564453125,
				"InhibKilled": "Barracks_T1_C1",
				"KillerName": "GhostBig"
			},
			{
				"Assisters": [
					"GrantReaper",
					"Japudi",
					"KubiK651"
				],
				"EventID": 78,
				"EventName": "ChampionKill",
				"EventTime": 1253.623046875,
				"KillerName": "GhostBig",
				"VictimName": "piperx"
			},
			{
				"Assisters": [
					"piperx",
					"mikehpoo"
				],
				"EventID": 79,
				"EventName": "ChampionKill",
				"EventTime": 1255.039306640625,
				"KillerName": "ntalice",
				"VictimName": "GrantReaper"
			},
			{
				"Assisters": [
					"piperx",
					"mikehpoo"
				],
				"EventID": 80,
				"EventName": "ChampionKill",
				"EventTime": 1256.0147705078126,
				"KillerName": "ntalice",
				"VictimName": "GhostBig"
			},
			{
				"EventID": 81,
				"EventName": "Multikill",
				"EventTime": 1256.0147705078126,
				"KillStreak": 2,
				"KillerName": "ntalice"
			},
			{
				"Assisters": [
					"piperx",
					"ntalice",
					"mikehpoo"
				],
				"EventID": 82,
				"EventName": "ChampionKill",
				"EventTime": 1258.62841796875,
				"KillerName": "wezznco",
				"VictimName": "KubiK651"
			},
			{
				"Assisters": [],
				"EventID": 83,
				"EventName": "ChampionKill",
				"EventTime": 1282.561767578125,
				"KillerName": "MoRice651",
				"VictimName": "wezznco"
			},
			{
				"Assisters": [],
				"EventID": 84,
				"EventName": "ChampionKill",
				"EventTime": 1288.7008056640626,
				"KillerName": "MoRice651",
				"VictimName": "Zero Damage ADC"
			},
			{
				"EventID": 85,
				"EventName": "Multikill",
				"EventTime": 1288.7008056640626,
				"KillStreak": 2,
				"KillerName": "MoRice651"
			},
			{
				"Assisters": [],
				"EventID": 86,
				"EventName": "ChampionKill",
				"EventTime": 1289.58984375,
				"KillerName": "MoRice651",
				"VictimName": "ntalice"
			},
			{
				"EventID": 87,
				"EventName": "Multikill",
				"EventTime": 1289.58984375,
				"KillStreak": 3,
				"KillerName": "MoRice651"
			},
			{
				"Assisters": [],
				"EventID": 88,
				"EventName": "ChampionKill",
				"EventTime": 1289.870849609375,
				"KillerName": "MoRice651",
				"VictimName": "mikehpoo"
			},
			{
				"EventID": 89,
				"EventName": "Multikill",
				"EventTime": 1289.870849609375,
				"KillStreak": 4,
				"KillerName": "MoRice651"
			},
			{
				"Acer": "MoRice651",
				"AcingTeam": "CHAOS",
				"EventID": 90,
				"EventName": "Ace",
				"EventTime": 1289.870849609375
			},
			{
				"Assisters": [],
				"EventID": 91,
				"EventName": "ChampionKill",
				"EventTime": 1311.781494140625,
				"KillerName": "MoRice651",
				"VictimName": "piperx"
			},
			{
				"EventID": 92,
				"EventName": "Multikill",
				"EventTime": 1311.781494140625,
				"KillStreak": 5,
				"KillerName": "MoRice651"
			},
			{
				"Acer": "MoRice651",
				"AcingTeam": "CHAOS",
				"EventID": 93,
				"EventName": "Ace",
				"EventTime": 1311.781494140625
			},
			{
				"Assisters": [],
				"EventID": 94,
				"EventName": "ChampionKill",
				"EventTime": 1316.6982421875,
				"KillerName": "piperx",
				"VictimName": "MoRice651"
			},
			{
				"Assisters": [
					"MoRice651"
				],
				"EventID": 95,
				"EventName": "TurretKilled",
				"EventTime": 1321.5599365234376,
				"KillerName": "Japudi",
				"TurretKilled": "Turret_T1_C_01_A"
			},
			{
				"Assisters": [
					"GrantReaper"
				],
				"EventID": 96,
				"EventName": "TurretKilled",
				"EventTime": 1330.56787109375,
				"KillerName": "Japudi",
				"TurretKilled": "Turret_T1_C_02_A"
			},
			{
				"Assisters": [
					"ntalice"
				],
				"EventID": 97,
				"EventName": "ChampionKill",
				"EventTime": 1336.7789306640626,
				"KillerName": "wezznco",
				"VictimName": "GrantReaper"
			},
			{
				"Assisters": [
					"wezznco",
					"ntalice"
				],
				"EventID": 98,
				"EventName": "ChampionKill",
				"EventTime": 1344.691650390625,
				"KillerName": "mikehpoo",
				"VictimName": "Japudi"
			},
			{
				"Assisters": [
					"Zero Damage ADC",
					"GhostBig",
					"Zero Damage ADC"
				],
				"DragonType": "Fire",
				"EventID": 99,
				"EventName": "DragonKill",
				"EventTime": 1392.5128173828126,
				"KillerName": "wezznco",
				"Stolen": "True"
			},
			{
				"Assisters": [
					"wezznco",
					"Zero Damage ADC",
					"ntalice",
					"mikehpoo"
				],
				"EventID": 100,
				"EventName": "ChampionKill",
				"EventTime": 1406.6220703125,
				"KillerName": "piperx",
				"VictimName": "KubiK651"
			},
			{
				"Assisters": [],
				"EventID": 101,
				"EventName": "BaronKill",
				"EventTime": 1408.58642578125,
				"KillerName": "MoRice651",
				"Stolen": "False"
			},
			{
				"Assisters": [
					"GhostBig",
					"KubiK651"
				],
				"EventID": 102,
				"EventName": "ChampionKill",
				"EventTime": 1422.88671875,
				"KillerName": "Japudi",
				"VictimName": "Zero Damage ADC"
			},
			{
				"Assisters": [],
				"EventID": 103,
				"EventName": "ChampionKill",
				"EventTime": 1430.2635498046876,
				"KillerName": "wezznco",
				"VictimName": "Japudi"
			},
			{
				"Assisters": [
					"Japudi"
				],
				"EventID": 104,
				"EventName": "ChampionKill",
				"EventTime": 1437.5771484375,
				"KillerName": "MoRice651",
				"VictimName": "wezznco"
			},
			{
				"Assisters": [
					"wezznco",
					"ntalice"
				],
				"EventID": 105,
				"EventName": "ChampionKill",
				"EventTime": 1453.95751953125,
				"KillerName": "mikehpoo",
				"VictimName": "MoRice651"
			},
			{
				"Assisters": [
					"ntalice",
					"mikehpoo"
				],
				"EventID": 106,
				"EventName": "ChampionKill",
				"EventTime": 1465.33740234375,
				"KillerName": "piperx",
				"VictimName": "GhostBig"
			},
			{
				"EventID": 107,
				"EventName": "InhibRespawningSoon",
				"EventTime": 1467.544189453125,
				"InhibRespawningSoon": "Barracks_T1_L1"
			},
			{
				"Assisters": [
					"piperx",
					"ntalice",
					"mikehpoo"
				],
				"EventID": 108,
				"EventName": "ChampionKill",
				"EventTime": 1477.2740478515626,
				"KillerName": "Zero Damage ADC",
				"VictimName": "GrantReaper"
			},
			{
				"EventID": 109,
				"EventName": "InhibRespawned",
				"EventTime": 1482.58837890625,
				"InhibRespawned": "Barracks_T1_L1"
			},
			{
				"Assisters": [],
				"EventID": 110,
				"EventName": "ChampionKill",
				"EventTime": 1513.5728759765626,
				"KillerName": "Japudi",
				"VictimName": "mikehpoo"
			},
			{
				"Assisters": [
					"GrantReaper",
					"MoRice651",
					"Japudi"
				],
				"EventID": 111,
				"EventName": "ChampionKill",
				"EventTime": 1531.2113037109376,
				"KillerName": "KubiK651",
				"VictimName": "ntalice"
			},
			{
				"EventID": 112,
				"EventName": "InhibRespawningSoon",
				"EventTime": 1532.7734375,
				"InhibRespawningSoon": "Barracks_T1_C1"
			},
			{
				"Assisters": [],
				"EventID": 113,
				"EventName": "ChampionKill",
				"EventTime": 1535.0859375,
				"KillerName": "MoRice651",
				"VictimName": "wezznco"
			},
			{
				"Assisters": [
					"GrantReaper",
					"KubiK651"
				],
				"EventID": 114,
				"EventName": "ChampionKill",
				"EventTime": 1536.937744140625,
				"KillerName": "Japudi",
				"VictimName": "piperx"
			},
			{
				"Assisters": [
					"wezznco",
					"piperx",
					"ntalice"
				],
				"EventID": 115,
				"EventName": "ChampionKill",
				"EventTime": 1539.6915283203126,
				"KillerName": "Zero Damage ADC",
				"VictimName": "KubiK651"
			},
			{
				"Assisters": [],
				"EventID": 116,
				"EventName": "ChampionKill",
				"EventTime": 1540.0699462890626,
				"KillerName": "MoRice651",
				"VictimName": "Zero Damage ADC"
			},
			{
				"EventID": 117,
				"EventName": "Multikill",
				"EventTime": 1540.0699462890626,
				"KillStreak": 2,
				"KillerName": "MoRice651"
			},
			{
				"Acer": "MoRice651",
				"AcingTeam": "CHAOS",
				"EventID": 118,
				"EventName": "Ace",
				"EventTime": 1540.0699462890626
			},
			{
				"Assisters": [
					"piperx",
					"ntalice"
				],
				"EventID": 119,
				"EventName": "ChampionKill",
				"EventTime": 1540.25634765625,
				"KillerName": "Zero Damage ADC",
				"VictimName": "Japudi"
			},
			{
				"EventID": 120,
				"EventName": "Multikill",
				"EventTime": 1540.25634765625,
				"KillStreak": 2,
				"KillerName": "Zero Damage ADC"
			},
			{
				"EventID": 121,
				"EventName": "InhibRespawned",
				"EventTime": 1547.851806640625,
				"InhibRespawned": "Barracks_T1_C1"
			},
			{
				"Assisters": [
					"GrantReaper"
				],
				"EventID": 122,
				"EventName": "InhibKilled",
				"EventTime": 1557.9013671875,
				"InhibKilled": "Barracks_T1_C1",
				"KillerName": "MoRice651"
			},
			{
				"Assisters": [],
				"EventID": 123,
				"EventName": "ChampionKill",
				"EventTime": 1569.3746337890626,
				"KillerName": "mikehpoo",
				"VictimName": "MoRice651"
			},
			{
				"Assisters": [
					"mikehpoo"
				],
				"EventID": 124,
				"EventName": "ChampionKill",
				"EventTime": 1580.808837890625,
				"KillerName": "ntalice",
				"VictimName": "GrantReaper"
			},
			{
				"Assisters": [],
				"EventID": 125,
				"EventName": "TurretKilled",
				"EventTime": 1589.3197021484376,
				"KillerName": "GhostBig",
				"TurretKilled": "Turret_T1_R_02_A"
			},
			{
				"EventID": 126,
				"EventName": "GameEnd",
				"EventTime": 1676.8541259765626,
				"Result": "Lose"
			}
		]
	},
	"gameData": {
		"gameMode": "CLASSIC",
		"gameTime": 1680.267333984375,
		"mapName": "Map11",
		"mapNumber": 11,
		"mapTerrain": "Infernal"
	}
}`

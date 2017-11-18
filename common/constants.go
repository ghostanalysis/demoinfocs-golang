package common

const (
	MaxEditctBits = 11
	IndexMask     = ((1 << MaxEditctBits) - 1)
)

const weaponPrefix = "weapon_"

/*
 */
type (
	RoundMVPReason   byte
	Hitgroup         byte
	RoundEndReason   byte
	Team             byte
	EquipmentElement int
	EquipmentClass   int
)

/*
 */
const (
	MVPReasonMostEliminations RoundMVPReason = iota + 1
	MVPReasonBombDefused
	MVPReasonBombPlanted
)

// MVPReasonStrings maps constant values to strings
var MVPReasonStrings = map[RoundMVPReason]string{
	MVPReasonMostEliminations: "Most Eliminations",
	MVPReasonBombDefused:      "Bomb Defused",
	MVPReasonBombPlanted:      "Bomb Planted",
}

func (c RoundMVPReason) String() string {
	return MVPReasonStrings[c]
}

/*
 */
const (
	HGGeneric  Hitgroup = 0
	HGHead     Hitgroup = 1
	HGChest    Hitgroup = 2
	HGStomach  Hitgroup = 3
	HGLeftArm  Hitgroup = 4
	HGRightArm Hitgroup = 5
	HGLeftLeg  Hitgroup = 6
	HGRightLeg Hitgroup = 7
	HGGear     Hitgroup = 10
)

// HGStrings maps constant values to strings
var HGStrings = map[Hitgroup]string{
	HGGeneric:  "Generic",
	HGHead:     "Head",
	HGChest:    "Chest",
	HGStomach:  "Stomach",
	HGLeftArm:  "Left Arm",
	HGRightArm: "Right Arm",
	HGLeftLeg:  "Left Leg",
	HGRightLeg: "Right leg",
	HGGear:     "Gear",
}

/*
 */
const (
	RERTargetBombed RoundEndReason = iota + 1
	RERVIPEscaped
	RERVIPKilled
	RERTerroristsEscaped
	RERCTStoppedEscape
	RERTerroristsStopped
	RERBombDefused
	RERCTWin
	RERTerroristsWin
	RERDraw
	RERHostagesRescued
	RERTargetSaved
	RERHostagesNotRescued
	RERTerroristsNotEscaped
	RERVIPNotEscaped
	RERGameStart
	RERTerroristsSurrender
	RERCTSurrender
)

// RERStrings maps constant values to strings
var RERStrings = map[RoundEndReason]string{
	RERTargetBombed:         "Target Bombed",
	RERVIPEscaped:           "VIP Escaped",
	RERVIPKilled:            "VIP Killed",
	RERTerroristsEscaped:    "Terrorists Escaped",
	RERCTStoppedEscape:      "CT Stopped Escape",
	RERTerroristsStopped:    "Terrorists Stopped",
	RERBombDefused:          "Bomb Defused",
	RERCTWin:                "CT Win",
	RERTerroristsWin:        "Terrorists Win",
	RERDraw:                 "Draw",
	RERHostagesRescued:      "Hostages Rescued",
	RERTargetSaved:          "Target Saved",
	RERHostagesNotRescued:   "Hostages Not Rescued",
	RERTerroristsNotEscaped: "Terrorists Not Escaped",
	RERVIPNotEscaped:        "VIP Not Escaped",
	RERGameStart:            "Game Start",
	RERTerroristsSurrender:  "Terrorists Surrender",
	RERCTSurrender:          "CT Surrender",
}

func (c RoundEndReason) String() string {
	return RERStrings[c]
}

/*
 */
const (
	TeamUnassigned Team = iota
	TeamSpectators
	TeamTerrorists
	TeamCounterTerrorists
)

// TeamStrings maps constant values to strings
var TeamStrings = map[Team]string{
	TeamUnassigned:        "Unassigned",
	TeamSpectators:        "Spectators",
	TeamTerrorists:        "Terrorists",
	TeamCounterTerrorists: "Counter Terrorists",
}

func (c Team) String() string {
	return TeamStrings[c]
}

/*
 */
const (
	EEUnknown EquipmentElement = 0

	// Pistols
	EEP2000        EquipmentElement = 1
	EEGlock        EquipmentElement = 2
	EEP250         EquipmentElement = 3
	EEDeagle       EquipmentElement = 4
	EEFiveSeven    EquipmentElement = 5
	EEDualBarettas EquipmentElement = 6
	EETec9         EquipmentElement = 7
	EECZ           EquipmentElement = 8
	EEUSP          EquipmentElement = 9
	EERevolver     EquipmentElement = 10

	// SMGs
	EEMP7   EquipmentElement = 101
	EEMP9   EquipmentElement = 102
	EEBizon EquipmentElement = 103
	EEMac10 EquipmentElement = 104
	EEUMP   EquipmentElement = 105
	EEP90   EquipmentElement = 106

	// Heavy
	EESawedOff EquipmentElement = 201
	EENova     EquipmentElement = 202
	EEMag7     EquipmentElement = 203
	EEXM1014   EquipmentElement = 204
	EEM249     EquipmentElement = 205
	EENegev    EquipmentElement = 206

	// Rifles
	EEGallil EquipmentElement = 301
	EEFamas  EquipmentElement = 302
	EEAK47   EquipmentElement = 303
	EEM4A4   EquipmentElement = 304
	EEM4A1   EquipmentElement = 305
	EEScout  EquipmentElement = 306
	EESG556  EquipmentElement = 307
	EEAUG    EquipmentElement = 308
	EEAWP    EquipmentElement = 309
	EEScar20 EquipmentElement = 310
	EEG3SG1  EquipmentElement = 311

	// Equipment
	EEZeus      EquipmentElement = 401
	EEKevlar    EquipmentElement = 402
	EEHelmet    EquipmentElement = 403
	EEBomb      EquipmentElement = 404
	EEKnife     EquipmentElement = 405
	EEDefuseKit EquipmentElement = 406
	EEWorld     EquipmentElement = 407

	// Grenades
	EEDecoy      EquipmentElement = 501
	EEMolotov    EquipmentElement = 502
	EEIncendiary EquipmentElement = 503
	EEFlash      EquipmentElement = 504
	EESmoke      EquipmentElement = 505
	EEHE         EquipmentElement = 506
)

// EquipmentElementStrings maps constant values to strings
var EquipmentElementStrings = map[EquipmentElement]string{
	EEUnknown: "Unknown",

	// Pistols
	EEP2000:        "P2000",
	EEGlock:        "Glock",
	EEP250:         "P250",
	EEDeagle:       "Desert Eagle",
	EEFiveSeven:    "Five-Seven",
	EEDualBarettas: "Dual Barettas",
	EETec9:         "Tec-9",
	EECZ:           "CZ75-Auto",
	EEUSP:          "USP-S",
	EERevolver:     "R8 Revolver",

	// SMGs
	EEMP7:   "MP7",
	EEMP9:   "MP9",
	EEBizon: "PP-Bizon",
	EEMac10: "Mac-10",
	EEUMP:   "UMP-45",
	EEP90:   "P90",

	// Heavy
	EESawedOff: "Sawed-Off",
	EENova:     "Nova",
	EEMag7:     "Mag-7",
	EEXM1014:   "XM1014",
	EEM249:     "M249",
	EENegev:    "Negev",

	// Rifles
	EEGallil: "Galil AR",
	EEFamas:  "Famas",
	EEAK47:   "AK-47",
	EEM4A4:   "M4A4",
	EEM4A1:   "M4A1-S",
	EEScout:  "SSG 08",
	EESG556:  "SG 553",
	EEAUG:    "Aug",
	EEAWP:    "AWP",
	EEScar20: "Scar-20",
	EEG3SG1:  "G3SG1",

	// Equipment
	EEZeus:      "Zeus",
	EEKevlar:    "Kevlar",
	EEHelmet:    "Helmet",
	EEBomb:      "C4",
	EEKnife:     "Knife",
	EEDefuseKit: "Defuse Kit",
	EEWorld:     "World",

	// Grenades
	EEDecoy:      "Decoy",
	EEMolotov:    "Molotov",
	EEIncendiary: "Incendiary",
	EEFlash:      "Flash",
	EESmoke:      "Smoke",
	EEHE:         "HE",
}

func (c EquipmentElement) String() string {
	return EquipmentElementStrings[c]
}

/*
 */
const (
	ECUnknown EquipmentClass = iota
	ECPistols
	ECSMG
	ECHeavy
	ECRifle
	ECEquipment
	ECGrenade
)

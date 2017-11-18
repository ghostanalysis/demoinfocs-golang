// Package common contains common types, constants and functions used over different demoinfocs packages.
// Some constants prefixes:
// MVPReason - the reason why someone got the MVP award.
// HG - HitGroup - where a bullet hit the player.
// EE - EquipmentElement - basically the weapon identifiers.
// RER - RoundEndReason - why the round ended (bomb exploded, defused, time ran out. . .).
// EC - EquipmentClass - type of equipment (pistol, smg, heavy. . .).
package common

import (
	"fmt"
	"os"
	"strings"
)

// MapEquipment creates an EquipmentElement from the name of the weapon / equipment.
func MapEquipment(eqName string) EquipmentElement {
	eqName = strings.TrimPrefix(eqName, weaponPrefix)

	wep := EEUnknown

	if strings.Contains(eqName, "knife") || strings.Contains(eqName, "bayonet") {
		wep = EEKnife
	} else {
		switch eqName {
		case "ak47":
			wep = EEAK47

		case "aug":
			wep = EEAUG

		case "awp":
			wep = EEAWP

		case "bizon":
			wep = EEBizon

		case "c4":
			wep = EEBomb

		case "deagle":
			wep = EEDeagle

		case "decoy":
			fallthrough
		case "decoygrenade":
			fallthrough
		case "decoy_projectile":
			wep = EEDecoy

		case "elite":
			wep = EEDualBarettas

		case "famas":
			wep = EEFamas

		case "fiveseven":
			wep = EEFiveSeven

		case "flashbang":
			wep = EEFlash

		case "g3sg1":
			wep = EEG3SG1

		case "galil":
			fallthrough
		case "galilar":
			wep = EEGallil

		case "glock":
			wep = EEGlock

		case "hegrenade":
			wep = EEHE

		case "hkp2000":
			wep = EEP2000

		case "incgrenade":
			fallthrough
		case "incendiarygrenade":
			wep = EEIncendiary

		case "m249":
			wep = EEM249

		case "m4a1":
			wep = EEM4A4

		case "mac10":
			wep = EEMac10

		case "mag7":
			wep = EEMag7

		case "molotov":
			fallthrough
		case "molotovgrenade":
			fallthrough
		case "molotov_projectile":
			wep = EEMolotov

		case "mp7":
			wep = EEMP7

		case "mp9":
			wep = EEMP9

		case "negev":
			wep = EENegev

		case "nova":
			wep = EENova

		case "p250":
			wep = EEP250

		case "p90":
			wep = EEP90

		case "sawedoff":
			wep = EESawedOff

		case "scar20":
			wep = EEScar20

		case "sg556":
			wep = EESG556

		case "smokegrenade":
			wep = EESmoke

		case "ssg08":
			wep = EEScout

		case "taser":
			wep = EEZeus

		case "tec9":
			wep = EETec9

		case "ump45":
			wep = EEUMP

		case "xm1014":
			wep = EEXM1014

		case "m4a1_silencer":
			fallthrough
		case "m4a1_silencer_off":
			wep = EEM4A1

		case "cz75a":
			wep = EECZ

		case "usp":
			fallthrough
		case "usp_silencer":
			fallthrough
		case "usp_silencer_off":
			wep = EEUSP

		case "world":
			wep = EEWorld

		case "inferno":
			wep = EEIncendiary

		case "revolver":
			wep = EERevolver

		case "vest":
			wep = EEKevlar

		case "vesthelm":
			wep = EEHelmet

		case "defuser":
			wep = EEDefuseKit

		case "sensorgrenade": // Only used in 'Co-op Strike' mode

		case "scar17": //These crash the game when given via give wep_[mp5navy|...], and cannot be purchased ingame.
		case "sg550": //yet the server-classes are networked, so we need to resolve them.
		case "mp5navy":
		case "p228":
		case "scout":
		case "sg552":
		case "tmp":
		case "worldspawn":

		default:
			fmt.Fprintf(os.Stderr, "WARNING: Unknown weapon / equipment %q\n", eqName)
		}
	}
	return wep
}

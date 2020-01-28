// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// LolQueueEligibilityEligibilityRestrictionCode lol queue eligibility eligibility restriction code
// swagger:model LolQueueEligibilityEligibilityRestrictionCode
type LolQueueEligibilityEligibilityRestrictionCode string

const (

	// LolQueueEligibilityEligibilityRestrictionCodeQueueDisabled captures enum value "QueueDisabled"
	LolQueueEligibilityEligibilityRestrictionCodeQueueDisabled LolQueueEligibilityEligibilityRestrictionCode = "QueueDisabled"

	// LolQueueEligibilityEligibilityRestrictionCodeQueueUnsupported captures enum value "QueueUnsupported"
	LolQueueEligibilityEligibilityRestrictionCodeQueueUnsupported LolQueueEligibilityEligibilityRestrictionCode = "QueueUnsupported"

	// LolQueueEligibilityEligibilityRestrictionCodePlayerLevelRestriction captures enum value "PlayerLevelRestriction"
	LolQueueEligibilityEligibilityRestrictionCodePlayerLevelRestriction LolQueueEligibilityEligibilityRestrictionCode = "PlayerLevelRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodePlayerTimedRestriction captures enum value "PlayerTimedRestriction"
	LolQueueEligibilityEligibilityRestrictionCodePlayerTimedRestriction LolQueueEligibilityEligibilityRestrictionCode = "PlayerTimedRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodePlayerBannedRestriction captures enum value "PlayerBannedRestriction"
	LolQueueEligibilityEligibilityRestrictionCodePlayerBannedRestriction LolQueueEligibilityEligibilityRestrictionCode = "PlayerBannedRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodePlayerAvailableChampionRestriction captures enum value "PlayerAvailableChampionRestriction"
	LolQueueEligibilityEligibilityRestrictionCodePlayerAvailableChampionRestriction LolQueueEligibilityEligibilityRestrictionCode = "PlayerAvailableChampionRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodeTeamDivisionRestriction captures enum value "TeamDivisionRestriction"
	LolQueueEligibilityEligibilityRestrictionCodeTeamDivisionRestriction LolQueueEligibilityEligibilityRestrictionCode = "TeamDivisionRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodeTeamMaxSizeRestriction captures enum value "TeamMaxSizeRestriction"
	LolQueueEligibilityEligibilityRestrictionCodeTeamMaxSizeRestriction LolQueueEligibilityEligibilityRestrictionCode = "TeamMaxSizeRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodeTeamMinSizeRestriction captures enum value "TeamMinSizeRestriction"
	LolQueueEligibilityEligibilityRestrictionCodeTeamMinSizeRestriction LolQueueEligibilityEligibilityRestrictionCode = "TeamMinSizeRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodePlayerBingeRestriction captures enum value "PlayerBingeRestriction"
	LolQueueEligibilityEligibilityRestrictionCodePlayerBingeRestriction LolQueueEligibilityEligibilityRestrictionCode = "PlayerBingeRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodePlayerDodgeRestriction captures enum value "PlayerDodgeRestriction"
	LolQueueEligibilityEligibilityRestrictionCodePlayerDodgeRestriction LolQueueEligibilityEligibilityRestrictionCode = "PlayerDodgeRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodePlayerInGameRestriction captures enum value "PlayerInGameRestriction"
	LolQueueEligibilityEligibilityRestrictionCodePlayerInGameRestriction LolQueueEligibilityEligibilityRestrictionCode = "PlayerInGameRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodePlayerLeaverBustedRestriction captures enum value "PlayerLeaverBustedRestriction"
	LolQueueEligibilityEligibilityRestrictionCodePlayerLeaverBustedRestriction LolQueueEligibilityEligibilityRestrictionCode = "PlayerLeaverBustedRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodePlayerLeaverTaintedWarningRestriction captures enum value "PlayerLeaverTaintedWarningRestriction"
	LolQueueEligibilityEligibilityRestrictionCodePlayerLeaverTaintedWarningRestriction LolQueueEligibilityEligibilityRestrictionCode = "PlayerLeaverTaintedWarningRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodePlayerMaxLevelRestriction captures enum value "PlayerMaxLevelRestriction"
	LolQueueEligibilityEligibilityRestrictionCodePlayerMaxLevelRestriction LolQueueEligibilityEligibilityRestrictionCode = "PlayerMaxLevelRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodePlayerMinLevelRestriction captures enum value "PlayerMinLevelRestriction"
	LolQueueEligibilityEligibilityRestrictionCodePlayerMinLevelRestriction LolQueueEligibilityEligibilityRestrictionCode = "PlayerMinLevelRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodePlayerMinorRestriction captures enum value "PlayerMinorRestriction"
	LolQueueEligibilityEligibilityRestrictionCodePlayerMinorRestriction LolQueueEligibilityEligibilityRestrictionCode = "PlayerMinorRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodePlayerTimePlayedRestriction captures enum value "PlayerTimePlayedRestriction"
	LolQueueEligibilityEligibilityRestrictionCodePlayerTimePlayedRestriction LolQueueEligibilityEligibilityRestrictionCode = "PlayerTimePlayedRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodePlayerRankedSuspensionRestriction captures enum value "PlayerRankedSuspensionRestriction"
	LolQueueEligibilityEligibilityRestrictionCodePlayerRankedSuspensionRestriction LolQueueEligibilityEligibilityRestrictionCode = "PlayerRankedSuspensionRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodeTeamHighMMRMaxSizeRestriction captures enum value "TeamHighMMRMaxSizeRestriction"
	LolQueueEligibilityEligibilityRestrictionCodeTeamHighMMRMaxSizeRestriction LolQueueEligibilityEligibilityRestrictionCode = "TeamHighMMRMaxSizeRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodeTeamSizeRestriction captures enum value "TeamSizeRestriction"
	LolQueueEligibilityEligibilityRestrictionCodeTeamSizeRestriction LolQueueEligibilityEligibilityRestrictionCode = "TeamSizeRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodePrerequisiteQueuesNotPlayedRestriction captures enum value "PrerequisiteQueuesNotPlayedRestriction"
	LolQueueEligibilityEligibilityRestrictionCodePrerequisiteQueuesNotPlayedRestriction LolQueueEligibilityEligibilityRestrictionCode = "PrerequisiteQueuesNotPlayedRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodeGameVersionMismatch captures enum value "GameVersionMismatch"
	LolQueueEligibilityEligibilityRestrictionCodeGameVersionMismatch LolQueueEligibilityEligibilityRestrictionCode = "GameVersionMismatch"

	// LolQueueEligibilityEligibilityRestrictionCodeUnknownRestriction captures enum value "UnknownRestriction"
	LolQueueEligibilityEligibilityRestrictionCodeUnknownRestriction LolQueueEligibilityEligibilityRestrictionCode = "UnknownRestriction"

	// LolQueueEligibilityEligibilityRestrictionCodeBanInfoNotAvailable captures enum value "BanInfoNotAvailable"
	LolQueueEligibilityEligibilityRestrictionCodeBanInfoNotAvailable LolQueueEligibilityEligibilityRestrictionCode = "BanInfoNotAvailable"

	// LolQueueEligibilityEligibilityRestrictionCodeMinorInfoNotAvailable captures enum value "MinorInfoNotAvailable"
	LolQueueEligibilityEligibilityRestrictionCodeMinorInfoNotAvailable LolQueueEligibilityEligibilityRestrictionCode = "MinorInfoNotAvailable"

	// LolQueueEligibilityEligibilityRestrictionCodeSummonerInfoNotAvailable captures enum value "SummonerInfoNotAvailable"
	LolQueueEligibilityEligibilityRestrictionCodeSummonerInfoNotAvailable LolQueueEligibilityEligibilityRestrictionCode = "SummonerInfoNotAvailable"

	// LolQueueEligibilityEligibilityRestrictionCodeLeaguesInfoNotAvailable captures enum value "LeaguesInfoNotAvailable"
	LolQueueEligibilityEligibilityRestrictionCodeLeaguesInfoNotAvailable LolQueueEligibilityEligibilityRestrictionCode = "LeaguesInfoNotAvailable"

	// LolQueueEligibilityEligibilityRestrictionCodeInventoryChampsInfoNotAvailable captures enum value "InventoryChampsInfoNotAvailable"
	LolQueueEligibilityEligibilityRestrictionCodeInventoryChampsInfoNotAvailable LolQueueEligibilityEligibilityRestrictionCode = "InventoryChampsInfoNotAvailable"

	// LolQueueEligibilityEligibilityRestrictionCodeInventoryQueuesInfoNotAvailable captures enum value "InventoryQueuesInfoNotAvailable"
	LolQueueEligibilityEligibilityRestrictionCodeInventoryQueuesInfoNotAvailable LolQueueEligibilityEligibilityRestrictionCode = "InventoryQueuesInfoNotAvailable"
)

// for schema
var lolQueueEligibilityEligibilityRestrictionCodeEnum []interface{}

func init() {
	var res []LolQueueEligibilityEligibilityRestrictionCode
	if err := json.Unmarshal([]byte(`["QueueDisabled","QueueUnsupported","PlayerLevelRestriction","PlayerTimedRestriction","PlayerBannedRestriction","PlayerAvailableChampionRestriction","TeamDivisionRestriction","TeamMaxSizeRestriction","TeamMinSizeRestriction","PlayerBingeRestriction","PlayerDodgeRestriction","PlayerInGameRestriction","PlayerLeaverBustedRestriction","PlayerLeaverTaintedWarningRestriction","PlayerMaxLevelRestriction","PlayerMinLevelRestriction","PlayerMinorRestriction","PlayerTimePlayedRestriction","PlayerRankedSuspensionRestriction","TeamHighMMRMaxSizeRestriction","TeamSizeRestriction","PrerequisiteQueuesNotPlayedRestriction","GameVersionMismatch","UnknownRestriction","BanInfoNotAvailable","MinorInfoNotAvailable","SummonerInfoNotAvailable","LeaguesInfoNotAvailable","InventoryChampsInfoNotAvailable","InventoryQueuesInfoNotAvailable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolQueueEligibilityEligibilityRestrictionCodeEnum = append(lolQueueEligibilityEligibilityRestrictionCodeEnum, v)
	}
}

func (m LolQueueEligibilityEligibilityRestrictionCode) validateLolQueueEligibilityEligibilityRestrictionCodeEnum(path, location string, value LolQueueEligibilityEligibilityRestrictionCode) error {
	if err := validate.Enum(path, location, value, lolQueueEligibilityEligibilityRestrictionCodeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol queue eligibility eligibility restriction code
func (m LolQueueEligibilityEligibilityRestrictionCode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolQueueEligibilityEligibilityRestrictionCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

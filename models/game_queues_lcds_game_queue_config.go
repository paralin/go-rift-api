// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GameQueuesLcdsGameQueueConfig game queues lcds game queue config
// swagger:model GameQueuesLcdsGameQueueConfig
type GameQueuesLcdsGameQueueConfig struct {

	// blocked minutes threshold
	BlockedMinutesThreshold int32 `json:"blockedMinutesThreshold,omitempty"`

	// bots can spawn on blue side
	BotsCanSpawnOnBlueSide bool `json:"botsCanSpawnOnBlueSide,omitempty"`

	// cache name
	CacheName string `json:"cacheName,omitempty"`

	// disallow free champions
	DisallowFreeChampions bool `json:"disallowFreeChampions,omitempty"`

	// game mode
	GameMode string `json:"gameMode,omitempty"`

	// game mutators
	GameMutators []string `json:"gameMutators"`

	// game type config Id
	GameTypeConfigID int32 `json:"gameTypeConfigId,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// last toggled off time
	LastToggledOffTime int64 `json:"lastToggledOffTime,omitempty"`

	// last toggled on time
	LastToggledOnTime int64 `json:"lastToggledOnTime,omitempty"`

	// map selection algorithm
	MapSelectionAlgorithm string `json:"mapSelectionAlgorithm,omitempty"`

	// matching throttle config
	MatchingThrottleConfig *GameQueuesLcdsMatchingThrottleConfig `json:"matchingThrottleConfig,omitempty"`

	// max level
	MaxLevel int32 `json:"maxLevel,omitempty"`

	// max summoner level for first win of the day
	MaxSummonerLevelForFirstWinOfTheDay int32 `json:"maxSummonerLevelForFirstWinOfTheDay,omitempty"`

	// maximum participant list size
	MaximumParticipantListSize int32 `json:"maximumParticipantListSize,omitempty"`

	// min level
	MinLevel int32 `json:"minLevel,omitempty"`

	// minimum participant list size
	MinimumParticipantListSize int32 `json:"minimumParticipantListSize,omitempty"`

	// minimum queue dodge delay time
	MinimumQueueDodgeDelayTime int32 `json:"minimumQueueDodgeDelayTime,omitempty"`

	// num players per team
	NumPlayersPerTeam int32 `json:"numPlayersPerTeam,omitempty"`

	// points config key
	PointsConfigKey string `json:"pointsConfigKey,omitempty"`

	// queue bonus key
	QueueBonusKey string `json:"queueBonusKey,omitempty"`

	// queue state
	QueueState string `json:"queueState,omitempty"`

	// queue state string
	QueueStateString string `json:"queueStateString,omitempty"`

	// randomize team sizes
	RandomizeTeamSizes bool `json:"randomizeTeamSizes,omitempty"`

	// ranked
	Ranked bool `json:"ranked,omitempty"`

	// removal from game allowed
	RemovalFromGameAllowed bool `json:"removalFromGameAllowed,omitempty"`

	// removal from game delay minutes
	RemovalFromGameDelayMinutes int32 `json:"removalFromGameDelayMinutes,omitempty"`

	// supported map ids
	SupportedMapIds []int32 `json:"supportedMapIds"`

	// team only
	TeamOnly bool `json:"teamOnly,omitempty"`

	// threshold enabled
	ThresholdEnabled bool `json:"thresholdEnabled,omitempty"`

	// threshold size
	ThresholdSize int64 `json:"thresholdSize,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// type string
	TypeString string `json:"typeString,omitempty"`
}

// Validate validates this game queues lcds game queue config
func (m *GameQueuesLcdsGameQueueConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatchingThrottleConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GameQueuesLcdsGameQueueConfig) validateMatchingThrottleConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.MatchingThrottleConfig) { // not required
		return nil
	}

	if m.MatchingThrottleConfig != nil {
		if err := m.MatchingThrottleConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("matchingThrottleConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GameQueuesLcdsGameQueueConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GameQueuesLcdsGameQueueConfig) UnmarshalBinary(b []byte) error {
	var res GameQueuesLcdsGameQueueConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

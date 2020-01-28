This directory contains resources extracted from the League of Legends client,
and may be subject to additional Riot license agreements.

The spec has the following issues:

- "paths./lol-clash/v1/tournament/{tournamentId}/player-honor-restricted.get.parameters"
  must validate one and only one schema (oneOf). Found none valid
- paths./lol-clash/v1/tournament/{tournamentId}/player-honor-restricted.get.parameters.in in body should be one of [header]
- paths./async/v1/result/{asyncToken}.get.parameters in body shouldn't contain duplicates
- paths./swagger/v1/api-docs/{api}.get.parameters in body shouldn't contain duplicates
- "paths./lol-ranked/v2/tiers.get.parameters" must validate one and only one schema (oneOf). Found none valid
- paths./lol-ranked/v2/tiers.get.parameters.items.$ref in body is a forbidden property
- paths./lol-ranked/v2/tiers.get.parameters.in in body should be one of [header]
- "paths./lol-loadouts/v4/loadouts/scope/{scope}/{scopeItemId}.get.parameters" must validate one and only one schema (oneOf). Found none valid
- paths./lol-loadouts/v4/loadouts/scope/{scope}/{scopeItemId}.get.parameters.in in body should be one of [header]
- paths./async/v1/status/{asyncToken}.delete.parameters in body shouldn't contain duplicates
- paths./async/v1/status/{asyncToken}.get.parameters in body shouldn't contain duplicates


package apiservermsgs

/*
Copyright 2018 Crunchy Data Solutions, Inc.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	crv1 "github.com/crunchydata/postgres-operator/apis/cr/v1"
)

// ShowUpgradeResponse ...
type ShowUpgradeResponse struct {
	UpgradeList crv1.PgupgradeList
	Status
}

// DeleteUpgradeResponse ...
type DeleteUpgradeResponse struct {
	Results []string
	Status
}

// CreateUpgradeRequest ...
type CreateUpgradeRequest struct {
	Args        []string
	Selector    string
	Namespace   string
	CCPImageTag string
	UpgradeType string
}

// CreateUpgradeResponse ...
type CreateUpgradeResponse struct {
	Results []string
	Status
}

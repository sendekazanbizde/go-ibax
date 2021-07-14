/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package model

import "github.com/IBAX-io/go-ibax/packages/converter"

// StateParameter is model
type StateParameter struct {
	ecosystem  int64
	ID         int64  `gorm:"primary_key;not null"`
	Name       string `gorm:"not null;size:100"`
	Value      string `gorm:"not null"`
	Conditions string `gorm:"not null"`
}

// TableName returns name of table
func (sp *StateParameter) TableName() string {
	if sp.ecosystem == 0 {
		sp.ecosystem = 1
	}
	return `1_parameters`
}

// SetTablePrefix is setting table prefix
func (sp *StateParameter) SetTablePrefix(prefix string) *StateParameter {
	sp.ecosystem = converter.StrToInt64(prefix)
	return sp
}

// Get is retrieving model from database
func (sp *StateParameter) Get(transaction *DbTransaction, name string) (bool, error) {
	return isFound(GetDB(transaction).Where("ecosystem = ? and name = ?", sp.ecosystem, name).First(sp))
}

// GetAllStateParameters is returning all state parameters

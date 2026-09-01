package testutil

import "github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"

func NewUserFixture() *domains.User    { return NewUserMock().User }
func NewContactFixture() *domains.Contact { return NewContactMock().Contact }
func NewGroupFixture() *domains.Group  { return NewGroupMock().Group }
func NewCostFixture() *domains.Cost    { return NewCostMock().Cost }

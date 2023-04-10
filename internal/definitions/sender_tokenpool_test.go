// Copyright © 2023 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package definitions

import (
	"context"
	"fmt"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/mocks/assetmocks"
	"github.com/hyperledger/firefly/mocks/broadcastmocks"
	"github.com/hyperledger/firefly/mocks/databasemocks"
	"github.com/hyperledger/firefly/mocks/datamocks"
	"github.com/hyperledger/firefly/mocks/identitymanagermocks"
	"github.com/hyperledger/firefly/mocks/syncasyncmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBroadcastTokenPoolInvalid(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = true

	mdm := ds.data.(*datamocks.Manager)

	pool := &core.TokenPoolDefinition{
		Pool: &core.TokenPool{
			ID:        fftypes.NewUUID(),
			Namespace: "",
			Name:      "",
			Type:      core.TokenTypeNonFungible,
			Locator:   "N1",
			Symbol:    "COIN",
			Published: true,
		},
	}

	err := ds.DefineTokenPool(context.Background(), pool, false)
	assert.Regexp(t, "FF10420", err)

	mdm.AssertExpectations(t)
}

func TestBroadcastTokenPoolInvalidNonMultiparty(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = false

	mdm := ds.data.(*datamocks.Manager)

	pool := &core.TokenPoolDefinition{
		Pool: &core.TokenPool{
			ID:        fftypes.NewUUID(),
			Namespace: "",
			Name:      "",
			Type:      core.TokenTypeNonFungible,
			Locator:   "N1",
			Symbol:    "COIN",
			Published: false,
		},
	}

	err := ds.DefineTokenPool(context.Background(), pool, false)
	assert.Regexp(t, "FF00140", err)

	mdm.AssertExpectations(t)
}

func TestBroadcastTokenPoolPublishNonMultiparty(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = false

	mdm := ds.data.(*datamocks.Manager)

	pool := &core.TokenPoolDefinition{
		Pool: &core.TokenPool{
			ID:        fftypes.NewUUID(),
			Namespace: "",
			Name:      "",
			Type:      core.TokenTypeNonFungible,
			Locator:   "N1",
			Symbol:    "COIN",
			Published: true,
		},
	}

	err := ds.DefineTokenPool(context.Background(), pool, false)
	assert.Regexp(t, "FF10414", err)

	mdm.AssertExpectations(t)
}

func TestBroadcastTokenPoolInvalidNameMultiparty(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = true

	mdm := ds.data.(*datamocks.Manager)

	pool := &core.TokenPoolDefinition{
		Pool: &core.TokenPool{
			ID:        fftypes.NewUUID(),
			Namespace: "",
			Name:      "",
			Type:      core.TokenTypeNonFungible,
			Locator:   "N1",
			Symbol:    "COIN",
			Connector: "connector1",
			Published: true,
		},
	}

	err := ds.DefineTokenPool(context.Background(), pool, false)
	assert.Regexp(t, "FF00140", err)

	mdm.AssertExpectations(t)
}

func TestDefineTokenPoolOk(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = true

	mdm := ds.data.(*datamocks.Manager)
	mim := ds.identity.(*identitymanagermocks.Manager)
	mbm := ds.broadcast.(*broadcastmocks.Manager)
	mms := &syncasyncmocks.Sender{}

	pool := &core.TokenPoolDefinition{
		Pool: &core.TokenPool{
			ID:        fftypes.NewUUID(),
			Namespace: "ns1",
			Name:      "mypool",
			Type:      core.TokenTypeNonFungible,
			Locator:   "N1",
			Symbol:    "COIN",
			Connector: "connector1",
			Published: true,
		},
	}

	mim.On("GetMultipartyRootOrg", ds.ctx).Return(&core.Identity{
		IdentityBase: core.IdentityBase{
			DID: "firefly:org1",
		},
	}, nil)
	mim.On("ResolveInputSigningIdentity", mock.Anything, mock.Anything).Return(nil)
	mbm.On("NewBroadcast", mock.Anything).Return(mms)
	mms.On("Send", ds.ctx).Return(nil)

	err := ds.DefineTokenPool(ds.ctx, pool, false)
	assert.NoError(t, err)

	mdm.AssertExpectations(t)
	mim.AssertExpectations(t)
	mbm.AssertExpectations(t)
	mms.AssertExpectations(t)
}

func TestDefineTokenPoolkONonMultiparty(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = false

	mdm := ds.data.(*datamocks.Manager)
	mdb := ds.database.(*databasemocks.Plugin)

	pool := &core.TokenPool{
		ID:        fftypes.NewUUID(),
		Namespace: "ns1",
		Name:      "mypool",
		Type:      core.TokenTypeNonFungible,
		Locator:   "N1",
		Symbol:    "COIN",
		Connector: "connector1",
		State:     core.TokenPoolStateConfirmed,
		Published: false,
	}
	definition := &core.TokenPoolDefinition{
		Pool: pool,
	}

	mdb.On("GetTokenPoolByID", mock.Anything, mock.Anything, mock.Anything).Return(pool, nil)

	err := ds.DefineTokenPool(context.Background(), definition, false)
	assert.NoError(t, err)

	mdm.AssertExpectations(t)
	mdb.AssertExpectations(t)
}

func TestDefineTokenPoolNonMultipartyTokenPoolFail(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()

	mdm := ds.data.(*datamocks.Manager)
	mbm := ds.broadcast.(*broadcastmocks.Manager)
	mdi := ds.database.(*databasemocks.Plugin)

	pool := &core.TokenPoolDefinition{
		Pool: &core.TokenPool{
			ID:        fftypes.NewUUID(),
			Namespace: "ns1",
			Name:      "mypool",
			Type:      core.TokenTypeNonFungible,
			Locator:   "N1",
			Symbol:    "COIN",
			Connector: "connector1",
			Published: false,
		},
	}

	mdi.On("GetTokenPoolByID", context.Background(), "ns1", pool.Pool.ID).Return(nil, fmt.Errorf("pop"))

	err := ds.DefineTokenPool(context.Background(), pool, false)
	assert.Regexp(t, "pop", err)

	mdm.AssertExpectations(t)
	mbm.AssertExpectations(t)
}

func TestDefineTokenPoolBadName(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = true

	pool := &core.TokenPoolDefinition{
		Pool: &core.TokenPool{
			ID:        fftypes.NewUUID(),
			Namespace: "ns1",
			Name:      "///bad/////",
			Type:      core.TokenTypeNonFungible,
			Locator:   "N1",
			Symbol:    "COIN",
			Connector: "connector1",
			Published: false,
		},
	}

	err := ds.DefineTokenPool(context.Background(), pool, false)
	assert.Regexp(t, "FF00140", err)
}

func TestPublishTokenPool(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = true

	mdi := ds.database.(*databasemocks.Plugin)
	mam := ds.assets.(*assetmocks.Manager)
	mim := ds.identity.(*identitymanagermocks.Manager)
	mbm := ds.broadcast.(*broadcastmocks.Manager)
	mms := &syncasyncmocks.Sender{}

	pool := &core.TokenPool{
		ID:        fftypes.NewUUID(),
		Namespace: "ns1",
		Name:      "pool1",
		Type:      core.TokenTypeNonFungible,
		Locator:   "N1",
		Symbol:    "COIN",
		Connector: "connector1",
		Published: false,
	}

	mam.On("GetTokenPoolByNameOrID", mock.Anything, "pool1").Return(pool, nil)
	mim.On("GetMultipartyRootOrg", context.Background()).Return(&core.Identity{
		IdentityBase: core.IdentityBase{
			DID: "firefly:org1",
		},
	}, nil)
	mim.On("ResolveInputSigningIdentity", mock.Anything, mock.Anything).Return(nil)
	mbm.On("NewBroadcast", mock.Anything).Return(mms)
	mms.On("Send", context.Background()).Return(nil)
	mdi.On("UpsertTokenPool", context.Background(), pool, database.UpsertOptimizationExisting).Return(nil)

	result, err := ds.PublishTokenPool(context.Background(), "pool1", "pool-shared", false)
	assert.NoError(t, err)
	assert.Equal(t, pool, result)
	assert.True(t, pool.Published)

	mdi.AssertExpectations(t)
	mam.AssertExpectations(t)
	mim.AssertExpectations(t)
	mbm.AssertExpectations(t)
	mms.AssertExpectations(t)
}

func TestPublishTokenPoolNonMultiparty(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = false

	_, err := ds.PublishTokenPool(context.Background(), "pool1", "pool-shared", false)
	assert.Regexp(t, "FF10414", err)
}

func TestPublishTokenPoolQueryFail(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = true

	mam := ds.assets.(*assetmocks.Manager)

	mam.On("GetTokenPoolByNameOrID", mock.Anything, "pool1").Return(nil, fmt.Errorf("pop"))

	_, err := ds.PublishTokenPool(context.Background(), "pool1", "pool-shared", false)
	assert.EqualError(t, err, "pop")

	mam.AssertExpectations(t)
}

func TestPublishTokenPoolSendFail(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = true

	mam := ds.assets.(*assetmocks.Manager)
	mim := ds.identity.(*identitymanagermocks.Manager)
	mbm := ds.broadcast.(*broadcastmocks.Manager)
	mms := &syncasyncmocks.Sender{}

	pool := &core.TokenPool{
		ID:        fftypes.NewUUID(),
		Namespace: "ns1",
		Name:      "pool1",
		Type:      core.TokenTypeNonFungible,
		Locator:   "N1",
		Symbol:    "COIN",
		Connector: "connector1",
		Published: false,
	}

	mam.On("GetTokenPoolByNameOrID", mock.Anything, "pool1").Return(pool, nil)
	mim.On("GetMultipartyRootOrg", context.Background()).Return(&core.Identity{
		IdentityBase: core.IdentityBase{
			DID: "firefly:org1",
		},
	}, nil)
	mim.On("ResolveInputSigningIdentity", mock.Anything, mock.Anything).Return(nil)
	mbm.On("NewBroadcast", mock.Anything).Return(mms)
	mms.On("Send", context.Background()).Return(fmt.Errorf("pop"))

	_, err := ds.PublishTokenPool(context.Background(), "pool1", "pool-shared", false)
	assert.EqualError(t, err, "pop")
	assert.True(t, pool.Published)

	mam.AssertExpectations(t)
	mim.AssertExpectations(t)
	mbm.AssertExpectations(t)
	mms.AssertExpectations(t)
}

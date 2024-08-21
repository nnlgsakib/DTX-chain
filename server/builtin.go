package server

import (
	"github.com/Dtx-validator/dtx-node/consensus"
	consensusDev "github.com/Dtx-validator/dtx-node/consensus/dev"
	consensusDummy "github.com/Dtx-validator/dtx-node/consensus/dummy"
	consensusIBFT "github.com/Dtx-validator/dtx-node/consensus/ibft"
	"github.com/Dtx-validator/dtx-node/secrets"
	"github.com/Dtx-validator/dtx-node/secrets/awsssm"
	"github.com/Dtx-validator/dtx-node/secrets/gcpssm"
	"github.com/Dtx-validator/dtx-node/secrets/hashicorpvault"
	"github.com/Dtx-validator/dtx-node/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}

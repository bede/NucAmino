package scorehandler

import (
	a "github.com/hivdb/viralign/types/amino"
	n "github.com/hivdb/viralign/types/nucleic"
)

const (
	GAPINS = true
	GAPDEL = false
)

type ScoreHandler interface {
	IsPositionalIndelScoreSupported() bool
	GetSubstitutionScore(
		/* refPosition */ int,
		/* base1 */ n.NucleicAcid,
		/* base2 */ n.NucleicAcid,
		/* base3 */ n.NucleicAcid,
		/* ref */ a.AminoAcid) int
	GetGapExtensionScore() int
	GetGapOpeningScore() int
	GetConstantIndelCodonScore() int
	GetPositionalIndelCodonScore(
		/* refPosition */ int,
		/* isInsertion */ bool) int
}

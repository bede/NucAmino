package main

//import "github.com/davecgh/go-spew/spew"
import (
	//"./scorehandler/blosum62"
	//a "./types/amino"
	"fmt"
	"time"
)

import (
	"github.com/davecheney/profile"
	"github.com/hivdb/viralign/alignment"
	d "github.com/hivdb/viralign/data"
	s "github.com/hivdb/viralign/scorehandler"
	"github.com/hivdb/viralign/scorehandler/hiv1b"
	n "github.com/hivdb/viralign/types/nucleic"
)

func test(scoreHandler s.ScoreHandler) *alignment.Alignment {
	cmp := alignment.NewAlignment(
		//[]n.NucleicAcid{n.G, n.C, n.G, n.A, n.T, n.G, n.A, n.G, n.C, n.T, n.T, n.G},
		//[]a.AminoAcid{a.M, a.R, a.A, a.C},
		//[]n.NucleicAcid{n.C, n.C, n.T, n.C, n.A, n.A, n.A, n.T, n.C, n.T, n.A, n.G, n.T, n.C, n.T, n.T, n.T, n.G, n.G, n.C},
		//[]a.AminoAcid{a.P, a.Q, a.I, a.T, a.L, a.W},
		//n.ReadString( /*"CCT*/ "CAAATCACTCTTTGGCAACGACCCATCGTCACAATAAAGATAGGGGGGCAGCTAARGGAAGCTCTATTAGATACAGGAGCAGATGATACAGTATTAGAAGATATAAATTTGCCAGGAAGATGGACACCAAAAATKATAGTGGGAATTGGAGGTTTTACCAAAGTAAGACAGTATGATCAGATACCTGTAGAAATTTGTGGACATAAAGCTATAGGTACAGTRTTAGTAGGACCTACACCTGCCAACATAATTGGAAGAAATCTGTTGACYCAGATTGGTTGCACTTTAAATTTT"+"CCNAT*AGTCCTATTGACACTGTACCAGTAAAATTAAAGCCAGGAATGGATGGCCCAAAAGTTAAACAATGGCCATTGACAGAAGAAAAAATAAAAGCATTAGTAGAAATTTGTGCAGAATTGGAASAGGACGGGAAAATTTCAAAAATTGGGCCTGAAAATCCATACAATACTCCAGTATTTGCCATAAAGAAAAAGAACAGYGATAAATGGAGAAAATTAGTAGATTTCAGAGAACTTAATAAGAGAACTCAAGACTTCTGGGAAGTTCAATTAGGAATACCACATCCCGGAGGGTTAAAAAAGAACAAATCAGTAACAGTACTGGATGTGGGTGATGCATATTTTTCARTTCCCTTAGATGAAGACTTCAGGAAGTATACTGCATTTACCATACCTAGTATAAACAATGAGACACCAGGGACTAGATATCAGTACAATGTGCTTCCACAGGGATGGAAAGGATCACCAGCAATATTCCAAAGTAGCATGACAAGAATCTTAGAACCTTTTAGAAAACAGAATCCAGACATAGTTATCTGTCAATAYGTGGATGATTTGTATGTAGGATCTGACTTAGAAATAGAGMAGCATAGAACAAAAGTAGAGGAACTGAGACAACATTTGTGGAAGTGGGGNACACACCAGACAAMAAACATCAGAAAGAACCTCCATTCCTTTGGATGGGTTATGAACTCCATCCTGATAAATGGACA"+"GCTTAATAGTGA"),
		//n.ReadString("CCTCAGATCACTCTTTGGCAACGACCCTTCGTCAAYATAAAGATAGGGGGGCAAACAATAGAAGCTCTATTAGATACAGGAGCAGATGATACAGTATTAGAAGACATAGATTTGCCAGGAAGATGGAAGCCAAAAATGATAGGGGGAATTGGAGGTTTTATCAAAGTAAAACAGTATGATCAGGTACCCATAGAAATCTGTGGACATAAAGTTATAGGTACAGTATTAGTAGGACCTACACCTGTCAACGTAATTGGAAGAAATCTGATGACTCGGATTGGTTGCACTTTAAATTTTC" /*C*/ +"G"+"CATTAGTCCTATTGAMACTGTACCAGTAAAATTAAAGCCAGGAATGGATGGGCCAAAAGTTAAACAATGGCCATTGACAGAAGAAAAAATAAAAGCATTAGAAGAAATTTGTGCAGAATTGGAAAAGGAAGGAAAAATTTCAAAAATTGGGCCTGAAAATCCATACAATACTCCAGTATTTGCCATAAAGAAAAAAGAAAGTAGTAGTGGTAAATGGAGAAAATTAGTAGATTTTAGAGAACTTAATAAGAGAACTCAAGATTTCTGTGAAGTTCAATTAGGAATACCACATCCCGCAGGGTTAAAAAAGAAAAAGTCAGTAACAGTACTGGATGTGGGTGATGCATATTTTTCAGTTCCCTTAGATGAAGACTTCAGGAAGTATACTGCATTTACCATACCTAGTACAAACAATGAGACACCAGGAACTAGGTATCAGTACAATGTGCTTCCACAGGGATGGAAA"+"AA"+"GGATCACCAGCAATATTCCAAGCTAGCATGACAAAAATCTTAGAGCCTTTCAGAAAGCAAAAYCCAGACATAGTTATCTATCAATACATGGATGATTTGTATGTAGGMTCTGACTTAGAAATAGGGCAGCATAGAACAAAAATAGAGGAATTAAGAGAACATCTGCTRAGGTGGGGATTTTACACACCAGACAAAAAACATCAGAAAGAACCTCCATTCGGATGGGCTATGAACTCCATCCTGATAAATGGACAGTGCAGCCTATWGTGCTGCCAGAAAAAGACAGCTGGA" /*C*/ +"A"+"TT"),
		///* 510101 */
		///* AB519485 */
		///* AF315244 */
		///* AF388132 */
		///* AF311159 */
		///* AB873924 */
		/* AF286238 */
		///* GU581690 */
		/* EF116817 */
		/* JX447300 */
		n.ReadString("tttctagatggaatagataaggcccaagaagagcatgaaaaatatcacaacaattggagagcaatggctagtgagtttaatttgccacccatagtagcaaaagaaatagtagccagctgtgataaatgtcagctaaaaggggaagccatgcatggacaagtagactgtagtccaggaatatggcaattagattgtacacatttagaaggaaaagtcatactggtagcagtccacgtagccagtggctacatagaagcagaggttatcccagcagaaacaggacaggaaacagcatactttatactaaagttagcagcacgatggcctgtcaaaataatacatacagacaatggcagtaatttcaccagtactgtagttaaggcagcctgttggtgggcaggtatccaacaagaatttgggattccctacaatccccaaagtcagggagtagtagaatccatgaataagaaattaaagaaaattatagggcaggtaagagatcaagctgagcaccttaagacagcagtacatatggcagtattcattcacaattttaaaagaaaaggggggattggggggtacagtgcagggaaaagaataggaatagacataaatttagcaacagacatacaaactaaagaattacaaaaacaaattataaaaattcaaaattttcgggtttattacagagacagcagagaccctatttggaaaggaccagccaaactactctggaaaggtgaaggggcagtagtaatacaagataacagtgacataaaggtagtaccaaggagaaaagcaaaaatcattagggattatggaaaacagatggcaggtgctgattgtgtggcaggtagacaggat"),
		d.HIV1BSEQ_IN,
		scoreHandler,
	)
	//cmp.CalcScore()
	return cmp
}

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	var cmp *alignment.Alignment
	scoreHandler := hiv1b.NewAsScoreHandler(
		/* gene                */ hiv1b.IN,
		/* indelCodonBonus	   */ 2,
		/* stopCodonPenalty    */ 4,
		/* gapOpenPenalty      */ 10,
		/* gapExtensionPenalty */ 2,
	)
	start := time.Now()
	for i := 0; i < 1; i++ {
		cmp = test(scoreHandler)
	}
	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s\n", elapsed)
	//calculated, all := cmp.GetCalcInfo()
	//fmt.Printf("Calculated nodes: %d (%f)\n", calculated, float64(calculated)/float64(all))
	cmp.GetReport()
	//table := alignment.GetMutationScoreTable()
	//spew.Print(table)
}

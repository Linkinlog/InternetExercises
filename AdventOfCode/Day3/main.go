package main

import (
	"bufio"
	"fmt"
	"strings"
)

func Day3Part1(s string) int {
	var priority int
	var stringHalves [][]string
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		text := scanner.Text()
		var str1, str2 string = text[:len(text)/2], text[len(text)/2:]
		stringHalves = append(stringHalves, []string{str1, str2})
	}
	for _, e := range stringHalves {
		half1, half2 := e[0], e[1]
		for _, j := range half2 {
			if strings.Contains(half1, string(j)) {
				priority += priorityMapper(j)
				break
			}
		}
	}
	return priority
}

func Day3Part2(s string) int {
	var priority int
	var previousLine string
	var matches []string
	var groupedMatches [][]string
	var iterator int = 1
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		text := scanner.Text()
		if previousLine == "" {
			previousLine = text
		} else {
			if iterator != 0 && iterator%3 == 0 {
				for _, e := range matches {
					if strings.Contains(text, string(e)) {
						priority += priorityMapper(rune(e[0]))
					}
				}
				groupedMatches = append(groupedMatches, matches)
				matches = []string{}
				previousLine = ""
			} else {
				dupeLetter := make(map[string]bool, 0)
				for _, e := range previousLine {
					strang := string(e)
					//fmt.Print(strang)
					if strings.Contains(text, strang) {
						if !dupeLetter[strang] {
							matches = append(matches, strang)
							dupeLetter[strang] = true
						}
					}
				}
				previousLine = text
			}
		}
		iterator++
	}
	for _, e := range groupedMatches {
		dupe := make(map[string]bool, 0)
		for _, j := range e {
			if !dupe[j] {
				// priority += priorityMapper(rune(j[0]))
				dupe[j] = true
			}
		}
	}
	return priority
}

func priorityMapper(r rune) int {
	var priorityMap = map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
		'i': 9,
		'j': 10,
		'k': 11,
		'l': 12,
		'm': 13,
		'n': 14,
		'o': 15,
		'p': 16,
		'q': 17,
		'r': 18,
		's': 19,
		't': 20,
		'u': 21,
		'v': 22,
		'w': 23,
		'x': 24,
		'y': 25,
		'z': 26,
		'A': 27,
		'B': 28,
		'C': 29,
		'D': 30,
		'E': 31,
		'F': 32,
		'G': 33,
		'H': 34,
		'I': 35,
		'J': 36,
		'K': 37,
		'L': 38,
		'M': 39,
		'N': 40,
		'O': 41,
		'P': 42,
		'Q': 43,
		'R': 44,
		'S': 45,
		'T': 46,
		'U': 47,
		'V': 48,
		'W': 49,
		'X': 50,
		'Y': 51,
		'Z': 52,
	}
	return priorityMap[r]
}

func main() {
	fmt.Print(Day3Part2(givenInput1))
}

var givenTestInput string = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

var givenInput1 string = `gvNbShZZgQfWdQhdPQmggLTFLwmwjFqjVVgM
CsJnHllcsnnnnJrGRnRwPPLVmFLHLBjFFVHmPT
JlnCtctJnJDcJlDCRpPrSSQWfphzWZfbZSvfNfbS
WjvRSdSQjvpjWzNlnZlNZqCCMzZZ
nJtJsbctPBPwLNcDZNNGLClC
tsFJHBgJwgJbnvSHHWVWHhVhpQ
zRzPhCCSHVZzfGHZ
qBsWBpqBwBcvqqWgdfZrprdggPHHVZ
WWmvwvBbnWmnwvWcbmWWnqNCRSDRRSSjjSDbPJbRjClLhC
rQrznfHHhrHzllzlzTGcJgtJ
jhhjjSZVPWVZvSFtTttLTglgFtvm
dqSCqVWdbDSSVqbVVSqhNdrnpnCnfsnnwfnsRpMpBMrf
rTnvrSSHvHtnDQVDtfmW
ZjCglhcCJschpZbZgbtGmRtbGmwtQtbmtN
FFZCghFJhlslFpjcdFTPLHHmrqTMTdzBLHLH
WzDfrDwwDCCDMnfDHHJjTHTtNdngdHtQ
FmScGPPPPmpspchbGHtRsjdJJJWJWNttgH
bvhFlVhPchPvqWLwqLWqvq
RcchVlCCvmhDRjNJJJjbjllZlJ
wPwFGFMPfpdtqbpwFFfGPQZTWTZjtzQWTBBQJzZjZZ
fbwPqqMPwdgvmDVDhmRCgR
SSQggtQVQQgQGmGVthVnWDGjCBLLzdDNCCjDGWCW
RcHcHbFqbPMZRFTvHFFFlcZZBNBtNjjLLjDdjLjDqWCpNLpj
sPPclZcsZtMPfVsnnwwVJhmQ
jlrwpVPjMwfzZfhfwddH
DbQBGgBJPPBBHfPZ
ggGSQTnTgQGbSRQRPrNlsjVjMrpmMLnrrW
MzNNbMSrZNSSvGWWznwGQlTG
ZCFCJHqJFjqTTqngPwPG
LCJLJCCFtfjsJCMbZSMSbbcDdBfd
LCdjljfJJBfLDCCdJqGqsGGtmBsWtbGmSS
gpgchRcPgRRQNZcpNhgNPzzmGrmrGwtSSswqrSGbbTzt
QgQvNppRPQcggZNccchvNZgCjlnsDLDflDdfjLfMDVsMLv
jfSfTWfwTJffQQNwTCHnGGGgHbgqGFvF
LmspctdsLDlBLPmFFgqbqFFbqlnSRS
PSZZsctsDDtdBmzBLBSmpJjMzWwjVfwJVMrNfwrfzf
zFZqbRNRHNHhqHFqzNnzzqDQVFMggMgDpmJMTgpVTQFJ
SwLVsSWBfStDGMfDTDGGGT
vlCBVsSvCRrHRjCqrn
ndGVddlVdrcGlGcljdVGjCgfffPCCLfCMvqWvPnqgn
FNtRZBDDSNSNStDhFRNQtLgCMTCBqvLMMTfJqBLPLP
bwQwZhSDFhZbFSNccwzwHmHHjHMwsd
BrSDgqrgWzWDwJBzMDWBLjpmVLFfFPLVJjVVPLFV
cNnRlvNRvQTcnZwmfQGPmfVFPP
dsCNllvTdnntTlTsHcdCRtNRBDWBDWDrWBhDBqMBDqbzWwCr
stLcchcTwRcsVTtftVVthchqmmCCFvWmmgSgQFCwGvWSFCrr
PljDbZbzjllHdQNWHMFGrMMHMFrmmM
njjjzbPdqRcRQQns
BTBLwzmqWNbpzqNzLppLBnhZJcJSHThhnnJZHHSQnr
RfQgDjgsjVfRFDdvddtvhhhJCnrSnMZCcMSCSJ
QRtjtRQQlzlqqzzzBW
FnzllplJMntnzptLFzsjTVSQcSjSjQFFmdQQ
wgwrBrwGHPGBqHrDcpmjddSQTmTGVmmG
DBhhNBrCrMvNvzptpN
LgSbLLRJQSzLCgJRJhWCCzRBdZdcMfcHsMdFFjfPjsCHBc
nrVTnmrrrlnGGmrlVwVZPFHmdsZHBmjMcdjfcZ
wrnTlNpvVVrqFrtnwGqwwrTpSgSQQLbQRzQSQRJbQpDbQz
pgvmcpVcpwJppwgwvMWfHbRGRHDBJbWGfGGb
hQqFqQNhqdCPNhTQCCFbbjmBHBbBRjGHBDBGFj
TtdCCNlPCmssnqgvzzvvtVrgzzrc
JltddTSgtLLvgvdldgvWPZPPSjWMjCfWNjNWPG
FrwpbwhbmfjGRpGPNM
cmBrBcBnmFBQqcFbscBhQfttznJgzlLLDlnglzLLlg
mDhRsDzWZzGhhhWpZwzwGdfbFsfFTQTLcNFfFrNvvcbT
ngMHjjgVgqHnPlnqtnClMHbfCJFbcLFFJFLbrQNbFQcQ
SlqHqVVqPnBHSHHljlMndBdZQphRhdGpzWmWWwBh
hVNhcfbSSNMfjjMlqRCqpdzpRJdnhl
sBHPPwQwttBBssLCHzDBlDDmRmpJdqJJJd
sHHwtsHTWwtZPHtZbbVSgMbGCNjNvN
SgMtSMVChFBHBFdCwhdqPcPsbRVQqsblDbRmlc
jlrLWfvWvzWWzmcQcqcNbmNDjR
JpLznrnfGrrfnWWzTpfvzvdHHCgSBwdlTMhCTMFhBBwg
fptnFPLsttddDdRhQShQzC
BqGHGlqmmGmlJmNmqlbbCzSCQHDzSjHhfSShjfVh
fbrrMqGGBlNMGNqrqglmrLZwFcwgFwsPcscctpWLgn
TCMMTtWBqCMMMWWMSWhNFhlNBNlNHwNDsvhw
frfdrbZJrVLdZLbnJcrmvhlhHshlgsZsHwwghZ
fLRVvcvRVpfbcfVVJTqpWCCSqCCTTtjjPj
vRCGzdTtvdBCTzCdmVsQjnVttngthgnQHs
rLNZlrwrjrFLMsggsQpsQngqMq
NfjZWcbffDrLZfNLbDczvmcTvvGCmCBBmSdP
DSqdzrHgJtSHMgvmnNGdcmCnvssN
ZwQLlVlRVVWZFZhbJZVjRJscvBGcpssLpBcnGpBsmpBs
TwwjjQJTjlhwMfTDgMrfgtqg
HhvVhgddvdvqTqVqHQgjgmCPWGrcWsGWPcvlsWrWmC
lbJBBFbMlJFRCmsRmRPtCtwm
FMLLpDLnFpFJbDfLnZZSzhqHVdVQjgSdlQTZ
mvBrszzsrrrGsBDvBvszDRQDhjCwWdLhJdCjjWZdqqdZdW
gMlfFVccTGMGPMSCdcLLdqhjLJWWdZ
VfggggPVltVSlSlHfgFtTpfRGmRrrrvQpQsrGmnvnQbzmz
HjtMgWbHBtbtcggVcHwMwcdzpZFGzGZFpvFFZZWnnpFG
RNNrfmSrSSzdGBzFfJFq
mNsSrrTPPRshmsBClRClPmSccVgQVHjHwtHwTMbTVwMHgg
pZmmLpQLDJBCPCssJTsCTc
SwlLwWvSwNSSSMSMfWLvvVdPFzjfTscdjcCsjzCPddFc
vbVMqGbLSVMhDQnRRQnBqH
lNmNwlmlbhhfFNgpJLnHfTHfTdnTML
BGvCwVGVPSDWDMMJTHZMpndH
rVcBrPSPCwWqPwGCBCSrqlNNQFhsQhtlhhhgNQtb
HqZCQRQgWjpmZHRHqQjttGGJQNQtwQGGGbGrJz
SlddfFVMLncbJJMzMzzMwD
ndSznsdVHsqgpZvH
sNStMtNtDztrvrRGPRRRRrNzqcchgJJLqnBScBcJThJggJBT
QCWWZbCbCVjVbFVFbVwbnhqLqhghBLwwTqJJcpgTBL
ZlHjfHHFWFjCdbHlQFnfdnZFMzPzrrlMrtrrGRMsDPrGvGPz
NnDHhNNldhNTdHllpptCRtftzmGBGmVRVGRH
wgZgZJLBwqgLSWrBvZVVzmffWzRmGsGmmssm
bgqPBBLbPbrrZBjjlpjccjDdDchFTNlc
WRWNNRWNfLZtsmtSWrtLmWgTcMcPrpMTcVMMPPBBqjcg
QJJDFbpnpJdlcwcgBTqlgPwc
GJFHJGhHnpWWShpsRS
DHDZvDQWdHJMHlJf
RhhnGLnBtGrnLjwwJTTMTdQccllSMQ
rmFrrBGQLQFLjbFnBhQLhrRGCPVDWzgmpsWvWWDzppWpzCvC
bsrJrWgWJvQRqbbRDNGGZDlwNfNfwS
ttCHThhhcBVpTBcdhdpfNGDSQLwDLfDGfQllCG
jFjzcFhzQTTjHFvJsrPPnnFgPb
SZggDgNLGCQzSgRsljscPvPPbHsBNc
TtFfTpMthzhFFTpdMMJPfjPPWjBbHbjPlPPBjj
dpdttMMTJttzMFhqFwGLgGRSrrQGSZrwRrDQ
JsJntWmghjgJTpfgFCczlztzrDZDtDzC
QddBBdBSPPPHBSVHVHvNdPVRClvlcclCDLCLFRFrrDLcZc
NVwBVSNlSBNHngnngsJWhwWh
hQVWJGWJlJQDGJHQWHpVWJVHmfPGPcjfdLdLSrcCmfjPjCPm
wgwnRwMvNvBvFZMngNvFZCjdSSfcbPLjCbcdCbSf
wgnRvMMvjMsTwngTgnnDVDWJDplHWhlhQWlWDT
WGswWFGsBFwHvjnbnnJjbG
gLDrSrPpLfmmDqqbbJZsjvmvRv
DTldTpTsDLgSsLpfPTslVdVVdNWVzcNwtWdWhzVt
sttWrhWZsLVtJVZrWhgbCPQQbDNVCHglnnlN
vvRRRMqpmTjwFqmfjFjmBMwNbCwQNgHgDnbNnHQHQCDg
THpvmBGfzWGzSzzS
nmZwwfPmNggwgPgNmdvPPhTRHTHBJTsvSP
WDccMzVQzGWcrllLdsGTvTsJJBLTdL
pWtzQrzDlrCpcQztlzCCCpCVFFFwtgmfbjmbqwZjZfmFwfdw
bmMNNMVSRWBWCSmVRdMRmMnvZZscvvszGcnZClpcGsvZ
gfLfDDtDgjgDtWTjFQjHsvvpZqqplfvvcllcnnvG
TFPjDtWPNPwBBBNB
QcTCHcTwdpQgcWzWwvMPRzSGwW
sDNjfhVmnlVffmLNfLbMPPDbqGtMSRWtDMtP
fZjJJshhhLNnrCJScHSccTQQ
qfhBhLLjvMqWWBWjsQGrdrGzpszpGrqd
JTSDHFtlTDbHTcFcDbTHzmrNRdRRbNzrmRRnRpnQ
PHcltTgJDlplgptScCfCgWMVfvfZvhMMBV
rttDPglHZZDDDPHgZtgWBswdsTfpfdplLqfsFfFdfd
cRVvpcmpmRSCGGCcMQmQFqNNqdVdqTbLwNbdFfFs
zhGjRhSvcMtnrpZjHZBg
STNrrzVdTJwsBggwFgBN
WbtMWJJpwpDsLBfP
cbvncCCmWtCGMcZHVVSlJVQrzvzVqQ
MtPVBHVbBBTwbvWgRgvF
LhZjZpnnMRwWfhSh
LQcQCLCCCcnmnqZpcrpnrjQlsBMHddtMGGzrJHPPsHsNPPJz
NbZtmZDmNNDBHBhDNtNGcvpmgLmmFrvprgrFGv
djMjffMdqnVVTdMfTPnQSFScFFVSLLFrvccvcllL
PPdPwMQwTMWDwBDZLHZN
DqpJFdhtmdSwpqSncljHsnfjqlbzlc
rQVZLGMWPCNrCGMwNVMcbzsVgfHHgfncfscssH
NBCrPLwMTrrMBLQLwCFhpFBDFdFFtdvDFvpD
QRwGBdGwqRTBMWRMMzBqQHHhngspPFFbHbzgbhZLLP
rCNmDjvrCFpvLZsgPb
lDrmpcfCtVDtCjlfNfQQqVdJBQWRVVRMMMJQ
zcgjlRcJCFtlnTvppn
BLSHDMVPVPHVHwFwMhTvgnFZMh
DsLSDdmPVNVdGgrzCCdbjc
tjjRbNQtvJVVsvqsvdTsGqBp
wlFLPZPzzzrzwZPmZfZPlrmpLDCpWDWCSTCpppDDBssdWG
rwwcFFghFmgTcNRJNnnNJtHV
RLhbzMhccqLzdjghdFddNWPN
DmsTZfvrfJLPPWwrdFgw
LsmDtQDfLLGbGQbSzcpM
qwhmmHlHmlwChcCWlpPCBzjPVdBzLsdLBVsT
DnJDMqRDsPsjRVQR
tSGSnNNSbrJJGfDSSfbchpvptFwmpcvlHqZlcF
rdQdBRPrQBBnfdBbzLLgLgmbzSLCPW
TcRTcFDTcjvswMWzzgLSzwwSLJbL
MNTRMNGTsFMMqqqcMDTTnfpBhQhNpHZBZHBfrptQ
ZZFwcQGwRGLTGLTl
rJtMrqBVvhBCnlCn
WWWMttbqqzfPWqWJVzmDQwFQlwZwjQfjgFcZ
RhgmdBbLLmRvhGZwhZZJDwGq
ldPpTtTPtVPpfTGzJQFGFqJsFG
MPpCnrlCVfPtrPVtntllnbdHSbnBvvgdgHgnWvSv
nbnfjQbQZBqBvjQdVFGHhsQhsFsFzGpD
RSJSTTcCVTTMlSNcSTRRwTccpwDFDDzFhHhtwppssFGGpsDp
clcTNJlLMgrcgRvBjfWVfWfBrrZn
LHQdpQLQDRcDBQccVQpnNRhllzMtRqzRztTNnt
mrvZPZvZvZbPSFJrSrggFvrJzbhlzttTlNtqtsMtTntntdqq
GPfJPJwfZFrmFvrPmgHHQLQWdHQBwWWCHVVC
TvTTLfflhZmwZLvtJmhdFhMpnWrpnbJcWMnpHzMrzbcn
VqRRNGCBjRmqPGqDzbHMcMccWnzngHnP
jNqNBjsDsDRCNVsmBRNQNRSSSddhZvLhddZhvhwvLTZwhSfT
vjpJvpgpNwjDWvvJLhFLLhNRTmPTRPRF
HHbHVrMHMmqzHlGmtGqMtdLCLQthQQPCSPTSFLLPRT
bZZzbHszMrqzGmrHMBzbHbqJjffcDgwWjDJpnjjcscwgnw
fBHDMNhhHMrBDBrfQqfwwvvLvfjLZZ
lmppmGJgPPpddJVdnwsQjLttqtjZTpcsvT
RndRmlGWJgdRnWQJQgWdQPlFFbhhbFbHbHMhhBbRDRFrrS
tVhwlMdMWlhlZMZMlzWrtRpBNpLPpmnmCPQBLppPPPBz
GvGTsgjcTLRQHmTQ
gGRGbfbgbcSbbcRggvfcwlZVhVMfltddVtrlZdhW
PwjPHwWssqrHqCqprpCPjwrrQnnDbdnQnGQHnbZbGQZbBQdD
WvNtLhfLndnBNbJJ
fvhvvWfTmtzTmjpsVVqsPzMjjq
GjHMMTMfZbnqHnJNCN
rlpQpthdlQbnlJqRzqVL
spphwcpBvcvdvFTqPDFfwwPZTG
mQfqmtpfBHJCHdlMDB
cbbPPsZqjqcSvVsCrDllsrdzrzzD
TSbVTvShqnVjPTPbggSPbNFQWGtfNLtFfWgGmGgGWR
HJtwgJJwJrcjRRjRnwhVdrVbmGhPmNNrGhdV
fWFvTDssTDMCBCWsvszWWMCQhZhLLNvNNbhVGGZdRZRdmhmd
zCTBzWzDDCDCfFCRqCJHcJJctjJSgHpqHqJg
dgnwgbcwVGbgDWnQbjdgDnBRBBNNqRLBNMLFRMRlqNsc
fTZHmTmzJMmSSCZChZvZTHRlBNBFqtqftlqssqNqftts
JTHJvhzCHpmPCSSTzPpPzJHgQVdQQjQdQwjgdpgdDMGgQM
qPGmWLJWBRRdvqvvNq
TnZHsTZsZsZZbDtbrNrdnDft
SShhQQZCHQCSTVlllVVQCjQWNJGFWpWgFpmmVmNgLgFpGw
TsSmttWMVstNfbMfBcpnFchFwFwZNrrwFp
DWWCGlHWdhddwpcrjh
gQvPqglRCHCHCPCPCJbbtSVVWtTVVJvTbs
ddHSZQHDStZStLtsdDpbCbBPcjGGCqCMBGdNPb
hWhvmglVznzglgzvzfznwfnfscswPjbMBMjNGBMPGMcjcPwP
sJWsJfVflWsnhrltrZLFDDRRFTSSLT
qqCqLNjnPPLZPjqnDjLNgnNrcGwtgddGrhprBJhhJtBGJG
lTMDsDHmlRRbQQbVfMHFmMBJhwcTddrcpBBGdGGJcJdB
lHFfQFssFFFMQRVWQVlssLnDSLLPvZZPqWCLWSqjvS
pJPZRpQtpPQpGPqWbqlWTdLVLJbT
nnnjCwgrgcCFBfnggBCsMTVbVsWVWdCTTdbb
cfhcnwnrcggrncrwcjBDcGHmNPRRGRZZQRRdDZzmzH
NqTrrzLpTNdLLJBBcBGGZGpFHB
sWbgfHSfSgVgZcmGtStZwBFG
VjQjjPjhbbfsCbbDQCgsDrRhzMdlNzHTRdRhqdNTrq
GlqfPdvdBqPgfqDgFMsMVDppbsNJNpMs
WnmmWCTCjZnmzRjTZRRnFBMNpbZZpFJMSFSNFSbS
TRRrnLBmzjttRRCmTRjjwLlvdhQvdrhqqlGdPqGggPlQ
McjMPgPgGPHJWjhPNgPqQbpQSfqffnWVnnnpnZ
BwwDmBTLRBBLzBTBTvtVqlfpqfbQqpqZbnQFlt
rDsDBsmvRTBRzCTZsTmCsCrLPgNcPHNhPccNHHrjMNrgdNdP
GqmLFDrcglvQGZwwzj
nlHbSNBhhslJfsBfvdZHdCvRzdCzppzR
JWShMfWMMthSJBBthJsgtcTTcFqcTrmTLrrTlr
CsBsShBWsBWhvFJWCSsJpbTZdTbdBRgHffRlHHHBgH
tqmnwwwDmVwfzZblzTRqzQ
DGtDjGMcntMGnnGhhjJJJsCpFRRSJJ
GQVVcGgFGcSFvfcFfJVnQmdbTwMLwTTnbnQMws
ZqZrDWDtNPPHRNPzGCdndLLMLwMWhwbmLddn
DRqRHCHlCtDrqtNqDplfpvGFjVjvBSFFcj
zfSfdGPdMVHHdcMThhpwqqrwrqJhcF
vlDNTWDDWWnngNQZtgwLrJZqLFZLFJpFJjJJ
NCtWmnvBmDnWQVTmTGSRmmRTPz
LLvppVLDsppGMGCLCCwNmJMNrNHrNmNNmHFjrh
lZZfvfgRPZQWNNmlWlFNHl
nnRtndRbcPttvTpswLTs
trtzWfszNPlhPlgrWglhZjSLDBvTfSDTmSvmLvfS
QMdVVGBQMTJDvSjdHJ
nBppGVVnQQgzgPsPzzhp
fVQzVrQVtVzHwjtMTjcTCjFgcDCGCd
JslDPWLbLbZlJBMgBTBPdcFTFd
JNJmZsDLnLDWmsVqSrHqnHqQwQSq
mLMZRMRmZLmHLLjNshrrJjRsNNsj
PnqnBVBPcpWrWfWzDhDsHW
qpBQgcQncqpBcVgtclHqcSMTLZZQvLLZdLdwCmLMSd
PZgMwZMjPgBCWgZFgmBGSTtwcccztTzbnGDcNT
vJLlvvlfdsVVslVHrrpLfRlDcSScmbSnnRGDnbDTSbSSDt
fvVHdLfLQljmmqqMQPhQ
DZRstNGrDWGRMlzLHLMpNSpp
bPgCrCvgnrCPgwplbMSjfHzHfj
CgnFnBVrFFVmdVggmPPVTTcZBGhRqGtcZqqJhGRGssWs
WrrCGrDlWhBBZtHB
ggdcnQpcPRSpzfvJBTfPJvtP
gSNnSBFnSMBNjpMssDmCrrGmCLqq
szZGZGGwNdVtpwvbzptzVBQQLFclTJQFlHJNcFBRRJ
CTrqSfWThWDjJCcLQRRRFLlQ
rggqfMqPWnDqnVbZvTtPZpssVw
GJHqPPcJnhthNHnL
rzzlVDVTQDzSSdsSlzpBspssLhmZNjhZWjjffWBLZhgBmnjt
TSTQTsnTbDsMvCGwcFvbJR
PPjtzdnCnSjSthPjWJgJWjgMNBsNrT
bZHfLLhvhpbfvLhHfQbqpNJTWWNsqNTrTNrgWN
ZFbZZFvlRvZwhQwLDSdGnmCGmnlmdznD
RFRRhzzsRsszRlhbrhgBcbTfHgBT
pJtNmSgwZNtSJLgSqHcbBDDTHPfrSqTc
ZNWdpNCJtLZpQJLLpmLtZNFGjVgFnsngRGllVsjjCnvs
sJVRJmmmwwlmzcszMzjRjJVzspCZBsZTTFpTTZvSZTFBFbtZ
GhNDgNWnGGqDvFSCZvbtFpvg
fWrCNNffGnhWdrPhWjwjzRlJmcjRcPzlmV
FdncmqmgcZztLWFvFFvjWLlT
RsMMJVpfhJVsDVMfJpvlTWTvhjQSQvCbLCSC
jDBjRrsVfNJgzdZrZwtmZt
THDDQLZTGQQLQDSSTBCZSZHfFrvvlVRRlPVFfVrvqnvnnH
VbWWcgbNpjgPrRvrNfqnzv
swggwpJWwsWscJsDSLTJTmSVZJTBDZ
wGsznPGGBHdzHhtHDD
WMVSWqVmrSBMpvWVMFtdClhDtqDDHhdtQt
RgWVRcmgnPPnjJBg
BmfPBHBFswLLjQQnLCnzzW
SRdvvRSGrdlRSvTgRrMVCjVWCZzMnVZhWzTj
vlSRJrRJvbGGrSSJgWvlScptFmwbmfHHBfsNFmBbfsFm`

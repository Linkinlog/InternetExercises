package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"whole shebang": {
			input: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`,
			want: 6440,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := part1(strings.Split(tt.input, "\n")); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"whole shebang": {
			input: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`,
			want: 5905,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := part2(strings.Split(tt.input, "\n")); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	tests := map[string]struct {
		input string
		want  Game
	}{
		"whole shebang": {
			input: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`,
			want: Game{
				hands: []*hand{
					{
						cards: []*card{
							{kind: "3"},
							{kind: "2"},
							{kind: "T"},
							{kind: "3"},
							{kind: "K"},
						},
						bet: 765,
					},
					{
						cards: []*card{
							{kind: "T"},
							{kind: "5"},
							{kind: "5"},
							{kind: "J"},
							{kind: "5"},
						},
						bet: 684,
					},
					{
						cards: []*card{
							{kind: "K"},
							{kind: "K"},
							{kind: "6"},
							{kind: "7"},
							{kind: "7"},
						},
						bet: 28,
					},
					{
						cards: []*card{
							{kind: "K"},
							{kind: "T"},
							{kind: "J"},
							{kind: "J"},
							{kind: "T"},
						},
						bet: 220,
					},
					{
						cards: []*card{
							{kind: "Q"},
							{kind: "Q"},
							{kind: "Q"},
							{kind: "J"},
							{kind: "A"},
						},
						bet: 483,
					},
				},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := parseInput(strings.Split(tt.input, "\n"))
			if len(got.hands) != len(tt.want.hands) {
				t.Fatalf("len(go.hands) = %v, want %v", len(got.hands), len(tt.want.hands))
			}
			for i := range got.hands {
				if got.hands[i].bet != tt.want.hands[i].bet {
					t.Fatalf("got.hands[i].bet = %v, want %v", got.hands[i].bet, tt.want.hands[i].bet)
				}
				if len(got.hands[i].cards) != len(tt.want.hands[i].cards) {
					t.Fatalf("got.hands[i].cards = %v, want %v", got.hands[i].cards, tt.want.hands[i].cards)
				}
				for j := range got.hands[i].cards {
					if got.hands[i].cards[j].kind != tt.want.hands[i].cards[j].kind {
						t.Fatalf("got.hands[i].cards[j].kind = %v, want %v", got.hands[i].cards[j].kind, tt.want.hands[i].cards[j].kind)
					}
				}
			}
		})
	}
}

func TestCardStrength(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"2": {
			input: "2",
			want:  2,
		},
		"3": {
			input: "3",
			want:  3,
		},
		"4": {
			input: "4",
			want:  4,
		},
		"5": {
			input: "5",
			want:  5,
		},
		"6": {
			input: "6",
			want:  6,
		},
		"7": {
			input: "7",
			want:  7,
		},
		"8": {
			input: "8",
			want:  8,
		},
		"9": {
			input: "9",
			want:  9,
		},
		"T": {
			input: "T",
			want:  10,
		},
		"J": {
			input: "J",
			want:  1,
		},
		"Q": {
			input: "Q",
			want:  11,
		},
		"K": {
			input: "K",
			want:  12,
		},
		"A": {
			input: "A",
			want:  13,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := card{kind: tt.input}.strength()
			if got != tt.want {
				t.Errorf("card.strength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetType(t *testing.T) {
	tests := map[string]struct {
		input string
		want  handType
	}{
		"five of a kind": {
			input: "JJJJJ",
			want:  fiveOfAKind,
		},
		"four of a kind": {
			input: "JJJJ5",
			want:  fourOfAKind,
		},
		"full house/three of a kind + two of a kind(same thing)": {
			input: "KKKQQ",
			want:  fullHouse,
		},
		"three of a kind": {
			input: "JJJ56",
			want:  threeOfAKind,
		},
		"two pair": {
			input: "KKQQ5",
			want:  twoPair,
		},
		"one pair": {
			input: "KKJ56",
			want:  onePair,
		},
		"high card": {
			input: "TJQ56",
			want:  highCard,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			cards := []*card{}
			for _, strCard := range strings.Split(tt.input, "") {
				cards = append(cards, &card{kind: strCard})
			}
			hand := hand{cards: cards}
			if got := hand.getType(); got != tt.want {
				t.Errorf("() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortHandsByRank(t *testing.T) {
	tests := map[string]struct {
		input []*hand
		want  []*hand
	}{
		"two hands": {
			input: []*hand{
				{
					cards: []*card{
						{kind: "T"},
						{kind: "5"},
						{kind: "5"},
						{kind: "J"},
						{kind: "5"},
					},
					bet: 684,
				},
				{
					cards: []*card{
						{kind: "3"},
						{kind: "2"},
						{kind: "T"},
						{kind: "3"},
						{kind: "K"},
					},
					bet: 765,
				},
			},
			want: []*hand{
				{
					cards: []*card{
						{kind: "3"},
						{kind: "2"},
						{kind: "T"},
						{kind: "3"},
						{kind: "K"},
					},
					bet: 765,
				},
				{
					cards: []*card{
						{kind: "T"},
						{kind: "5"},
						{kind: "5"},
						{kind: "J"},
						{kind: "5"},
					},
					bet: 684,
				},
			},
		},
		"whole shebang": {
			input: []*hand{
				{
					cards: []*card{
						{kind: "3"},
						{kind: "2"},
						{kind: "T"},
						{kind: "3"},
						{kind: "K"},
					},
					bet: 765,
				},
				{
					cards: []*card{
						{kind: "T"},
						{kind: "5"},
						{kind: "5"},
						{kind: "J"},
						{kind: "5"},
					},
					bet: 684,
				},
				{
					cards: []*card{
						{kind: "K"},
						{kind: "K"},
						{kind: "6"},
						{kind: "7"},
						{kind: "7"},
					},
					bet: 28,
				},
				{
					cards: []*card{
						{kind: "K"},
						{kind: "T"},
						{kind: "J"},
						{kind: "J"},
						{kind: "T"},
					},
					bet: 220,
				},
				{
					cards: []*card{
						{kind: "Q"},
						{kind: "Q"},
						{kind: "Q"},
						{kind: "J"},
						{kind: "A"},
					},
					bet: 483,
				},
			},
			want: []*hand{
				{
					cards: []*card{
						{kind: "3"},
						{kind: "2"},
						{kind: "T"},
						{kind: "3"},
						{kind: "K"},
					},
					bet: 765,
				},
				{
					cards: []*card{
						{kind: "K"},
						{kind: "T"},
						{kind: "J"},
						{kind: "J"},
						{kind: "T"},
					},
					bet: 220,
				},
				{
					cards: []*card{
						{kind: "K"},
						{kind: "K"},
						{kind: "6"},
						{kind: "7"},
						{kind: "7"},
					},
					bet: 28,
				},
				{
					cards: []*card{
						{kind: "T"},
						{kind: "5"},
						{kind: "5"},
						{kind: "J"},
						{kind: "5"},
					},
					bet: 684,
				},
				{
					cards: []*card{
						{kind: "Q"},
						{kind: "Q"},
						{kind: "Q"},
						{kind: "J"},
						{kind: "A"},
					},
					bet: 483,
				},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			g := Game{hands: tt.input}
			g.rankHands(false)
			for i := range g.hands {
				if g.hands[i].bet != tt.want[i].bet {
					t.Fatalf("g.hands[i].bet = %v, want %v", g.hands[i].bet, tt.want[i].bet)
				}
				if len(g.hands[i].cards) != len(tt.want[i].cards) {
					t.Fatalf("len(g.hands[i].cards) = %v, want %v", len(g.hands[i].cards), len(tt.want[i].cards))
				}
				for j := range g.hands[i].cards {
					if g.hands[i].cards[j].kind != tt.want[i].cards[j].kind {
						t.Fatalf("g.hands[i].cards[j].kind = %v, want %v", g.hands[i].cards[j].kind, tt.want[i].cards[j].kind)
					}
				}
			}
		})
	}
}

func TestSortHandsByRank2(t *testing.T) {
	tests := map[string]struct {
		input []*hand
		want  []*hand
	}{
		"two hands": {
			input: []*hand{
				{
					cards: []*card{
						{kind: "T"},
						{kind: "5"},
						{kind: "5"},
						{kind: "J"},
						{kind: "5"},
					},
					bet: 684,
				},
				{
					cards: []*card{
						{kind: "3"},
						{kind: "2"},
						{kind: "T"},
						{kind: "3"},
						{kind: "K"},
					},
					bet: 765,
				},
			},
			want: []*hand{
				{
					cards: []*card{
						{kind: "3"},
						{kind: "2"},
						{kind: "T"},
						{kind: "3"},
						{kind: "K"},
					},
					bet: 765,
				},
				{
					cards: []*card{
						{kind: "T"},
						{kind: "5"},
						{kind: "5"},
						{kind: "J"},
						{kind: "5"},
					},
					bet: 684,
				},
			},
		},
		"whole shebang": {
			input: []*hand{
				{
					cards: []*card{
						{kind: "3"},
						{kind: "2"},
						{kind: "T"},
						{kind: "3"},
						{kind: "K"},
					},
					bet: 765,
				},
				{
					cards: []*card{
						{kind: "T"},
						{kind: "5"},
						{kind: "5"},
						{kind: "J"},
						{kind: "5"},
					},
					bet: 684,
				},
				{
					cards: []*card{
						{kind: "K"},
						{kind: "K"},
						{kind: "6"},
						{kind: "7"},
						{kind: "7"},
					},
					bet: 28,
				},
				{
					cards: []*card{
						{kind: "K"},
						{kind: "T"},
						{kind: "J"},
						{kind: "J"},
						{kind: "T"},
					},
					bet: 220,
				},
				{
					cards: []*card{
						{kind: "Q"},
						{kind: "Q"},
						{kind: "Q"},
						{kind: "J"},
						{kind: "A"},
					},
					bet: 483,
				},
			},
			want: []*hand{
				{
					cards: []*card{
						{kind: "3"},
						{kind: "2"},
						{kind: "T"},
						{kind: "3"},
						{kind: "K"},
					},
					bet: 765,
				},
				{
					cards: []*card{
						{kind: "K"},
						{kind: "K"},
						{kind: "6"},
						{kind: "7"},
						{kind: "7"},
					},
					bet: 28,
				},
				{
					cards: []*card{
						{kind: "T"},
						{kind: "5"},
						{kind: "5"},
						{kind: "J"},
						{kind: "5"},
					},
					bet: 684,
				},
				{
					cards: []*card{
						{kind: "Q"},
						{kind: "Q"},
						{kind: "Q"},
						{kind: "J"},
						{kind: "A"},
					},
					bet: 483,
				},
				{
					cards: []*card{
						{kind: "K"},
						{kind: "T"},
						{kind: "J"},
						{kind: "J"},
						{kind: "T"},
					},
					bet: 220,
				},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			g := Game{hands: tt.input}
			g.rankHands(true)
			for i := range g.hands {
				if g.hands[i].bet != tt.want[i].bet {
					t.Fatalf("g.hands[i].bet = %v, want %v", g.hands[i].bet, tt.want[i].bet)
				}
				if len(g.hands[i].cards) != len(tt.want[i].cards) {
					t.Fatalf("len(g.hands[i].cards) = %v, want %v", len(g.hands[i].cards), len(tt.want[i].cards))
				}
				for j := range g.hands[i].cards {
					if g.hands[i].cards[j].kind != tt.want[i].cards[j].kind {
						t.Fatalf("g.hands[i].cards[j].kind = %v, want %v", g.hands[i].cards[j].kind, tt.want[i].cards[j].kind)
					}
				}
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tests := map[string]struct {
		a    []*card
		b    []*card
		want []*card
	}{
		"same hand": {
			a: []*card{
				{kind: "3"},
				{kind: "2"},
				{kind: "T"},
				{kind: "3"},
				{kind: "K"},
			},
			b: []*card{
				{kind: "3"},
				{kind: "2"},
				{kind: "T"},
				{kind: "3"},
				{kind: "K"},
			},
			want: []*card{
				{kind: "3"},
				{kind: "2"},
				{kind: "T"},
				{kind: "3"},
				{kind: "K"},
			},
		},
		"b is better": {
			a: []*card{
				{kind: "3"},
				{kind: "2"},
				{kind: "T"},
				{kind: "3"},
				{kind: "K"},
			},
			b: []*card{
				{kind: "4"},
				{kind: "2"},
				{kind: "T"},
				{kind: "3"},
				{kind: "K"},
			},
			want: []*card{
				{kind: "4"},
				{kind: "2"},
				{kind: "T"},
				{kind: "3"},
				{kind: "K"},
			},
		},
		"a is better": {
			a: []*card{
				{kind: "3"},
				{kind: "3"},
				{kind: "T"},
				{kind: "3"},
				{kind: "K"},
			},
			b: []*card{
				{kind: "3"},
				{kind: "2"},
				{kind: "T"},
				{kind: "3"},
				{kind: "K"},
			},
			want: []*card{
				{kind: "3"},
				{kind: "3"},
				{kind: "T"},
				{kind: "3"},
				{kind: "K"},
			},
		},
		"failing case idk why but i should write it down": {
			a: []*card{
				{kind: "T"},
				{kind: "5"},
				{kind: "5"},
				{kind: "J"},
				{kind: "5"},
			},
			b: []*card{
				{kind: "Q"},
				{kind: "Q"},
				{kind: "Q"},
				{kind: "J"},
				{kind: "A"},
			},
			want: []*card{
				{kind: "Q"},
				{kind: "Q"},
				{kind: "Q"},
				{kind: "J"},
				{kind: "A"},
			},
		},
		"failing case idk why but i should write it down part 2": {
			a: []*card{
				{kind: "J"},
				{kind: "K"},
				{kind: "K"},
				{kind: "K"},
				{kind: "2"},
			},
			b: []*card{
				{kind: "Q"},
				{kind: "Q"},
				{kind: "Q"},
				{kind: "Q"},
				{kind: "2"},
			},
			want: []*card{
				{kind: "Q"},
				{kind: "Q"},
				{kind: "Q"},
				{kind: "Q"},
				{kind: "2"},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := compare(tt.a, tt.b)
			if !equal(got, tt.want) {
				for i := range got {
					if got[i].kind != tt.want[i].kind {
						t.Fatalf("got[i].kind = %v, want %v", got[i].kind, tt.want[i].kind)
					}
				}
			}
		})
	}
}

func TestMaxPossibleType(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		input hand
		want  handType
	}{
		"non-j returns normal, 3 of a kind": {
			input: hand{
				cards: []*card{
					{kind: "T"},
					{kind: "5"},
					{kind: "5"},
					{kind: "K"},
					{kind: "5"},
				},
			},
			want: threeOfAKind,
		},
		"one joker, random cards": {
			input: hand{
				cards: []*card{
					{kind: "2"},
					{kind: "3"},
					{kind: "4"},
					{kind: "5"},
					{kind: "J"},
				},
			},
			want: onePair,
		},
		"two joker, random cards": {
			input: hand{
				cards: []*card{
					{kind: "2"},
					{kind: "3"},
					{kind: "4"},
					{kind: "J"},
					{kind: "J"},
				},
			},
			want: threeOfAKind,
		},
		"three joker, random cards": {
			input: hand{
				cards: []*card{
					{kind: "2"},
					{kind: "3"},
					{kind: "J"},
					{kind: "J"},
					{kind: "J"},
				},
			},
			want: fourOfAKind,
		},
		"four joker, random cards": {
			input: hand{
				cards: []*card{
					{kind: "2"},
					{kind: "J"},
					{kind: "J"},
					{kind: "J"},
					{kind: "J"},
				},
			},
			want: fiveOfAKind,
		},
		"five joker, random cards": {
			input: hand{
				cards: []*card{
					{kind: "J"},
					{kind: "J"},
					{kind: "J"},
					{kind: "J"},
					{kind: "J"},
				},
			},
			want: fiveOfAKind,
		},
		"one joker, 3 of a kind goes to 4": {
			input: hand{
				cards: []*card{
					{kind: "T"},
					{kind: "5"},
					{kind: "5"},
					{kind: "J"},
					{kind: "5"},
				},
			},
			want: fourOfAKind,
		},
		"one joker, makes full house": {
			input: hand{
				cards: []*card{
					{kind: "5"},
					{kind: "5"},
					{kind: "J"},
					{kind: "9"},
					{kind: "9"},
				},
			},
			want: fullHouse,
		},
		"omg please make it stop": {
			input: hand{
				cards: []*card{
					{kind: "J"},
					{kind: "5"},
					{kind: "J"},
					{kind: "J"},
					{kind: "J"},
				},
			},
			want: fiveOfAKind,
		},
		"existence is truly pain": {
			input: hand{
				cards: []*card{
					{kind: "J"},
					{kind: "J"},
					{kind: "J"},
					{kind: "J"},
					{kind: "J"},
				},
			},
			want: fiveOfAKind,
		},
		"please I need this to pass, my kids are hungry": {
			input: hand{
				cards: []*card{
					{kind: "T"},
					{kind: "5"},
					{kind: "5"},
					{kind: "J"},
					{kind: "J"},
				},
			},
			want: fourOfAKind,
		},
		"okay my kids are dogs but still": {
			input: hand{
				cards: []*card{
					{kind: "J"},
					{kind: "5"},
					{kind: "5"},
					{kind: "J"},
					{kind: "J"},
				},
			},
			want: fiveOfAKind,
		},
		"my eyes hurt": {
			input: hand{
				cards: []*card{
					// QJJQ2
					{kind: "Q"},
					{kind: "J"},
					{kind: "J"},
					{kind: "Q"},
					{kind: "2"},
				},
			},
			want: fourOfAKind,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.input.getType2(); got != tt.want {
				t.Fatalf("() = %v, want %v", got, tt.want)
			}
		})
	}
}

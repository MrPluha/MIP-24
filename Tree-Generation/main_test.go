package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_generateStartNr(t *testing.T) {
	type args struct {
		seed int64
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"+valid",
			args{1},
			[]int{46140, 49680, 46860, 40620, 45060},
		},
		{
			"+valid2",
			args{2},
			[]int{41220, 49560, 49680, 43380, 45900},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateStartNr(tt.args.seed)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("generateStartNr() = %s", diff)
			}
		})
	}
}

func Test_prepareGameStart(t *testing.T) {
	type args struct {
		startNumbers []int
	}
	tests := []struct {
		name string
		args args
		want *gameState
	}{
		{
			"+valid",
			args{[]int{46140, 49680, 46860, 40620, 45060}},
			&gameState{
				current: []*number{
					{value: 46140}, {value: 49680}, {value: 46860}, {value: 40620}, {value: 45060},
				},
			},
		},
		{
			"+withZero",
			args{[]int{46140, 49680, 46860, 0, 45060}},
			&gameState{
				current: []*number{
					{value: 46140}, {value: 49680}, {value: 46860}, nil, {value: 45060},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := prepareGameStart(tt.args.startNumbers)
			if diff := cmp.Diff(tt.want, got, cmp.AllowUnexported(gameState{}), cmp.AllowUnexported(number{})); diff != "" {
				t.Fatalf("prepareGameStart() = %s", diff)
			}
		})
	}
}

func Test_gameState_findNextState(t *testing.T) {
	tests := []struct {
		name string
		gs   *gameState
		want *gameState
	}{

		{
			"+valid",
			&gameState{
				current: []*number{
					{
						value: 769,
					},
				},
			},
			&gameState{
				current: []*number{
					{
						value: 769,
					},
				},
				final: true,
			},
		},
		{
			"+singleStartNumber",
			&gameState{
				current: []*number{
					{
						value: 46140,
					},
				},
			},
			&gameState{
				current: []*number{
					{
						value: 46140,
					},
				},
				nextStates: []*gameState{
					{
						current: []*number{
							{
								value:  15380,
								points: 1,
								bank:   1,
							},
							{
								value:  11535,
								points: -1,
								bank:   1,
							},
							{
								value:  9228,
								points: 1,
								bank:   0,
							},
						},
						nextStates: []*gameState{
							{
								current: []*number{
									{
										value:  3845,
										points: 0,
										bank:   2,
									},
									{
										value:  3076,
										points: 2,
										bank:   1,
									},
								},
								nextStates: []*gameState{
									{
										current: []*number{
											{
												value:  769,
												points: -1,
												bank:   2,
											},
										},
										final: true,
									},
									{
										current: []*number{
											{
												value:  769,
												points: 1,
												bank:   1,
											},
										},
										final: true,
									},
								},
							},
							{
								current: []*number{
									{
										value:  3845,
										points: -2,
										bank:   2,
									},
									{
										value:  2307,
										points: -2,
										bank:   1,
									},
								},
								nextStates: []*gameState{
									{
										current: []*number{
											{
												value:  769,
												points: -3,
												bank:   2,
											},
										},
										final: true,
									},
									{
										current: []*number{
											{
												value:  769,
												points: -3,
												bank:   1,
											},
										},
										final: true,
									},
								},
							},
							{
								current: []*number{
									{
										value:  3076,
										points: 2,
									},
									{
										value: 2307,
									},
								},
								nextStates: []*gameState{
									{
										current: []*number{
											{
												value:  769,
												points: 1,
											},
										},
										final: true,
									},
									{
										current: []*number{
											{
												value:  769,
												points: -1,
											},
										},
										final: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.gs.findNextState()
			if diff := cmp.Diff(tt.want, tt.gs, cmp.AllowUnexported(gameState{}), cmp.AllowUnexported(number{})); diff != "" {
				t.Fatalf("gs.findNextState() = %s", diff)
			}
		})
	}
}

func Test_calculate(t *testing.T) {
	type args struct {
		nr *number
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"+valid",
			args{&number{value: 46140}},
			[]int{15380, 11535, 9228},
		},
		{
			"+44580",
			args{&number{value: 44580}},
			[]int{14860, 11145, 8916},
		},
		{
			"+11145",
			args{&number{value: 11145}},
			[]int{3715, 2229},
		},
		{
			"+8916",
			args{&number{value: 8916}},
			[]int{2972, 2229},
		},
		{
			"+2972",
			args{&number{value: 2972}},
			[]int{743},
		},
		{
			"+2229",
			args{&number{value: 2229}},
			[]int{743},
		}, {
			"+743",
			args{&number{value: 743}},
			nil,
		},
		{
			"+14860",
			args{&number{value: 14860}},
			[]int{3715, 2972},
		},
		{
			"+3715",
			args{&number{value: 3715}},
			[]int{743},
		},
		{
			"+2972",
			args{&number{value: 2972}},
			[]int{743},
		},
		{
			"+last",
			args{&number{value: 769}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.nr.calculateNextNumbers()
			if diff := cmp.Diff(tt.want, got, cmp.AllowUnexported(gameState{}), cmp.AllowUnexported(number{})); diff != "" {
				t.Fatalf("number.calculate() number = %s", diff)
			}
		})
	}
}

func Test_number_calculatePointsAndBank(t *testing.T) {
	tests := []struct {
		name       string
		nr         *number
		wantPoints int
		wantBank   int
	}{
		{
			"+valid",
			&number{value: 800},
			1,
			1,
		},
		{
			"+negativePoint",
			&number{value: 885},
			-1,
			1,
		},
		{
			"+onlyPoints",
			&number{value: 769},
			-1,
			0,
		},
		{
			"+existingPointsAndBank",
			&number{value: 769, points: 2, bank: 2},
			1,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.nr.calculatePointsAndBank()

			if diff := cmp.Diff(tt.wantPoints, tt.nr.points); diff != "" {
				t.Fatalf("number.calculatePointsAndBank() points = %s", diff)
			}

			if diff := cmp.Diff(tt.wantBank, tt.nr.bank); diff != "" {
				t.Fatalf("number.calculatePointsAndBank() bank = %s", diff)
			}
		})
	}
}

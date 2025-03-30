package tree

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
		want *GameState
	}{
		{
			"+valid",
			args{[]int{46140, 49680, 46860, 40620, 45060}},
			&GameState{
				Current: []*Number{
					{Value: 46140}, {Value: 49680}, {Value: 46860}, {Value: 40620}, {Value: 45060},
				},
			},
		},
		{
			"+withZero",
			args{[]int{46140, 49680, 46860, 0, 45060}},
			&GameState{
				Current: []*Number{
					{Value: 46140}, {Value: 49680}, {Value: 46860}, nil, {Value: 45060},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := PrepareGameStart(tt.args.startNumbers)
			if diff := cmp.Diff(tt.want, got, cmp.AllowUnexported(GameState{}), cmp.AllowUnexported(Number{})); diff != "" {
				t.Fatalf("prepareGameStart() = %s", diff)
			}
		})
	}
}

func Test_gameState_findNextState(t *testing.T) {
	tests := []struct {
		name string
		gs   *GameState
		want *GameState
	}{
		{
			"print",
			&GameState{
				Current: []*Number{
					{
						Value: 40020,
					},
				},
			},
			nil,
		},
		{
			"+valid",
			&GameState{
				Current: []*Number{
					{
						Value: 769,
					},
				},
			},
			&GameState{
				Current: []*Number{
					{
						Value: 769,
					},
				},
				Final: true,
			},
		},
		{
			"+singleStartNumber",
			&GameState{
				Current: []*Number{
					{
						Value: 46140,
					},
				},
			},
			&GameState{
				Current: []*Number{
					{
						Value: 46140,
					},
				},
				NextStates: []*GameState{
					{
						Current: []*Number{
							{
								Value:  15380,
								Points: 1,
								Bank:   1,
							},
							{
								Value:  11535,
								Points: -1,
								Bank:   1,
							},
							{
								Value:  9228,
								Points: 1,
								Bank:   0,
							},
						},
						NextStates: []*GameState{
							{
								Current: []*Number{
									{
										Value:  3845,
										Points: 0,
										Bank:   2,
									},
									{
										Value:  3076,
										Points: 2,
										Bank:   1,
									},
								},
								NextStates: []*GameState{
									{
										Current: []*Number{
											{
												Value:  769,
												Points: -1,
												Bank:   2,
											},
										},
										Final: true,
									},
									{
										Current: []*Number{
											{
												Value:  769,
												Points: 1,
												Bank:   1,
											},
										},
										Final: true,
									},
								},
							},
							{
								Current: []*Number{
									{
										Value:  3845,
										Points: -2,
										Bank:   2,
									},
									{
										Value:  2307,
										Points: -2,
										Bank:   1,
									},
								},
								NextStates: []*GameState{
									{
										Current: []*Number{
											{
												Value:  769,
												Points: -3,
												Bank:   2,
											},
										},
										Final: true,
									},
									{
										Current: []*Number{
											{
												Value:  769,
												Points: -3,
												Bank:   1,
											},
										},
										Final: true,
									},
								},
							},
							{
								Current: []*Number{
									{
										Value:  3076,
										Points: 2,
									},
									{
										Value: 2307,
									},
								},
								NextStates: []*GameState{
									{
										Current: []*Number{
											{
												Value:  769,
												Points: 1,
											},
										},
										Final: true,
									},
									{
										Current: []*Number{
											{
												Value:  769,
												Points: -1,
											},
										},
										Final: true,
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
			tt.gs.FindNextState()
			if diff := cmp.Diff(tt.want, tt.gs, cmp.AllowUnexported(GameState{}), cmp.AllowUnexported(Number{})); diff != "" {
				t.Fatalf("gs.findNextState() = %s", diff)
			}
		})
	}
}

func Test_calculate(t *testing.T) {
	type args struct {
		nr *Number
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"+valid",
			args{&Number{Value: 46140}},
			[]int{15380, 11535, 9228},
		},
		{
			"+44580",
			args{&Number{Value: 44580}},
			[]int{14860, 11145, 8916},
		},
		{
			"+11145",
			args{&Number{Value: 11145}},
			[]int{3715, 2229},
		},
		{
			"+8916",
			args{&Number{Value: 8916}},
			[]int{2972, 2229},
		},
		{
			"+2972",
			args{&Number{Value: 2972}},
			[]int{743},
		},
		{
			"+2229",
			args{&Number{Value: 2229}},
			[]int{743},
		}, {
			"+743",
			args{&Number{Value: 743}},
			nil,
		},
		{
			"+14860",
			args{&Number{Value: 14860}},
			[]int{3715, 2972},
		},
		{
			"+3715",
			args{&Number{Value: 3715}},
			[]int{743},
		},
		{
			"+2972",
			args{&Number{Value: 2972}},
			[]int{743},
		},
		{
			"+last",
			args{&Number{Value: 769}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.nr.calculateNextNumbers()
			if diff := cmp.Diff(tt.want, got, cmp.AllowUnexported(GameState{}), cmp.AllowUnexported(Number{})); diff != "" {
				t.Fatalf("Number.calculate() Number = %s", diff)
			}
		})
	}
}

func Test_number_calculatePointsAndBank(t *testing.T) {
	tests := []struct {
		name       string
		nr         *Number
		wantPoints int
		wantBank   int
	}{
		{
			"+valid",
			&Number{Value: 800},
			1,
			1,
		},
		{
			"+negativePoint",
			&Number{Value: 885},
			-1,
			1,
		},
		{
			"+onlyPoints",
			&Number{Value: 769},
			-1,
			0,
		},
		{
			"+existingPointsAndBank",
			&Number{Value: 769, Points: 2, Bank: 2},
			1,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.nr.calculatePointsAndBank()

			if diff := cmp.Diff(tt.wantPoints, tt.nr.Points); diff != "" {
				t.Fatalf("Number.calculatePointsAndBank() Points = %s", diff)
			}

			if diff := cmp.Diff(tt.wantBank, tt.nr.Bank); diff != "" {
				t.Fatalf("Number.calculatePointsAndBank() Bank = %s", diff)
			}
		})
	}
}

package main

import (
	tree "game/Tree-Generation"
	"testing"
)

// func Test_minMaxTest(t *testing.T) {
// 	type args struct {
// 		startNumbers []int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			"+valid",
// 			args{
// 				[]int{40020},
// 			},
// 			-1,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := minMaxTest(tt.args.startNumbers); got != tt.want {
// 				t.Errorf("minMaxTest() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_minMax(t *testing.T) {
	type args struct {
		state *tree.GameState
		max   bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	"min_win_1",
		// 	args{
		// 		&tree.GameState{
		// 			Final: true,
		// 			Current: []*tree.Number{
		// 				{
		// 					Value:       667,
		// 					Points:      -1,
		// 					Bank:        2,
		// 					TotalPoints: 1,
		// 					// WinPlayer:   -1,
		// 				},
		// 			},
		// 			NextStates: nil,
		// 		},
		// 		false,
		// 	},
		// 	-1,
		// },
		// {
		// 	"max_win_1",
		// 	args{
		// 		&tree.GameState{
		// 			Final: true,
		// 			Current: []*tree.Number{
		// 				{
		// 					Value:       667,
		// 					Points:      1,
		// 					Bank:        1,
		// 					TotalPoints: 2,
		// 					// WinPlayer:   1,
		// 				},
		// 			},
		// 			NextStates: nil,
		// 		},
		// 		false,
		// 	},
		// 	1,
		// },
		// {
		// 	"min_win_2",
		// 	args{
		// 		&tree.GameState{
		// 			Final: true,
		// 			Current: []*tree.Number{
		// 				{
		// 					Value:       667,
		// 					Points:      -3,
		// 					Bank:        2,
		// 					TotalPoints: -1,
		// 					// WinPlayer:   -1,
		// 				},
		// 			},
		// 			NextStates: nil,
		// 		},
		// 		false,
		// 	},
		// 	-1,
		// },
		// {
		// 	"min_win_3",
		// 	args{
		// 		&tree.GameState{
		// 			Final: false,
		// 			Current: []*tree.Number{
		// 				{
		// 					Value:  3335,
		// 					Points: 0,
		// 					Bank:   2,
		// 					// WinPlayer:   -1,
		// 				},
		// 			},
		// 			NextStates: []*tree.GameState{
		// 				{
		// 					Final: true,
		// 					Current: []*tree.Number{
		// 						{
		// 							Value:       667,
		// 							Points:      -1,
		// 							Bank:        2,
		// 							TotalPoints: 1,
		// 							// WinPlayer:   -1,
		// 						},
		// 					},
		// 					NextStates: nil,
		// 				},
		// 			},
		// 		},
		// 		true,
		// 	},
		// 	-1,
		// },
		// {
		// 	"max_win_3",
		// 	args{
		// 		&tree.GameState{
		// 			Final: false,
		// 			Current: []*tree.Number{
		// 				{
		// 					Value:  2668,
		// 					Points: 2,
		// 					Bank:   1,
		// 					// WinPlayer:   -1,
		// 				},
		// 			},
		// 			NextStates: []*tree.GameState{
		// 				{
		// 					Final: true,
		// 					Current: []*tree.Number{
		// 						{
		// 							Value:       667,
		// 							Points:      1,
		// 							Bank:        1,
		// 							TotalPoints: 2,
		// 							// WinPlayer:   1,
		// 						},
		// 					},
		// 					NextStates: nil,
		// 				},
		// 			},
		// 		},
		// 		true,
		// 	},
		// 	1,
		// },
		// {
		// 	"min_win_4",
		// 	args{
		// 		state: &tree.GameState{
		// 			Final: false,
		// 			Current: []*tree.Number{
		// 				{
		// 					Value:  13340,
		// 					Points: 1,
		// 					Bank:   1,
		// 					// WinPlayer:   -1,
		// 				},
		// 			},
		// 			NextStates: []*tree.GameState{
		// 				{
		// 					Final: false,
		// 					Current: []*tree.Number{
		// 						{
		// 							Value:  3335,
		// 							Points: 0,
		// 							Bank:   2,
		// 							// WinPlayer:   -1,
		// 						},
		// 						{
		// 							Value:  2668,
		// 							Points: 2,
		// 							Bank:   1,
		// 							// WinPlayer:   -1,
		// 						},
		// 					},
		// 					NextStates: []*tree.GameState{
		// 						{
		// 							Final: true,
		// 							Current: []*tree.Number{
		// 								{
		// 									Value:       667,
		// 									Points:      -1,
		// 									Bank:        2,
		// 									TotalPoints: 1,
		// 									// WinPlayer:   1,
		// 								},
		// 							},
		// 							NextStates: nil,
		// 						},
		// 						{
		// 							Final: true,
		// 							Current: []*tree.Number{
		// 								{
		// 									Value:       667,
		// 									Points:      1,
		// 									Bank:        1,
		// 									TotalPoints: 2,
		// 									// WinPlayer:   1,
		// 								},
		// 							},
		// 							NextStates: nil,
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		max: false,
		// 	},
		// 	-1,
		// },
		// {
		// 	"max_win_3",
		// 	args{
		// 		&tree.GameState{
		// 			Final: false,
		// 			Current: []*tree.Number{
		// 				{
		// 					Value:  2668,
		// 					Points: 2,
		// 					Bank:   1,
		// 					// WinPlayer:   -1,
		// 				},
		// 			},
		// 			NextStates: []*tree.GameState{
		// 				{
		// 					Final: true,
		// 					Current: []*tree.Number{
		// 						{
		// 							Value:       667,
		// 							Points:      1,
		// 							Bank:        1,
		// 							TotalPoints: 2,
		// 							// WinPlayer:   1,
		// 						},
		// 					},
		// 					NextStates: nil,
		// 				},
		// 			},
		// 		},
		// 		true,
		// 	},
		// 	1,
		// },
		// {
		// 	"min_win_5",
		// 	args{
		// 		state: &tree.GameState{
		// 			Final: false,
		// 			Current: []*tree.Number{
		// 				{
		// 					Value:  10005,
		// 					Points: -1,
		// 					Bank:   1,
		// 					// WinPlayer:   -1,
		// 				},
		// 			},
		// 			NextStates: []*tree.GameState{
		// 				{
		// 					Final: false,
		// 					Current: []*tree.Number{
		// 						{
		// 							Value:  3335,
		// 							Points: -2,
		// 							Bank:   2,
		// 							// WinPlayer:   -1,
		// 						},
		// 						{
		// 							Value:  2001,
		// 							Points: -2,
		// 							Bank:   1,
		// 							// WinPlayer:   -1,
		// 						},
		// 					},
		// 					NextStates: []*tree.GameState{
		// 						{
		// 							Final: true,
		// 							Current: []*tree.Number{
		// 								{
		// 									Value:       667,
		// 									Points:      -3,
		// 									Bank:        2,
		// 									TotalPoints: -1,
		// 									// WinPlayer:   1,
		// 								},
		// 							},
		// 							NextStates: nil,
		// 						},
		// 						{
		// 							Final: true,
		// 							Current: []*tree.Number{
		// 								{
		// 									Value:       667,
		// 									Points:      -3,
		// 									Bank:        1,
		// 									TotalPoints: -2,
		// 									// WinPlayer:   1,
		// 								},
		// 							},
		// 							NextStates: nil,
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		max: false,
		// 	},
		// 	-1,
		// },
		// {
		// 	"min_win_6",
		// 	args{
		// 		state: &tree.GameState{
		// 			Final: false,
		// 			Current: []*tree.Number{
		// 				{
		// 					Value:  8004,
		// 					Points: 1,
		// 					Bank:   0,
		// 					// WinPlayer:   -1,
		// 				},
		// 			},
		// 			NextStates: []*tree.GameState{
		// 				{
		// 					Final: false,
		// 					Current: []*tree.Number{
		// 						{
		// 							Value:  2668,
		// 							Points: 2,
		// 							Bank:   0,
		// 							// WinPlayer:   -1,
		// 						},
		// 						{
		// 							Value:  2001,
		// 							Points: 0,
		// 							Bank:   0,
		// 							// WinPlayer:   -1,
		// 						},
		// 					},
		// 					NextStates: []*tree.GameState{
		// 						{
		// 							Final: true,
		// 							Current: []*tree.Number{
		// 								{
		// 									Value:       667,
		// 									Points:      1,
		// 									Bank:        0,
		// 									TotalPoints: 1,
		// 									// WinPlayer:   1,
		// 								},
		// 							},
		// 							NextStates: nil,
		// 						},
		// 						{
		// 							Final: true,
		// 							Current: []*tree.Number{
		// 								{
		// 									Value:       667,
		// 									Points:      -1,
		// 									Bank:        0,
		// 									TotalPoints: -1,
		// 									// WinPlayer:   1,
		// 								},
		// 							},
		// 							NextStates: nil,
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		max: false,
		// 	},
		// 	-1,
		// },
		{
			"full_40020",
			args{
				&tree.GameState{
					Current: []*tree.Number{
						{Value: 40020},
					},
					NextStates: []*tree.GameState{
						{
							Final: false,
							Current: []*tree.Number{
								{
									Value:  13340,
									Points: 1,
									Bank:   1,
									// WinPlayer:   -1,
								},
							},
							NextStates: []*tree.GameState{
								{
									Final: false,
									Current: []*tree.Number{
										{
											Value:  3335,
											Points: 0,
											Bank:   2,
											// WinPlayer:   -1,
										},
										{
											Value:  2668,
											Points: 2,
											Bank:   1,
											// WinPlayer:   -1,
										},
									},
									NextStates: []*tree.GameState{
										{
											Final: true,
											Current: []*tree.Number{
												{
													Value:       667,
													Points:      -1,
													Bank:        2,
													TotalPoints: 1,
													// WinPlayer:   1,
												},
											},
											NextStates: nil,
										},
										{
											Final: true,
											Current: []*tree.Number{
												{
													Value:       667,
													Points:      1,
													Bank:        1,
													TotalPoints: 2,
													// WinPlayer:   1,
												},
											},
											NextStates: nil,
										},
									},
								},
							},
						},
						{
							Final: false,
							Current: []*tree.Number{
								{
									Value:  10005,
									Points: -1,
									Bank:   1,
									// WinPlayer:   -1,
								},
							},
							NextStates: []*tree.GameState{
								{
									Final: false,
									Current: []*tree.Number{
										{
											Value:  3335,
											Points: -2,
											Bank:   2,
											// WinPlayer:   -1,
										},
										{
											Value:  2001,
											Points: -2,
											Bank:   1,
											// WinPlayer:   -1,
										},
									},
									NextStates: []*tree.GameState{
										{
											Final: true,
											Current: []*tree.Number{
												{
													Value:       667,
													Points:      -3,
													Bank:        2,
													TotalPoints: -1,
													// WinPlayer:   1,
												},
											},
											NextStates: nil,
										},
										{
											Final: true,
											Current: []*tree.Number{
												{
													Value:       667,
													Points:      -3,
													Bank:        1,
													TotalPoints: -2,
													// WinPlayer:   1,
												},
											},
											NextStates: nil,
										},
									},
								},
							},
						},
						{
							Final: false,
							Current: []*tree.Number{
								{
									Value:  8004,
									Points: 1,
									Bank:   0,
									// WinPlayer:   -1,
								},
							},
							NextStates: []*tree.GameState{
								{
									Final: false,
									Current: []*tree.Number{
										{
											Value:  2668,
											Points: 2,
											Bank:   0,
											// WinPlayer:   -1,
										},
										{
											Value:  2001,
											Points: 0,
											Bank:   0,
											// WinPlayer:   -1,
										},
									},
									NextStates: []*tree.GameState{
										{
											Final: true,
											Current: []*tree.Number{
												{
													Value:       667,
													Points:      1,
													Bank:        0,
													TotalPoints: 1,
													// WinPlayer:   1,
												},
											},
											NextStates: nil,
										},
										{
											Final: true,
											Current: []*tree.Number{
												{
													Value:       667,
													Points:      -1,
													Bank:        0,
													TotalPoints: -1,
													// WinPlayer:   1,
												},
											},
											NextStates: nil,
										},
									},
								},
							},
						},
					},
				},
				true,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMax(tt.args.state, tt.args.max); got != tt.want {
				t.Errorf("minMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getValue(t *testing.T) {
	type args struct {
		in  []int
		max bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"max1",
			args{
				[]int{1, 1, 1},
				true,
			},
			1,
		},
		{
			"min1",
			args{
				[]int{1, 1, 1},
				false,
			},
			1,
		},
		{
			"max2",
			args{
				[]int{1, -1, -1},
				true,
			},
			1,
		},
		{
			"min2",
			args{
				[]int{-1, 1, -1},
				false,
			},
			-1,
		},
		{
			"max3",
			args{
				[]int{-1, -1, -1},
				true,
			},
			-1,
		},
		{
			"min3",
			args{
				[]int{-1, -1, -1},
				false,
			},
			-1,
		},
		{
			"max4",
			args{
				[]int{},
				true,
			},
			0,
		},
		{
			"min4",
			args{
				[]int{},
				false,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValue(tt.args.in, tt.args.max); got != tt.want {
				t.Errorf("getValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

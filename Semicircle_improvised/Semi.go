package main

import (
	"fmt"
	"y/Generic"
	"y/Structs_for_patterns"
	"y/pattern_up"
	"y/utils"
)

type Lettersemicricle struct {
	Date_struct []pattern_up.Date
	Patternn    []string
}

//1.changing the method to find peaks and troughs.
//2.after matching he pattern upward and then downward
//3.then matching starting and ending point's value are same or euivalent and if yes then declaring the points as semicircle.
func (l *Lettersemicricle) Pattern() ([]int, []int, int, bool) {
	Peaks := Peaks(l.Date_struct)
	Troughs := Troughs(l.Date_struct)
	fmt.Println("Peaks", Peaks)
	fmt.Println("Troughs", Troughs)

	var All_trends_string []string
	var mixed_unsorted []int
	for i := range Peaks {
		mixed_unsorted = append(mixed_unsorted, Peaks[i])
	}
	for i := range Troughs {
		mixed_unsorted = append(mixed_unsorted, Troughs[i])
	}
	mixed_sorted := utils.Sort(mixed_unsorted) //mixed values
	var All_trends_values []int
	for _, i := range mixed_sorted {
		if pattern_up.P_up(l.Date_struct[i : i+2]) {
			All_trends_string = append(All_trends_string, "Upward")
			All_trends_values = append(All_trends_values, i)

		} else {
			All_trends_string = append(All_trends_string, "Downward")
			All_trends_values = append(All_trends_values, i)
		}
	}
	var bol bool //just to check if the pattern is semicircle or not,this is boolean that will tell us that.
	fmt.Println("Mix =", mixed_sorted)
	for _, j := range mixed_sorted {
		fmt.Println("Mix =", l.Date_struct[j])
	}
	fmt.Println("main Pattern =", l.Patternn)
	fmt.Println("Pattern =", All_trends_string)
	fmt.Println("Values =", l.Date_struct[All_trends_values[0]:All_trends_values[len(All_trends_values)-1]])

	var start []int
	var end []int
	var count int
	for i := range All_trends_string {
		if All_trends_string[i] == "Upward" && i != len(All_trends_string)-1 {

			if All_trends_string[i+1] == "Downward" && i+1 != len(All_trends_string)-1 { //Task:=>check here for first and last point to be in the align manner.
				var ara1 []int
				var ara2 []int

				var x int
				for i1 := 0; i1 < 250; i1++ {
					x = l.Date_struct[All_trends_values[i]].Value + i1
					ara1 = append(ara1, x)
				}
				for i2 := 0; i2 < 250; i2++ {
					x = l.Date_struct[All_trends_values[i]].Value - i2

					ara1 = append(ara1, x)
				}
				for i1 := 0; i1 < 250; i1++ {
					x = l.Date_struct[All_trends_values[i+1]].Value + i1
					ara2 = append(ara2, x)
				}
				for i2 := 0; i2 < 250; i2++ {
					x = l.Date_struct[All_trends_values[i+1]].Value - i2

					ara2 = append(ara2, x)
				}
				fmt.Println("ara1 =", ara1[0], "ara2:=", ara2[0])
				fmt.Println(l.Date_struct[All_trends_values[i+1]].Value)
				//IsValidCategory1(l.Date_struct[All_trends_values[i+1]].Value, ara1) &&
				if IsValidCategory1(l.Date_struct[All_trends_values[i]].Value, ara1) && IsValidCategory1(l.Date_struct[All_trends_values[i]].Value, ara2) {

					count = count + 1
					// fmt.Println("Count =", count)
					// fmt.Println("Starting Point =", l.T[All_trends_values[i]])
					// fmt.Println("Ending Point =", l.T[All_trends_values[i+1]+1])
					start = append(start, All_trends_values[i])
					end = append(end, All_trends_values[i+2]+1)
					bol = true
					fmt.Println("Just checking", start)
					//|| || IsValidCategory1(l.Date_struct[All_trends_values[i]].Index+2, ara1) || IsValidCategory1(l.Date_struct[All_trends_values[i]].Index+2, ara2)
				} else if IsValidCategory1(l.Date_struct[l.Date_struct[All_trends_values[i]].Index+1].Value, ara1) || IsValidCategory1(l.Date_struct[l.Date_struct[All_trends_values[i]].Index+1].Value, ara2) {
					count = count + 1
					fmt.Println("Ohh mil gaya")
					// fmt.Println("Starting Point =", l.T[All_trends_values[i]])
					// fmt.Println("Ending Point =", l.T[All_trends_values[i+1]+1])
					start = append(start, l.Date_struct[All_trends_values[i]].Index-1)
					end = append(end, All_trends_values[i+2]+1)
					bol = true
					fmt.Println("Just checking", start)

				} else if IsValidCategory1(l.Date_struct[l.Date_struct[All_trends_values[i]].Index+2].Value, ara1) || IsValidCategory1(l.Date_struct[l.Date_struct[All_trends_values[i]].Index+2].Value, ara2) {
					count = count + 1
					fmt.Println("Ohh mil gaya second")
					// fmt.Println("Starting Point =", l.T[All_trends_values[i]])
					// fmt.Println("Ending Point =", l.T[All_trends_values[i+1]+1])
					start = append(start, l.Date_struct[All_trends_values[i]].Index+1)
					end = append(end, All_trends_values[i+2]+1)
					bol = true
					fmt.Println("Just checking", start)
				} else if IsValidCategory1(l.Date_struct[l.Date_struct[All_trends_values[i]].Index+3].Value, ara1) || IsValidCategory1(l.Date_struct[l.Date_struct[All_trends_values[i]].Index+3].Value, ara2) {
					count = count + 1
					fmt.Println("Ohh mil gaya third")
					// fmt.Println("Starting Point =", l.T[All_trends_values[i]])
					// fmt.Println("Ending Point =", l.T[All_trends_values[i+1]+1])
					start = append(start, l.Date_struct[All_trends_values[i]].Index+2)
					end = append(end, All_trends_values[i+2]+1)
					bol = true
					fmt.Println("Just checking", start)
				} else if IsValidCategory1(l.Date_struct[l.Date_struct[All_trends_values[i]].Index+4].Value, ara1) || IsValidCategory1(l.Date_struct[l.Date_struct[All_trends_values[i]].Index+4].Value, ara2) {
					count = count + 1
					fmt.Println("Ohh mil gaya fourth")
					// fmt.Println("Starting Point =", l.T[All_trends_values[i]])
					// fmt.Println("Ending Point =", l.T[All_trends_values[i+1]+1])
					start = append(start, l.Date_struct[All_trends_values[i]].Index+3)
					end = append(end, All_trends_values[i+2]+1)
					bol = true
					fmt.Println("Just checking", start)
				} else {
					bol = false
				}
			}

		}

	}
	// fmt.Println("Count =", count)
	return start, end, count, bol
}

func IsValidCategory1(category int, n []int) bool {
	switch category {
	case
		n[0],
		n[1],
		n[2],
		n[3],
		n[4],
		n[5],
		n[6],
		n[7],
		n[8],
		n[9],
		n[10],
		n[11],
		n[12],
		n[13],
		n[14],
		n[15],
		n[16],
		n[17],
		n[18],
		n[19],
		n[20],
		n[21],
		n[22],
		n[23],
		n[24],
		n[25],
		n[26],
		n[27],
		n[28],
		n[29],
		n[30],
		n[31],
		n[32],
		n[33],
		n[34],
		n[35],
		n[36],
		n[37],
		n[38],
		n[39],
		n[40],
		n[41],
		n[42],
		n[43],
		n[44],
		n[45],
		n[46],
		n[47],
		n[48],
		n[49],
		n[50],
		n[51],
		n[52],
		n[53],
		n[54],
		n[55],
		n[56],
		n[57],
		n[58],
		n[59],
		n[60],
		n[61],
		n[62],
		n[63],
		n[64],
		n[65],
		n[66],
		n[67],
		n[68],
		n[69],
		n[70],
		n[71],
		n[72],
		n[73],
		n[74],
		n[75],
		n[76],
		n[77],
		n[78],
		n[79],
		n[80],
		n[81],
		n[82],
		n[83],
		n[84],
		n[85],
		n[86],
		n[87],
		n[88],
		n[89],
		n[90],
		n[91],
		n[92],
		n[93],
		n[94],
		n[95],
		n[96],
		n[97],
		n[98],
		n[99],
		n[100],
		n[101],
		n[102],
		n[103],
		n[104],
		n[105],
		n[106],
		n[107],
		n[108],
		n[109],
		n[110],
		n[111],
		n[112],
		n[113],
		n[114],
		n[115],
		n[116],
		n[117],
		n[118],
		n[119],
		n[120],
		n[121],
		n[122],
		n[123],
		n[124],
		n[125],
		n[126],
		n[127],
		n[128],
		n[129],
		n[130],
		n[131],
		n[132],
		n[133],
		n[134],
		n[135],
		n[136],
		n[137],
		n[138],
		n[139],
		n[140],
		n[141],
		n[142],
		n[143],
		n[144],
		n[145],
		n[146],
		n[147],
		n[148],
		n[149],
		n[150],
		n[151],
		n[152],
		n[153],
		n[154],
		n[155],
		n[156],
		n[157],
		n[158],
		n[159],
		n[160],
		n[161],
		n[162],
		n[163],
		n[164],
		n[165],
		n[166],
		n[167],
		n[168],
		n[169],
		n[170],
		n[171],
		n[172],
		n[173],
		n[174],
		n[175],
		n[176],
		n[177],
		n[178],
		n[179],
		n[180],
		n[181],
		n[182],
		n[183],
		n[184],
		n[185],
		n[186],
		n[187],
		n[188],
		n[189],
		n[190],
		n[191],
		n[192],
		n[193],
		n[194],
		n[195],
		n[196],
		n[197],
		n[198],
		n[199],
		n[199],
		n[200],
		n[201],
		n[202],
		n[203],
		n[204],
		n[205],
		n[206],
		n[207],
		n[208],
		n[209],
		n[210],
		n[211],
		n[212],
		n[213],
		n[214],
		n[215],
		n[216],
		n[217],
		n[218],
		n[219],
		n[220],
		n[221],
		n[222],
		n[223],
		n[224],
		n[225],
		n[226],
		n[227],
		n[228],
		n[229],
		n[230],
		n[231],
		n[232],
		n[233],
		n[234],
		n[235],
		n[236],
		n[237],
		n[238],
		n[239],
		n[240],
		n[241],
		n[242],
		n[243],
		n[244],
		n[245],
		n[246],
		n[247],
		n[248],
		n[249],
		n[250],
		n[251],
		n[252],
		n[253],
		n[254],
		n[255],
		n[256],
		n[257],
		n[258],
		n[259],
		n[260],
		n[261],
		n[262],
		n[263],
		n[264],
		n[265],
		n[266],
		n[267],
		n[268],
		n[269],
		n[270],
		n[271],
		n[272],
		n[273],
		n[274],
		n[275],
		n[276],
		n[277],
		n[278],
		n[279],
		n[280],
		n[281],
		n[282],
		n[283],
		n[284],
		n[285],
		n[286],
		n[287],
		n[288],
		n[289],
		n[290],
		n[291],
		n[292],
		n[293],
		n[294],
		n[295],
		n[296],
		n[297],
		n[298],
		n[299],
		n[300],
		n[301],
		n[302],
		n[303],
		n[304],
		n[305],
		n[306],
		n[307],
		n[308],
		n[309],
		n[310],
		n[311],
		n[312],
		n[313],
		n[314],
		n[315],
		n[316],
		n[317],
		n[318],
		n[319],
		n[320],
		n[321],
		n[322],
		n[323],
		n[324],
		n[325],
		n[326],
		n[327],
		n[328],
		n[329],
		n[330],
		n[331],
		n[332],
		n[333],
		n[334],
		n[335],
		n[336],
		n[337],
		n[338],
		n[339],
		n[340],
		n[341],
		n[342],
		n[343],
		n[344],
		n[345],
		n[346],
		n[347],
		n[348],
		n[349],
		n[350],
		n[351],
		n[352],
		n[353],
		n[354],
		n[355],
		n[356],
		n[357],
		n[358],
		n[359],
		n[360],
		n[361],
		n[362],
		n[363],
		n[364],
		n[365],
		n[366],
		n[367],
		n[368],
		n[369],
		n[370],
		n[371],
		n[372],
		n[373],
		n[374],
		n[375],
		n[376],
		n[377],
		n[378],
		n[379],
		n[380],
		n[381],
		n[382],
		n[383],
		n[384],
		n[385],
		n[386],
		n[387],
		n[388],
		n[389],
		n[390],
		n[391],
		n[392],
		n[393],
		n[394],
		n[395],
		n[396],
		n[397],
		n[398],
		n[399],
		n[400],
		n[401],
		n[402],
		n[403],
		n[404],
		n[405],
		n[406],
		n[407],
		n[408],
		n[409],
		n[410],
		n[411],
		n[412],
		n[413],
		n[414],
		n[415],
		n[416],
		n[417],
		n[418],
		n[419],
		n[420],
		n[421],
		n[422],
		n[423],
		n[424],
		n[425],
		n[426],
		n[427],

		n[428],
		n[429],
		n[430],
		n[431],
		n[432],
		n[433],
		n[434],
		n[435],
		n[436],
		n[437],
		n[438],
		n[439],
		n[440],
		n[441],
		n[442],
		n[443],
		n[444],
		n[445],
		n[446],
		n[447],
		n[448],
		n[449],
		n[450],
		n[451],
		n[452],
		n[453],
		n[454],
		n[455],

		n[456],
		n[457],
		n[458],
		n[459],
		n[460],
		n[461],
		n[462],
		n[463],
		n[464],
		n[465],
		n[466],
		n[467],
		n[468],

		n[469],
		n[470],
		n[471],
		n[472],
		n[473],
		n[474],
		n[475],
		n[476],
		n[477],
		n[478],
		n[479],
		n[480],
		n[481],
		n[482],
		n[483],
		n[484],
		n[485],
		n[486],
		n[487],
		n[488],
		n[489],
		n[490],
		n[491],
		n[492],
		n[493],
		n[494],

		n[495],
		n[496],

		n[497],
		n[498]:
		// n[499],
		// n[500]:
		// n[501],
		// n[502],

		// n[503],
		// n[504],
		// n[505],

		// n[506],
		// n[507],
		// n[508],
		// n[509],
		// n[510],
		// n[511],
		// n[512],
		// n[513],
		// n[514],
		// n[515],
		// n[516],
		// n[517],
		// n[518],
		// n[519],
		// n[520],
		// n[521],
		// n[522],
		// n[523],
		// n[524],

		// n[525],
		// n[526],
		// n[527],
		// n[528],
		// n[529],
		// n[530],
		// n[531],
		// n[532],
		// n[533],
		// n[534],
		// n[535],
		// n[536],
		// n[537],
		// n[538],
		// n[539],
		// n[540],

		// n[541],
		// n[542],
		// n[543],

		// n[544],
		// n[545],
		// n[546],
		// n[547],
		// n[548],
		// n[549],
		// n[550],
		// n[551],
		// n[552],
		// n[553],
		// n[554],
		// n[555],
		// n[556],
		// n[557],
		// n[558],
		// n[559],
		// n[560],
		// n[561],
		// n[562],
		// n[563],
		// n[564],
		// n[565],
		// n[566],
		// n[567],
		// n[568],
		// n[569],
		// n[570],
		// n[571],
		// n[572],
		// n[573],
		// n[574],
		// n[575],
		// n[576],
		// n[577],
		// n[578],
		// n[579],
		// n[580],

		// n[581],
		// n[582],
		// n[583],
		// n[584],
		// n[585],
		// n[586],
		// n[587],
		// n[588],
		// n[589],
		// n[590],
		// n[591],
		// n[592],
		// n[593],
		// n[594],
		// n[595],
		// n[596],
		// n[597],
		// n[598]:

		return true
	}
	return false
}
func Peaks(pts []pattern_up.Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value == pts[0].Value {
			continue
		} else if pts[i].Value == pts[len(pts)-1].Value {
			continue
		} else if pts[i-1].Value < pts[i].Value && pts[i+1].Value < pts[i].Value && pts[i-1] != pts[0] && pts[i+1] != pts[len(pts)-1] {
			// if pts[i-2].Value < pts[i].Value && pts[i+2].Value < pts[i].Value && pts[i-2] != pts[0] && pts[i+2] != pts[len(pts)-1] {
			// 	if pts[i-3].Value < pts[i].Value && pts[i+3].Value < pts[i].Value && pts[i-3] != pts[0] && pts[i+3] != pts[len(pts)-1] {

			b = append(b, i)
			continue
			// 	} else {
			// 		continue
			// 	}
			// } else {
			// 	continue
			// }
		} else {
			continue
		}
	}
	return b
}

func Troughs(pts []pattern_up.Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value == pts[0].Value {
			continue
		} else if pts[i].Value == pts[len(pts)-1].Value {
			continue
		} else if pts[i-1].Value > pts[i].Value && pts[i+1].Value > pts[i].Value && pts[i-1] != pts[0] && pts[i+1] != pts[len(pts)-1] {
			if pts[i-2].Value > pts[i].Value && pts[i+2].Value > pts[i].Value && pts[i-2] != pts[0] && pts[i+2] != pts[len(pts)-1] {
				if pts[i-3].Value > pts[i].Value && pts[i+3].Value > pts[i].Value && pts[i-3] != pts[0] && pts[i+3] != pts[len(pts)-1] {
					if pts[i-4].Value > pts[i].Value && pts[i+4].Value > pts[i].Value && pts[i-4] != pts[0] && pts[i+4] != pts[len(pts)-1] {
						b = append(b, i)
						continue
					} else {
						continue
					}
				} else {
					continue
				}
			} else {
				continue
			}
			// } else {
			// 	continue
			// }
		} else {
			continue
		}
	}
	return b

}

func Adjustpoints(start, end []int, count int, date []pattern_up.Date) ([]int, []int, int) { //
	//1->make these variables as many times as count variable c.

	nearby_points := make([][]int, count)
	//var nearby_points2 []int
	t := make([][]int, count)  //this variable for the index of the first pattern
	t1 := make([][]int, count) //this variable for the index of the second pattern

	// var t1 [][]int

	//1
	//2->make automatic isvalidcategory function and loop thorough the points until the last point.
	// for i := 0; i < count; i++ {
	// 	for _, All_trends_string := range date[start[i] : end[i]+1] {
	// 		if IsValidCategory2(All_trends_string.Value) {
	// 			nearby_points[i] = append(nearby_points[i], All_trends_string.Index)
	// 		}
	// 	}

	// for _, All_trends_values := range date[start[i] : end[i]+1] {
	// 	if IsValidCategory1(All_trends_values.Value) {
	// 		nearby_points[i] = append(nearby_points[i], All_trends_values.Index)
	// 	}

	// }
	// }
	fmt.Println("nearby_points1:", nearby_points[0])
	fmt.Println("nearby_points2:", nearby_points[1])
	//2
	//3-->adjust the points for evry start-end pair of points.
	// for i := 0; i < count; i++ {
	// 	for _, All_trends_string := range date[end[i]-2 : end[i]+1] {
	// 		if IsValidCategory2(All_trends_string.Value) {
	// 			fmt.Println("Foundpoint nearby the ending point 1:", All_trends_string.Index)
	// 			t[i] = append(t[i], All_trends_string.Index)
	// 		}
	// 	}

	// for _, All_trends_values := range date[end[i]-4 : end[i]+1] {
	// 	if IsValidCategory1(All_trends_values.Value) {
	// 		fmt.Println("Foundpoint nearby the ending point 2:", All_trends_values.Index)
	// 		t1[i] = append(t1[i], All_trends_values.Index)
	// 	}
	// }
	// }
	fmt.Println("nearby_points:", nearby_points)
	fmt.Println("t:", t)
	fmt.Println("t1:", t1)

	//3
	//4->according to the number of couunts work this out ,the showing of variable and all condition.
	for i := 0; i < count; i++ {
		if len(t[i]) == 0 || len(t[i+1]) == 0 {
			fmt.Println("No nearby points found for:")
			fmt.Println("elemenating the first semicircle")
			start = start[1:]
			end = end[1:]
			break
			// fmt.Println(len(a), len(b))
		} else {
			fmt.Println("nearby points found for:")
			fmt.Println(len(t))
			end = end[1:]
			end = append(end, t1[1][1])
			break
		}
	}
	for i := 0; i < count; i++ {
		if len(t1[i]) == 0 || len(t1[i+1]) == 0 {
			fmt.Println(" no nearby points found for pair of points:=>")
			start = start[1:]
			end = end[1:]
			fmt.Println(start, end)
			break
		} else {
			fmt.Println("nearby points found for:")
			fmt.Println(len(t1))
			end = end[1:]
			end = append(end, t1[1][1])
			break
		}
	}

	// fmt.Println(len(end))

	count = len(start)
	fmt.Println(start, end, count)
	//4
	return start, end, count
}

//This is kind of like in operator in python language.

func main() {
	date := Generic.Genericcsv("/home/tatva.j@ah.zymrinc.com/Desktop/go-graph/Generic/data.csv")
	pattern := []string{"Downward", "Upward"}
	Semicircle_data_struct := &Lettersemicricle{date, pattern}
	Start1, End1, Count1, _ := Structs_for_patterns.Find(Semicircle_data_struct) //start,end,count,bool
	fmt.Println("Start1:", Start1)
	fmt.Println("End1:", End1)
	fmt.Println("Count1:", Count1)
	// Start, End, Count := Semicircle.Adjustpoints(Start1, End1, Count1, date) //Adjusting the output to meet the semicircle atleast requirements.
	// fmt.Println("Start:", Start)
	// fmt.Println("End:", End)
	// fmt.Println("Count:", Count)
	if Count1 == 0 {
		fmt.Println("No Inverted_semicircle found")
	} else {
		utils.Plot1(date, Count1, Start1, End1)
	}

}

// var a1 []pattern_up.Date
// var b1 []pattern_up.Date

// for _, i := range Peaks {
// 	a1 = append(a1, l.Date_struct[i])
// }
// for _, i := range Troughs {
// 	b1 = append(b1, l.Date_struct[i])
// }
// func IsValidCategory1(category int, n []int) bool {
// 	switch category {
// 	case
// 		n[0],
// 		n[1],
// 		n[2],
// 		n[3],
// 		n[4],
// 		n[5],
// 		n[6],
// 		n[7],
// 		n[8],
// 		n[9],
// 		n[10],
// 		n[11],
// 		n[12],
// 		n[13],
// 		n[14],
// 		n[15],
// 		n[16],
// 		n[17],
// 		n[18],
// 		n[19],
// 		n[20],
// 		n[21],
// 		n[22],
// 		n[23],
// 		n[24],
// 		n[25],
// 		n[26],
// 		n[27],
// 		n[28],
// 		n[29],
// 		n[30],
// 		n[31],
// 		n[32],
// 		n[33],
// 		n[34],
// 		n[35],
// 		n[36],
// 		n[37],
// 		n[38],
// 		n[39],
// 		n[40],
// 		n[41],
// 		n[42],
// 		n[43],
// 		n[44],
// 		n[45],
// 		n[46],
// 		n[47],
// 		n[48],
// 		n[49],
// 		n[50],
// 		n[51],
// 		n[52],
// 		n[53],
// 		n[54],
// 		n[55],
// 		n[56],
// 		n[57],
// 		n[58],
// 		n[59],
// 		n[60],
// 		n[61],
// 		n[62],
// 		n[63],
// 		n[64],
// 		n[65],
// 		n[66],
// 		n[67],
// 		n[68],
// 		n[69],
// 		n[70],
// 		n[71],
// 		n[72],
// 		n[73],
// 		n[74],
// 		n[75],
// 		n[76],
// 		n[77],
// 		n[78],
// 		n[79],
// 		n[80],
// 		n[81],
// 		n[82],
// 		n[83],
// 		n[84],
// 		n[85],
// 		n[86],
// 		n[87],
// 		n[88],
// 		n[89],
// 		n[90],
// 		n[91],
// 		n[92],
// 		n[93],
// 		n[94],
// 		n[95],
// 		n[96],
// 		n[97],
// 		n[98],
// 		n[99],
// 		n[100],
// 		n[101],
// 		n[102],
// 		n[103],
// 		n[104],
// 		n[105],
// 		n[106],
// 		n[107],
// 		n[108],
// 		n[109],
// 		n[110],
// 		n[111],
// 		n[112],
// 		n[113],
// 		n[114],
// 		n[115],
// 		n[116],
// 		n[117],
// 		n[118],
// 		n[119],
// 		n[120],
// 		n[121],
// 		n[122],
// 		n[123],
// 		n[124],
// 		n[125],
// 		n[126],
// 		n[127],
// 		n[128],
// 		n[129],
// 		n[130],
// 		n[131],
// 		n[132],
// 		n[133],
// 		n[134],
// 		n[135],
// 		n[136],
// 		n[137],
// 		n[138],
// 		n[139],
// 		n[140],
// 		n[141],
// 		n[142],
// 		n[143],
// 		n[144],
// 		n[145],
// 		n[146],
// 		n[147],
// 		n[148],
// 		n[149],
// 		n[150],
// 		n[151],
// 		n[152],
// 		n[153],
// 		n[154],
// 		n[155],
// 		n[156],
// 		n[157],
// 		n[158],
// 		n[159],
// 		n[160],
// 		n[161],
// 		n[162],
// 		n[163],
// 		n[164],
// 		n[165],
// 		n[166],
// 		n[167],
// 		n[168],
// 		n[169],
// 		n[170],
// 		n[171],
// 		n[172],
// 		n[173],
// 		n[174],
// 		n[175],
// 		n[176],
// 		n[177],
// 		n[178],
// 		n[179],
// 		n[180],
// 		n[181],
// 		n[182],
// 		n[183],
// 		n[184],
// 		n[185],
// 		n[186],
// 		n[187],
// 		n[188],
// 		n[189],
// 		n[190],
// 		n[191],
// 		n[192],
// 		n[193],
// 		n[194],
// 		n[195],
// 		n[196],
// 		n[197],
// 		n[198]:

// 		return true
// 	}
// 	return false
// }

package main

// 5 * 6 = 30
func makeGlyphs() map[rune][]int {
	return map[rune][]int{
		'?': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			0, 0, 1, 1, 1,
			0, 0, 1, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 1, 0, 0,
		},
		' ': []int{
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
		},
		// кириллица
		'А': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
		},
		'Б': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		'В': []int{
			1, 1, 1, 1, 0,
			1, 0, 0, 1, 0,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		'Г': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
		},
		'Д': []int{
			0, 1, 1, 1, 0,
			0, 1, 0, 1, 0,
			0, 1, 0, 1, 0,
			0, 1, 0, 1, 0,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
		},
		'Е': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
		},
		'Ё': []int{
			0, 1, 0, 1, 0,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
		},
		'Ж': []int{
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 1, 1, 1, 1,
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
		},
		'З': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			0, 0, 1, 1, 1,
			0, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		'И': []int{
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 1, 1,
			1, 0, 1, 0, 1,
			1, 1, 0, 0, 1,
			1, 0, 0, 0, 1,
		},
		'Й': []int{
			0, 1, 1, 1, 0,
			1, 0, 0, 0, 1,
			1, 0, 0, 1, 1,
			1, 0, 1, 0, 1,
			1, 1, 0, 0, 1,
			1, 0, 0, 0, 1,
		},
		'К': []int{
			1, 0, 0, 0, 1,
			1, 0, 0, 1, 0,
			1, 0, 1, 0, 0,
			1, 1, 0, 0, 0,
			1, 0, 1, 0, 0,
			1, 0, 0, 1, 0,
		},
		'Л': []int{
			0, 0, 0, 0, 1,
			0, 0, 0, 1, 1,
			0, 0, 1, 0, 1,
			0, 1, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
		},
		'М': []int{
			1, 0, 0, 0, 1,
			1, 1, 0, 1, 1,
			1, 0, 1, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
		},
		'Н': []int{
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
		},
		'О': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		'П': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
		},
		'Р': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
		},
		'С': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		'Т': []int{
			1, 1, 1, 1, 1,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
		},
		'У': []int{
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
			0, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		'Ф': []int{
			1, 1, 1, 1, 1,
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 1, 1, 1, 1,
			0, 0, 1, 0, 0,
		},
		'Х': []int{
			1, 0, 0, 0, 1,
			0, 1, 0, 1, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 1, 0, 1, 0,
			1, 0, 0, 0, 1,
		},
		'Ц': []int{
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
			0, 0, 0, 0, 1,
		},
		'Ч': []int{
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
			0, 0, 0, 0, 1,
			0, 0, 0, 0, 1,
		},
		'Ш': []int{
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 1, 1, 1, 1,
		},
		'Щ': []int{
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 1, 1, 1, 1,
			0, 0, 0, 0, 1,
		},
		'Ъ': []int{
			1, 1, 0, 0, 0,
			1, 1, 0, 0, 0,
			0, 1, 1, 1, 1,
			0, 1, 0, 0, 1,
			0, 1, 0, 0, 1,
			0, 1, 1, 1, 1,
		},
		'Ь': []int{
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		'Ы': []int{
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 1, 1, 0, 1,
		},
		'Э': []int{
			1, 1, 1, 1, 1,
			0, 0, 0, 0, 1,
			0, 0, 0, 0, 1,
			0, 0, 1, 1, 1,
			0, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		'Ю': []int{
			1, 0, 1, 1, 1,
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 1, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 0, 1, 1, 1,
		},
		'Я': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
			0, 0, 1, 0, 1,
			0, 1, 0, 0, 1,
		},

		// Latin
		'A': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
		},
		'B': []int{
			1, 1, 1, 1, 0,
			1, 0, 0, 1, 0,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		'C': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		'D': []int{
			1, 1, 1, 0, 0,
			1, 0, 0, 1, 0,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 1, 0,
			1, 1, 1, 0, 0,
		},
		'E': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
		},
		'F': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
		},
		'G': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 0,
			1, 0, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		'H': []int{
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
		},
		'I': []int{
			1, 1, 1, 1, 1,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			1, 1, 1, 1, 1,
		},
		'J': []int{
			0, 0, 0, 0, 1,
			0, 0, 0, 0, 1,
			0, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		'K': []int{
			1, 0, 0, 0, 1,
			1, 0, 0, 1, 0,
			1, 0, 1, 0, 0,
			1, 1, 0, 0, 0,
			1, 0, 1, 0, 0,
			1, 0, 0, 1, 0,
		},
		'L': []int{
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
		},
		'M': []int{
			1, 0, 0, 0, 1,
			1, 1, 0, 1, 1,
			1, 0, 1, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
		},
		'N': []int{
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 0, 0, 1,
			1, 0, 1, 0, 1,
			1, 0, 0, 1, 1,
			1, 0, 0, 0, 1,
		},
		'O': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		'P': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
		},
		'Q': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 1, 0, 1,
			1, 0, 0, 1, 1,
			1, 1, 1, 1, 1,
		},
		'R': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
			1, 0, 1, 0, 0,
			1, 0, 0, 1, 0,
		},
		'S': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
			0, 0, 0, 0, 1,
			0, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		'T': []int{
			1, 1, 1, 1, 1,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
		},
		'U': []int{
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		'V': []int{
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			0, 1, 0, 1, 0,
			0, 0, 1, 0, 0,
		},
		'W': []int{
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
			1, 0, 1, 0, 1,
			0, 1, 0, 1, 0,
		},
		'X': []int{
			1, 0, 0, 0, 1,
			0, 1, 0, 1, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 1, 0, 1, 0,
			1, 0, 0, 0, 1,
		},
		'Y': []int{
			1, 0, 0, 0, 1,
			0, 1, 0, 1, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
		},
		'Z': []int{
			1, 1, 1, 1, 1,
			0, 0, 0, 1, 0,
			0, 0, 1, 0, 0,
			0, 1, 0, 0, 0,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
		},
		// Digits
		'0': []int{
			0, 1, 1, 1, 0,
			1, 0, 0, 0, 1,
			1, 0, 0, 1, 1,
			1, 0, 1, 0, 1,
			1, 1, 0, 0, 1,
			0, 1, 1, 1, 0,
		},
		'1': []int{
			0, 0, 1, 0, 0,
			0, 1, 1, 0, 0,
			1, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
		},
		'2': []int{
			0, 1, 1, 1, 0,
			1, 0, 0, 0, 1,
			0, 0, 0, 1, 0,
			0, 0, 1, 0, 0,
			0, 1, 0, 0, 0,
			1, 1, 1, 1, 1,
		},
		'3': []int{
			0, 1, 1, 1, 0,
			1, 0, 0, 0, 1,
			0, 0, 0, 1, 0,
			0, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			0, 1, 1, 1, 0,
		},
		'4': []int{
			0, 0, 0, 1, 0,
			0, 0, 1, 1, 0,
			0, 1, 0, 1, 0,
			1, 1, 1, 1, 0,
			0, 0, 0, 1, 0,
			0, 0, 0, 1, 0,
		},
		'5': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 0,
			0, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			0, 1, 1, 1, 0,
		},
		'6': []int{
			0, 1, 1, 1, 0,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 0,
			1, 0, 0, 0, 1,
			0, 1, 1, 1, 0,
		},
		'7': []int{
			1, 1, 1, 1, 1,
			0, 0, 0, 0, 1,
			0, 0, 0, 1, 0,
			0, 0, 1, 0, 0,
			0, 1, 0, 0, 0,
			1, 0, 0, 0, 0,
		},
		'8': []int{
			0, 1, 1, 1, 0,
			1, 0, 0, 0, 1,
			0, 1, 1, 1, 0,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			0, 1, 1, 1, 0,
		},
		'9': []int{
			0, 1, 1, 1, 0,
			1, 0, 0, 0, 1,
			1, 0, 0, 0, 1,
			0, 1, 1, 1, 1,
			0, 0, 0, 0, 1,
			0, 1, 1, 1, 0,
		},
		// Sings
		'.': []int{
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 1, 1, 0, 0,
			0, 1, 1, 0, 0,
		},
		',': []int{
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 1, 1, 0, 0,
			0, 0, 1, 0, 0,
		},
		';': []int{
			0, 0, 0, 0, 0,
			0, 1, 1, 0, 0,
			0, 1, 1, 0, 0,
			0, 0, 0, 0, 0,
			0, 1, 1, 0, 0,
			0, 0, 1, 0, 0,
		},
		'!': []int{
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 1, 0, 0,
		},
		'-': []int{
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
		},
		'+': []int{
			0, 0, 0, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			1, 1, 1, 1, 1,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
		},
		'/': []int{
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 1,
			0, 0, 0, 1, 0,
			0, 0, 1, 0, 0,
			0, 1, 0, 0, 0,
			1, 0, 0, 0, 0,
		},
		'\\': []int{
			0, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
			0, 1, 0, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 0, 1, 0,
			0, 0, 0, 0, 1,
		},
		':': []int{
			0, 0, 0, 0, 0,
			0, 1, 1, 0, 0,
			0, 1, 1, 0, 0,
			0, 0, 0, 0, 0,
			0, 1, 1, 0, 0,
			0, 1, 1, 0, 0,
		},
		'"': []int{
			0, 1, 0, 1, 0,
			0, 1, 0, 1, 0,
			0, 1, 0, 1, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
		},
		'\'': []int{
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
		},
		'&': []int{
			0, 1, 1, 1, 0,
			0, 1, 0, 1, 0,
			0, 0, 1, 0, 0,
			0, 1, 0, 1, 0,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
		},
		'@': []int{
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 0, 1, 1, 1,
			1, 0, 1, 1, 1,
			1, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
		},
		'*': []int{
			0, 0, 1, 0, 0,
			0, 1, 1, 1, 0,
			0, 0, 1, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
		},
		'(': []int{
			0, 0, 0, 1, 0,
			0, 0, 1, 0, 0,
			0, 1, 0, 0, 0,
			0, 1, 0, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 0, 1, 0,
		},
		')': []int{
			0, 1, 0, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 0, 1, 0,
			0, 0, 0, 1, 0,
			0, 0, 1, 0, 0,
			0, 1, 0, 0, 0,
		},
		'[': []int{
			0, 1, 1, 0, 0,
			0, 1, 0, 0, 0,
			0, 1, 0, 0, 0,
			0, 1, 0, 0, 0,
			0, 1, 0, 0, 0,
			0, 1, 1, 0, 0,
		},
		']': []int{
			0, 0, 1, 1, 0,
			0, 0, 0, 1, 0,
			0, 0, 0, 1, 0,
			0, 0, 0, 1, 0,
			0, 0, 0, 1, 0,
			0, 0, 1, 1, 0,
		},
		'{': []int{
			0, 0, 0, 1, 1,
			0, 0, 1, 0, 0,
			0, 1, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 0, 1, 1,
		},
		'}': []int{
			1, 1, 0, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 1, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			1, 1, 0, 0, 0,
		},
		'^': []int{
			0, 0, 1, 0, 0,
			0, 1, 0, 1, 0,
			1, 0, 0, 0, 1,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
		},
		'#': []int{
			0, 1, 0, 1, 0,
			1, 1, 1, 1, 1,
			0, 1, 0, 1, 0,
			0, 1, 0, 1, 0,
			1, 1, 1, 1, 1,
			0, 1, 0, 1, 0,
		},
		'~': []int{
			0, 0, 0, 0, 0,
			0, 1, 0, 1, 0,
			1, 0, 1, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
		},
		'=': []int{
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
			0, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
			0, 0, 0, 0, 0,
		},
		'%': []int{
			0, 0, 0, 0, 0,
			1, 0, 0, 0, 1,
			0, 0, 0, 1, 0,
			0, 0, 1, 0, 0,
			0, 1, 0, 0, 0,
			1, 0, 0, 0, 1,
		},
		'>': []int{
			0, 0, 0, 0, 0,
			0, 1, 0, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 0, 1, 0,
			0, 0, 1, 0, 0,
			0, 1, 0, 0, 0,
		},
		'<': []int{
			0, 0, 0, 0, 0,
			0, 0, 0, 1, 0,
			0, 0, 1, 0, 0,
			0, 1, 0, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 0, 1, 0,
		},
		'_': []int{
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			1, 1, 1, 1, 1,
		},
		'|': []int{
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
		},
	}
}

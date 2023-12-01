package internal

type Tree struct {
	Next map[rune]*Tree
}

func NewTree() map[rune]*Tree {
	treeMap := map[rune]*Tree{}

	treeMap['o'] = &Tree{
		map[rune]*Tree{
			'n': {
				map[rune]*Tree{
					'e': nil,
				},
			},
		},
	}

	treeMap['t'] = &Tree{
		map[rune]*Tree{
			'w': {
				map[rune]*Tree{
					'o': nil,
				},
			},
			'h': {
				map[rune]*Tree{
					'r': {
						map[rune]*Tree{
							'e': {
								map[rune]*Tree{
									'e': nil,
								},
							},
						},
					},
				},
			},
		},
	}

	treeMap['f'] = &Tree{
		map[rune]*Tree{
			'o': {
				map[rune]*Tree{
					'u': {
						map[rune]*Tree{
							'r': nil,
						},
					},
				},
			},
			'i': {
				map[rune]*Tree{
					'v': {
						map[rune]*Tree{
							'e': nil,
						},
					},
				},
			},
		},
	}

	treeMap['s'] = &Tree{
		map[rune]*Tree{
			'i': {
				map[rune]*Tree{
					'x': nil,
				},
			},
			'e': {
				map[rune]*Tree{
					'v': {
						map[rune]*Tree{
							'e': {
								map[rune]*Tree{
									'n': nil,
								},
							},
						},
					},
				},
			},
		},
	}

	treeMap['e'] = &Tree{
		map[rune]*Tree{
			'i': {
				map[rune]*Tree{
					'g': {
						map[rune]*Tree{
							'h': {
								map[rune]*Tree{
									't': nil,
								},
							},
						},
					},
				},
			},
		},
	}

	treeMap['n'] = &Tree{
		map[rune]*Tree{
			'i': {
				map[rune]*Tree{
					'n': {
						map[rune]*Tree{
							'e': nil,
						},
					},
				},
			},
		},
	}
	return treeMap
}

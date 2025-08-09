package scales

var (
	MandolinChordShapes                 = map[string][]ChordShape{}
	allMandolinChordShapes []ChordShape = []ChordShape{
		mandolin_6_0_ChordShape, mandolin_7_0_ChordShape, mandolin_7b5_0_ChordShape,

		// TODO
		mandolin_7sus_0_ChordShape, mandolin__0_ChordShape, mandolin_add9_0_ChordShape, mandolin_aug6_0_ChordShape, mandolin_aug7_0_ChordShape, mandolin_aug_0_ChordShape,

		// TODO
		mandolin_dim6_0_ChordShape, mandolin_dim7_0_ChordShape, mandolin_dim_0_ChordShape, mandolin_m7_0_ChordShape, mandolin_m7_sharp_5_0_ChordShape,

		// TODO
		mandolin_m7b5_0_ChordShape, mandolin_m_0_ChordShape, mandolin_m_add9_0_ChordShape, mandolin_m_maj7_0_ChordShape, mandolin_maj7_0_ChordShape, mandolin_maj7_sharp_5_0_ChordShape,

		// TODO
		mandolin_maj7b5_0_ChordShape, mandolin_sus2_0_ChordShape, mandolin_sus_0_ChordShape,
	}
)

var (
	mandolin_m_maj7_0_Schema     = []int{}
	mandolin_m_maj7_0_ChordShape = newChordShape("mandolin", []int{0, 3, 7, 11}, mandolin_m_maj7_0_Schema, -1)
)

var (
	mandolin_aug7_0_Schema     = []int{}
	mandolin_aug7_0_ChordShape = newChordShape("mandolin", []int{0, 4, 8, 10}, mandolin_aug7_0_Schema, -1)
)

var (
	mandolin_m_add9_0_Schema     = []int{}
	mandolin_m_add9_0_ChordShape = newChordShape("mandolin", []int{0, 3, 7, 14}, mandolin_m_add9_0_Schema, -1)
)

var (
	mandolin_7_0_Schema     = []int{}
	mandolin_7_0_ChordShape = newChordShape("mandolin", []int{0, 4, 7, 10}, mandolin_7_0_Schema, -1)
)

var (
	mandolin_sus_0_Schema     = []int{} // TODO
	mandolin_sus_0_ChordShape = newChordShape("mandolin", []int{0, 5, 7}, mandolin_sus_0_Schema, -1)
)

var (
	mandolin_m7b5_0_Schema     = []int{} // TODO
	mandolin_m7b5_0_ChordShape = newChordShape("mandolin", []int{0, 3, 6, 10}, mandolin_m7b5_0_Schema, -1)
)

var (
	mandolin_maj7b5_0_Schema     = []int{} // TODO
	mandolin_maj7b5_0_ChordShape = newChordShape("mandolin", []int{0, 4, 6, 11}, mandolin_maj7b5_0_Schema, -1)
)

var (
	mandolin_dim_0_Schema     = []int{} // TODO
	mandolin_dim_0_ChordShape = newChordShape("mandolin", []int{0, 3, 6}, mandolin_dim_0_Schema, -1)
)

var (
	mandolin_m7_0_Schema     = []int{} // TODO
	mandolin_m7_0_ChordShape = newChordShape("mandolin", []int{0, 3, 7, 10}, mandolin_m7_0_Schema, -1)
)

var (
	mandolin_7b5_0_Schema     = []int{} // TODO
	mandolin_7b5_0_ChordShape = newChordShape("mandolin", []int{0, 4, 6, 10}, mandolin_7b5_0_Schema, -1)
)

var (
	mandolin_6_0_Schema     = []int{} // TODO
	mandolin_6_0_ChordShape = newChordShape("mandolin", []int{0, 4, 7, 9}, mandolin_6_0_Schema, -1)
)

var (
	mandolin_dim7_0_Schema     = []int{} // TODO
	mandolin_dim7_0_ChordShape = newChordShape("mandolin", []int{0, 3, 6, 9}, mandolin_dim7_0_Schema, -1)
)

var (
	mandolin_add9_0_Schema     = []int{} // TODO
	mandolin_add9_0_ChordShape = newChordShape("mandolin", []int{0, 4, 7, 14}, mandolin_add9_0_Schema, -1)
)

var (
	mandolin_maj7_sharp_5_0_Schema     = []int{} // TODO
	mandolin_maj7_sharp_5_0_ChordShape = newChordShape("mandolin", []int{0, 4, 8, 11}, mandolin_maj7_sharp_5_0_Schema, -1)
)

var (
	mandolin_aug6_0_Schema     = []int{} // TODO
	mandolin_aug6_0_ChordShape = newChordShape("mandolin", []int{0, 4, 8, 9}, mandolin_aug6_0_Schema, -1)
)

var (
	mandolin_m_0_Schema     = []int{} // TODO
	mandolin_m_0_ChordShape = newChordShape("mandolin", []int{0, 3, 7}, mandolin_m_0_Schema, -1)
)

var (
	mandolin_m7_sharp_5_0_Schema     = []int{} // TODO
	mandolin_m7_sharp_5_0_ChordShape = newChordShape("mandolin", []int{0, 3, 8, 10}, mandolin_m7_sharp_5_0_Schema, -1)
)

var (
	mandolin__0_Schema     = []int{} // TODO
	mandolin__0_ChordShape = newChordShape("mandolin", []int{0, 4, 7}, mandolin__0_Schema, -1)
)

var (
	mandolin_aug_0_Schema     = []int{} // TODO
	mandolin_aug_0_ChordShape = newChordShape("mandolin", []int{0, 4, 8}, mandolin_aug_0_Schema, -1)
)

var (
	mandolin_7sus_0_Schema     = []int{} // TODO
	mandolin_7sus_0_ChordShape = newChordShape("mandolin", []int{0, 5, 7, 10}, mandolin_7sus_0_Schema, -1)
)

var (
	mandolin_maj7_0_Schema     = []int{} // TODO
	mandolin_maj7_0_ChordShape = newChordShape("mandolin", []int{0, 4, 7, 11}, mandolin_maj7_0_Schema, -1)
)

var (
	mandolin_sus2_0_Schema     = []int{} // TODO
	mandolin_sus2_0_ChordShape = newChordShape("mandolin", []int{0, 2, 7}, mandolin_sus2_0_Schema, -1)
)

var (
	mandolin_dim6_0_Schema     = []int{} // TODO
	mandolin_dim6_0_ChordShape = newChordShape("mandolin", []int{0, 3, 6, 8}, mandolin_dim6_0_Schema, -1)
)

package pkg

import "fmt"

const (
	letterWidth  = 5
	letterHeight = 6
)

type Letter [5]byte

// tested letters: ABCEFGHJKLPRUYZ

var (
	LetterA = Letter{0x3E, 0x09, 0x09, 0x3E} //A
	LetterB = Letter{0x3F, 0x25, 0x25, 0x1A}
	LetterC = Letter{0x1E, 0x21, 0x21, 0x12}
	LetterD = Letter{0x3F, 0x21, 0x21, 0x1E}
	LetterE = Letter{0x3F, 0x25, 0x25, 0x21}
	LetterF = Letter{0x3F, 0x05, 0x05, 0x01}
	LetterG = Letter{0x1E, 0x21, 0x29, 0x3A}
	LetterH = Letter{0x3F, 0x04, 0x04, 0x3F}
	LetterI = Letter{0x21, 0x3F, 0x21, 0x00}
	LetterJ = Letter{0x10, 0x20, 0x21, 0x1F}
	LetterK = Letter{0x3F, 0x04, 0x1A, 0x21}
	LetterL = Letter{0x3F, 0x20, 0x20, 0x20}
	LetterM = Letter{0x3F, 0x02, 0x02, 0x3F}
	LetterN = Letter{0x3F, 0x02, 0x04, 0x3F}
	LetterO = Letter{0x1E, 0x21, 0x21, 0x1E}
	LetterP = Letter{0x3F, 0x09, 0x09, 0x06}
	LetterQ = Letter{0x1E, 0x21, 0x11, 0x2E}
	LetterR = Letter{0x3F, 0x09, 0x19, 0x26}
	LetterS = Letter{0x22, 0x25, 0x25, 0x19}
	LetterT = Letter{0x01, 0x3F, 0x01, 0x01}
	LetterU = Letter{0x1F, 0x20, 0x20, 0x1F}
	LetterV = Letter{0x0F, 0x10, 0x20, 0x1F}
	LetterW = Letter{0x3F, 0x10, 0x10, 0x3F}
	LetterX = Letter{0x3B, 0x04, 0x04, 0x3B}
	LetterY = Letter{0x03, 0x04, 0x38, 0x04, 0x03} // so far only the Y letter use the 5th column
	LetterZ = Letter{0x31, 0x29, 0x25, 0x23}       // Z

	AllLetters = map[Letter]string{
		LetterA: "A",
		LetterB: "B",
		LetterC: "C",
		LetterD: "D",
		LetterE: "E",
		LetterF: "F",
		LetterG: "G",
		LetterH: "H",
		LetterI: "I",
		LetterJ: "J",
		LetterK: "K",
		LetterL: "L",
		LetterM: "M",
		LetterN: "N",
		LetterO: "O",
		LetterP: "P",
		LetterQ: "Q",
		LetterR: "R",
		LetterS: "S",
		LetterT: "T",
		LetterU: "U",
		LetterV: "V",
		LetterW: "W",
		LetterX: "X",
		LetterY: "Y",
		LetterZ: "Z",
	}
)

func (l *Letter) SetColor(i, j int) {
	l[i] |= 1 << j
}

func (l *Letter) Color(i, j int) bool {
	return l[i]&(1<<j) > 0
}

func debugLetter(l Letter) string {
	s := ""
	for j := 0; j < letterHeight; j++ {
		for i := range l {
			if l.Color(i, j) {
				s += "1 "
			} else {
				s += "0 "
			}
		}
		s += "\n"
	}
	return s
}

type Word []Letter

// screen is a slice of 1 and 0
// blackColor is the color black (0 or 1)
func NewWordFromScreen(screen []int, wide, blackColor int) Word {
	letterCount := wide / letterWidth
	letters := make(Word, letterCount)
	for i, pixel := range screen {
		if pixel == blackColor {
			letters[(i/letterWidth)%letterCount].SetColor(i%letterWidth, i/wide)
		}
	}
	return letters
}

func (w Word) String() string {
	s := ""
	for _, l := range w {
		letter, exist := AllLetters[l]
		if !exist {
			panic(fmt.Errorf("letter does not exist:\n%v", debugLetter(l)))
		}
		s += letter
	}
	return s
}

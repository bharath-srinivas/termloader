package termloader

// Predefined charsets for the loader.
var Charsets = map[int][]string{
	0: {"▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒", "█▒▒▒▒▒▒▒▒▒▒▒▒▒▒", "████▒▒▒▒▒▒▒▒▒▒▒", "███████▒▒▒▒▒▒▒▒", "██████████▒▒▒▒▒", "█████████████▒▒", "███████████████", "▒▒█████████████", "▒▒▒▒▒██████████", "▒▒▒▒▒▒▒▒███████", "▒▒▒▒▒▒▒▒▒▒▒████", "▒▒▒▒▒▒▒▒▒▒▒▒▒▒█"},
	1: {">>--->", " >>--->", "  >>--->", "   >>--->", "    >>--->", "    <---<<", "   <---<<", "  <---<<", " <---<<", "<---<<"},
	2: {"|", "||", "|||", "||||", "|||||", "|||||||", "||||||||", "|||||||", "||||||", "|||||", "||||", "|||", "||", "|"},
	3: {"(*----------)", "(-*---------)", "(--*--------)", "(---*-------)", "(----*------)", "(-----*-----)", "(------*----)", "(-------*---)", "(--------*--)", "(---------*-)", "(----------*)"},
}

package termloader

import "time"

// Charset config type
type CharsetConfig struct {
	Charset []string
	Delay   time.Duration
}

// Predefined charset configurations for the loader.
var CharsetConfigs = map[string]CharsetConfig{
	"default": {
		[]string{"▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒", "█▒▒▒▒▒▒▒▒▒▒▒▒▒▒", "████▒▒▒▒▒▒▒▒▒▒▒", "███████▒▒▒▒▒▒▒▒", "██████████▒▒▒▒▒", "█████████████▒▒", "███████████████", "▒▒█████████████", "▒▒▒▒▒██████████", "▒▒▒▒▒▒▒▒███████", "▒▒▒▒▒▒▒▒▒▒▒████", "▒▒▒▒▒▒▒▒▒▒▒▒▒▒█"},
		100 * time.Millisecond,
	},
	"dots": {
		[]string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		80 * time.Millisecond,
	},
	"dots2": {
		[]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
		80 * time.Millisecond,
	},
	"dots3": {
		[]string{"⠋", "⠙", "⠚", "⠞", "⠖", "⠦", "⠴", "⠲", "⠳", "⠓"},
		80 * time.Millisecond,
	},
	"dots4": {
		[]string{"⠄", "⠆", "⠇", "⠋", "⠙", "⠸", "⠰", "⠠", "⠰", "⠸", "⠙", "⠋", "⠇", "⠆"},
		80 * time.Millisecond,
	},
	"dots5": {
		[]string{"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"},
		80 * time.Millisecond,
	},
	"dots6": {
		[]string{"⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠴", "⠲", "⠒", "⠂", "⠂", "⠒", "⠚", "⠙", "⠉", "⠁"},
		80 * time.Millisecond,
	},
	"dots7": {
		[]string{"⠈", "⠉", "⠋", "⠓", "⠒", "⠐", "⠐", "⠒", "⠖", "⠦", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈"},
		80 * time.Millisecond,
	},
	"dots8": {
		[]string{"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"},
		80 * time.Millisecond,
	},
	"dots9": {
		[]string{"⢹", "⢺", "⢼", "⣸", "⣇", "⡧", "⡗", "⡏"},
		80 * time.Millisecond,
	},
	"dots10": {
		[]string{"⢄", "⢂", "⢁", "⡁", "⡈", "⡐", "⡠"},
		80 * time.Millisecond,
	},
	"dots11": {
		[]string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
		100 * time.Millisecond,
	},
	"dots12": {
		[]string{"⢀⠀", "⡀⠀", "⠄⠀", "⢂⠀", "⡂⠀", "⠅⠀", "⢃⠀", "⡃⠀", "⠍⠀", "⢋⠀", "⡋⠀", "⠍⠁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⢈⠩", "⡀⢙", "⠄⡙", "⢂⠩", "⡂⢘", "⠅⡘", "⢃⠨", "⡃⢐", "⠍⡐", "⢋⠠", "⡋⢀", "⠍⡁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⠈⠩", "⠀⢙", "⠀⡙", "⠀⠩", "⠀⢘", "⠀⡘", "⠀⠨", "⠀⢐", "⠀⡐", "⠀⠠", "⠀⢀", "⠀⡀"},
		80 * time.Millisecond,
	},
	"line": {
		[]string{"-", "\\", "|", "/"},
		130 * time.Millisecond,
	},
	"line2": {
		[]string{"⠂", "-", "–", "—", "–", "-"},
		100 * time.Millisecond,
	},
	"line3": {
		[]string{"|", "||", "|||", "||||", "|||||", "|||||||", "||||||||", "|||||||", "||||||", "|||||", "||||", "|||", "||", "|"},
		100 * time.Millisecond,
	},
	"pipe": {
		[]string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"},
		100 * time.Millisecond,
	},
	"simpleDots": {
		[]string{".  ", ".. ", "...", "   "},
		400 * time.Millisecond,
	},
	"simpleDotsScrolling": {
		[]string{".  ", ".. ", "...", " ..", "  .", "   "},
		200 * time.Millisecond,
	},
	"star": {
		[]string{"✶", "✸", "✹", "✺", "✹", "✷"},
		70 * time.Millisecond,
	},
	"star2": {
		[]string{"+", "x", "*"},
		80 * time.Millisecond,
	},
	"star3": {
		[]string{"(*----------)", "(-*---------)", "(--*--------)", "(---*-------)", "(----*------)", "(-----*-----)", "(------*----)", "(-------*---)", "(--------*--)", "(---------*-)", "(----------*)"},
		100 * time.Millisecond,
	},
	"flip": {
		[]string{"_", "_", "_", "-", "`", "`", "'", "´", "-", "_", "_", "_"},
		70 * time.Millisecond,
	},
	"hamburger": {
		[]string{"☱", "☲", "☴"},
		100 * time.Millisecond,
	},
	"growVertical": {
		[]string{"▁", "▃", "▄", "▅", "▆", "▇", "▆", "▅", "▄", "▃"},
		120 * time.Millisecond,
	},
	"growHorizontal": {
		[]string{"▏", "▎", "▍", "▌", "▋", "▊", "▉", "▊", "▋", "▌", "▍", "▎"},
		120 * time.Millisecond,
	},
	"balloon": {
		[]string{" ", ".", "o", "O", "@", "*", " "},
		120 * time.Millisecond,
	},
	"balloon2": {
		[]string{".", "o", "O", "°", "O", "o", "."},
		120 * time.Millisecond,
	},
	"noise": {
		[]string{"▓", "▒", "░"},
		100 * time.Millisecond,
	},
	"bounce": {
		[]string{"⠁", "⠂", "⠄", "⠂"},
		120 * time.Millisecond,
	},
	"boxBounce": {
		[]string{"▖", "▘", "▝", "▗"},
		120 * time.Millisecond,
	},
	"boxBounce2": {
		[]string{"▌", "▀", "▐", "▄"},
		100 * time.Millisecond,
	},
	"triangle": {
		[]string{"◢", "◣", "◤", "◥"},
		50 * time.Millisecond,
	},
	"arc": {
		[]string{"◜", "◠", "◝", "◞", "◡", "◟"},
		100 * time.Millisecond,
	},
	"circle": {
		[]string{"◡", "⊙", "◠"},
		120 * time.Millisecond,
	},
	"squareCorners": {
		[]string{"◰", "◳", "◲", "◱"},
		180 * time.Millisecond,
	},
	"circleQuarters": {
		[]string{"◴", "◷", "◶", "◵"},
		120 * time.Millisecond,
	},
	"circleHalves": {
		[]string{"◐", "◓", "◑", "◒"},
		50 * time.Millisecond,
	},
	"squish": {
		[]string{"╫", "╪"},
		100 * time.Millisecond,
	},
	"toggle": {
		[]string{"⊶", "⊷"},
		250 * time.Millisecond,
	},
	"toggle2": {
		[]string{"▫", "▪"},
		80 * time.Millisecond,
	},
	"toggle3": {
		[]string{"□", "■"},
		120 * time.Millisecond,
	},
	"toggle4": {
		[]string{"■", "□", "▪", "▫"},
		100 * time.Millisecond,
	},
	"toggle5": {
		[]string{"▮", "▯"},
		100 * time.Millisecond,
	},
	"toggle6": {
		[]string{"ဝ", "၀"},
		300 * time.Millisecond,
	},
	"toggle7": {
		[]string{"⦾", "⦿"},
		80 * time.Millisecond,
	},
	"toggle8": {
		[]string{"◍", "◌"},
		100 * time.Millisecond,
	},
	"toggle9": {
		[]string{"◉", "◎"},
		100 * time.Millisecond,
	},
	"toggle10": {
		[]string{"㊂", "㊀", "㊁"},
		100 * time.Millisecond,
	},
	"toggle11": {
		[]string{"⧇", "⧆"},
		50 * time.Millisecond,
	},
	"toggle12": {
		[]string{"☗", "☖"},
		120 * time.Millisecond,
	},
	"toggle13": {
		[]string{"=", "*", "-"},
		80 * time.Millisecond,
	},
	"arrow": {
		[]string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
		100 * time.Millisecond,
	},
	"arrow2": {
		[]string{"⬆️ ", "↗️ ", "➡️ ", "↘️ ", "⬇️ ", "↙️ ", "⬅️ ", "↖️ "},
		80 * time.Millisecond,
	},
	"arrow3": {
		[]string{"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"},
		120 * time.Millisecond,
	},
	"arrow4": {
		[]string{">>--->", " >>--->", "  >>--->", "   >>--->", "    >>--->", "    <---<<", "   <---<<", "  <---<<", " <---<<", "<---<<"},
		100 * time.Millisecond,
	},
	"bouncingBar": {
		[]string{"[    ]", "[=   ]", "[==  ]", "[=== ]", "[ ===]", "[  ==]", "[   =]", "[    ]", "[   =]", "[  ==]", "[ ===]", "[====]", "[=== ]", "[==  ]", "[=   ]"},
		80 * time.Millisecond,
	},
	"bouncingBall": {
		[]string{"( ●    )", "(  ●   )", "(   ●  )", "(    ● )", "(     ●)", "(    ● )", "(   ●  )", "(  ●   )", "( ●    )", "(●     )"},
		80 * time.Millisecond,
	},
	"smiley": {
		[]string{"😄 ", "😝 "},
		200 * time.Millisecond,
	},
	"moon": {
		[]string{"🌑 ", "🌒 ", "🌓 ", "🌔 ", "🌕 ", "🌖 ", "🌗 ", "🌘 "},
		80 * time.Millisecond,
	},
	"pong": {
		[]string{"▐⠂       ▌", "▐⠈       ▌", "▐ ⠂      ▌", "▐ ⠠      ▌", "▐  ⡀     ▌", "▐  ⠠     ▌", "▐   ⠂    ▌", "▐   ⠈    ▌", "▐    ⠂   ▌", "▐    ⠠   ▌", "▐     ⡀  ▌", "▐     ⠠  ▌", "▐      ⠂ ▌", "▐      ⠈ ▌", "▐       ⠂▌", "▐       ⠠▌", "▐       ⡀▌", "▐      ⠠ ▌", "▐      ⠂ ▌", "▐     ⠈  ▌", "▐     ⠂  ▌", "▐    ⠠   ▌", "▐    ⡀   ▌", "▐   ⠠    ▌", "▐   ⠂    ▌", "▐  ⠈     ▌", "▐  ⠂     ▌", "▐ ⠠      ▌", "▐ ⡀      ▌", "▐⠠       ▌"},
		80 * time.Millisecond,
	},
	"shark": {
		[]string{"▐|\\____________▌", "▐_|\\___________▌", "▐__|\\__________▌", "▐___|\\_________▌", "▐____|\\________▌", "▐_____|\\_______▌", "▐______|\\______▌", "▐_______|\\_____▌", "▐________|\\____▌", "▐_________|\\___▌", "▐__________|\\__▌", "▐___________|\\_▌", "▐____________|\\▌", "▐____________/|▌", "▐___________/|_▌", "▐__________/|__▌", "▐_________/|___▌", "▐________/|____▌", "▐_______/|_____▌", "▐______/|______▌", "▐_____/|_______▌", "▐____/|________▌", "▐___/|_________▌", "▐__/|__________▌", "▐_/|___________▌", "▐/|____________▌"},
		120 * time.Millisecond,
	},
	"dqpb": {
		[]string{"d", "q", "p", "b"},
		100 * time.Millisecond,
	},
	"grenade": {
		[]string{"،   ", "′   ", " ´ ", " ‾ ", "  ⸌", "  ⸊", "  |", "  ⁎", "  ⁕", " ෴ ", "  ⁓", "   ", "   ", "   "},
		80 * time.Millisecond,
	},
	"point": {
		[]string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"},
		125 * time.Millisecond,
	},
	"layer": {
		[]string{"-", "=", "≡"},
		150 * time.Millisecond,
	},
	"betaWave": {
		[]string{"ρββββββ", "βρβββββ", "ββρββββ", "βββρβββ", "ββββρββ", "βββββρβ", "ββββββρ"},
		80 * time.Millisecond,
	},
}

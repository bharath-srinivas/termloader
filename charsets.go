package termloader

import "time"

// Character set type
type Charset []string

// Predefined charsets for the loader.
var Charsets = map[string]Charset{
	"default":             {"▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒", "█▒▒▒▒▒▒▒▒▒▒▒▒▒▒", "████▒▒▒▒▒▒▒▒▒▒▒", "███████▒▒▒▒▒▒▒▒", "██████████▒▒▒▒▒", "█████████████▒▒", "███████████████", "▒▒█████████████", "▒▒▒▒▒██████████", "▒▒▒▒▒▒▒▒███████", "▒▒▒▒▒▒▒▒▒▒▒████", "▒▒▒▒▒▒▒▒▒▒▒▒▒▒█"},
	"dots":                {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	"dots2":               {"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
	"dots3":               {"⠋", "⠙", "⠚", "⠞", "⠖", "⠦", "⠴", "⠲", "⠳", "⠓"},
	"dots4":               {"⠄", "⠆", "⠇", "⠋", "⠙", "⠸", "⠰", "⠠", "⠰", "⠸", "⠙", "⠋", "⠇", "⠆"},
	"dots5":               {"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"},
	"dots6":               {"⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠴", "⠲", "⠒", "⠂", "⠂", "⠒", "⠚", "⠙", "⠉", "⠁"},
	"dots7":               {"⠈", "⠉", "⠋", "⠓", "⠒", "⠐", "⠐", "⠒", "⠖", "⠦", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈"},
	"dots8":               {"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"},
	"dots9":               {"⢹", "⢺", "⢼", "⣸", "⣇", "⡧", "⡗", "⡏"},
	"dots10":              {"⢄", "⢂", "⢁", "⡁", "⡈", "⡐", "⡠"},
	"dots11":              {"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
	"dots12":              {"⢀⠀", "⡀⠀", "⠄⠀", "⢂⠀", "⡂⠀", "⠅⠀", "⢃⠀", "⡃⠀", "⠍⠀", "⢋⠀", "⡋⠀", "⠍⠁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⢈⠩", "⡀⢙", "⠄⡙", "⢂⠩", "⡂⢘", "⠅⡘", "⢃⠨", "⡃⢐", "⠍⡐", "⢋⠠", "⡋⢀", "⠍⡁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⠈⠩", "⠀⢙", "⠀⡙", "⠀⠩", "⠀⢘", "⠀⡘", "⠀⠨", "⠀⢐", "⠀⡐", "⠀⠠", "⠀⢀", "⠀⡀"},
	"line":                {"-", "\\", "|", "/"},
	"line2":               {"⠂", "-", "–", "—", "–", "-"},
	"line3":               {"|", "||", "|||", "||||", "|||||", "|||||||", "||||||||", "|||||||", "||||||", "|||||", "||||", "|||", "||", "|"},
	"pipe":                {"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"},
	"simpleDots":          {".  ", ".. ", "...", "   "},
	"simpleDotsScrolling": {".  ", ".. ", "...", " ..", "  .", "   "},
	"star":                {"✶", "✸", "✹", "✺", "✹", "✷"},
	"star2":               {"+", "x", "*"},
	"star3":               {"(*----------)", "(-*---------)", "(--*--------)", "(---*-------)", "(----*------)", "(-----*-----)", "(------*----)", "(-------*---)", "(--------*--)", "(---------*-)", "(----------*)"},
	"flip":                {"_", "_", "_", "-", "`", "`", "'", "´", "-", "_", "_", "_"},
	"hamburger":           {"☱", "☲", "☴"},
	"growVertical":        {"▁", "▃", "▄", "▅", "▆", "▇", "▆", "▅", "▄", "▃"},
	"growHorizontal":      {"▏", "▎", "▍", "▌", "▋", "▊", "▉", "▊", "▋", "▌", "▍", "▎"},
	"balloon":             {" ", ".", "o", "O", "@", "*", " "},
	"balloon2":            {".", "o", "O", "°", "O", "o", "."},
	"noise":               {"▓", "▒", "░"},
	"bounce":              {"⠁", "⠂", "⠄", "⠂"},
	"boxBounce":           {"▖", "▘", "▝", "▗"},
	"boxBounce2":          {"▌", "▀", "▐", "▄"},
	"triangle":            {"◢", "◣", "◤", "◥"},
	"arc":                 {"◜", "◠", "◝", "◞", "◡", "◟"},
	"circle":              {"◡", "⊙", "◠"},
	"squareCorners":       {"◰", "◳", "◲", "◱"},
	"circleQuarters":      {"◴", "◷", "◶", "◵"},
	"circleHalves":        {"◐", "◓", "◑", "◒"},
	"squish":              {"╫", "╪"},
	"toggle":              {"⊶", "⊷"},
	"toggle2":             {"▫", "▪"},
	"toggle3":             {"□", "■"},
	"toggle4":             {"■", "□", "▪", "▫"},
	"toggle5":             {"▮", "▯"},
	"toggle6":             {"ဝ", "၀"},
	"toggle7":             {"⦾", "⦿"},
	"toggle8":             {"◍", "◌"},
	"toggle9":             {"◉", "◎"},
	"toggle10":            {"㊂", "㊀", "㊁"},
	"toggle11":            {"⧇", "⧆"},
	"toggle12":            {"☗", "☖"},
	"toggle13":            {"=", "*", "-"},
	"arrow":               {"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
	"arrow2":              {"⬆️ ", "↗️ ", "➡️ ", "↘️ ", "⬇️ ", "↙️ ", "⬅️ ", "↖️ "},
	"arrow3":              {"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"},
	"arrow4":              {">>--->", " >>--->", "  >>--->", "   >>--->", "    >>--->", "    <---<<", "   <---<<", "  <---<<", " <---<<", "<---<<"},
	"bouncingBar":         {"[    ]", "[=   ]", "[==  ]", "[=== ]", "[ ===]", "[  ==]", "[   =]", "[    ]", "[   =]", "[  ==]", "[ ===]", "[====]", "[=== ]", "[==  ]", "[=   ]"},
	"bouncingBall":        {"( ●    )", "(  ●   )", "(   ●  )", "(    ● )", "(     ●)", "(    ● )", "(   ●  )", "(  ●   )", "( ●    )", "(●     )"},
	"smiley":              {"😄 ", "😝 "},
	"moon":                {"🌑 ", "🌒 ", "🌓 ", "🌔 ", "🌕 ", "🌖 ", "🌗 ", "🌘 "},
	"pong":                {"▐⠂       ▌", "▐⠈       ▌", "▐ ⠂      ▌", "▐ ⠠      ▌", "▐  ⡀     ▌", "▐  ⠠     ▌", "▐   ⠂    ▌", "▐   ⠈    ▌", "▐    ⠂   ▌", "▐    ⠠   ▌", "▐     ⡀  ▌", "▐     ⠠  ▌", "▐      ⠂ ▌", "▐      ⠈ ▌", "▐       ⠂▌", "▐       ⠠▌", "▐       ⡀▌", "▐      ⠠ ▌", "▐      ⠂ ▌", "▐     ⠈  ▌", "▐     ⠂  ▌", "▐    ⠠   ▌", "▐    ⡀   ▌", "▐   ⠠    ▌", "▐   ⠂    ▌", "▐  ⠈     ▌", "▐  ⠂     ▌", "▐ ⠠      ▌", "▐ ⡀      ▌", "▐⠠       ▌"},
	"shark":               {"▐|\\____________▌", "▐_|\\___________▌", "▐__|\\__________▌", "▐___|\\_________▌", "▐____|\\________▌", "▐_____|\\_______▌", "▐______|\\______▌", "▐_______|\\_____▌", "▐________|\\____▌", "▐_________|\\___▌", "▐__________|\\__▌", "▐___________|\\_▌", "▐____________|\\▌", "▐____________/|▌", "▐___________/|_▌", "▐__________/|__▌", "▐_________/|___▌", "▐________/|____▌", "▐_______/|_____▌", "▐______/|______▌", "▐_____/|_______▌", "▐____/|________▌", "▐___/|_________▌", "▐__/|__________▌", "▐_/|___________▌", "▐/|____________▌"},
	"dqpb":                {"d", "q", "p", "b"},
	"grenade":             {"،   ", "′   ", " ´ ", " ‾ ", "  ⸌", "  ⸊", "  |", "  ⁎", "  ⁕", " ෴ ", "  ⁓", "   ", "   ", "   "},
	"point":               {"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"},
	"layer":               {"-", "=", "≡"},
	"betaWave":            {"ρββββββ", "βρβββββ", "ββρββββ", "βββρβββ", "ββββρββ", "βββββρβ", "ββββββρ"},
}

// Delay type
type Delay time.Duration

// Predefined delays for the charsets.
var Delays = map[string]time.Duration{
	"default":             100 * time.Millisecond,
	"dots":                80 * time.Millisecond,
	"dots2":               80 * time.Millisecond,
	"dots3":               80 * time.Millisecond,
	"dots4":               80 * time.Millisecond,
	"dots5":               80 * time.Millisecond,
	"dots6":               80 * time.Millisecond,
	"dots7":               80 * time.Millisecond,
	"dots8":               80 * time.Millisecond,
	"dots9":               80 * time.Millisecond,
	"dots10":              80 * time.Millisecond,
	"dots11":              100 * time.Millisecond,
	"dots12":              80 * time.Millisecond,
	"line":                130 * time.Millisecond,
	"line2":               100 * time.Millisecond,
	"line3":               100 * time.Millisecond,
	"pipe":                100 * time.Millisecond,
	"simpleDots":          400 * time.Millisecond,
	"simpleDotsScrolling": 200 * time.Millisecond,
	"star":                70 * time.Millisecond,
	"star2":               80 * time.Millisecond,
	"star3":               100 * time.Millisecond,
	"flip":                70 * time.Millisecond,
	"hamburger":           100 * time.Millisecond,
	"growVertical":        120 * time.Millisecond,
	"growHorizontal":      120 * time.Millisecond,
	"balloon":             120 * time.Millisecond,
	"balloon2":            120 * time.Millisecond,
	"noise":               100 * time.Millisecond,
	"bounce":              120 * time.Millisecond,
	"boxBounce":           120 * time.Millisecond,
	"boxBounce2":          100 * time.Millisecond,
	"triangle":            50 * time.Millisecond,
	"arc":                 100 * time.Millisecond,
	"circle":              120 * time.Millisecond,
	"squareCorners":       180 * time.Millisecond,
	"circleQuarters":      120 * time.Millisecond,
	"circleHalves":        50 * time.Millisecond,
	"squish":              100 * time.Millisecond,
	"toggle":              250 * time.Millisecond,
	"toggle2":             80 * time.Millisecond,
	"toggle3":             120 * time.Millisecond,
	"toggle4":             100 * time.Millisecond,
	"toggle5":             100 * time.Millisecond,
	"toggle6":             300 * time.Millisecond,
	"toggle7":             80 * time.Millisecond,
	"toggle8":             100 * time.Millisecond,
	"toggle9":             100 * time.Millisecond,
	"toggle10":            100 * time.Millisecond,
	"toggle11":            50 * time.Millisecond,
	"toggle12":            120 * time.Millisecond,
	"toggle13":            80 * time.Millisecond,
	"arrow":               100 * time.Millisecond,
	"arrow2":              80 * time.Millisecond,
	"arrow3":              120 * time.Millisecond,
	"arrow4":              100 * time.Millisecond,
	"bouncingBar":         80 * time.Millisecond,
	"bouncingBall":        80 * time.Millisecond,
	"smiley":              200 * time.Millisecond,
	"moon":                80 * time.Millisecond,
	"pong":                80 * time.Millisecond,
	"shark":               120 * time.Millisecond,
	"dqpb":                100 * time.Millisecond,
	"grenade":             80 * time.Millisecond,
	"point":               125 * time.Millisecond,
	"layer":               150 * time.Millisecond,
	"betaWave":            80 * time.Millisecond,
}

package folder

type RGB struct {
	R, G, B int
}

func ColorPicker(colorName string) RGB {
	colors := map[string]RGB{
		"red":         {255, 0, 0},
		"green":       {0, 255, 0},
		"blue":        {0, 0, 255},
		"yellow":      {255, 255, 0},
		"cyan":        {0, 255, 255},
		"magenta":     {255, 0, 255},
		"white":       {255, 255, 255},
		"black":       {0, 0, 0},
		"orange":      {255, 165, 0},
		"purple":      {128, 0, 128},
		"pink":        {255, 192, 203},
		"brown":       {165, 42, 42},
		"teal":        {0, 128, 128},
		"maroon":      {128, 0, 0},
		"navy":        {0, 0, 128},
		"olive":       {128, 128, 0},
		"lime":        {0, 255, 0},
		"skyblue":     {135, 206, 235},
		"salmon":      {250, 128, 114},
		"lightgrey":   {211, 211, 211},
		"darkgrey":    {169, 169, 169},
		"lightblue":   {173, 216, 230},
		"orchid":      {218, 112, 214},
		"tan":         {210, 180, 140},
		"khaki":       {240, 230, 140},
		"lavender":    {230, 230, 250},
		"indigo":      {75, 0, 130},
		"crimson":     {220, 20, 60},
		"gold":        {255, 215, 0},
		"silver":      {192, 192, 192},
		"turquoise":   {64, 224, 208},
		"violet":      {238, 130, 238},
		"darkcyan":    {0, 139, 139},
		"darkmagenta": {139, 0, 139},
		"darkolive":   {85, 107, 47},
		"darkorange":  {255, 140, 0},
		"darkorchid":  {153, 50, 204},
		"darksalmon":  {233, 150, 122},
		"darksilver":  {169, 169, 169},
		"darktan":     {181, 101, 29},
		"darkviolet":  {148, 0, 211},
		"deeppink":    {255, 20, 147},
		"lightcoral":  {240, 128, 128},
		"lightgold":   {255, 223, 0},
		"lightgreen":  {144, 238, 144},
		"lightpink":   {255, 182, 193},
		"lightsilver": {192, 192, 192},
		"lighttan":    {245, 222, 179},
		"lightviolet": {229, 204, 255},
		"navajowhite": {255, 222, 173},
		"palegreen":   {152, 251, 152},
		"peachpuff":   {255, 218, 185},
		"rosybrown":   {188, 143, 143},
		"saddlebrown": {139, 69, 19},
		"seagreen":    {46, 139, 87},
		"slateblue":   {106, 90, 205},
		"slategray":   {112, 128, 144},
		"springgreen": {0, 255, 127},
		"steelblue":   {70, 130, 180},
		"thistle":     {216, 191, 216},
		"tomato":      {255, 99, 71},
		"wheat":       {245, 222, 179},
		"yellowgreen": {154, 205, 50},
		"azure":       {240, 255, 255},
		"beige":       {245, 245, 220},
		"bisque":      {255, 228, 196},
		"coral":       {255, 127, 80},
		"cornflower":  {100, 149, 237},
		"cornsilk":    {255, 248, 220},
		"darkblue":    {0, 0, 139},
		"darkgreen":   {0, 100, 0},
		"darkred":     {139, 0, 0},
		"fuchsia":     {255, 0, 255},
		"ivory":       {255, 255, 240},
		"lightcyan":   {224, 255, 255},
		"lightsalmon": {255, 160, 122},
		"lightyellow": {255, 255, 224},
		"moccasin":    {255, 228, 181},
		"oldlace":     {253, 245, 230},
		"orangered":   {255, 69, 0},
		"papayawhip":  {255, 239, 213},
		"peach":       {255, 229, 180},
		"peru":        {205, 133, 63},
		"plum":        {221, 160, 221},
		"powderblue":  {176, 224, 230},
		"sandybrown":  {244, 164, 96},
		"seashell":    {255, 245, 238},
		"sienna":      {160, 82, 45},
		"snow":        {255, 250, 250},
	}

	return colors[colorName]
}

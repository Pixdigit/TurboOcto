package turboOcto

import "github.com/veandco/go-sdl2/sdl"

type scancodeLookup struct {
	//The Code below is a heavily modified code from go-sdl2 at version https://github.com/veandco/go-sdl2/commit/25677014b163d4ff223cb0de5694fa48169a5b77

	/*
		Copyright (c) 2013, Go-SDL2 Authors
		All rights reserved.

		Redistribution and use in source and binary forms, with or without
		modification, are permitted provided that the following conditions are met:

			* Redistributions of source code must retain the above copyright notice,
		this list of conditions and the following disclaimer.
			* Redistributions in binary form must reproduce the above copyright
		notice, this list of conditions and the following disclaimer in the
		documentation and/or other materials provided with the distribution.
			* Neither the name of Go-SDL2 nor the names of its contributors may be
		used to endorse or promote products derived from this software without specific
		prior written permission.

		THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
		ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
		WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
		DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
		ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
		(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
		LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
		ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
		(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
		SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
	*/

	// The SDL keyboard scancode representation.
	// (https://wiki.libsdl.org/SDL_Scancode)
	// (https://wiki.libsdl.org/SDLScancodeLookup)
	Unknown sdl.Scancode // "" (no name, empty string)

	A sdl.Scancode // "A"
	B sdl.Scancode // "B"
	C sdl.Scancode // "C"
	D sdl.Scancode // "D"
	E sdl.Scancode // "E"
	F sdl.Scancode // "F"
	G sdl.Scancode // "G"
	H sdl.Scancode // "H"
	I sdl.Scancode // "I"
	J sdl.Scancode // "J"
	K sdl.Scancode // "K"
	L sdl.Scancode // "L"
	M sdl.Scancode // "M"
	N sdl.Scancode // "N"
	O sdl.Scancode // "O"
	P sdl.Scancode // "P"
	Q sdl.Scancode // "Q"
	R sdl.Scancode // "R"
	S sdl.Scancode // "S"
	T sdl.Scancode // "T"
	U sdl.Scancode // "U"
	V sdl.Scancode // "V"
	W sdl.Scancode // "W"
	X sdl.Scancode // "X"
	Y sdl.Scancode // "Y"
	Z sdl.Scancode // "Z"

	One   sdl.Scancode // "1"
	Two   sdl.Scancode // "2"
	Three sdl.Scancode // "3"
	Four  sdl.Scancode // "4"
	Five  sdl.Scancode // "5"
	Six   sdl.Scancode // "6"
	Seven sdl.Scancode // "7"
	Eight sdl.Scancode // "8"
	Nine  sdl.Scancode // "9"
	Zero  sdl.Scancode // "0"

	Return    sdl.Scancode // "Return"
	Escape    sdl.Scancode // "Escape" (the Esc key)
	Backspace sdl.Scancode // "Backspace"
	Tab       sdl.Scancode // "Tab" (the Tab key)
	Space     sdl.Scancode // "Space" (the Space Bar key(s))

	Minus        sdl.Scancode // "-"
	Equals       sdl.Scancode // "="
	Leftbracket  sdl.Scancode // "["
	Rightbracket sdl.Scancode // "]"
	Backslash    sdl.Scancode // "\"
	Nonushash    sdl.Scancode // "#" (ISO USB keyboards actually use this code instead of 49 for the same key, but all OSes I've seen treat the two codes identically. So, as an implementor, unless your keyboard generates both of those codes and your OS treats them differently, you should generate SDL_SCANCODE_BACKSLASH instead of this code. As a user, you should not rely on this code because SDL will never generate it with most (all?) keyboards.)
	Semicolon    sdl.Scancode // ";"
	Apostrophe   sdl.Scancode // "'"
	Grave        sdl.Scancode // "`"
	Comma        sdl.Scancode // ","
	Period       sdl.Scancode // "."
	Slash        sdl.Scancode // "/"
	Capslock     sdl.Scancode // "CapsLock"
	F1           sdl.Scancode // "F1"
	F2           sdl.Scancode // "F2"
	F3           sdl.Scancode // "F3"
	F4           sdl.Scancode // "F4"
	F5           sdl.Scancode // "F5"
	F6           sdl.Scancode // "F6"
	F7           sdl.Scancode // "F7"
	F8           sdl.Scancode // "F8"
	F9           sdl.Scancode // "F9"
	F10          sdl.Scancode // "F10"
	F11          sdl.Scancode // "F11"
	F12          sdl.Scancode // "F12"
	Printscreen  sdl.Scancode // "PrintScreen"
	Scrolllock   sdl.Scancode // "ScrollLock"
	Pause        sdl.Scancode // "Pause" (the Pause / Break key)
	Insert       sdl.Scancode // "Insert" (insert on PC, help on some Mac keyboards (but does send code 73, not 117))
	Home         sdl.Scancode // "Home"
	Pageup       sdl.Scancode // "PageUp"
	Delete       sdl.Scancode // "Delete"
	End          sdl.Scancode // "End"
	Pagedown     sdl.Scancode // "PageDown"
	Right        sdl.Scancode // "Right" (the Right arrow key (navigation keypad))
	Left         sdl.Scancode // "Left" (the Left arrow key (navigation keypad))
	Down         sdl.Scancode // "Down" (the Down arrow key (navigation keypad))
	Up           sdl.Scancode // "Up" (the Up arrow key (navigation keypad))

	Numlockclear sdl.Scancode // "Numlock" (the Num Lock key (PC) / the Clear key (Mac))
	KP_divide    sdl.Scancode // "Keypad /" (the / key (numeric keypad))
	KP_multiply  sdl.Scancode // "Keypad *" (the * key (numeric keypad))
	KP_minus     sdl.Scancode // "Keypad -" (the - key (numeric keypad))
	KP_plus      sdl.Scancode // "Keypad +" (the + key (numeric keypad))
	KP_enter     sdl.Scancode // "Keypad Enter" (the Enter key (numeric keypad))
	KP_1         sdl.Scancode // "Keypad 1" (the 1 key (numeric keypad))
	KP_2         sdl.Scancode // "Keypad 2" (the 2 key (numeric keypad))
	KP_3         sdl.Scancode // "Keypad 3" (the 3 key (numeric keypad))
	KP_4         sdl.Scancode // "Keypad 4" (the 4 key (numeric keypad))
	KP_5         sdl.Scancode // "Keypad 5" (the 5 key (numeric keypad))
	KP_6         sdl.Scancode // "Keypad 6" (the 6 key (numeric keypad))
	KP_7         sdl.Scancode // "Keypad 7" (the 7 key (numeric keypad))
	KP_8         sdl.Scancode // "Keypad 8" (the 8 key (numeric keypad))
	KP_9         sdl.Scancode // "Keypad 9" (the 9 key (numeric keypad))
	KP_0         sdl.Scancode // "Keypad 0" (the 0 key (numeric keypad))
	KP_period    sdl.Scancode // "Keypad ." (the . key (numeric keypad))

	Nonusbackslash sdl.Scancode // "" (no name, empty string; This is the additional key that ISO keyboards have over ANSI ones, located between left shift and Y. Produces GRAVE ACCENT and TILDE in a US or UK Mac layout, REVERSE SOLIDUS (backslash) and VERTICAL LINE in a US or UK Windows layout, and LESS-THAN SIGN and GREATER-THAN SIGN in a Swiss German, German, or French layout.)
	Application    sdl.Scancode // "Application" (the Application / Compose / Context Menu (Windows) key)
	Power          sdl.Scancode // "Power" (The USB document says this is a status flag, not a physical key - but some Mac keyboards do have a power key.)
	KP_equals      sdl.Scancode // "Keypad =" (the = key (numeric keypad))
	F13            sdl.Scancode // "F13"
	F14            sdl.Scancode // "F14"
	F15            sdl.Scancode // "F15"
	F16            sdl.Scancode // "F16"
	F17            sdl.Scancode // "F17"
	F18            sdl.Scancode // "F18"
	F19            sdl.Scancode // "F19"
	F20            sdl.Scancode // "F20"
	F21            sdl.Scancode // "F21"
	F22            sdl.Scancode // "F22"
	F23            sdl.Scancode // "F23"
	F24            sdl.Scancode // "F24"
	Execute        sdl.Scancode // "Execute"
	Help           sdl.Scancode // "Help"
	Menu           sdl.Scancode // "Menu"
	Select         sdl.Scancode // "Select"
	Stop           sdl.Scancode // "Stop"
	Again          sdl.Scancode // "Again" (the Again key (Redo))
	Undo           sdl.Scancode // "Undo"
	Cut            sdl.Scancode // "Cut"
	Copy           sdl.Scancode // "Copy"
	Paste          sdl.Scancode // "Paste"
	Find           sdl.Scancode // "Find"
	Mute           sdl.Scancode // "Mute"
	Volumeup       sdl.Scancode // "VolumeUp"
	Volumedown     sdl.Scancode // "VolumeDown"
	KP_comma       sdl.Scancode // "Keypad ," (the Comma key (numeric keypad))
	KP_equalsas400 sdl.Scancode // "Keypad = (AS400)" (the Equals AS400 key (numeric keypad))

	International1 sdl.Scancode // "" (no name, empty string; used on Asian keyboards, see footnotes in USB doc)
	International2 sdl.Scancode // "" (no name, empty string)
	International3 sdl.Scancode // "" (no name, empty string; Yen)
	International4 sdl.Scancode // "" (no name, empty string)
	International5 sdl.Scancode // "" (no name, empty string)
	International6 sdl.Scancode // "" (no name, empty string)
	International7 sdl.Scancode // "" (no name, empty string)
	International8 sdl.Scancode // "" (no name, empty string)
	International9 sdl.Scancode // "" (no name, empty string)
	Lang1          sdl.Scancode // "" (no name, empty string; Hangul/English toggle)
	Lang2          sdl.Scancode // "" (no name, empty string; Hanja conversion)
	Lang3          sdl.Scancode // "" (no name, empty string; Katakana)
	Lang4          sdl.Scancode // "" (no name, empty string; Hiragana)
	Lang5          sdl.Scancode // "" (no name, empty string; Zenkaku/Hankaku)
	Lang6          sdl.Scancode // "" (no name, empty string; reserved)
	Lang7          sdl.Scancode // "" (no name, empty string; reserved)
	Lang8          sdl.Scancode // "" (no name, empty string; reserved)
	Lang9          sdl.Scancode // "" (no name, empty string; reserved)

	Alterase   sdl.Scancode // "AltErase" (Erase-Eaze)
	Sysreq     sdl.Scancode // "SysReq" (the SysReq key)
	Cancel     sdl.Scancode // "Cancel"
	Clear      sdl.Scancode // "Clear"
	Prior      sdl.Scancode // "Prior"
	Return2    sdl.Scancode // "Return"
	Separator  sdl.Scancode // "Separator"
	Out        sdl.Scancode // "Out"
	Oper       sdl.Scancode // "Oper"
	Clearagain sdl.Scancode // "Clear / Again"
	Crsel      sdl.Scancode // "CrSel"
	Exsel      sdl.Scancode // "ExSel"

	KP_00              sdl.Scancode // "Keypad 00" (the 00 key (numeric keypad))
	KP_000             sdl.Scancode // "Keypad 000" (the 000 key (numeric keypad))
	Thousandsseparator sdl.Scancode // "ThousandsSeparator" (the Thousands Separator key)
	Decimalseparator   sdl.Scancode // "DecimalSeparator" (the Decimal Separator key)
	Currencyunit       sdl.Scancode // "CurrencyUnit" (the Currency Unit key)
	Currencysubunit    sdl.Scancode // "CurrencySubUnit" (the Currency Subunit key)
	KP_leftparen       sdl.Scancode // "Keypad (" (the Left Parenthesis key (numeric keypad))
	KP_rightparen      sdl.Scancode // "Keypad )" (the Right Parenthesis key (numeric keypad))
	KP_leftbrace       sdl.Scancode // "Keypad {" (the Left Brace key (numeric keypad))
	KP_rightbrace      sdl.Scancode // "Keypad }" (the Right Brace key (numeric keypad))
	KP_tab             sdl.Scancode // "Keypad Tab" (the Tab key (numeric keypad))
	KP_backspace       sdl.Scancode // "Keypad Backspace" (the Backspace key (numeric keypad))
	KP_a               sdl.Scancode // "Keypad A" (the A key (numeric keypad))
	KP_b               sdl.Scancode // "Keypad B" (the B key (numeric keypad))
	KP_c               sdl.Scancode // "Keypad C" (the C key (numeric keypad))
	KP_d               sdl.Scancode // "Keypad D" (the D key (numeric keypad))
	KP_e               sdl.Scancode // "Keypad E" (the E key (numeric keypad))
	KP_f               sdl.Scancode // "Keypad F" (the F key (numeric keypad))
	KP_xor             sdl.Scancode // "Keypad XOR" (the XOR key (numeric keypad))
	KP_power           sdl.Scancode // "Keypad ^" (the Power key (numeric keypad))
	KP_percent         sdl.Scancode // "Keypad %" (the Percent key (numeric keypad))
	KP_less            sdl.Scancode // "Keypad <" (the Less key (numeric keypad))
	KP_greater         sdl.Scancode // "Keypad >" (the Greater key (numeric keypad))
	KP_ampersand       sdl.Scancode // "Keypad &" (the & key (numeric keypad))
	KP_dblampersand    sdl.Scancode // "Keypad &&" (the && key (numeric keypad))
	KP_verticalbar     sdl.Scancode // "Keypad |" (the | key (numeric keypad))
	KP_dblverticalbar  sdl.Scancode // "Keypad ||" (the || key (numeric keypad))
	KP_colon           sdl.Scancode // "Keypad :" (the : key (numeric keypad))
	KP_hash            sdl.Scancode // "Keypad #" (the # key (numeric keypad))
	KP_space           sdl.Scancode // "Keypad Space" (the Space key (numeric keypad))
	KP_at              sdl.Scancode // "Keypad @" (the @ key (numeric keypad))
	KP_exclam          sdl.Scancode // "Keypad !" (the ! key (numeric keypad))
	KP_memstore        sdl.Scancode // "Keypad MemStore" (the Mem Store key (numeric keypad))
	KP_memrecall       sdl.Scancode // "Keypad MemRecall" (the Mem Recall key (numeric keypad))
	KP_memclear        sdl.Scancode // "Keypad MemClear" (the Mem Clear key (numeric keypad))
	KP_memadd          sdl.Scancode // "Keypad MemAdd" (the Mem Add key (numeric keypad))
	KP_memsubtract     sdl.Scancode // "Keypad MemSubtract" (the Mem Subtract key (numeric keypad))
	KP_memmultiply     sdl.Scancode // "Keypad MemMultiply" (the Mem Multiply key (numeric keypad))
	KP_memdivide       sdl.Scancode // "Keypad MemDivide" (the Mem Divide key (numeric keypad))
	KP_plusminus       sdl.Scancode // "Keypad +/-" (the +/- key (numeric keypad))
	KP_clear           sdl.Scancode // "Keypad Clear" (the Clear key (numeric keypad))
	KP_clearentry      sdl.Scancode // "Keypad ClearEntry" (the Clear Entry key (numeric keypad))
	KP_binary          sdl.Scancode // "Keypad Binary" (the Binary key (numeric keypad))
	KP_octal           sdl.Scancode // "Keypad Octal" (the Octal key (numeric keypad))
	KP_decimal         sdl.Scancode // "Keypad Decimal" (the Decimal key (numeric keypad))
	KP_hexadecimal     sdl.Scancode // "Keypad Hexadecimal" (the Hexadecimal key (numeric keypad))

	Lctrl          sdl.Scancode // "Left Ctrl"
	Lshift         sdl.Scancode // "Left Shift"
	Lalt           sdl.Scancode // "Left Alt" (alt, option)
	Lgui           sdl.Scancode // "Left GUI" (windows, command (apple), meta)
	Rctrl          sdl.Scancode // "Right Ctrl"
	Rshift         sdl.Scancode // "Right Shift"
	Ralt           sdl.Scancode // "Right Alt" (alt gr, option)
	Rgui           sdl.Scancode // "Right GUI" (windows, command (apple), meta)
	Mode           sdl.Scancode // "ModeSwitch" (I'm not sure if this is really not covered by any of the above, but since there's a special KMOD_MODE for it I'm adding it here)
	Audionext      sdl.Scancode // "AudioNext" (the Next Track media key)
	Audioprev      sdl.Scancode // "AudioPrev" (the Previous Track media key)
	Audiostop      sdl.Scancode // "AudioStop" (the Stop media key)
	Audioplay      sdl.Scancode // "AudioPlay" (the Play media key)
	Audiomute      sdl.Scancode // "AudioMute" (the Mute volume key)
	Mediaselect    sdl.Scancode // "MediaSelect" (the Media Select key)
	Www            sdl.Scancode // "WWW" (the WWW/World Wide Web key)
	Mail           sdl.Scancode // "Mail" (the Mail/eMail key)
	Calculator     sdl.Scancode // "Calculator" (the Calculator key)
	Computer       sdl.Scancode // "Computer" (the My Computer key)
	AC_search      sdl.Scancode // "AC Search" (the Search key (application control keypad))
	AC_home        sdl.Scancode // "AC Home" (the Home key (application control keypad))
	AC_back        sdl.Scancode // "AC Back" (the Back key (application control keypad))
	AC_forward     sdl.Scancode // "AC Forward" (the Forward key (application control keypad))
	AC_stop        sdl.Scancode // "AC Stop" (the Stop key (application control keypad))
	AC_refresh     sdl.Scancode // "AC Refresh" (the Refresh key (application control keypad))
	AC_bookmarks   sdl.Scancode // "AC Bookmarks" (the Bookmarks key (application control keypad))
	Brightnessdown sdl.Scancode // "BrightnessDown" (the Brightness Down key)
	Brightnessup   sdl.Scancode // "BrightnessUp" (the Brightness Up key)
	Displayswitch  sdl.Scancode // "DisplaySwitch" (display mirroring/dual display switch, video mode switch)
	Kbdillumtoggle sdl.Scancode // "KBDIllumToggle" (the Keyboard Illumination Toggle key)
	Kbdillumdown   sdl.Scancode // "KBDIllumDown" (the Keyboard Illumination Down key)
	Kbdillumup     sdl.Scancode // "KBDIllumUp" (the Keyboard Illumination Up key)
	Eject          sdl.Scancode // "Eject" (the Eject key)
	Sleep          sdl.Scancode // "Sleep" (the Sleep key)
	App1           sdl.Scancode
	App2           sdl.Scancode
}

var Scancodes scancodeLookup

func init() {
	Scancodes = scancodeLookup{
		Unknown: 0, //  (no name, empty string)

		A: sdl.SCANCODE_A, // A
		B: sdl.SCANCODE_B, // B
		C: sdl.SCANCODE_C, // C
		D: sdl.SCANCODE_D, // D
		E: sdl.SCANCODE_E, // E
		F: sdl.SCANCODE_F, // F
		G: sdl.SCANCODE_G, // G
		H: sdl.SCANCODE_H, // H
		I: sdl.SCANCODE_I, // I
		J: sdl.SCANCODE_J, // J
		K: sdl.SCANCODE_K, // K
		L: sdl.SCANCODE_L, // L
		M: sdl.SCANCODE_M, // M
		N: sdl.SCANCODE_N, // N
		O: sdl.SCANCODE_O, // O
		P: sdl.SCANCODE_P, // P
		Q: sdl.SCANCODE_Q, // Q
		R: sdl.SCANCODE_R, // R
		S: sdl.SCANCODE_S, // S
		T: sdl.SCANCODE_T, // T
		U: sdl.SCANCODE_U, // U
		V: sdl.SCANCODE_V, // V
		W: sdl.SCANCODE_W, // W
		X: sdl.SCANCODE_X, // X
		Y: sdl.SCANCODE_Y, // Y
		Z: sdl.SCANCODE_Z, // Z

		One:   sdl.SCANCODE_1, // 1
		Two:   sdl.SCANCODE_2, // 2
		Three: sdl.SCANCODE_3, // 3
		Four:  sdl.SCANCODE_4, // 4
		Five:  sdl.SCANCODE_5, // 5
		Six:   sdl.SCANCODE_6, // 6
		Seven: sdl.SCANCODE_7, // 7
		Eight: sdl.SCANCODE_8, // 8
		Nine:  sdl.SCANCODE_9, // 9
		Zero:  sdl.SCANCODE_0, // 0

		Return:    sdl.SCANCODE_RETURN,    // Return
		Escape:    sdl.SCANCODE_ESCAPE,    // Escape (the Esc key)
		Backspace: sdl.SCANCODE_BACKSPACE, // Backspace
		Tab:       sdl.SCANCODE_TAB,       // Tab (the Tab key)
		Space:     sdl.SCANCODE_SPACE,     // Space (the Space Bar key(s))

		Minus:        sdl.SCANCODE_MINUS,        // -
		Equals:       sdl.SCANCODE_EQUALS,       // :
		Leftbracket:  sdl.SCANCODE_LEFTBRACKET,  // [
		Rightbracket: sdl.SCANCODE_RIGHTBRACKET, // ]
		Backslash:    sdl.SCANCODE_BACKSLASH,    // \
		Nonushash:    sdl.SCANCODE_NONUSHASH,    // # (ISO USB keyboards actually use this code instead of 49 for the same key, but all OSes I've seen treat the two codes identically. So, as an implementor, unless your keyboard generates both of those codes and your OS treats them differently, you should generate SDL_SCANCODE_BACKSLASH instead of this code. As a user, you should not rely on this code because SDL will never generate it with most (all?) keyboards.)
		Semicolon:    sdl.SCANCODE_SEMICOLON,    // ;
		Apostrophe:   sdl.SCANCODE_APOSTROPHE,   // '
		Grave:        sdl.SCANCODE_GRAVE,        // `
		Comma:        sdl.SCANCODE_COMMA,        // ,
		Period:       sdl.SCANCODE_PERIOD,       // .
		Slash:        sdl.SCANCODE_SLASH,        // /
		Capslock:     sdl.SCANCODE_CAPSLOCK,     // CapsLock
		F1:           sdl.SCANCODE_F1,           // F1
		F2:           sdl.SCANCODE_F2,           // F2
		F3:           sdl.SCANCODE_F3,           // F3
		F4:           sdl.SCANCODE_F4,           // F4
		F5:           sdl.SCANCODE_F5,           // F5
		F6:           sdl.SCANCODE_F6,           // F6
		F7:           sdl.SCANCODE_F7,           // F7
		F8:           sdl.SCANCODE_F8,           // F8
		F9:           sdl.SCANCODE_F9,           // F9
		F10:          sdl.SCANCODE_F10,          // F10
		F11:          sdl.SCANCODE_F11,          // F11
		F12:          sdl.SCANCODE_F12,          // F12
		Printscreen:  sdl.SCANCODE_PRINTSCREEN,  // PrintScreen
		Scrolllock:   sdl.SCANCODE_SCROLLLOCK,   // ScrollLock
		Pause:        sdl.SCANCODE_PAUSE,        // Pause (the Pause / Break key)
		Insert:       sdl.SCANCODE_INSERT,       // Insert (insert on PC, help on some Mac keyboards (but does send code 73, not 117))
		Home:         sdl.SCANCODE_HOME,         // Home
		Pageup:       sdl.SCANCODE_PAGEUP,       // PageUp
		Delete:       sdl.SCANCODE_DELETE,       // Delete
		End:          sdl.SCANCODE_END,          // End
		Pagedown:     sdl.SCANCODE_PAGEDOWN,     // PageDown
		Right:        sdl.SCANCODE_RIGHT,        // Right (the Right arrow key (navigation keypad))
		Left:         sdl.SCANCODE_LEFT,         // Left (the Left arrow key (navigation keypad))
		Down:         sdl.SCANCODE_DOWN,         // Down (the Down arrow key (navigation keypad))
		Up:           sdl.SCANCODE_UP,           // Up (the Up arrow key (navigation keypad))

		Numlockclear: sdl.SCANCODE_NUMLOCKCLEAR, // Numlock (the Num Lock key (PC) / the Clear key (Mac))
		KP_divide:    sdl.SCANCODE_KP_DIVIDE,    // Keypad / (the / key (numeric keypad))
		KP_multiply:  sdl.SCANCODE_KP_MULTIPLY,  // Keypad * (the * key (numeric keypad))
		KP_minus:     sdl.SCANCODE_KP_MINUS,     // Keypad - (the - key (numeric keypad))
		KP_plus:      sdl.SCANCODE_KP_PLUS,      // Keypad + (the + key (numeric keypad))
		KP_enter:     sdl.SCANCODE_KP_ENTER,     // Keypad Enter (the Enter key (numeric keypad))
		KP_1:         sdl.SCANCODE_KP_1,         // Keypad 1 (the 1 key (numeric keypad))
		KP_2:         sdl.SCANCODE_KP_2,         // Keypad 2 (the 2 key (numeric keypad))
		KP_3:         sdl.SCANCODE_KP_3,         // Keypad 3 (the 3 key (numeric keypad))
		KP_4:         sdl.SCANCODE_KP_4,         // Keypad 4 (the 4 key (numeric keypad))
		KP_5:         sdl.SCANCODE_KP_5,         // Keypad 5 (the 5 key (numeric keypad))
		KP_6:         sdl.SCANCODE_KP_6,         // Keypad 6 (the 6 key (numeric keypad))
		KP_7:         sdl.SCANCODE_KP_7,         // Keypad 7 (the 7 key (numeric keypad))
		KP_8:         sdl.SCANCODE_KP_8,         // Keypad 8 (the 8 key (numeric keypad))
		KP_9:         sdl.SCANCODE_KP_9,         // Keypad 9 (the 9 key (numeric keypad))
		KP_0:         sdl.SCANCODE_KP_0,         // Keypad 0 (the 0 key (numeric keypad))
		KP_period:    sdl.SCANCODE_KP_PERIOD,    // Keypad . (the . key (numeric keypad))

		Nonusbackslash: sdl.SCANCODE_NONUSBACKSLASH, //  (no name, empty string; This is the additional key that ISO keyboards have over ANSI ones, located between left shift and Y. Produces GRAVE ACCENT and TILDE in a US or UK Mac layout, REVERSE SOLIDUS (backslash) and VERTICAL LINE in a US or UK Windows layout, and LESS-THAN SIGN and GREATER-THAN SIGN in a Swiss German, German, or French layout.)
		Application:    sdl.SCANCODE_APPLICATION,    // Application (the Application / Compose / Context Menu (Windows) key)
		Power:          sdl.SCANCODE_POWER,          // Power (The USB document says this is a status flag, not a physical key - but some Mac keyboards do have a power key.)
		KP_equals:      sdl.SCANCODE_KP_EQUALS,      // Keypad : (the : key (numeric keypad))
		F13:            sdl.SCANCODE_F13,            // F13
		F14:            sdl.SCANCODE_F14,            // F14
		F15:            sdl.SCANCODE_F15,            // F15
		F16:            sdl.SCANCODE_F16,            // F16
		F17:            sdl.SCANCODE_F17,            // F17
		F18:            sdl.SCANCODE_F18,            // F18
		F19:            sdl.SCANCODE_F19,            // F19
		F20:            sdl.SCANCODE_F20,            // F20
		F21:            sdl.SCANCODE_F21,            // F21
		F22:            sdl.SCANCODE_F22,            // F22
		F23:            sdl.SCANCODE_F23,            // F23
		F24:            sdl.SCANCODE_F24,            // F24
		Execute:        sdl.SCANCODE_EXECUTE,        // Execute
		Help:           sdl.SCANCODE_HELP,           // Help
		Menu:           sdl.SCANCODE_MENU,           // Menu
		Select:         sdl.SCANCODE_SELECT,         // Select
		Stop:           sdl.SCANCODE_STOP,           // Stop
		Again:          sdl.SCANCODE_AGAIN,          // Again (the Again key (Redo))
		Undo:           sdl.SCANCODE_UNDO,           // Undo
		Cut:            sdl.SCANCODE_CUT,            // Cut
		Copy:           sdl.SCANCODE_COPY,           // Copy
		Paste:          sdl.SCANCODE_PASTE,          // Paste
		Find:           sdl.SCANCODE_FIND,           // Find
		Mute:           sdl.SCANCODE_MUTE,           // Mute
		Volumeup:       sdl.SCANCODE_VOLUMEUP,       // VolumeUp
		Volumedown:     sdl.SCANCODE_VOLUMEDOWN,     // VolumeDown
		KP_comma:       sdl.SCANCODE_KP_COMMA,       // Keypad , (the Comma key (numeric keypad))
		KP_equalsas400: sdl.SCANCODE_KP_EQUALSAS400, // Keypad : (AS400) (the Equals AS400 key (numeric keypad))

		International1: sdl.SCANCODE_INTERNATIONAL1, //  (no name, empty string; used on Asian keyboards, see footnotes in USB doc)
		International2: sdl.SCANCODE_INTERNATIONAL2, //  (no name, empty string)
		International3: sdl.SCANCODE_INTERNATIONAL3, //  (no name, empty string; Yen)
		International4: sdl.SCANCODE_INTERNATIONAL4, //  (no name, empty string)
		International5: sdl.SCANCODE_INTERNATIONAL5, //  (no name, empty string)
		International6: sdl.SCANCODE_INTERNATIONAL6, //  (no name, empty string)
		International7: sdl.SCANCODE_INTERNATIONAL7, //  (no name, empty string)
		International8: sdl.SCANCODE_INTERNATIONAL8, //  (no name, empty string)
		International9: sdl.SCANCODE_INTERNATIONAL9, //  (no name, empty string)
		Lang1:          sdl.SCANCODE_LANG1,          //  (no name, empty string; Hangul/English toggle)
		Lang2:          sdl.SCANCODE_LANG2,          //  (no name, empty string; Hanja conversion)
		Lang3:          sdl.SCANCODE_LANG3,          //  (no name, empty string; Katakana)
		Lang4:          sdl.SCANCODE_LANG4,          //  (no name, empty string; Hiragana)
		Lang5:          sdl.SCANCODE_LANG5,          //  (no name, empty string; Zenkaku/Hankaku)
		Lang6:          sdl.SCANCODE_LANG6,          //  (no name, empty string; reserved)
		Lang7:          sdl.SCANCODE_LANG7,          //  (no name, empty string; reserved)
		Lang8:          sdl.SCANCODE_LANG8,          //  (no name, empty string; reserved)
		Lang9:          sdl.SCANCODE_LANG9,          //  (no name, empty string; reserved)

		Alterase:   sdl.SCANCODE_ALTERASE,   // AltErase (Erase-Eaze)
		Sysreq:     sdl.SCANCODE_SYSREQ,     // SysReq (the SysReq key)
		Cancel:     sdl.SCANCODE_CANCEL,     // Cancel
		Clear:      sdl.SCANCODE_CLEAR,      // Clear
		Prior:      sdl.SCANCODE_PRIOR,      // Prior
		Return2:    sdl.SCANCODE_RETURN2,    // Return
		Separator:  sdl.SCANCODE_SEPARATOR,  // Separator
		Out:        sdl.SCANCODE_OUT,        // Out
		Oper:       sdl.SCANCODE_OPER,       // Oper
		Clearagain: sdl.SCANCODE_CLEARAGAIN, // Clear / Again
		Crsel:      sdl.SCANCODE_CRSEL,      // CrSel
		Exsel:      sdl.SCANCODE_EXSEL,      // ExSel

		KP_00:              sdl.SCANCODE_KP_00,              // Keypad 00 (the 00 key (numeric keypad))
		KP_000:             sdl.SCANCODE_KP_000,             // Keypad 000 (the 000 key (numeric keypad))
		Thousandsseparator: sdl.SCANCODE_THOUSANDSSEPARATOR, // ThousandsSeparator (the Thousands Separator key)
		Decimalseparator:   sdl.SCANCODE_DECIMALSEPARATOR,   // DecimalSeparator (the Decimal Separator key)
		Currencyunit:       sdl.SCANCODE_CURRENCYUNIT,       // CurrencyUnit (the Currency Unit key)
		Currencysubunit:    sdl.SCANCODE_CURRENCYSUBUNIT,    // CurrencySubUnit (the Currency Subunit key)
		KP_leftparen:       sdl.SCANCODE_KP_LEFTPAREN,       // Keypad ( (the Left Parenthesis key (numeric keypad))
		KP_rightparen:      sdl.SCANCODE_KP_RIGHTPAREN,      // Keypad ) (the Right Parenthesis key (numeric keypad))
		KP_leftbrace:       sdl.SCANCODE_KP_LEFTBRACE,       // Keypad { (the Left Brace key (numeric keypad))
		KP_rightbrace:      sdl.SCANCODE_KP_RIGHTBRACE,      // Keypad } (the Right Brace key (numeric keypad))
		KP_tab:             sdl.SCANCODE_KP_TAB,             // Keypad Tab (the Tab key (numeric keypad))
		KP_backspace:       sdl.SCANCODE_KP_BACKSPACE,       // Keypad Backspace (the Backspace key (numeric keypad))
		KP_a:               sdl.SCANCODE_KP_A,               // Keypad A (the A key (numeric keypad))
		KP_b:               sdl.SCANCODE_KP_B,               // Keypad B (the B key (numeric keypad))
		KP_c:               sdl.SCANCODE_KP_C,               // Keypad C (the C key (numeric keypad))
		KP_d:               sdl.SCANCODE_KP_D,               // Keypad D (the D key (numeric keypad))
		KP_e:               sdl.SCANCODE_KP_E,               // Keypad E (the E key (numeric keypad))
		KP_f:               sdl.SCANCODE_KP_F,               // Keypad F (the F key (numeric keypad))
		KP_xor:             sdl.SCANCODE_KP_XOR,             // Keypad XOR (the XOR key (numeric keypad))
		KP_power:           sdl.SCANCODE_KP_POWER,           // Keypad ^ (the Power key (numeric keypad))
		KP_percent:         sdl.SCANCODE_KP_PERCENT,         // Keypad % (the Percent key (numeric keypad))
		KP_less:            sdl.SCANCODE_KP_LESS,            // Keypad < (the Less key (numeric keypad))
		KP_greater:         sdl.SCANCODE_KP_GREATER,         // Keypad > (the Greater key (numeric keypad))
		KP_ampersand:       sdl.SCANCODE_KP_AMPERSAND,       // Keypad & (the & key (numeric keypad))
		KP_dblampersand:    sdl.SCANCODE_KP_DBLAMPERSAND,    // Keypad && (the && key (numeric keypad))
		KP_verticalbar:     sdl.SCANCODE_KP_VERTICALBAR,     // Keypad | (the | key (numeric keypad))
		KP_dblverticalbar:  sdl.SCANCODE_KP_DBLVERTICALBAR,  // Keypad || (the || key (numeric keypad))
		KP_colon:           sdl.SCANCODE_KP_COLON,           // Keypad : (the : key (numeric keypad))
		KP_hash:            sdl.SCANCODE_KP_HASH,            // Keypad # (the # key (numeric keypad))
		KP_space:           sdl.SCANCODE_KP_SPACE,           // Keypad Space (the Space key (numeric keypad))
		KP_at:              sdl.SCANCODE_KP_AT,              // Keypad @ (the @ key (numeric keypad))
		KP_exclam:          sdl.SCANCODE_KP_EXCLAM,          // Keypad ! (the ! key (numeric keypad))
		KP_memstore:        sdl.SCANCODE_KP_MEMSTORE,        // Keypad MemStore (the Mem Store key (numeric keypad))
		KP_memrecall:       sdl.SCANCODE_KP_MEMRECALL,       // Keypad MemRecall (the Mem Recall key (numeric keypad))
		KP_memclear:        sdl.SCANCODE_KP_MEMCLEAR,        // Keypad MemClear (the Mem Clear key (numeric keypad))
		KP_memadd:          sdl.SCANCODE_KP_MEMADD,          // Keypad MemAdd (the Mem Add key (numeric keypad))
		KP_memsubtract:     sdl.SCANCODE_KP_MEMSUBTRACT,     // Keypad MemSubtract (the Mem Subtract key (numeric keypad))
		KP_memmultiply:     sdl.SCANCODE_KP_MEMMULTIPLY,     // Keypad MemMultiply (the Mem Multiply key (numeric keypad))
		KP_memdivide:       sdl.SCANCODE_KP_MEMDIVIDE,       // Keypad MemDivide (the Mem Divide key (numeric keypad))
		KP_plusminus:       sdl.SCANCODE_KP_PLUSMINUS,       // Keypad +/- (the +/- key (numeric keypad))
		KP_clear:           sdl.SCANCODE_KP_CLEAR,           // Keypad Clear (the Clear key (numeric keypad))
		KP_clearentry:      sdl.SCANCODE_KP_CLEARENTRY,      // Keypad ClearEntry (the Clear Entry key (numeric keypad))
		KP_binary:          sdl.SCANCODE_KP_BINARY,          // Keypad Binary (the Binary key (numeric keypad))
		KP_octal:           sdl.SCANCODE_KP_OCTAL,           // Keypad Octal (the Octal key (numeric keypad))
		KP_decimal:         sdl.SCANCODE_KP_DECIMAL,         // Keypad Decimal (the Decimal key (numeric keypad))
		KP_hexadecimal:     sdl.SCANCODE_KP_HEXADECIMAL,     // Keypad Hexadecimal (the Hexadecimal key (numeric keypad))

		Lctrl:          sdl.SCANCODE_LCTRL,          // Left Ctrl
		Lshift:         sdl.SCANCODE_LSHIFT,         // Left Shift
		Lalt:           sdl.SCANCODE_LALT,           // Left Alt (alt, option)
		Lgui:           sdl.SCANCODE_LGUI,           // Left GUI (windows, command (apple), meta)
		Rctrl:          sdl.SCANCODE_RCTRL,          // Right Ctrl
		Rshift:         sdl.SCANCODE_RSHIFT,         // Right Shift
		Ralt:           sdl.SCANCODE_RALT,           // Right Alt (alt gr, option)
		Rgui:           sdl.SCANCODE_RGUI,           // Right GUI (windows, command (apple), meta)
		Mode:           sdl.SCANCODE_MODE,           // ModeSwitch (I'm not sure if this is really not covered by any of the above, but since there's a special KMOD_MODE for it I'm adding it here)
		Audionext:      sdl.SCANCODE_AUDIONEXT,      // AudioNext (the Next Track media key)
		Audioprev:      sdl.SCANCODE_AUDIOPREV,      // AudioPrev (the Previous Track media key)
		Audiostop:      sdl.SCANCODE_AUDIOSTOP,      // AudioStop (the Stop media key)
		Audioplay:      sdl.SCANCODE_AUDIOPLAY,      // AudioPlay (the Play media key)
		Audiomute:      sdl.SCANCODE_AUDIOMUTE,      // AudioMute (the Mute volume key)
		Mediaselect:    sdl.SCANCODE_MEDIASELECT,    // MediaSelect (the Media Select key)
		Www:            sdl.SCANCODE_WWW,            // WWW (the WWW/World Wide Web key)
		Mail:           sdl.SCANCODE_MAIL,           // Mail (the Mail/eMail key)
		Calculator:     sdl.SCANCODE_CALCULATOR,     // Calculator (the Calculator key)
		Computer:       sdl.SCANCODE_COMPUTER,       // Computer (the My Computer key)
		AC_search:      sdl.SCANCODE_AC_SEARCH,      // AC Search (the Search key (application control keypad))
		AC_home:        sdl.SCANCODE_AC_HOME,        // AC Home (the Home key (application control keypad))
		AC_back:        sdl.SCANCODE_AC_BACK,        // AC Back (the Back key (application control keypad))
		AC_forward:     sdl.SCANCODE_AC_FORWARD,     // AC Forward (the Forward key (application control keypad))
		AC_stop:        sdl.SCANCODE_AC_STOP,        // AC Stop (the Stop key (application control keypad))
		AC_refresh:     sdl.SCANCODE_AC_REFRESH,     // AC Refresh (the Refresh key (application control keypad))
		AC_bookmarks:   sdl.SCANCODE_AC_BOOKMARKS,   // AC Bookmarks (the Bookmarks key (application control keypad))
		Brightnessdown: sdl.SCANCODE_BRIGHTNESSDOWN, // BrightnessDown (the Brightness Down key)
		Brightnessup:   sdl.SCANCODE_BRIGHTNESSUP,   // BrightnessUp (the Brightness Up key)
		Displayswitch:  sdl.SCANCODE_DISPLAYSWITCH,  // DisplaySwitch (display mirroring/dual display switch, video mode switch)
		Kbdillumtoggle: sdl.SCANCODE_KBDILLUMTOGGLE, // KBDIllumToggle (the Keyboard Illumination Toggle key)
		Kbdillumdown:   sdl.SCANCODE_KBDILLUMDOWN,   // KBDIllumDown (the Keyboard Illumination Down key)
		Kbdillumup:     sdl.SCANCODE_KBDILLUMUP,     // KBDIllumUp (the Keyboard Illumination Up key)
		Eject:          sdl.SCANCODE_EJECT,          // Eject (the Eject key)
		Sleep:          sdl.SCANCODE_SLEEP,          // Sleep (the Sleep key)
		App1:           sdl.SCANCODE_APP1,
		App2:           sdl.SCANCODE_APP2,
	}
}

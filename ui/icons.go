package ui

import (
	"fmt"
	"os"
)

// IconPack represents a set of themed icons
type IconPack struct {
	Name        string
	Description string
	Icons       IconSet
}

// IconSet contains all the icons used in WipeOs
type IconSet struct {
	// Main operations
	Wipe        string
	Clean       string
	Forensic    string
	Start       string
	Version     string
	Status      string
	Help        string
	Exit        string
	Clear       string

	// Results
	Success     string
	Error       string
	Warning     string
	Info        string
	Progress    string

	// Forensic operations
	SystemLogs  string
	Registry    string
	Prefetch    string
	Thumbnails  string
	EventLogs   string
	MFT         string
	Shadows     string
	Memory      string
	Swap        string
	FreeSpace   string

	// File operations
	File        string
	Directory   string
	Browser     string
	Temp        string
	Cache       string
	Downloads   string

	// Security
	Shield      string
	Lock        string
	Key         string
	Danger      string
	Safe        string

	// UI Elements
	Prompt      string
	Banner      string
	Summary     string
	Interactive string
}

// Available icon packs
var (
	// Classic pack (original emojis)
	ClassicPack = IconPack{
		Name:        "classic",
		Description: "Original emoji style",
		Icons: IconSet{
			Wipe:        "🗑️",
			Clean:       "🧽",
			Forensic:    "🔍",
			Start:       "🚀",
			Version:     "ℹ️",
			Status:      "📊",
			Help:        "❓",
			Exit:        "🚪",
			Clear:       "🧹",
			Success:     "✅",
			Error:       "❌",
			Warning:     "⚠️",
			Info:        "💡",
			Progress:    "⚡",
			SystemLogs:  "🗂️",
			Registry:    "📋",
			Prefetch:    "⚡",
			Thumbnails:  "🖼️",
			EventLogs:   "📋",
			MFT:         "🗃️",
			Shadows:     "👥",
			Memory:      "🧠",
			Swap:        "💾",
			FreeSpace:   "🗂️",
			File:        "📄",
			Directory:   "📁",
			Browser:     "🌐",
			Temp:        "📂",
			Cache:       "💾",
			Downloads:   "⬇️",
			Shield:      "🛡️",
			Lock:        "🔒",
			Key:         "🔑",
			Danger:      "🚨",
			Safe:        "🛡️",
			Prompt:      "▶️",
			Banner:      "🧹",
			Summary:     "📊",
			Interactive: "🔥",
		},
	}

	// Cyber pack (cyberpunk/hacker theme)
	CyberPack = IconPack{
		Name:        "cyber",
		Description: "Cyberpunk hacker aesthetic",
		Icons: IconSet{
			Wipe:        "⟨⟩",
			Clean:       "⦿",
			Forensic:    "◉",
			Start:       "▲",
			Version:     "◈",
			Status:      "◐",
			Help:        "◈",
			Exit:        "⬢",
			Clear:       "⟐",
			Success:     "◉",
			Error:       "⬢",
			Warning:     "◭",
			Info:        "◒",
			Progress:    "⟪⟫",
			SystemLogs:  "⟨╱⟩",
			Registry:    "◉◉",
			Prefetch:    "⟪⟫",
			Thumbnails:  "◫",
			EventLogs:   "◐◑",
			MFT:         "⬡",
			Shadows:     "◎◎",
			Memory:      "◯",
			Swap:        "⬢⬡",
			FreeSpace:   "⟨⟩",
			File:        "◧",
			Directory:   "⬢",
			Browser:     "◉",
			Temp:        "◈",
			Cache:       "⬢",
			Downloads:   "▼",
			Shield:      "◈",
			Lock:        "⬢",
			Key:         "◎",
			Danger:      "◯",
			Safe:        "◉",
			Prompt:      "▶",
			Banner:      "⟨╱⟩",
			Summary:     "◐",
			Interactive: "◯",
		},
	}

	// Military pack (tactical/military theme)
	MilitaryPack = IconPack{
		Name:        "military",
		Description: "Military tactical symbols",
		Icons: IconSet{
			Wipe:        "⚔",
			Clean:       "⊗",
			Forensic:    "⊙",
			Start:       "▲",
			Version:     "◎",
			Status:      "⊕",
			Help:        "⊖",
			Exit:        "◈",
			Clear:       "⊗",
			Success:     "⊕",
			Error:       "⊘",
			Warning:     "⚠",
			Info:        "◈",
			Progress:    "◉",
			SystemLogs:  "⚔",
			Registry:    "⊗",
			Prefetch:    "◉",
			Thumbnails:  "⊙",
			EventLogs:   "⊕",
			MFT:         "◎",
			Shadows:     "⚔⚔",
			Memory:      "⊗",
			Swap:        "⊙⊙",
			FreeSpace:   "⚔",
			File:        "◦",
			Directory:   "◈",
			Browser:     "⊙",
			Temp:        "⊗",
			Cache:       "⊕",
			Downloads:   "▼",
			Shield:      "◈",
			Lock:        "⊗",
			Key:         "⊙",
			Danger:      "⚠",
			Safe:        "⊕",
			Prompt:      "▶",
			Banner:      "⚔",
			Summary:     "⊕",
			Interactive: "◉",
		},
	}

	// Minimal pack (clean minimal design)
	MinimalPack = IconPack{
		Name:        "minimal",
		Description: "Clean minimal symbols",
		Icons: IconSet{
			Wipe:        "×",
			Clean:       "○",
			Forensic:    "●",
			Start:       "▲",
			Version:     "◇",
			Status:      "◯",
			Help:        "?",
			Exit:        "◯",
			Clear:       "○",
			Success:     "✓",
			Error:       "✗",
			Warning:     "!",
			Info:        "•",
			Progress:    "→",
			SystemLogs:  "×",
			Registry:    "○",
			Prefetch:    "→",
			Thumbnails:  "◇",
			EventLogs:   "○",
			MFT:         "●",
			Shadows:     "××",
			Memory:      "○",
			Swap:        "○○",
			FreeSpace:   "×",
			File:        "•",
			Directory:   "◇",
			Browser:     "○",
			Temp:        "×",
			Cache:       "●",
			Downloads:   "↓",
			Shield:      "◇",
			Lock:        "●",
			Key:         "○",
			Danger:      "!",
			Safe:        "✓",
			Prompt:      "→",
			Banner:      "×",
			Summary:     "◯",
			Interactive: "●",
		},
	}

	// Matrix pack (Matrix movie theme)
	MatrixPack = IconPack{
		Name:        "matrix",
		Description: "Matrix digital rain style",
		Icons: IconSet{
			Wipe:        "⊗",
			Clean:       "◯",
			Forensic:    "◉",
			Start:       "▶",
			Version:     "◎",
			Status:      "◐",
			Help:        "◈",
			Exit:        "◯",
			Clear:       "⊗",
			Success:     "◉",
			Error:       "⊘",
			Warning:     "⚠",
			Info:        "◈",
			Progress:    "◐",
			SystemLogs:  "⊗",
			Registry:    "◯",
			Prefetch:    "◉",
			Thumbnails:  "◎",
			EventLogs:   "◐",
			MFT:         "◈",
			Shadows:     "◯◯",
			Memory:      "◉",
			Swap:        "◎◎",
			FreeSpace:   "⊗",
			File:        "◦",
			Directory:   "◈",
			Browser:     "◉",
			Temp:        "⊗",
			Cache:       "◎",
			Downloads:   "↓",
			Shield:      "◈",
			Lock:        "⊗",
			Key:         "◉",
			Danger:      "⚠",
			Safe:        "◎",
			Prompt:      "▶",
			Banner:      "⊗",
			Summary:     "◐",
			Interactive: "◉",
		},
	}

	// Retro pack (80s/90s gaming aesthetic)
	RetroPack = IconPack{
		Name:        "retro",
		Description: "80s/90s retro gaming vibes",
		Icons: IconSet{
			Wipe:        "⚡",
			Clean:       "★",
			Forensic:    "◆",
			Start:       "►",
			Version:     "♦",
			Status:      "♠",
			Help:        "?",
			Exit:        "◄",
			Clear:       "※",
			Success:     "♥",
			Error:       "☠",
			Warning:     "!",
			Info:        "♪",
			Progress:    "♫",
			SystemLogs:  "⚡",
			Registry:    "♦",
			Prefetch:    "►",
			Thumbnails:  "♠",
			EventLogs:   "♣",
			MFT:         "♥",
			Shadows:     "☠☠",
			Memory:      "★",
			Swap:        "♠♠",
			FreeSpace:   "⚡",
			File:        "•",
			Directory:   "♦",
			Browser:     "★",
			Temp:        "⚡",
			Cache:       "♠",
			Downloads:   "▼",
			Shield:      "♦",
			Lock:        "★",
			Key:         "♥",
			Danger:      "☠",
			Safe:        "♠",
			Prompt:      "►",
			Banner:      "⚡",
			Summary:     "♠",
			Interactive: "★",
		},
	}

	// Neon pack (bright electronic aesthetic)
	NeonPack = IconPack{
		Name:        "neon",
		Description: "Bright neon electronic style",
		Icons: IconSet{
			Wipe:        "◢",
			Clean:       "◣",
			Forensic:    "◤",
			Start:       "◥",
			Version:     "◢",
			Status:      "◣",
			Help:        "◤",
			Exit:        "◥",
			Clear:       "◢",
			Success:     "◣",
			Error:       "◤",
			Warning:     "◥",
			Info:        "◢",
			Progress:    "◣",
			SystemLogs:  "◢",
			Registry:    "◣",
			Prefetch:    "◤",
			Thumbnails:  "◥",
			EventLogs:   "◢",
			MFT:         "◣",
			Shadows:     "◤◤",
			Memory:      "◥",
			Swap:        "◢◢",
			FreeSpace:   "◣",
			File:        "◦",
			Directory:   "◢",
			Browser:     "◣",
			Temp:        "◤",
			Cache:       "◥",
			Downloads:   "▼",
			Shield:      "◢",
			Lock:        "◣",
			Key:         "◤",
			Danger:      "◥",
			Safe:        "◢",
			Prompt:      "►",
			Banner:      "◣",
			Summary:     "◤",
			Interactive: "◥",
		},
	}
)

// currentIconPack holds the active icon pack
var currentIconPack = ClassicPack

// SetIconPack changes the active icon pack
func SetIconPack(packName string) error {
	switch packName {
	case "classic":
		currentIconPack = ClassicPack
	case "cyber":
		currentIconPack = CyberPack
	case "military":
		currentIconPack = MilitaryPack
	case "minimal":
		currentIconPack = MinimalPack
	case "matrix":
		currentIconPack = MatrixPack
	case "retro":
		currentIconPack = RetroPack
	case "neon":
		currentIconPack = NeonPack
	default:
		return fmt.Errorf("unknown icon pack: %s", packName)
	}
	
	// Save preference to environment or config file
	os.Setenv("WIPEOS_ICON_PACK", packName)
	return nil
}

// GetCurrentIconPack returns the currently active icon pack
func GetCurrentIconPack() IconPack {
	// Try to load from environment
	if packName := os.Getenv("WIPEOS_ICON_PACK"); packName != "" {
		SetIconPack(packName) // This will validate and set if valid
	}
	return currentIconPack
}

// GetAllIconPacks returns all available icon packs
func GetAllIconPacks() []IconPack {
	return []IconPack{
		ClassicPack,
		CyberPack,
		MilitaryPack,
		MinimalPack,
		MatrixPack,
		RetroPack,
		NeonPack,
	}
}

// Icon helper functions for easy access
func Icon() *IconSet {
	return &currentIconPack.Icons
}

// Specific icon getters
func IconWipe() string        { return currentIconPack.Icons.Wipe }
func IconClean() string       { return currentIconPack.Icons.Clean }
func IconForensic() string    { return currentIconPack.Icons.Forensic }
func IconStart() string       { return currentIconPack.Icons.Start }
func IconSuccess() string     { return currentIconPack.Icons.Success }
func IconError() string       { return currentIconPack.Icons.Error }
func IconWarning() string     { return currentIconPack.Icons.Warning }
func IconInfo() string        { return currentIconPack.Icons.Info }
func IconShield() string      { return currentIconPack.Icons.Shield }
func IconDanger() string      { return currentIconPack.Icons.Danger }
func IconBanner() string      { return currentIconPack.Icons.Banner }
func IconInteractive() string { return currentIconPack.Icons.Interactive } 
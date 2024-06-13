package ui

import (
	"github.com/charmbracelet/lipgloss"
)

func Canvas(width, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#fff")).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)
}

func Button() lipgloss.Style {
	textStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(White()).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(Purple()).
		Align(lipgloss.Center).
		PaddingTop(1).
		Width(22).
		Height(3)

	return textStyle
}

func White() lipgloss.Color {
	return lipgloss.Color("#FAFAFA")
}

func Neutral() lipgloss.Color {
	return lipgloss.Color("#383838")
}

func Purple() lipgloss.Color {
	return lipgloss.Color("#7D56F4")
}

func Pink() lipgloss.Color {
	return lipgloss.Color("#f23cc7")
}

func Red() lipgloss.Color {
	return lipgloss.Color("#ff5291")
}

func Orange() lipgloss.Color {
	return lipgloss.Color("#ff8c64")
}

func Gold() lipgloss.Color {
	return lipgloss.Color("#ffc652")
}

func Yellow() lipgloss.Color {
	return lipgloss.Color("#f9f871")
}

//func Container() lipgloss.Style {
//    return lipgloss.NewStyle().
//        Background(lipgloss.Color("#FAFAFA")).
//        Foreground(lipgloss.Color("#7D56F4")).
//        Padding(2, 4, 2, 4).
//        Border(lipgloss.RoundedBorder())
//}

//func Grid() string {
//    buttonStyle := lipgloss.NewStyle().
//        Foreground(lipgloss.Color("#FFF7DB")).
//        Background(lipgloss.Color("#888B7E")).
//        Padding(0, 3).
//        MarginTop(1)

//    activeButtonStyle := buttonStyle.
//        Foreground(lipgloss.Color("#FFF7DB")).
//        Background(lipgloss.Color("#F25D94")).
//        MarginRight(2).
//        Underline(true)
//    okButton := activeButtonStyle.Render("Yes")
//    cancelButton := buttonStyle.Render("Maybe")

//    question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render("Are you sure you want to eat marmalade?")
//    buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, cancelButton)
//    ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

//    dialogBoxStyle := lipgloss.NewStyle().
//        Border(lipgloss.RoundedBorder()).
//        BorderForeground(lipgloss.Color("#874BFD")).
//        Padding(1, 0).
//        BorderTop(true).
//        BorderLeft(true).
//        BorderRight(true).
//        BorderBottom(true)

//    dialog := lipgloss.Place(80, 40,
//        lipgloss.Center,
//        lipgloss.Center,
//        dialogBoxStyle.Render(ui),
//        lipgloss.WithWhitespaceChars("猫"),
//        lipgloss.WithWhitespaceForeground(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}),
//    )

//    return dialog
//}

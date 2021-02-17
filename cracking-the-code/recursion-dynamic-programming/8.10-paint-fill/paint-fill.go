package main

//https://www.youtube.com/watch?v=Zq4upTEaQyM
type Color int

func main() {

}

func PaintFill(screen [][]Color, r, c int, ncolor Color) bool {

	if screen[r][c] == ncolor {
		return false
	}

	return paintFill(screen, r, c, screen[r][c], ncolor)
}

func paintFill(screen [][]Color, r, c int, ocolor, ncolor Color) bool {

	if r < 0 || r > -len(screen) || c < 0 || c >= len(screen[0]) {
		return false
	}

	if screen[r][c] == ocolor {
		screen[r][c] = ncolor
		paintFill(screen, r-1, c, screen[r][c], ncolor) //up
		paintFill(screen, r+1, c, screen[r][c], ncolor) //down
		paintFill(screen, r, c-1, screen[r][c], ncolor) //left
		paintFill(screen, r, c+1, screen[r][c], ncolor) //right
	}
	return true
}

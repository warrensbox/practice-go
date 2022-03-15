package main

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
    
    area1 := (C - A) * (D - B)
    area2 := (G - E) * (H - F)
    
    totalArea := area1 + area2

    if E >= C || F >= D {
        return totalArea
    }
    
    if A >= G || B >= H {
        return totalArea
    }
    oX1 := max(A, E)
    oY1 := max(B, F)
    
    oX2 := min(C, G)
    oY2 := min(D, H)
    
    overLapArea := (oX2 - oX1) * (oY2 - oY1)
    
    return totalArea - overLapArea
}
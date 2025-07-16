// v0.08
// 07/09/25
// by sayzenberg

// VFD Display:
/*
       x_=_0_______________________________x_=_280
y = 0  |                                         |
       |                                         |
y = 32 |_________________________________________|

*/

package main

import (
	"fmt"
	"net"
	"time"
)

var sb_conn net.Dialer

// IIRC?
var MAX_X = 279
var MAX_Y = 15 // ????

func main() {
	fmt.Println("Roku SoundBridge GFX Library")

	// Connect to the SB and enter sketch mode
	// This is the "net.Conn" part AFTER the "net.Dial/Dialer" is set up!! (likew "new_conn" in sb_connect()
	sb_conn := sb_connect()
	
	// Draw
	print_text("hello!", sb_conn)

	// Clear
	clear_screen(sb_conn)

	// Line
	draw_line(0, 0, 20, 20, sb_conn)

	// Some tests
	clear_screen(sb_conn)
	draw_point(0, 0, sb_conn)
        draw_point(0, MAX_Y, sb_conn)
        time.Sleep(1 * time.Second)
        draw_point(MAX_X, 0, sb_conn)
        draw_point(MAX_X, MAX_Y, sb_conn)
        time.Sleep(1 * time.Second)

	// Square
	// Args: Filled Rectangle, X, Y, Width, Height, conn ID
	draw_rectangle(true, 10, 0, 10, 10, sb_conn)

	// Exit
	sb_disconnect(sb_conn)
}

// Connectivity
func sb_connect() (butts net.Conn) {
	new_conn, err := sb_conn.Dial("tcp", "192.168.0.228:4444")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	time.Sleep(1 * time.Second) // heh
	fmt.Fprintf(new_conn, "sketch\r\n")
	fmt.Println("Connected!")

	// Return the net.Conn "handler"(?)
	return new_conn
}

func sb_disconnect(butts net.Conn) () {
        fmt.Println("Exiting...")
        fmt.Fprintf(butts, "exit\r\nexit\r\n")
	butts.Close()
}

func clear_screen(butts net.Conn) {
	fmt.Fprintf(butts, "clear\r\n")
	time.Sleep(1 * time.Second)
}

func print_text(text string, butts net.Conn) () {
	fmt.Println("Sending sketch")
	fmt.Fprintf(butts, "text 5 5 %s\r\n", text)
	time.Sleep(1 * time.Second)
}

func draw_line(x1 int, y1 int, x2 int, y2 int, butts net.Conn) {
	fmt.Fprintf(butts, "line %d %d %d %d\r\n", x1, y1, x2, y2)
	time.Sleep(1 * time.Second)
}

func draw_point(x int, y int, butts net.Conn) () {
        fmt.Fprintf(butts, "point %d %d\r\n", x, y)
        time.Sleep(1 * time.Second)
}

func draw_rectangle(filled bool, x int, y int, width int, height int, butts net.Conn) {
	var tool string
	if filled {
		tool = "rect"
	} else {
		tool = "framerect"
	}
	fmt.Fprintf(butts, "%s %d %d %d %d\r\n", tool, x, y, width, height)
	time.Sleep(1 * time.Second)
}


// v0.06
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

func main() {
	fmt.Println("Roku SoundBridge GFX Library")

	// Connect to the SB and enter sketch mode
	// This is the "net.Conn" part AFTER the "net.Dial/Dialer" is set up!! (likew "new_conn" in sb_connect()
	sb_conn := sb_connect()
	
	// Draw
	print_text("hello!", sb_conn)

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

func print_text(text string, butts net.Conn) () {
	fmt.Println("Sending sketch")
	fmt.Fprintf(butts, "text 5 5 %s\r\n", text)
	time.Sleep(1 * time.Second)
}
